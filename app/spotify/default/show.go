package spotifydefault

import (
	"context"
	"net/http"
	"time"

	"github.com/Barben360/spotify-tools/app/spotify"
	"github.com/Barben360/spotify-tools/app/spotify/default/spotifyclient"
)

func (s *Spotify) GetShow(ctx context.Context, showID string) (*spotify.Show, error) {
	var show *spotifyclient.ShowObject
	if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
		var err error
		var httpResp *http.Response
		show, httpResp, err = s.httpOpenAPIClient.ShowsAPI.GetAShow(ctx, showID).Execute()
		return httpResp, err
	}); err != nil {
		return nil, err
	}
	var episodes *spotifyclient.PagingSimplifiedEpisodeObject
	if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
		var err error
		var httpResp *http.Response
		episodes, httpResp, err = s.httpOpenAPIClient.ShowsAPI.GetAShowsEpisodes(ctx, showID).Execute()
		return httpResp, err
	}); err != nil {
		return nil, err
	}
	itemsRet := make([]*spotify.Item, 0, episodes.GetTotal())
	for _, item := range episodes.GetItems() {
		itemsRet = append(itemsRet, &spotify.Item{
			ID:          item.GetId(),
			URI:         item.GetUri(),
			Name:        item.GetName(),
			Type:        spotify.ItemTypeEpisode,
			Duration:    time.Duration(item.GetDurationMs()) * time.Millisecond,
			ReleaseDate: item.GetReleaseDate(),
		})
	}
	next := episodes.GetNext()
	offset := episodes.GetOffset()
	limit := episodes.GetLimit()
	for next != "" {
		var resp *spotifyclient.PagingSimplifiedEpisodeObject
		if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
			var err error
			var httpResp *http.Response
			resp, httpResp, err = s.httpOpenAPIClient.ShowsAPI.GetAShowsEpisodes(ctx, showID).Offset(offset + limit).Limit(limit).Execute()
			return httpResp, err
		}); err != nil {
			return nil, err
		}
		next = resp.GetNext()
		offset = resp.GetOffset()
		// Using the default limit
		for _, item := range resp.GetItems() {
			itemsRet = append(itemsRet, &spotify.Item{
				ID:          item.GetId(),
				URI:         item.GetUri(),
				Name:        item.GetName(),
				Type:        spotify.ItemTypeEpisode,
				Duration:    time.Duration(item.GetDurationMs()) * time.Millisecond,
				ReleaseDate: item.GetReleaseDate(),
			})
		}
	}
	for _, item := range itemsRet {
		item.TryResolveReleaseDate(ctx)
	}
	return &spotify.Show{
		ID:          show.GetId(),
		Name:        show.GetName(),
		Description: show.GetDescription(),
		Episodes:    itemsRet,
	}, nil
}
