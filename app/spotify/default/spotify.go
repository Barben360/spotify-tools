package spotifydefault

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/Barben360/spotify-tools/app/spotify"
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
