package spotifydefault

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Barben360/spotify-tools/app/spotify"
	"github.com/Barben360/spotify-tools/app/spotify/default/spotifyclient"
)

func (s *Spotify) GetPlaylist(ctx context.Context, playlistID string) (*spotify.Playlist, error) {
	var resp *spotifyclient.PlaylistObject
	if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
		var err error
		var httpResp *http.Response
		resp, httpResp, err = s.httpOpenAPIClient.PlaylistsApi.GetPlaylist(ctx, playlistID).Execute()
		return httpResp, err
	}); err != nil {
		return nil, err
	}
	tracks := resp.GetTracks()
	itemsRet := make([]*spotify.Item, 0, tracks.GetTotal())
	items := tracks.GetItems()
	for _, item := range items {
		track := item.GetTrack()
		if track.TrackObject != nil {
			itemsRet = append(itemsRet, &spotify.Item{
				ID:       track.TrackObject.GetId(),
				Name:     track.TrackObject.GetName(),
				Type:     spotify.ItemTypeTrack,
				Duration: time.Duration(track.TrackObject.GetDurationMs()) * time.Millisecond,
				AddedAt:  item.GetAddedAt(),
			})
		} else if track.EpisodeObject != nil {
			itemsRet = append(itemsRet, &spotify.Item{
				ID:       track.EpisodeObject.GetId(),
				Name:     track.EpisodeObject.GetName(),
				Type:     spotify.ItemTypeEpisode,
				Duration: time.Duration(track.EpisodeObject.GetDurationMs()) * time.Millisecond,
				AddedAt:  item.GetAddedAt(),
			})
		} else {
			return nil, fmt.Errorf("unknown track type")
		}
	}
	next := tracks.GetNext()
	offset := tracks.GetOffset()
	limit := tracks.GetLimit()
	for next != "" {
		var resp *spotifyclient.PagingPlaylistTrackObject
		if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
			var err error
			var httpResp *http.Response
			resp, httpResp, err = s.httpOpenAPIClient.PlaylistsApi.GetPlaylistsTracks(ctx, playlistID).Offset(offset + limit).Limit(limit).Execute()
			return httpResp, err
		}); err != nil {
			return nil, err
		}
		next = resp.GetNext()
		offset = resp.GetOffset()
		// Using the default limit
		for _, item := range resp.GetItems() {
			track := item.GetTrack()
			if track.TrackObject != nil {
				itemsRet = append(itemsRet, &spotify.Item{
					ID:       track.TrackObject.GetId(),
					Name:     track.TrackObject.GetName(),
					Type:     spotify.ItemTypeTrack,
					Duration: time.Duration(track.TrackObject.GetDurationMs()) * time.Millisecond,
					AddedAt:  item.GetAddedAt(),
				})
			} else if track.EpisodeObject != nil {
				itemsRet = append(itemsRet, &spotify.Item{
					ID:       track.EpisodeObject.GetId(),
					Name:     track.EpisodeObject.GetName(),
					Type:     spotify.ItemTypeEpisode,
					Duration: time.Duration(track.EpisodeObject.GetDurationMs()) * time.Millisecond,
					AddedAt:  item.GetAddedAt(),
				})
			} else {
				return nil, fmt.Errorf("unknown track type")
			}
		}
	}

	return &spotify.Playlist{
		ID:          resp.GetId(),
		Name:        resp.GetName(),
		Description: resp.GetDescription(),
		IsPublic:    resp.GetPublic(),
		Items:       itemsRet,
	}, nil
}
