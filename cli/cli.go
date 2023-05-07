package cli

import (
	"context"

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
	// configJSONFilePath string
	// configJSON         string
}

func New(ctx context.Context) *CLI {
	c := &CLI{}

	c.cmdRoot = &cobra.Command{
		Use:   "spotify-tools",
		Short: "A Spotify automation tool",
	}

	cmdRun := &cobra.Command{
		Use:   "run",
		Short: "run a command",
	}
	c.cmdRoot.AddCommand(cmdRun)

	c.cmdRoot.AddCommand(&cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("spotify-tools %s\n", version)
		},
	})

	cmdRun.PersistentFlags().StringVarP(&c.spotifyAppClientID, "spotify-app-client-id", "i", "", "Spotify App Client ID")
	err := cmdRun.MarkPersistentFlagRequired("spotify-app-client-id")
	if err != nil {
		panic(err)
	}
	cmdRun.PersistentFlags().StringVarP(&c.spotifyAppClientSecret, "spotify-app-client-secret", "s", "", "Spotify App Client Secret")
	err = cmdRun.MarkPersistentFlagRequired("spotify-app-client-secret")
	if err != nil {
		panic(err)
	}
	cmdRun.PersistentFlags().BoolVar(&c.devMode, "dev", false, "Dev mode")
	cmdRun.PersistentFlags().StringVar(&c.publicAPIEndpoint, "public-api-endpoint", "http://localhost:8080", "Public API endpoint")
	cmdRun.PersistentFlags().Uint16Var(&c.serverListenPort, "server-listen-port", 8080, "Server listen port")

	cmdRun.AddCommand(&cobra.Command{
		Use:   "auth-test",
		Short: "Test Spotify user authentication and token refresh",
		RunE:  c.runApp(ctx, c.authTestHandler),
	})

	cmdRun.AddCommand(&cobra.Command{
		Use:   "reset",
		Short: "Reset Spotify user authentication",
		Long:  "Reset Spotify user authentication by deleting the config cache file. Use this command if you want to change the Spotify user or if you have a more general issue with authentication.",
		RunE:  c.runApp(ctx, c.resetHandler),
	})

	cmdRun.AddCommand(&cobra.Command{
		Use:   "get-playlist",
		Short: "Get basic info about a Spotify playlist",
		Args:  cobra.ExactArgs(1),
		RunE:  c.runApp(ctx, c.getPlaylistHandler),
	})

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
		app, err := app.New(ctx, &cfg)
		if err != nil {
			return err
		}
		return handler(ctx, app, args)
	}
}
