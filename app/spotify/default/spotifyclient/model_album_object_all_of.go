/*
Spotify Web API with fixes and improvements from sonallux

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 2023.2.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spotifyclient

import (
	"encoding/json"
)

// checks if the AlbumObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumObjectAllOf{}

// AlbumObjectAllOf struct for AlbumObjectAllOf
type AlbumObjectAllOf struct {
	// The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist. 
	Artists []SimplifiedArtistObject `json:"artists,omitempty"`
	Tracks *AlbumObjectAllOfTracks `json:"tracks,omitempty"`
	// The popularity of the album, with 100 being the most popular. The popularity is calculated from the popularity of the album's individual tracks.
	Popularity *int32 `json:"popularity,omitempty"`
	// The label for the album.
	Label *string `json:"label,omitempty"`
	ExternalIds *AlbumObjectAllOfExternalIds `json:"external_ids,omitempty"`
	// A list of the genres used to classify the album. (If not yet classified, the array is empty.)
	Genres []string `json:"genres,omitempty"`
	// The copyright statements of the album.
	Copyrights []CopyrightObject `json:"copyrights,omitempty"`
}

// NewAlbumObjectAllOf instantiates a new AlbumObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumObjectAllOf() *AlbumObjectAllOf {
	this := AlbumObjectAllOf{}
	return &this
}

// NewAlbumObjectAllOfWithDefaults instantiates a new AlbumObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumObjectAllOfWithDefaults() *AlbumObjectAllOf {
	this := AlbumObjectAllOf{}
	return &this
}

// GetArtists returns the Artists field value if set, zero value otherwise.
func (o *AlbumObjectAllOf) GetArtists() []SimplifiedArtistObject {
	if o == nil || IsNil(o.Artists) {
		var ret []SimplifiedArtistObject
		return ret
	}
	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumObjectAllOf) GetArtistsOk() ([]SimplifiedArtistObject, bool) {
	if o == nil || IsNil(o.Artists) {
		return nil, false
	}
	return o.Artists, true
}

// HasArtists returns a boolean if a field has been set.
func (o *AlbumObjectAllOf) HasArtists() bool {
	if o != nil && !IsNil(o.Artists) {
		return true
	}

	return false
}

// SetArtists gets a reference to the given []SimplifiedArtistObject and assigns it to the Artists field.
func (o *AlbumObjectAllOf) SetArtists(v []SimplifiedArtistObject) {
	o.Artists = v
}

// GetTracks returns the Tracks field value if set, zero value otherwise.
func (o *AlbumObjectAllOf) GetTracks() AlbumObjectAllOfTracks {
	if o == nil || IsNil(o.Tracks) {
		var ret AlbumObjectAllOfTracks
		return ret
	}
	return *o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumObjectAllOf) GetTracksOk() (*AlbumObjectAllOfTracks, bool) {
	if o == nil || IsNil(o.Tracks) {
		return nil, false
	}
	return o.Tracks, true
}

// HasTracks returns a boolean if a field has been set.
func (o *AlbumObjectAllOf) HasTracks() bool {
	if o != nil && !IsNil(o.Tracks) {
		return true
	}

	return false
}

// SetTracks gets a reference to the given AlbumObjectAllOfTracks and assigns it to the Tracks field.
func (o *AlbumObjectAllOf) SetTracks(v AlbumObjectAllOfTracks) {
	o.Tracks = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *AlbumObjectAllOf) GetPopularity() int32 {
	if o == nil || IsNil(o.Popularity) {
		var ret int32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumObjectAllOf) GetPopularityOk() (*int32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *AlbumObjectAllOf) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given int32 and assigns it to the Popularity field.
func (o *AlbumObjectAllOf) SetPopularity(v int32) {
	o.Popularity = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AlbumObjectAllOf) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumObjectAllOf) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AlbumObjectAllOf) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AlbumObjectAllOf) SetLabel(v string) {
	o.Label = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *AlbumObjectAllOf) GetExternalIds() AlbumObjectAllOfExternalIds {
	if o == nil || IsNil(o.ExternalIds) {
		var ret AlbumObjectAllOfExternalIds
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumObjectAllOf) GetExternalIdsOk() (*AlbumObjectAllOfExternalIds, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *AlbumObjectAllOf) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given AlbumObjectAllOfExternalIds and assigns it to the ExternalIds field.
func (o *AlbumObjectAllOf) SetExternalIds(v AlbumObjectAllOfExternalIds) {
	o.ExternalIds = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *AlbumObjectAllOf) GetGenres() []string {
	if o == nil || IsNil(o.Genres) {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumObjectAllOf) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *AlbumObjectAllOf) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *AlbumObjectAllOf) SetGenres(v []string) {
	o.Genres = v
}

// GetCopyrights returns the Copyrights field value if set, zero value otherwise.
func (o *AlbumObjectAllOf) GetCopyrights() []CopyrightObject {
	if o == nil || IsNil(o.Copyrights) {
		var ret []CopyrightObject
		return ret
	}
	return o.Copyrights
}

// GetCopyrightsOk returns a tuple with the Copyrights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumObjectAllOf) GetCopyrightsOk() ([]CopyrightObject, bool) {
	if o == nil || IsNil(o.Copyrights) {
		return nil, false
	}
	return o.Copyrights, true
}

// HasCopyrights returns a boolean if a field has been set.
func (o *AlbumObjectAllOf) HasCopyrights() bool {
	if o != nil && !IsNil(o.Copyrights) {
		return true
	}

	return false
}

// SetCopyrights gets a reference to the given []CopyrightObject and assigns it to the Copyrights field.
func (o *AlbumObjectAllOf) SetCopyrights(v []CopyrightObject) {
	o.Copyrights = v
}

func (o AlbumObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Artists) {
		toSerialize["artists"] = o.Artists
	}
	if !IsNil(o.Tracks) {
		toSerialize["tracks"] = o.Tracks
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["external_ids"] = o.ExternalIds
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.Copyrights) {
		toSerialize["copyrights"] = o.Copyrights
	}
	return toSerialize, nil
}

type NullableAlbumObjectAllOf struct {
	value *AlbumObjectAllOf
	isSet bool
}

func (v NullableAlbumObjectAllOf) Get() *AlbumObjectAllOf {
	return v.value
}

func (v *NullableAlbumObjectAllOf) Set(val *AlbumObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumObjectAllOf(val *AlbumObjectAllOf) *NullableAlbumObjectAllOf {
	return &NullableAlbumObjectAllOf{value: val, isSet: true}
}

func (v NullableAlbumObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


