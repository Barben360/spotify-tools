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

// CurrentlyPlayingObjectItem - The currently playing track or episode. Can be `null`.
type CurrentlyPlayingObjectItem struct {
	EpisodeObject *EpisodeObject
	TrackObject *TrackObject
}

// EpisodeObjectAsCurrentlyPlayingObjectItem is a convenience function that returns EpisodeObject wrapped in CurrentlyPlayingObjectItem
func EpisodeObjectAsCurrentlyPlayingObjectItem(v *EpisodeObject) CurrentlyPlayingObjectItem {
	return CurrentlyPlayingObjectItem{
		EpisodeObject: v,
	}
}

// TrackObjectAsCurrentlyPlayingObjectItem is a convenience function that returns TrackObject wrapped in CurrentlyPlayingObjectItem
func TrackObjectAsCurrentlyPlayingObjectItem(v *TrackObject) CurrentlyPlayingObjectItem {
	return CurrentlyPlayingObjectItem{
		TrackObject: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CurrentlyPlayingObjectItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EpisodeObject
	err = newStrictDecoder(data).Decode(&dst.EpisodeObject)
	if err == nil {
		jsonEpisodeObject, _ := json.Marshal(dst.EpisodeObject)
		if string(jsonEpisodeObject) == "{}" { // empty struct
			dst.EpisodeObject = nil
		} else {
			match++
		}
	} else {
		dst.EpisodeObject = nil
	}

	// try to unmarshal data into TrackObject
	err = newStrictDecoder(data).Decode(&dst.TrackObject)
	if err == nil {
		jsonTrackObject, _ := json.Marshal(dst.TrackObject)
		if string(jsonTrackObject) == "{}" { // empty struct
			dst.TrackObject = nil
		} else {
			match++
		}
	} else {
		dst.TrackObject = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EpisodeObject = nil
		dst.TrackObject = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CurrentlyPlayingObjectItem)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CurrentlyPlayingObjectItem)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CurrentlyPlayingObjectItem) MarshalJSON() ([]byte, error) {
	if src.EpisodeObject != nil {
		return json.Marshal(&src.EpisodeObject)
	}

	if src.TrackObject != nil {
		return json.Marshal(&src.TrackObject)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CurrentlyPlayingObjectItem) GetActualInstance() (interface{}) {
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

type NullableCurrentlyPlayingObjectItem struct {
	value *CurrentlyPlayingObjectItem
	isSet bool
}

func (v NullableCurrentlyPlayingObjectItem) Get() *CurrentlyPlayingObjectItem {
	return v.value
}

func (v *NullableCurrentlyPlayingObjectItem) Set(val *CurrentlyPlayingObjectItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentlyPlayingObjectItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentlyPlayingObjectItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentlyPlayingObjectItem(val *CurrentlyPlayingObjectItem) *NullableCurrentlyPlayingObjectItem {
	return &NullableCurrentlyPlayingObjectItem{value: val, isSet: true}
}

func (v NullableCurrentlyPlayingObjectItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentlyPlayingObjectItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


