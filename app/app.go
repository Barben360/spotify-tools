package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	authtestdefault "github.com/Barben360/spotify-tools/app/authtest/default"
	"github.com/Barben360/spotify-tools/app/services/logger"
	"github.com/Barben360/spotify-tools/app/spotify"
	spotifydefault "github.com/Barben360/spotify-tools/app/spotify/default"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type App struct {
	log      *logger.Logger
	features *features
}

func New(ctx context.Context, cfg *AppConfig) (*App, error) {
	// Validate config
	validate := validator.New()
	err := validate.Struct(cfg)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			errs := make([]error, len(validationErrors))
			for i, validationError := range validationErrors {
				errs[i] = validationError
			}
			return nil, errors.Join(errs...)
		}
		return nil, err
	}

	// Initialize logger
	var log *logger.Logger
	if cfg.DevMode {
		logZap, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		log = &logger.Logger{Logger: logZap}
	} else {
		logZap, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		log = &logger.Logger{Logger: logZap}
	}

	ctx = log.ToContext(ctx)

	// Initialize features
	spotify := spotifydefault.New(
		ctx,
		cfg.SpotifyAppClientID,
		cfg.SpotifyAppClientSecret,
		cfg.PublicAPIEndpoint,
		cfg.ServerListenPort,
	)
	features := &features{
		spotify:  spotify,
		authTest: authtestdefault.New(ctx, spotify),
	}

	return &App{
		log:      log,
		features: features,
	}, nil
}

func (a *App) RunAuthTest(ctx context.Context) error {
	ctx = a.log.ToContext(ctx)
	return a.features.authTest.TestUserAuthAndTokenRefresh(ctx)
}

func (a *App) RunReset(ctx context.Context) error {
	ctx = a.log.ToContext(ctx)
	a.features.spotify.ResetUserTokens(ctx)
	return nil
}

func (a *App) GetPlaylist(ctx context.Context, playlistID string) error {
	ctx = a.log.ToContext(ctx)
	resp, err := a.features.spotify.GetPlaylist(ctx, playlistID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	fmt.Println(string(respJSON))
	return nil
}

func (a *App) GetShow(ctx context.Context, showID string) error {
	ctx = a.log.ToContext(ctx)
	resp, err := a.features.spotify.GetShow(ctx, showID)
	if err != nil {
		return err
	}
	respJSON, _ := json.Marshal(resp)
	fmt.Println(string(respJSON))
	return nil
}

func (a *App) PlaylistFilter(ctx context.Context, config *spotify.PlaylistFilterConfig) error {
	ctx = a.log.ToContext(ctx)
	err := a.features.spotify.UpdatePlaylistFilter(ctx, config)
	if err != nil {
		return err
	}
	return nil
}
