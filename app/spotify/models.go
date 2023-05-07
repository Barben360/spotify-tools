package spotify

import (
	"context"
	"time"
)

type PlaylistFilterConfigSlice []*PlaylistFilterConfig

type PlaylistFilterConfig struct {
	TargetPlaylistID string                  `json:"target_playlist_id" validate:"required"`
	Description      string                  `json:"description" validate:"omitempty"`
	Sources          []*PlaylistFilterSource `json:"sources" validate:"required,dive,required"`
	// OrderBy is the order of the items in the target playlist
	// For items having no added time (e.g. shows), the release date is used as fallback
	OrderBy PlaylistOrder `json:"order_by" validate:"omitempty,oneof=added_at release_date"`
}

type PlaylistOrder string

const (
	PlaylistOrderAddedAt     PlaylistOrder = "added_at"
	PlaylistOrderReleaseDate PlaylistOrder = "release_date"
)

type PlaylistFilterSource struct {
	PlaylistID                       string           `json:"playlist_id" validate:"required_without=ShowID"`
	ShowID                           string           `json:"show_id" validate:"required_without=PlaylistID"`
	Filters                          *PlaylistFilters `json:"filters" validate:"required"`
	AddLatestUpdateDateToDescription bool             `json:"add_latest_update_date_to_description" validate:"-"`
}

type PlaylistFilters struct {
	ItemNameRegexp string        `json:"item_name_regexp" validate:"omitempty"`
	MinDuration    time.Duration `json:"duration_from" validate:"omitempty"`
	MaxDuration    time.Duration `json:"duration_to" validate:"omitempty"`
}

type Playlist struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	IsPublic    bool    `json:"is_public,omitempty"`
	Items       []*Item `json:"items,omitempty"`
}

type ItemType string

const (
	ItemTypeTrack   ItemType = "track"
	ItemTypeEpisode ItemType = "episode"
)

type Item struct {
	ID       string        `json:"id,omitempty"`
	URI      string        `json:"uri,omitempty"`
	Name     string        `json:"name,omitempty"`
	Type     ItemType      `json:"type,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
	AddedAt  time.Time     `json:"added_at,omitempty"`
	// Only for episode type
	ReleaseDate string `json:"release_date,omitempty"`
	// For comparison
	ReleaseDateTime time.Time `json:"-"`
}

func (i *Item) TryResolveReleaseDate(ctx context.Context) {
	if i.ReleaseDate == "" {
		return
	}
	t, err := time.Parse("2006-01-02", i.ReleaseDate)
	if err == nil {
		i.ReleaseDateTime = t
		return
	}
	t, err = time.Parse("2006-01", i.ReleaseDate)
	if err == nil {
		i.ReleaseDateTime = t
		return
	}
	t, err = time.Parse("2006", i.ReleaseDate)
	if err == nil {
		i.ReleaseDateTime = t
	}
}

func (i *Item) GetFallbackDate() time.Time {
	if i.AddedAt.IsZero() {
		return i.ReleaseDateTime
	}
	return i.AddedAt
}

type Show struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Episodes    []*Item `json:"episodes,omitempty"`
}
