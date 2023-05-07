package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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
			os.Exit(1)
		}
	case <-sigTerm:
		fmt.Print("Received SIGTERM, shutting down")
		cancel()
		<-errChan
	}
}
