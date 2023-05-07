package cli

import (
	"context"

	"github.com/Barben360/spotify-tools/app"
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
