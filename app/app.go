package app

import (
	"context"
	"errors"

	authtestdefault "github.com/Barben360/spotify-playlist-filters/app/authtest/default"
	"github.com/Barben360/spotify-playlist-filters/app/services/logger"
	spotifydefault "github.com/Barben360/spotify-playlist-filters/app/spotify/default"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type App struct {
	log      *logger.Logger
	features *Features
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
	features := &Features{
		AuthTest: authtestdefault.New(ctx, spotify),
	}

	return &App{
		log:      log,
		features: features,
	}, nil
}

func (a *App) RunAuthTest(ctx context.Context) error {
	ctx = a.log.ToContext(ctx)
	return a.features.AuthTest.TestUserAuthAndTokenRefresh(ctx)
}
