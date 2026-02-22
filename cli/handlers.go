package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Barben360/spotify-tools/app"
	"github.com/Barben360/spotify-tools/app/spotify"
	"github.com/spf13/cobra"
)

func (c *CLI) authStatusHandler(ctx context.Context, appInstance *app.App, args []string) error {
	return appInstance.AuthStatus(ctx)
}

func (c *CLI) authLoginHandler(ctx context.Context, appInstance *app.App, args []string) error {
	return appInstance.Login(ctx)
}

// authLogoutHandler removes cached tokens directly, without needing app initialization.
// It never fails, even if the user is not logged in or credentials are not provided.
func (c *CLI) authLogoutHandler(cmd *cobra.Command, args []string) error {
	app.LogoutFromDisk(c.configCacheFilePath)
	fmt.Println("Logged out successfully")
	return nil
}

func (c *CLI) getPlaylistHandler(ctx context.Context, appInstance *app.App, args []string) error {
	return appInstance.GetPlaylist(ctx, args[0])
}

func (c *CLI) getShowHandler(ctx context.Context, appInstance *app.App, args []string) error {
	return appInstance.GetShow(ctx, args[0])
}

func (c *CLI) filterPlaylistsHandler(ctx context.Context, appInstance *app.App, args []string) error {
	cfg := spotify.PlaylistFilterConfigSlice{}
	err := c.getConfig(ctx, &cfg)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	go func() {
		for {
			errs := []error{}
			for _, config := range cfg {
				err := appInstance.PlaylistFilter(ctx, config)
				if err != nil {
					errs = append(errs, err)
				}
			}
			if len(errs) > 0 {
				errChan <- errors.Join(errs...)
				return
			}
			if !c.daemonMode {
				errChan <- nil
				return
			}
			select {
			case <-ctx.Done():
				errChan <- nil
			case <-time.After(c.period):
				// continue
			}
		}
	}()
	return <-errChan
}
