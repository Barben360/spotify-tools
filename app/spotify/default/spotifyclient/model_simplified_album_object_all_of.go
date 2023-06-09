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

// checks if the SimplifiedAlbumObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedAlbumObjectAllOf{}

// SimplifiedAlbumObjectAllOf struct for SimplifiedAlbumObjectAllOf
type SimplifiedAlbumObjectAllOf struct {
	// The field is present when getting an artist's albums. Compare to album_type this field represents relationship between the artist and the album. 
	AlbumGroup *string `json:"album_group,omitempty"`
	// The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist. 
	Artists []SimplifiedArtistObject `json:"artists"`
}

// NewSimplifiedAlbumObjectAllOf instantiates a new SimplifiedAlbumObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedAlbumObjectAllOf(artists []SimplifiedArtistObject) *SimplifiedAlbumObjectAllOf {
	this := SimplifiedAlbumObjectAllOf{}
	this.Artists = artists
	return &this
}

// NewSimplifiedAlbumObjectAllOfWithDefaults instantiates a new SimplifiedAlbumObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedAlbumObjectAllOfWithDefaults() *SimplifiedAlbumObjectAllOf {
	this := SimplifiedAlbumObjectAllOf{}
	return &this
}

// GetAlbumGroup returns the AlbumGroup field value if set, zero value otherwise.
func (o *SimplifiedAlbumObjectAllOf) GetAlbumGroup() string {
	if o == nil || IsNil(o.AlbumGroup) {
		var ret string
		return ret
	}
	return *o.AlbumGroup
}

// GetAlbumGroupOk returns a tuple with the AlbumGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedAlbumObjectAllOf) GetAlbumGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AlbumGroup) {
		return nil, false
	}
	return o.AlbumGroup, true
}

// HasAlbumGroup returns a boolean if a field has been set.
func (o *SimplifiedAlbumObjectAllOf) HasAlbumGroup() bool {
	if o != nil && !IsNil(o.AlbumGroup) {
		return true
	}

	return false
}

// SetAlbumGroup gets a reference to the given string and assigns it to the AlbumGroup field.
func (o *SimplifiedAlbumObjectAllOf) SetAlbumGroup(v string) {
	o.AlbumGroup = &v
}

// GetArtists returns the Artists field value
func (o *SimplifiedAlbumObjectAllOf) GetArtists() []SimplifiedArtistObject {
	if o == nil {
		var ret []SimplifiedArtistObject
		return ret
	}

	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value
// and a boolean to check if the value has been set.
func (o *SimplifiedAlbumObjectAllOf) GetArtistsOk() ([]SimplifiedArtistObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artists, true
}

// SetArtists sets field value
func (o *SimplifiedAlbumObjectAllOf) SetArtists(v []SimplifiedArtistObject) {
	o.Artists = v
}

func (o SimplifiedAlbumObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedAlbumObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlbumGroup) {
		toSerialize["album_group"] = o.AlbumGroup
	}
	toSerialize["artists"] = o.Artists
	return toSerialize, nil
}

type NullableSimplifiedAlbumObjectAllOf struct {
	value *SimplifiedAlbumObjectAllOf
	isSet bool
}

func (v NullableSimplifiedAlbumObjectAllOf) Get() *SimplifiedAlbumObjectAllOf {
	return v.value
}

func (v *NullableSimplifiedAlbumObjectAllOf) Set(val *SimplifiedAlbumObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedAlbumObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedAlbumObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedAlbumObjectAllOf(val *SimplifiedAlbumObjectAllOf) *NullableSimplifiedAlbumObjectAllOf {
	return &NullableSimplifiedAlbumObjectAllOf{value: val, isSet: true}
}

func (v NullableSimplifiedAlbumObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedAlbumObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


