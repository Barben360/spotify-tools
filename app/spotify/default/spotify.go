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
	spotifyclient "github.com/Barben360/spotify-tools/app/spotify/default/spotifyclient"
	"go.uber.org/zap"
)

//go:generate openapi-generator-cli generate -g go --additional-properties=disallowAdditionalPropertiesIfNotPresent=false,packageName=spotifyclient,isGoSubmodule=true,structPrefix=true -o spotifyclient -i assets/spotify-web-api/fixed-spotify-open-api.yml --git-repo-id=spotify-tools/app/default --git-user-id Barben360
//go:generate rm -rf spotifyclient/go.mod spotifyclient/go.sum spotifyclient/git_push.sh spotifyclient/.openapi-generator/VERSION spotifyclient/test

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
	configCacheFilePath    string
	httpClient             *http.Client
	spotifyClient          *spotifyclient.APIClient
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
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		configCacheFilePath: filepath.Join(os.TempDir(), ".spotify-tools-cache.json"),
	}

	spotifyClientConfig := spotifyclient.NewConfiguration()
	spotifyClientConfig.HTTPClient = ret.httpClient
	ret.spotifyClient = spotifyclient.NewAPIClient(spotifyClientConfig)

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
