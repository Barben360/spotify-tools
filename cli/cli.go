package cli

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Barben360/spotify-tools/app"

	"github.com/spf13/cobra"
)

var version = "dev"

type CLI struct {
	cmdRoot                *cobra.Command
	spotifyAppClientID     string
	spotifyAppClientSecret string
	devMode                bool
	publicAPIEndpoint      string
	serverListenPort       uint16
	configJSONFilePath     string
	configJSON             string
	daemonMode             bool
	period                 time.Duration
}

func New(ctx context.Context) *CLI {
	c := &CLI{}

	c.cmdRoot = &cobra.Command{
		Use:     "spotify-tools",
		Version: version,
		Short:   "A Spotify automation tool",
	}

	c.cmdRoot.PersistentFlags().StringVarP(&c.spotifyAppClientID, "spotify-app-client-id", "i", "", "Spotify App Client ID")
	c.cmdRoot.PersistentFlags().StringVarP(&c.spotifyAppClientSecret, "spotify-app-client-secret", "s", "", "Spotify App Client Secret")

	c.cmdRoot.PersistentFlags().BoolVar(&c.devMode, "dev", false, "Dev mode")
	c.cmdRoot.PersistentFlags().StringVar(&c.publicAPIEndpoint, "public-api-endpoint", "http://localhost:8080", "Public API endpoint")
	c.cmdRoot.PersistentFlags().Uint16Var(&c.serverListenPort, "server-listen-port", 8080, "Server listen port")

	c.cmdRoot.AddCommand(&cobra.Command{
		Use:   "reset",
		Short: "Reset Spotify user authentication",
		Long:  "Reset Spotify user authentication by deleting the config cache file. Use this command if you want to change the Spotify user or if you have a more general issue with authentication.",
		RunE:  c.runApp(ctx, c.resetHandler),
	})

	c.cmdRoot.AddCommand(&cobra.Command{
		Use:   "get-playlist",
		Short: "Get basic info about a Spotify playlist and its items",
		Args:  cobra.ExactArgs(1),
		RunE:  c.runApp(ctx, c.getPlaylistHandler),
	})

	c.cmdRoot.AddCommand(&cobra.Command{
		Use:   "get-show",
		Short: "Get basic info of a Spotify show and its episodes",
		Args:  cobra.ExactArgs(1),
		RunE:  c.runApp(ctx, c.getShowHandler),
	})

	cmdFilterPlaylists := &cobra.Command{
		Use:   "filter-playlists",
		Short: "Update Spotify playlists based on filters on other playlists/shows",
		RunE:  c.runApp(ctx, c.filterPlaylistsHandler),
	}

	cmdFilterPlaylists.PersistentFlags().StringVarP(&c.configJSONFilePath, "config-file", "f", "", "Path to the config JSON file")
	cmdFilterPlaylists.PersistentFlags().StringVarP(&c.configJSON, "config", "c", "", "Config JSON")
	cmdFilterPlaylists.MarkFlagsMutuallyExclusive("config-file", "config")
	cmdFilterPlaylists.PersistentFlags().BoolVar(&c.daemonMode, "daemon", false, "Run in daemon mode")
	cmdFilterPlaylists.PersistentFlags().DurationVarP(&c.period, "period", "p", 1*time.Hour, "Period between each run in daemon mode")
	c.cmdRoot.AddCommand(cmdFilterPlaylists)

	return c
}

func (c *CLI) Run(ctx context.Context) error {
	return c.cmdRoot.ExecuteContext(ctx)
}

func (c *CLI) runApp(ctx context.Context, handler func(ctx context.Context, appInstance *app.App, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cfg := app.AppConfig{
			DevMode:                c.devMode,
			SpotifyAppClientID:     c.spotifyAppClientID,
			SpotifyAppClientSecret: c.spotifyAppClientSecret,
			PublicAPIEndpoint:      c.publicAPIEndpoint,
			ServerListenPort:       c.serverListenPort,
		}
		if c.spotifyAppClientID == "" {
			cfg.SpotifyAppClientID = os.Getenv("SPOTIFY_TOOLS_CLIENT_ID")
		}
		if c.spotifyAppClientSecret == "" {
			cfg.SpotifyAppClientSecret = os.Getenv("SPOTIFY_TOOLS_CLIENT_SECRET")
		}
		if cfg.SpotifyAppClientID == "" || cfg.SpotifyAppClientSecret == "" {
			return fmt.Errorf("missing Spotify App Client ID and/or Client Secret")
		}
		app, err := app.New(ctx, &cfg)
		if err != nil {
			return err
		}
		return handler(ctx, app, args)
	}
}
