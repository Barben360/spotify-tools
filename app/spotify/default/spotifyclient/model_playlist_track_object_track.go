/*
Spotify Web API with fixes and improvements from sonallux

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 2023.2.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spotifyclient

import (
	"encoding/json"
	"fmt"
)

// PlaylistTrackObjectTrack - Information about the track or episode.
type PlaylistTrackObjectTrack struct {
	EpisodeObject *EpisodeObject
	TrackObject *TrackObject
}

// EpisodeObjectAsPlaylistTrackObjectTrack is a convenience function that returns EpisodeObject wrapped in PlaylistTrackObjectTrack
func EpisodeObjectAsPlaylistTrackObjectTrack(v *EpisodeObject) PlaylistTrackObjectTrack {
	return PlaylistTrackObjectTrack{
		EpisodeObject: v,
	}
}

// TrackObjectAsPlaylistTrackObjectTrack is a convenience function that returns TrackObject wrapped in PlaylistTrackObjectTrack
func TrackObjectAsPlaylistTrackObjectTrack(v *TrackObject) PlaylistTrackObjectTrack {
	return PlaylistTrackObjectTrack{
		TrackObject: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PlaylistTrackObjectTrack) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal data into EpisodeObject
	err = json.Unmarshal(data, &dst.EpisodeObject)
	if err == nil && dst.EpisodeObject.Type == "episode" { 
			return nil
	} else {
		dst.EpisodeObject = nil
		// try to unmarshal data into TrackObject
		err = json.Unmarshal(data, &dst.TrackObject)
		if err == nil {
			jsonTrackObject, _ := json.Marshal(dst.TrackObject)
			if string(jsonTrackObject) == "{}" { // empty struct
				return fmt.Errorf("data failed to match schemas in oneOf(PlaylistTrackObjectTrack)")
			} 
		} else {
			return fmt.Errorf("data failed to match schemas in oneOf(PlaylistTrackObjectTrack)")
		}
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PlaylistTrackObjectTrack) MarshalJSON() ([]byte, error) {
	if src.EpisodeObject != nil {
		return json.Marshal(&src.EpisodeObject)
	}

	if src.TrackObject != nil {
		return json.Marshal(&src.TrackObject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PlaylistTrackObjectTrack) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EpisodeObject != nil {
		return obj.EpisodeObject
	}

	if obj.TrackObject != nil {
		return obj.TrackObject
	}

	// all schemas are nil
	return nil
}

type NullablePlaylistTrackObjectTrack struct {
	value *PlaylistTrackObjectTrack
	isSet bool
}

func (v NullablePlaylistTrackObjectTrack) Get() *PlaylistTrackObjectTrack {
	return v.value
}

func (v *NullablePlaylistTrackObjectTrack) Set(val *PlaylistTrackObjectTrack) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistTrackObjectTrack) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistTrackObjectTrack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistTrackObjectTrack(val *PlaylistTrackObjectTrack) *NullablePlaylistTrackObjectTrack {
	return &NullablePlaylistTrackObjectTrack{value: val, isSet: true}
}

func (v NullablePlaylistTrackObjectTrack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistTrackObjectTrack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


