package app

import (
	"github.com/Barben360/spotify-tools/app/authtest"
	"github.com/Barben360/spotify-tools/app/spotify"
)

type AppConfig struct {
	DevMode                bool
	SpotifyAppClientID     string `validate:"required"`
	SpotifyAppClientSecret string `validate:"required"`
	PublicAPIEndpoint      string `validate:"required"`
	ServerListenPort       uint16 `validate:"required"`
}

type features struct {
	spotify  spotify.Spotifier
	authTest authtest.AuthTester
}
