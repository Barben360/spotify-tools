package cli

import (
	"context"

	"github.com/Barben360/spotify-playlist-filters/app"
)

func authTestHandler(ctx context.Context, appInstance *app.App) error {
	return appInstance.RunAuthTest(ctx)
}
