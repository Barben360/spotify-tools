package cli

import (
	"context"
	"errors"

	"github.com/Barben360/spotify-tools/app"
	"github.com/Barben360/spotify-tools/app/spotify"
)

func (c *CLI) authTestHandler(ctx context.Context, appInstance *app.App, args []string) error {
	return appInstance.RunAuthTest(ctx)
}

func (c *CLI) resetHandler(ctx context.Context, appInstance *app.App, args []string) error {
	return appInstance.RunReset(ctx)
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
		}
	}()
	return <-errChan
}
