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

// checks if the AddTracksToPlaylistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddTracksToPlaylistRequest{}

// AddTracksToPlaylistRequest struct for AddTracksToPlaylistRequest
type AddTracksToPlaylistRequest struct {
	// A JSON array of the [Spotify URIs](/documentation/web-api/concepts/spotify-uris-ids) to add. For example: `{\"uris\": [\"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\",\"spotify:track:1301WleyT98MSxVHPZCA6M\", \"spotify:episode:512ojhOuo1ktJprKbVcKyQ\"]}`<br/>A maximum of 100 items can be added in one request. _**Note**: if the `uris` parameter is present in the query string, any URIs listed here in the body will be ignored._ 
	Uris []string `json:"uris,omitempty"`
	// The position to insert the items, a zero-based index. For example, to insert the items in the first position: `position=0` ; to insert the items in the third position: `position=2`. If omitted, the items will be appended to the playlist. Items are added in the order they appear in the uris array. For example: `{\"uris\": [\"spotify:track:4iV5W9uYEdYUVa79Axb7Rh\",\"spotify:track:1301WleyT98MSxVHPZCA6M\"], \"position\": 3}` 
	Position *int32 `json:"position,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AddTracksToPlaylistRequest AddTracksToPlaylistRequest

// NewAddTracksToPlaylistRequest instantiates a new AddTracksToPlaylistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTracksToPlaylistRequest() *AddTracksToPlaylistRequest {
	this := AddTracksToPlaylistRequest{}
	return &this
}

// NewAddTracksToPlaylistRequestWithDefaults instantiates a new AddTracksToPlaylistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTracksToPlaylistRequestWithDefaults() *AddTracksToPlaylistRequest {
	this := AddTracksToPlaylistRequest{}
	return &this
}

// GetUris returns the Uris field value if set, zero value otherwise.
func (o *AddTracksToPlaylistRequest) GetUris() []string {
	if o == nil || IsNil(o.Uris) {
		var ret []string
		return ret
	}
	return o.Uris
}

// GetUrisOk returns a tuple with the Uris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTracksToPlaylistRequest) GetUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.Uris) {
		return nil, false
	}
	return o.Uris, true
}

// HasUris returns a boolean if a field has been set.
func (o *AddTracksToPlaylistRequest) HasUris() bool {
	if o != nil && !IsNil(o.Uris) {
		return true
	}

	return false
}

// SetUris gets a reference to the given []string and assigns it to the Uris field.
func (o *AddTracksToPlaylistRequest) SetUris(v []string) {
	o.Uris = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *AddTracksToPlaylistRequest) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTracksToPlaylistRequest) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *AddTracksToPlaylistRequest) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *AddTracksToPlaylistRequest) SetPosition(v int32) {
	o.Position = &v
}

func (o AddTracksToPlaylistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddTracksToPlaylistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uris) {
		toSerialize["uris"] = o.Uris
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddTracksToPlaylistRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAddTracksToPlaylistRequest := _AddTracksToPlaylistRequest{}

	if err = json.Unmarshal(bytes, &varAddTracksToPlaylistRequest); err == nil {
		*o = AddTracksToPlaylistRequest(varAddTracksToPlaylistRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "uris")
		delete(additionalProperties, "position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddTracksToPlaylistRequest struct {
	value *AddTracksToPlaylistRequest
	isSet bool
}

func (v NullableAddTracksToPlaylistRequest) Get() *AddTracksToPlaylistRequest {
	return v.value
}

func (v *NullableAddTracksToPlaylistRequest) Set(val *AddTracksToPlaylistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTracksToPlaylistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTracksToPlaylistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTracksToPlaylistRequest(val *AddTracksToPlaylistRequest) *NullableAddTracksToPlaylistRequest {
	return &NullableAddTracksToPlaylistRequest{value: val, isSet: true}
}

func (v NullableAddTracksToPlaylistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTracksToPlaylistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


