package spotifydefault

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Barben360/spotify-tools/app/services/logger"
	"github.com/Barben360/spotify-tools/app/spotify"
	"go.uber.org/zap"
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
	configCacheFilePath    string
}

type spotifyConfig struct {
	Auth *authConfig `json:"auth"`
}

type authConfig struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func New(
	ctx context.Context,
	spotifyAppClientID string,
	spotifyAppClientSecret string,
	publicAPIEndpoint string,
	serverListenPort uint16,
) spotify.Spotifier {

	ret := &Spotify{
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
		configCacheFilePath: filepath.Join(os.TempDir(), ".spotify-tools-cache.json"),
	}

	// Attempting to get the access and refresh tokens from disk
	if f, err := os.Open(ret.configCacheFilePath); err != nil {
		logger.FromContext(ctx).Info("No cache config file detected", zap.String("path", ret.configCacheFilePath))
	} else {
		var config spotifyConfig
		if err := json.NewDecoder(f).Decode(&config); err != nil {
			logger.FromContext(ctx).Warn("Error while decoding cache config file, it might be outdated", zap.Error(err))
		} else {
			ret.accessToken = config.Auth.AccessToken
			ret.refreshToken = config.Auth.RefreshToken
		}
	}
	return ret
}
