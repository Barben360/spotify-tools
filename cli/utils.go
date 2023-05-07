package cli

import (
	"context"
	"encoding/json"
	"os"
)

func (c *CLI) getConfig(ctx context.Context, v any) error {
	if c.configJSONFilePath != "" {
		f, err := os.Open(c.configJSONFilePath)
		if err != nil {
			return err
		}
		defer f.Close()
		return json.NewDecoder(f).Decode(v)
	}
	return json.Unmarshal([]byte(c.configJSON), v)
}
