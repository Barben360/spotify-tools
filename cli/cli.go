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
	configCacheFilePath    string
}

func New(ctx context.Context) *CLI {
	c := &CLI{}

	c.cmdRoot = &cobra.Command{
		Use:           "spotify-tools",
		Version:       version,
		Short:         "A Spotify automation tool",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	c.cmdRoot.PersistentFlags().StringVarP(&c.spotifyAppClientID, "spotify-app-client-id", "i", "", "Spotify App Client ID")
	c.cmdRoot.PersistentFlags().StringVarP(&c.spotifyAppClientSecret, "spotify-app-client-secret", "s", "", "Spotify App Client Secret")

	c.cmdRoot.PersistentFlags().BoolVar(&c.devMode, "dev", false, "Dev mode")
	c.cmdRoot.PersistentFlags().StringVar(&c.publicAPIEndpoint, "public-api-endpoint", "http://127.0.0.1:8080", "Public API endpoint")
	c.cmdRoot.PersistentFlags().Uint16Var(&c.serverListenPort, "server-listen-port", 8080, "Server listen port")
	c.cmdRoot.PersistentFlags().StringVar(&c.configCacheFilePath, "config-cache-file", "", "Path to auth token cache file (env: SPOTIFY_TOOLS_CACHE_FILE)")

	// auth command group
	cmdAuth := &cobra.Command{
		Use:   "auth",
		Short: "Manage Spotify authentication",
	}

	cmdAuth.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Check Spotify authentication status",
		RunE:  c.runApp(ctx, c.authStatusHandler),
	})

	cmdAuth.AddCommand(&cobra.Command{
		Use:   "login",
		Short: "Log in to Spotify via OAuth authorization flow",
		RunE:  c.runApp(ctx, c.authLoginHandler),
	})

	cmdAuth.AddCommand(&cobra.Command{
		Use:   "logout",
		Short: "Log out of Spotify by removing cached tokens",
		Long:  "Log out of Spotify by removing cached tokens from disk. Never fails, even if not currently logged in. Does not require Spotify app credentials.",
		RunE:  c.authLogoutHandler,
	})

	c.cmdRoot.AddCommand(cmdAuth)

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

func (c *CLI) effectiveCacheFilePath() string {
	if c.configCacheFilePath != "" {
		return c.configCacheFilePath
	}
	if v := os.Getenv("SPOTIFY_TOOLS_CACHE_FILE"); v != "" {
		return v
	}
	return app.DefaultConfigCacheFilePath
}

func (c *CLI) runApp(ctx context.Context, handler func(ctx context.Context, appInstance *app.App, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cfg := app.AppConfig{
			DevMode:                c.devMode,
			SpotifyAppClientID:     c.spotifyAppClientID,
			SpotifyAppClientSecret: c.spotifyAppClientSecret,
			PublicAPIEndpoint:      c.publicAPIEndpoint,
			ServerListenPort:       c.serverListenPort,
			ConfigCacheFilePath:    c.effectiveCacheFilePath(),
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
