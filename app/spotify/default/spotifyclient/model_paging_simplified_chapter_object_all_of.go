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

// checks if the PagingSimplifiedChapterObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingSimplifiedChapterObjectAllOf{}

// PagingSimplifiedChapterObjectAllOf struct for PagingSimplifiedChapterObjectAllOf
type PagingSimplifiedChapterObjectAllOf struct {
	Items []SimplifiedChapterObject `json:"items,omitempty"`
}

// NewPagingSimplifiedChapterObjectAllOf instantiates a new PagingSimplifiedChapterObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingSimplifiedChapterObjectAllOf() *PagingSimplifiedChapterObjectAllOf {
	this := PagingSimplifiedChapterObjectAllOf{}
	return &this
}

// NewPagingSimplifiedChapterObjectAllOfWithDefaults instantiates a new PagingSimplifiedChapterObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingSimplifiedChapterObjectAllOfWithDefaults() *PagingSimplifiedChapterObjectAllOf {
	this := PagingSimplifiedChapterObjectAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PagingSimplifiedChapterObjectAllOf) GetItems() []SimplifiedChapterObject {
	if o == nil || IsNil(o.Items) {
		var ret []SimplifiedChapterObject
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingSimplifiedChapterObjectAllOf) GetItemsOk() ([]SimplifiedChapterObject, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PagingSimplifiedChapterObjectAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SimplifiedChapterObject and assigns it to the Items field.
func (o *PagingSimplifiedChapterObjectAllOf) SetItems(v []SimplifiedChapterObject) {
	o.Items = v
}

func (o PagingSimplifiedChapterObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingSimplifiedChapterObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullablePagingSimplifiedChapterObjectAllOf struct {
	value *PagingSimplifiedChapterObjectAllOf
	isSet bool
}

func (v NullablePagingSimplifiedChapterObjectAllOf) Get() *PagingSimplifiedChapterObjectAllOf {
	return v.value
}

func (v *NullablePagingSimplifiedChapterObjectAllOf) Set(val *PagingSimplifiedChapterObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingSimplifiedChapterObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingSimplifiedChapterObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingSimplifiedChapterObjectAllOf(val *PagingSimplifiedChapterObjectAllOf) *NullablePagingSimplifiedChapterObjectAllOf {
	return &NullablePagingSimplifiedChapterObjectAllOf{value: val, isSet: true}
}

func (v NullablePagingSimplifiedChapterObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingSimplifiedChapterObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


