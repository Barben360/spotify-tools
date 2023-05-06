package spotifydefault

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Barben360/spotify-playlist-filters/app/services/logger"
	"github.com/Barben360/spotify-playlist-filters/app/spotify"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Spotify struct {
	accessToken            string
	refreshToken           string
	authLock               sync.Mutex
	spotifyAppClientID     string
	spotifyAppClientSecret string
	authHeader             string
	redirectURI            string
	publicAPIEndpoint      string
	serverListenPort       uint16
	client                 *http.Client
}

func New(
	ctx context.Context,
	spotifyAppClientID string,
	spotifyAppClientSecret string,
	publicAPIEndpoint string,
	serverListenPort uint16,
) spotify.Spotifier {
	return &Spotify{
		authLock:               sync.Mutex{},
		spotifyAppClientID:     spotifyAppClientID,
		spotifyAppClientSecret: spotifyAppClientSecret,
		authHeader: "Basic " + base64.StdEncoding.EncodeToString(
			[]byte(
				fmt.Sprintf("%s:%s",
					spotifyAppClientID,
					spotifyAppClientSecret,
				),
			),
		),
		redirectURI:       strings.TrimSuffix(publicAPIEndpoint, "/") + "/authorize",
		publicAPIEndpoint: publicAPIEndpoint,
		serverListenPort:  serverListenPort,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *Spotify) GetUserToken(ctx context.Context) (string, error) {
	if s.accessToken != "" {
		return s.accessToken, nil
	}
	return s.GetNewUserToken(ctx)
}

func (s *Spotify) GetNewUserToken(ctx context.Context) (string, error) {
	s.authLock.Lock()
	defer s.authLock.Unlock()
	if s.refreshToken != "" {
		logger.FromContext(ctx).Debug("refreshing token")
		err := s.requestTokenRefresh(ctx)
		if err != nil {
			return "", err
		}
		return s.accessToken, nil
	}
	logger.FromContext(ctx).Debug("no refresh token, doing authorization flow")
	code, err := s.requestAuthorizationCode(ctx)
	if err != nil {
		return "", err
	}
	logger.FromContext(ctx).Debug("got authorization code, requesting token")
	err = s.requestToken(ctx, code)
	if err != nil {
		return "", err
	}
	return s.accessToken, nil
}

func (s *Spotify) requestAuthorizationCode(ctx context.Context) (string, error) {
	state := uuid.New().String()

	url := fmt.Sprintf("https://accounts.spotify.com/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
		url.PathEscape(s.spotifyAppClientID),
		url.PathEscape(strings.Join([]string{
			"playlist-read-private",
			"playlist-read-collaborative",
			"playlist-modify-private",
			"playlist-modify-public",
		}, " ")),
		url.PathEscape(s.redirectURI),
		url.PathEscape(state),
	)

	timeout := 1 * time.Minute
	logger.FromContext(ctx).Info("Please open this URL in a web browser to authorize the app to access your Spotify account", zap.String("url", url))

	ctxAuth, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Start a web server to listen for the callback
	var code string
	var errReturn string
	errChan := make(chan error, 1)

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{s.publicAPIEndpoint},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/authorize", func(c echo.Context) error {
		// Reading query parameters "code" and "state"
		code = c.QueryParam("code")
		errReturn = c.QueryParam("error")
		receivedState := c.QueryParam("state")
		if errReturn != "" {
			defer func() {
				errChan <- fmt.Errorf("authorization failed with error: %q", errReturn)
			}()
			return c.String(http.StatusOK, fmt.Sprintf("Authorization failed: error received from Spotify: %q. Please retry", errReturn))
		}
		if receivedState != state {
			defer func() {
				errChan <- errors.New("invalid state received")
			}()
			return c.String(http.StatusOK, "Authorization failed: inconsistent state received. Please retry")
		}
		defer func() {
			errChan <- nil
		}()
		return c.String(http.StatusOK, "Authorization successful! You can close this window now.")
	})

	serverErrChan := make(chan error, 1)
	go func() {
		logger.FromContext(ctx).Info("starting authentication server")
		serverErrChan <- e.Start(fmt.Sprintf(":%d", s.serverListenPort))
		logger.FromContext(ctx).Info("authentication server stopped")
	}()

	stopServer := func(ctx context.Context) {
		// Stopping server
		logger.FromContext(ctx).Info("stopping authentication server")
		err := e.Shutdown(ctx)
		if err != nil {
			logger.FromContext(ctx).Error("error while shutting down server, port may not be left free", zap.Error(err))
		}
		err = <-serverErrChan
		if err != nil && err != http.ErrServerClosed {
			logger.FromContext(ctx).Error("error while shutting down server, port may not be left free", zap.Error(err))
		}
	}

	select {
	case err := <-errChan:
		logger.FromContext(ctx).Debug("authorization flow finished")
		ctx, cancel := context.WithTimeout(logger.FromContext(ctx).ToContext(context.Background()), 5*time.Second)
		defer cancel()
		defer stopServer(ctx)
		if err != nil {
			return "", err
		}
		if code == "" {
			return "", errors.New("no authorization code received")
		}
		return code, nil
	case err := <-serverErrChan:
		logger.FromContext(ctx).Debug("authorization flow finished (could not start server)")
		return "", fmt.Errorf("could not start server: %w", err)
	case <-ctxAuth.Done():
		logger.FromContext(ctx).Debug("authorization flow finished (timeout or cancel)")
		ctx, cancel := context.WithTimeout(logger.FromContext(ctx).ToContext(context.Background()), 5*time.Second)
		defer cancel()
		defer stopServer(ctx)
		return "", fmt.Errorf("timeout while waiting for authorization code: %w", ctx.Err())
	}
}

type SpotifyRequestTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Spotify) requestToken(ctx context.Context, authorizationCode string) error {
	logger.FromContext(ctx).Info("requesting access and refresh tokens")
	// create request body
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", authorizationCode)
	data.Set("redirect_uri", s.redirectURI)

	// create request
	req, err := http.NewRequestWithContext(ctx, "POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	// add headers
	req.Header.Set("Authorization", s.authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// send request
	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	// read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		logger.FromContext(ctx).Error("got response body", zap.String("body", string(respBody)))
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// parse response body
	var tokenResponse SpotifyRequestTokenResponse
	err = json.Unmarshal(respBody, &tokenResponse)
	if err != nil {
		return fmt.Errorf("could not parse response body: %w", err)
	}

	logger.FromContext(ctx).Debug("got token response",
		zap.String("token-type", tokenResponse.TokenType),
		zap.Int("expires-in", tokenResponse.ExpiresIn),
		zap.String("scope", tokenResponse.Scope),
	)

	// Checking that access and refresh tokens are present
	if tokenResponse.AccessToken == "" {
		return errors.New("access token is empty")
	}
	s.accessToken = tokenResponse.AccessToken
	if tokenResponse.RefreshToken == "" {
		return errors.New("refresh token is empty")
	}
	s.refreshToken = tokenResponse.RefreshToken
	logger.FromContext(ctx).Info("got access and refresh tokens")
	return nil
}

func (s *Spotify) requestTokenRefresh(ctx context.Context) error {
	logger.FromContext(ctx).Info("refreshing access token")
	// create request body
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", s.refreshToken)

	// create request
	req, err := http.NewRequestWithContext(ctx, "POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}

	// add headers
	req.Header.Set("Authorization", s.authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// send request
	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	// read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		logger.FromContext(ctx).Error("got response body", zap.String("body", string(respBody)))
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	// parse response body
	var tokenResponse SpotifyRequestTokenResponse
	err = json.Unmarshal(respBody, &tokenResponse)
	if err != nil {
		return fmt.Errorf("could not parse response body: %w", err)
	}

	logger.FromContext(ctx).Debug("got token response",
		zap.String("token-type", tokenResponse.TokenType),
		zap.Int("expires-in", tokenResponse.ExpiresIn),
		zap.String("scope", tokenResponse.Scope),
	)

	// Checking that access and refresh tokens are present
	if tokenResponse.AccessToken == "" {
		return errors.New("access token is empty")
	}
	s.accessToken = tokenResponse.AccessToken
	logger.FromContext(ctx).Info("got new access token")
	return nil
}

func (s *Spotify) ResetUserTokens(ctx context.Context) {
	logger.FromContext(ctx).Info("resetting user tokens")
	s.accessToken = ""
	s.refreshToken = ""
	logger.FromContext(ctx).Info("user tokens were reset")
}
