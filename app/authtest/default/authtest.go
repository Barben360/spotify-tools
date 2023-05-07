package authtestdefault

import (
	"context"

	"github.com/Barben360/spotify-tools/app/authtest"
	"github.com/Barben360/spotify-tools/app/services/logger"
	"github.com/Barben360/spotify-tools/app/spotify"
)

type AuthTest struct {
	spotify spotify.Spotifier
}

func New(ctx context.Context, spotify spotify.Spotifier) authtest.AuthTester {
	return &AuthTest{
		spotify: spotify,
	}
}

func (a *AuthTest) TestUserAuthAndTokenRefresh(ctx context.Context) error {
	logger.FromContext(ctx).Info("Testing user auth and token refresh")
	a.spotify.ResetUserTokens(ctx)
	_, err := a.spotify.GetUserToken(ctx)
	if err != nil {
		return err
	}
	_, err = a.spotify.GetNewUserToken(ctx)
	if err != nil {
		return err
	}
	logger.FromContext(ctx).Info("User auth and token refresh test successful")
	return nil
}
