package spotifydefault

import (
	"context"
	"net/http"
	"time"

	"github.com/Barben360/spotify-tools/app/spotify"
	"github.com/Barben360/spotify-tools/app/spotify/default/spotifyclient"
)

func (s *Spotify) GetShow(ctx context.Context, showID string) (*spotify.Show, error) {
	var resp *spotifyclient.ShowObject
	if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
		var err error
		var httpResp *http.Response
		resp, httpResp, err = s.httpOpenAPIClient.ShowsApi.GetAShow(ctx, showID).Execute()
		return httpResp, err
	}); err != nil {
		return nil, err
	}
	episodes := resp.GetEpisodes()
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
			resp, httpResp, err = s.httpOpenAPIClient.ShowsApi.GetAShowsEpisodes(ctx, showID).Offset(offset + limit).Limit(limit).Execute()
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
		ID:          resp.GetId(),
		Name:        resp.GetName(),
		Description: resp.GetDescription(),
		Episodes:    itemsRet,
	}, nil
}
