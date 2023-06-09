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

// checks if the PlaylistOwnerObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaylistOwnerObjectAllOf{}

// PlaylistOwnerObjectAllOf struct for PlaylistOwnerObjectAllOf
type PlaylistOwnerObjectAllOf struct {
	// The name displayed on the user's profile. `null` if not available. 
	DisplayName NullableString `json:"display_name,omitempty"`
}

// NewPlaylistOwnerObjectAllOf instantiates a new PlaylistOwnerObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistOwnerObjectAllOf() *PlaylistOwnerObjectAllOf {
	this := PlaylistOwnerObjectAllOf{}
	return &this
}

// NewPlaylistOwnerObjectAllOfWithDefaults instantiates a new PlaylistOwnerObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistOwnerObjectAllOfWithDefaults() *PlaylistOwnerObjectAllOf {
	this := PlaylistOwnerObjectAllOf{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaylistOwnerObjectAllOf) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaylistOwnerObjectAllOf) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PlaylistOwnerObjectAllOf) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *PlaylistOwnerObjectAllOf) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *PlaylistOwnerObjectAllOf) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *PlaylistOwnerObjectAllOf) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o PlaylistOwnerObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaylistOwnerObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	return toSerialize, nil
}

type NullablePlaylistOwnerObjectAllOf struct {
	value *PlaylistOwnerObjectAllOf
	isSet bool
}

func (v NullablePlaylistOwnerObjectAllOf) Get() *PlaylistOwnerObjectAllOf {
	return v.value
}

func (v *NullablePlaylistOwnerObjectAllOf) Set(val *PlaylistOwnerObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistOwnerObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistOwnerObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistOwnerObjectAllOf(val *PlaylistOwnerObjectAllOf) *NullablePlaylistOwnerObjectAllOf {
	return &NullablePlaylistOwnerObjectAllOf{value: val, isSet: true}
}

func (v NullablePlaylistOwnerObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistOwnerObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


