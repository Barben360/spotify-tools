package cli

import (
	"context"

	"github.com/Barben360/spotify-playlist-filters/app"
	"github.com/spf13/cobra"
)

type CLI struct {
	cmdRoot                *cobra.Command
	spotifyAppClientID     string
	spotifyAppClientSecret string
	devMode                bool
	publicAPIEndpoint      string
	serverListenPort       uint16
	// configJSONFilePath string
	// configJSON         string
}

func New(ctx context.Context) *CLI {
	c := &CLI{}

	c.cmdRoot = &cobra.Command{
		Use:   "spotify-playlist-filters",
		Short: "spotify-playlist-filters is a Spotify playlist automation tool",
	}

	c.cmdRoot.PersistentFlags().StringVarP(&c.spotifyAppClientID, "spotify-app-client-id", "i", "", "Spotify App Client ID")
	err := c.cmdRoot.MarkPersistentFlagRequired("spotify-app-client-id")
	if err != nil {
		panic(err)
	}
	c.cmdRoot.PersistentFlags().StringVarP(&c.spotifyAppClientSecret, "spotify-app-client-secret", "s", "", "Spotify App Client Secret")
	err = c.cmdRoot.MarkPersistentFlagRequired("spotify-app-client-secret")
	if err != nil {
		panic(err)
	}
	c.cmdRoot.PersistentFlags().BoolVar(&c.devMode, "dev", false, "Dev mode")
	c.cmdRoot.PersistentFlags().StringVar(&c.publicAPIEndpoint, "public-api-endpoint", "http://localhost:8080", "Public API endpoint")
	c.cmdRoot.PersistentFlags().Uint16Var(&c.serverListenPort, "server-listen-port", 8080, "Server listen port")

	c.cmdRoot.AddCommand(&cobra.Command{
		Use:   "auth-test",
		Short: "Test Spotify user authentication and token refresh",
		RunE:  c.runApp(ctx, authTestHandler),
	})
	return c
}

func (c *CLI) Run(ctx context.Context) error {
	return c.cmdRoot.ExecuteContext(ctx)
}

func (c *CLI) runApp(ctx context.Context, handler func(ctx context.Context, appInstance *app.App) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cfg := app.AppConfig{
			DevMode:                c.devMode,
			SpotifyAppClientID:     c.spotifyAppClientID,
			SpotifyAppClientSecret: c.spotifyAppClientSecret,
			PublicAPIEndpoint:      c.publicAPIEndpoint,
			ServerListenPort:       c.serverListenPort,
		}
		app, err := app.New(ctx, &cfg)
		if err != nil {
			return err
		}
		return handler(ctx, app)
	}
}
