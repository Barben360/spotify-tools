package app

import (
	"github.com/Barben360/spotify-tools/app/authtest"
	"github.com/Barben360/spotify-tools/app/playlistfilter"
)

type AppConfig struct {
	DevMode                bool
	SpotifyAppClientID     string                                `validate:"required"`
	SpotifyAppClientSecret string                                `validate:"required"`
	PublicAPIEndpoint      string                                `validate:"required"`
	ServerListenPort       uint16                                `validate:"required"`
	PlaylistFiltersConfig  *playlistfilter.PlaylistFiltersConfig `validate:"omitempty"`
}

type Features struct {
	AuthTest authtest.AuthTester
}
