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

// checks if the GetNewReleases200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNewReleases200Response{}

// GetNewReleases200Response struct for GetNewReleases200Response
type GetNewReleases200Response struct {
	Albums PagingSimplifiedAlbumObject `json:"albums"`
}

// NewGetNewReleases200Response instantiates a new GetNewReleases200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNewReleases200Response(albums PagingSimplifiedAlbumObject) *GetNewReleases200Response {
	this := GetNewReleases200Response{}
	this.Albums = albums
	return &this
}

// NewGetNewReleases200ResponseWithDefaults instantiates a new GetNewReleases200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNewReleases200ResponseWithDefaults() *GetNewReleases200Response {
	this := GetNewReleases200Response{}
	return &this
}

// GetAlbums returns the Albums field value
func (o *GetNewReleases200Response) GetAlbums() PagingSimplifiedAlbumObject {
	if o == nil {
		var ret PagingSimplifiedAlbumObject
		return ret
	}

	return o.Albums
}

// GetAlbumsOk returns a tuple with the Albums field value
// and a boolean to check if the value has been set.
func (o *GetNewReleases200Response) GetAlbumsOk() (*PagingSimplifiedAlbumObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Albums, true
}

// SetAlbums sets field value
func (o *GetNewReleases200Response) SetAlbums(v PagingSimplifiedAlbumObject) {
	o.Albums = v
}

func (o GetNewReleases200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNewReleases200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["albums"] = o.Albums
	return toSerialize, nil
}

type NullableGetNewReleases200Response struct {
	value *GetNewReleases200Response
	isSet bool
}

func (v NullableGetNewReleases200Response) Get() *GetNewReleases200Response {
	return v.value
}

func (v *NullableGetNewReleases200Response) Set(val *GetNewReleases200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNewReleases200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNewReleases200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNewReleases200Response(val *GetNewReleases200Response) *NullableGetNewReleases200Response {
	return &NullableGetNewReleases200Response{value: val, isSet: true}
}

func (v NullableGetNewReleases200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNewReleases200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


