package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Barben360/spotify-tools/app/services/logger"
	"github.com/Barben360/spotify-tools/app/spotify"
	spotifydefault "github.com/Barben360/spotify-tools/app/spotify/default"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// DefaultConfigCacheFilePath is the default path where auth tokens are cached on disk.
var DefaultConfigCacheFilePath = spotifydefault.DefaultConfigCacheFilePath

// ExitCodeError is an error that requests a specific process exit code.
// The error message was already printed to stdout before this error is returned,
// so callers should exit with Code without printing anything further.
type ExitCodeError struct {
	Code int
	msg  string
}

func (e *ExitCodeError) Error() string {
	return e.msg
}

type App struct {
	log      *logger.Logger
	features *features
}

func New(ctx context.Context, cfg *AppConfig) (*App, error) {
	// Validate config
	validate := validator.New()
	err := validate.Struct(cfg)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			errs := make([]error, len(validationErrors))
			for i, validationError := range validationErrors {
				errs[i] = validationError
			}
			return nil, errors.Join(errs...)
		}
		return nil, err
	}

	// Initialize logger
	var log *logger.Logger
	if cfg.DevMode {
		logZap, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		log = &logger.Logger{Logger: logZap}
	} else {
		logZap, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		log = &logger.Logger{Logger: logZap}
	}

	ctx = log.ToContext(ctx)

	// Initialize features
	spotify := spotifydefault.New(
		ctx,
		cfg.SpotifyAppClientID,
		cfg.SpotifyAppClientSecret,
		cfg.PublicAPIEndpoint,
		cfg.ServerListenPort,
		cfg.ConfigCacheFilePath,
	)
	features := &features{
		spotify: spotify,
	}

	return &App{
		log:      log,
		features: features,
	}, nil
}

// LogoutFromDisk removes the cached Spotify auth tokens from disk without requiring
// an initialized App instance. Useful for logout when credentials may not be available.
// It never fails even if the user is not logged in.
func LogoutFromDisk(configCacheFilePath string) {
	_ = os.Remove(configCacheFilePath)
}

func (a *App) Login(ctx context.Context) error {
	ctx = a.log.ToContext(ctx)
	return a.features.spotify.Login(ctx)
}

func (a *App) Logout(ctx context.Context) {
	ctx = a.log.ToContext(ctx)
	a.features.spotify.Logout(ctx)
}

// AuthStatus prints the current authentication status to stdout and returns nil on success.
// Returns an *ExitCodeError with Code=2 if the user is not authenticated or tokens are invalid.
func (a *App) AuthStatus(ctx context.Context) error {
	ctx = a.log.ToContext(ctx)
	loggedIn, err := a.features.spotify.AuthStatus(ctx)
	if !loggedIn {
		fmt.Println("Not authenticated (no tokens found)")
		return &ExitCodeError{Code: 2, msg: "not authenticated"}
	}
	if err != nil {
		fmt.Printf("Not authenticated (token refresh failed: %v)\n", err)
		return &ExitCodeError{Code: 2, msg: "not authenticated"}
	}
	fmt.Println("Authenticated (tokens valid)")
	return nil
}

func (a *App) GetPlaylist(ctx context.Context, playlistID string) error {
	ctx = a.log.ToContext(ctx)
	resp, err := a.features.spotify.GetPlaylist(ctx, playlistID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	fmt.Println(string(respJSON))
	return nil
}

func (a *App) GetShow(ctx context.Context, showID string) error {
	ctx = a.log.ToContext(ctx)
	resp, err := a.features.spotify.GetShow(ctx, showID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	fmt.Println(string(respJSON))
	return nil
}

func (a *App) PlaylistFilter(ctx context.Context, config *spotify.PlaylistFilterConfig) error {
	ctx = a.log.ToContext(ctx)
	err := a.features.spotify.UpdatePlaylistFilter(ctx, config)
	if err != nil {
		return err
	}
	return nil
}
