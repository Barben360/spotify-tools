package cli

import "context"

type CLI struct {
	configJSONFilePath string
	configJSON         string
}

func New(ctx context.Context) *CLI {
	return &CLI{}
}
