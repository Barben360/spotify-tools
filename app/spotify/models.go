package spotify

import (
	"regexp"
	"time"
)

type PlaylistFilterConfig []*PlaylistFilterConfigElement

type PlaylistFilterConfigElement struct {
	TargetPlaylistID string `json:"target_playlist_id" validate:"required"`
	Sources          []*PlaylistFilterSource
}

type PlaylistFilterSource struct {
	PlaylistID                       string           `json:"playlist_id" validate:"required_without=ShowID"`
	ShowID                           string           `json:"show_id" validate:"required_without=PlaylistID"`
	Filters                          *PlaylistFilters `json:"filters" validate:"required"`
	AddLatestUpdateDateToDescription bool             `json:"add_latest_update_date_to_description" validate:"-"`
	DateLocation                     string           `json:"date_location" validate:"-"`
}

type PlaylistFilters struct {
	TrackNameRegex         string         `json:"name_regex" validate:"omitempty"`
	TrackNameRegexResolved *regexp.Regexp `json:"-" validate:"-"`
	DurationFrom           time.Duration  `json:"duration_from" validate:"omitempty"`
	DurationTo             time.Duration  `json:"duration_to" validate:"omitempty"`
}

type Playlist struct {
	ID          string
	Name        string
	Description string
	IsPublic    bool
	Items       []*Item
}

type ItemType string

const (
	ItemTypeTrack   ItemType = "track"
	ItemTypeEpisode ItemType = "episode"
)

type Item struct {
	ID       string
	Name     string
	Type     ItemType
	Duration time.Duration
	AddedAt  time.Time
}
