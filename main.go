package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Barben360/spotify-tools/app"
	"github.com/Barben360/spotify-tools/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := cli.New(ctx)
	errChan := make(chan error, 1)
	go func() {
		errChan <- c.Run(ctx)
	}()
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, syscall.SIGTERM)
	select {
	case err := <-errChan:
		if err != nil {
			var exitCodeErr *app.ExitCodeError
			if errors.As(err, &exitCodeErr) {
				os.Exit(exitCodeErr.Code)
			}
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	case <-sigTerm:
		fmt.Print("Received SIGTERM, shutting down")
		cancel()
		<-errChan
	}
}
