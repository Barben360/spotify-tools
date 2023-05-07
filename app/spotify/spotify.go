package spotify

import "context"

type Spotifier interface {
	SpotifiyAuther
	SpotifyPlaylister
	SpotifyShower
}

type SpotifiyAuther interface {
	// GetUserToken returns the latest user token, which may be expired.
	// If the token has never been fetched, this will run authorization process.
	GetUserToken(ctx context.Context) (string, error)
	// GetNewUserToken returns a new user token, which is never expired, using the refresh token if present
	// or running the authorization process otherwise.
	GetNewUserToken(ctx context.Context) (string, error)
	// ResetUserTokens resets the user tokens, including the refresh token.
	// This should be used only for testing
	ResetUserTokens(ctx context.Context)
}

type SpotifyPlaylister interface {
	GetPlaylist(ctx context.Context, playlistID string) (*Playlist, error)
	UpdatePlaylistFilter(ctx context.Context, filterConfig *PlaylistFilterConfig) error
}

type SpotifyShower interface {
	GetShow(ctx context.Context, showID string) (*Show, error)
}
