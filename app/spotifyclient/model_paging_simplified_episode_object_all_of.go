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

// checks if the PagingSimplifiedEpisodeObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingSimplifiedEpisodeObjectAllOf{}

// PagingSimplifiedEpisodeObjectAllOf struct for PagingSimplifiedEpisodeObjectAllOf
type PagingSimplifiedEpisodeObjectAllOf struct {
	Items []SimplifiedEpisodeObject `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PagingSimplifiedEpisodeObjectAllOf PagingSimplifiedEpisodeObjectAllOf

// NewPagingSimplifiedEpisodeObjectAllOf instantiates a new PagingSimplifiedEpisodeObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingSimplifiedEpisodeObjectAllOf() *PagingSimplifiedEpisodeObjectAllOf {
	this := PagingSimplifiedEpisodeObjectAllOf{}
	return &this
}

// NewPagingSimplifiedEpisodeObjectAllOfWithDefaults instantiates a new PagingSimplifiedEpisodeObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingSimplifiedEpisodeObjectAllOfWithDefaults() *PagingSimplifiedEpisodeObjectAllOf {
	this := PagingSimplifiedEpisodeObjectAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PagingSimplifiedEpisodeObjectAllOf) GetItems() []SimplifiedEpisodeObject {
	if o == nil || IsNil(o.Items) {
		var ret []SimplifiedEpisodeObject
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingSimplifiedEpisodeObjectAllOf) GetItemsOk() ([]SimplifiedEpisodeObject, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PagingSimplifiedEpisodeObjectAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SimplifiedEpisodeObject and assigns it to the Items field.
func (o *PagingSimplifiedEpisodeObjectAllOf) SetItems(v []SimplifiedEpisodeObject) {
	o.Items = v
}

func (o PagingSimplifiedEpisodeObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingSimplifiedEpisodeObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PagingSimplifiedEpisodeObjectAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPagingSimplifiedEpisodeObjectAllOf := _PagingSimplifiedEpisodeObjectAllOf{}

	if err = json.Unmarshal(bytes, &varPagingSimplifiedEpisodeObjectAllOf); err == nil {
		*o = PagingSimplifiedEpisodeObjectAllOf(varPagingSimplifiedEpisodeObjectAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePagingSimplifiedEpisodeObjectAllOf struct {
	value *PagingSimplifiedEpisodeObjectAllOf
	isSet bool
}

func (v NullablePagingSimplifiedEpisodeObjectAllOf) Get() *PagingSimplifiedEpisodeObjectAllOf {
	return v.value
}

func (v *NullablePagingSimplifiedEpisodeObjectAllOf) Set(val *PagingSimplifiedEpisodeObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingSimplifiedEpisodeObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingSimplifiedEpisodeObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingSimplifiedEpisodeObjectAllOf(val *PagingSimplifiedEpisodeObjectAllOf) *NullablePagingSimplifiedEpisodeObjectAllOf {
	return &NullablePagingSimplifiedEpisodeObjectAllOf{value: val, isSet: true}
}

func (v NullablePagingSimplifiedEpisodeObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingSimplifiedEpisodeObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


