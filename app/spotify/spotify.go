package spotify

import "context"

type Spotifier interface {
	SpotifiyAuther
	SpotifyPlaylister
	SpotifyShower
}

type SpotifiyAuther interface {
	// Login forces a new OAuth authorization flow, clearing any existing tokens first.
	Login(ctx context.Context) error
	// Logout removes all tokens from memory and the cache file. It never fails,
	// even if the user is not logged in.
	Logout(ctx context.Context)
	// AuthStatus checks if the user is logged in and whether the tokens are valid.
	// Returns (loggedIn bool, err error):
	//   - loggedIn=false, err=nil: not logged in (no refresh token present)
	//   - loggedIn=true, err=nil: logged in and refresh token is valid
	//   - loggedIn=true, err!=nil: tokens present but refresh failed
	AuthStatus(ctx context.Context) (bool, error)
}

type SpotifyPlaylister interface {
	GetPlaylist(ctx context.Context, playlistID string) (*Playlist, error)
	UpdatePlaylistFilter(ctx context.Context, filterConfig *PlaylistFilterConfig) error
}

type SpotifyShower interface {
	GetShow(ctx context.Context, showID string) (*Show, error)
}
