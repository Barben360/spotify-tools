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
	"github.com/Barben360/spotify-tools/app/utils"
	"go.uber.org/zap"
)

//go:generate rm -rf spotifyclient
//go:generate openapi-generator-cli generate -g go --additional-properties=packageName=spotifyclient,isGoSubmodule=true,structPrefix=true -o spotifyclient -i assets/spotify-open-api.yml --git-repo-id=spotify-tools/app/default --git-user-id Barben360
//go:generate rm -rf openapitools.json spotifyclient/go.mod spotifyclient/go.sum spotifyclient/git_push.sh spotifyclient/.openapi-generator/VERSION spotifyclient/test
//go:generate git apply --whitespace=nowarn assets/patches/0001-Make-oneof-less-strict.patch

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
	httpAuthClient         *http.Client
	httpOpenAPIClient      *spotifyclient.APIClient
}

type spotifyConfig struct {
	Auth authConfig `json:"auth,omitempty"`
}

type authConfig struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
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
		httpAuthClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: utils.NewTransport(ctx,
				utils.WithLog(logger.FromContext(ctx), zap.DebugLevel)),
		},
		configCacheFilePath: filepath.Join(os.TempDir(), ".spotify-tools-cache.json"),
	}

	spotifyClientConfig := spotifyclient.NewConfiguration()
	spotifyClientConfig.HTTPClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: utils.NewTransport(ctx,
			utils.WithLog(logger.FromContext(ctx), zap.DebugLevel),
			utils.WithReqTransformer(func(r *http.Request) {
				r.Header.Add("Authorization", "Bearer "+ret.accessToken)
			}),
		),
	}
	ret.httpOpenAPIClient = spotifyclient.NewAPIClient(spotifyClientConfig)

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
