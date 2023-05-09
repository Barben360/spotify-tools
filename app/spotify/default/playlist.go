package spotifydefault

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/Barben360/spotify-tools/app/services/logger"
	"github.com/Barben360/spotify-tools/app/spotify"
	"github.com/Barben360/spotify-tools/app/spotify/default/spotifyclient"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
	for _, item := range tracks.GetItems() {
		track, ok := item.GetTrackOk()
		if !ok {
			logger.FromContext(ctx).Warn("track is undefined, this element may not be available any more, skipping.")
			continue
		}
		if track.TrackObject != nil {
			itemsRet = append(itemsRet, &spotify.Item{
				ID:          track.TrackObject.GetId(),
				URI:         track.TrackObject.GetUri(),
				Name:        track.TrackObject.GetName(),
				Type:        spotify.ItemTypeTrack,
				Duration:    time.Duration(track.TrackObject.GetDurationMs()) * time.Millisecond,
				AddedAt:     item.GetAddedAt(),
				ReleaseDate: track.TrackObject.Album.GetReleaseDate(),
			})
		} else if track.EpisodeObject != nil {
			itemsRet = append(itemsRet, &spotify.Item{
				ID:          track.EpisodeObject.GetId(),
				URI:         track.EpisodeObject.GetUri(),
				Name:        track.EpisodeObject.GetName(),
				Type:        spotify.ItemTypeEpisode,
				Duration:    time.Duration(track.EpisodeObject.GetDurationMs()) * time.Millisecond,
				AddedAt:     item.GetAddedAt(),
				ReleaseDate: track.EpisodeObject.GetReleaseDate(),
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
			track, ok := item.GetTrackOk()
			if !ok {
				logger.FromContext(ctx).Warn("track is undefined, this element may not be available any more, skipping.")
				continue
			}
			if track.TrackObject != nil {
				itemsRet = append(itemsRet, &spotify.Item{
					ID:       track.TrackObject.GetId(),
					URI:      track.TrackObject.GetUri(),
					Name:     track.TrackObject.GetName(),
					Type:     spotify.ItemTypeTrack,
					Duration: time.Duration(track.TrackObject.GetDurationMs()) * time.Millisecond,
					AddedAt:  item.GetAddedAt(),
				})
			} else if track.EpisodeObject != nil {
				itemsRet = append(itemsRet, &spotify.Item{
					ID:       track.EpisodeObject.GetId(),
					URI:      track.EpisodeObject.GetUri(),
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
	for _, item := range itemsRet {
		item.TryResolveReleaseDate(ctx)
	}
	return &spotify.Playlist{
		ID:          resp.GetId(),
		Name:        resp.GetName(),
		Description: resp.GetDescription(),
		IsPublic:    resp.GetPublic(),
		Items:       itemsRet,
	}, nil
}

var descriptionLatestUpdateDateRegexp = regexp.MustCompile(`-- Latest update: (?:.*)$`)

func (s *Spotify) UpdatePlaylistFilter(ctx context.Context, filterConfig *spotify.PlaylistFilterConfig) error {
	// Validate filter config
	validate := validator.New()
	err := validate.Struct(filterConfig)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			errs := make([]error, len(validationErrors))
			for i, validationError := range validationErrors {
				errs[i] = validationError
			}
			return errors.Join(errs...)
		}
		return err
	}
	var location *time.Location
	if filterConfig.LatestUpdateDateLocation != "" {
		location, err = time.LoadLocation(filterConfig.LatestUpdateDateLocation)
		if err != nil {
			return fmt.Errorf("invalid location %q: %w", filterConfig.LatestUpdateDateLocation, err)
		}
	} else {
		location = time.UTC
	}

	logger.FromContext(ctx).Info("applying filter", zap.String("target_playlist_id", filterConfig.TargetPlaylistID), zap.String("filter_description", filterConfig.Description))

	// Resolve regexp from filters
	regExpMap := make(map[string]*regexp.Regexp)
	for _, src := range filterConfig.Sources {
		if src.Filters.ItemNameRegexp != "" {
			regExp, err := regexp.Compile(src.Filters.ItemNameRegexp)
			if err != nil {
				return fmt.Errorf("invalid regexp %q: %w", src.Filters.ItemNameRegexp, err)
			}
			regExpMap[src.Filters.ItemNameRegexp] = regExp
		}
	}

	// Get target playlist
	targetPlaylist, err := s.GetPlaylist(ctx, filterConfig.TargetPlaylistID)
	if err != nil {
		return err
	}

	// Making target playlist items map
	targetPlaylistItemsMap := make(map[string]*spotify.Item)
	for _, item := range targetPlaylist.Items {
		targetPlaylistItemsMap[item.ID] = item
	}

	itemsToAdd := make([]*spotify.Item, 0)
	for _, src := range filterConfig.Sources {
		// Copying context in loop scope for contextual logging
		ctx := ctx
		// Get source playlist or show with all episodes
		var items []*spotify.Item
		if src.PlaylistID != "" {
			playlist, err := s.GetPlaylist(ctx, src.PlaylistID)
			if err != nil {
				return err
			}
			l := &logger.Logger{Logger: logger.FromContext(ctx).With(zap.String("source", playlist.Name), zap.String("source_type", "playlist"))}
			ctx = l.ToContext(ctx)
			items = playlist.Items
		} else if src.ShowID != "" {
			show, err := s.GetShow(ctx, src.ShowID)
			if err != nil {
				return err
			}
			l := &logger.Logger{Logger: logger.FromContext(ctx).With(zap.String("source", show.Name), zap.String("source_type", "show"))}
			ctx = l.ToContext(ctx)
			items = show.Episodes
		} else {
			return fmt.Errorf("either playlist or show ID must be set")
		}

		// Filter episodes and skip episodes already in target playlist
		for _, item := range items {
			l := &logger.Logger{Logger: logger.FromContext(ctx).With(zap.String("item_id", item.ID), zap.String("item_name", item.Name), zap.String("item_type", string(item.Type)))}
			ctx := l.ToContext(ctx)
			// Is it already in target playlist?
			if _, ok := targetPlaylistItemsMap[item.ID]; ok {
				logger.FromContext(ctx).Debug("skipping item already in target playlist")
				continue
			}
			// Is it matching filters?
			if src.Filters.ItemNameRegexp != "" {
				regExp := regExpMap[src.Filters.ItemNameRegexp]
				if !regExp.MatchString(item.Name) {
					logger.FromContext(ctx).Debug("skipping item not matching regexp")
					continue
				}
			}
			// Is it matching duration?
			if src.Filters.MinDuration != 0 && item.Duration < time.Duration(src.Filters.MinDuration)*time.Second {
				logger.FromContext(ctx).Debug("skipping item with too short duration")
				continue
			}
			if src.Filters.MaxDuration != 0 && item.Duration > time.Duration(src.Filters.MaxDuration)*time.Second {
				logger.FromContext(ctx).Debug("skipping item with too long duration")
				continue
			}
			// Filter passed, adding item to the list
			logger.FromContext(ctx).Info("adding item to the list")
			itemsToAdd = append(itemsToAdd, item)
		}
	}

	// Note: even if there is no item to add, we continue to reorder items, they could have been manually reordered by error

	// Ordering items to add by release date
	if filterConfig.OrderBy == spotify.PlaylistOrderReleaseDate {
		sort.Slice(itemsToAdd, func(i, j int) bool {
			return itemsToAdd[i].ReleaseDateTime.Before(itemsToAdd[j].ReleaseDateTime)
		})
	} else {
		sort.Slice(itemsToAdd, func(i, j int) bool {
			return itemsToAdd[i].GetFallbackDate().Before(itemsToAdd[j].GetFallbackDate())
		})
	}

	// Add items at the end of the target playlist
	err = s.playlistAppendItems(ctx, filterConfig.TargetPlaylistID, itemsToAdd)
	if err != nil {
		return err
	}

	// Reorder items if needed
	allItems := append(targetPlaylist.Items, itemsToAdd...)
	allItemsOriginal := make([]*spotify.Item, len(allItems))
	copy(allItemsOriginal, allItems)
	if filterConfig.OrderBy == spotify.PlaylistOrderReleaseDate {
		sort.Slice(allItems, func(i, j int) bool {
			return allItems[i].ReleaseDateTime.Before(allItems[j].ReleaseDateTime)
		})
	} else {
		sort.Slice(allItems, func(i, j int) bool {
			return allItems[i].GetFallbackDate().Before(allItems[j].GetFallbackDate())
		})
	}
	// If order changed, reorder playlist
	err = s.playlistReorder(ctx, filterConfig.TargetPlaylistID, allItems)
	if err != nil {
		return err
	}

	// If requested, update target playlist description
	if filterConfig.AddLatestUpdateDateToDescription {
		description := targetPlaylist.Description
		descriptionTimestampString := "-- Latest update: " + time.Now().In(location).Format("2006-01-02 15:04")
		// Is there already a timestamp in the description?
		if descriptionLatestUpdateDateRegexp.MatchString(description) {
			// If yes, we replace it
			description = descriptionLatestUpdateDateRegexp.ReplaceAllString(description, descriptionTimestampString)
		} else {
			// If no, we append it
			description = description + " " + descriptionTimestampString
		}

		// Update description
		err = s.playlistUpdateDescription(ctx, filterConfig.TargetPlaylistID, description)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Spotify) playlistAppendItems(ctx context.Context, playlistID string, items []*spotify.Item) error {
	const batchSize = 100
	urisBatches := make([][]string, 0)
	currURIs := make([]string, 0, batchSize)
	for _, item := range items {
		currURIs = append(currURIs, item.URI)
		if len(currURIs) == batchSize {
			urisBatches = append(urisBatches, currURIs)
			currURIs = make([]string, 0, batchSize)
		}
	}
	if len(currURIs) > 0 {
		urisBatches = append(urisBatches, currURIs)
	}
	for _, uris := range urisBatches {
		if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
			var httpResp *http.Response
			_, httpResp, err := s.httpOpenAPIClient.PlaylistsApi.AddTracksToPlaylist(ctx, playlistID).RequestBody(map[string]interface{}{"uris": uris}).Execute()
			return httpResp, err
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Spotify) playlistReorder(ctx context.Context, playlistID string, items []*spotify.Item) error {
	const batchSize = 100
	urisBatches := make([][]string, 0)
	currURIs := make([]string, 0, batchSize)
	for _, item := range items {
		currURIs = append(currURIs, item.URI)
		if len(currURIs) == batchSize {
			urisBatches = append(urisBatches, currURIs)
			currURIs = make([]string, 0, batchSize)
		}
	}
	if len(currURIs) > 0 {
		urisBatches = append(urisBatches, currURIs)
	}
	for _, uris := range urisBatches {
		if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
			var httpResp *http.Response
			_, httpResp, err := s.httpOpenAPIClient.PlaylistsApi.ReorderOrReplacePlaylistsTracks(ctx, playlistID).RequestBody(map[string]interface{}{"uris": uris}).Execute()
			return httpResp, err
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *Spotify) playlistUpdateDescription(ctx context.Context, playlistID string, description string) error {
	description = strings.TrimSpace(description)
	if _, err := s.authExec(ctx, func(ctx context.Context) (*http.Response, error) {
		httpResp, err := s.httpOpenAPIClient.PlaylistsApi.ChangePlaylistDetails(ctx, playlistID).RequestBody(map[string]interface{}{
			"description": description,
		}).Execute()
		return httpResp, err
	}); err != nil {
		return err
	}
	return nil
}
