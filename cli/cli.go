package cli

import "context"

type CLI struct {
	spotifyAppClientID string
	publicAPIEndpoint  string
	configJSONFilePath string
	configJSON         string
}

func New(ctx context.Context) *CLI {
	return &CLI{}
}
