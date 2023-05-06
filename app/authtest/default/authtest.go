package authtestdefault

import (
	"context"

	"github.com/Barben360/spotify-playlist-filters/app/authtest"
	"github.com/Barben360/spotify-playlist-filters/app/spotify"
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
	a.spotify.ResetUserTokens(ctx)
	_, err := a.spotify.GetUserToken(ctx)
	if err != nil {
		return err
	}
	// TODO: test token refresh
	return nil
}
