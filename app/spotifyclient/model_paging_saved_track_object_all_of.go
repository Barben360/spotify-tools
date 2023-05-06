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

// checks if the PagingSavedTrackObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingSavedTrackObjectAllOf{}

// PagingSavedTrackObjectAllOf struct for PagingSavedTrackObjectAllOf
type PagingSavedTrackObjectAllOf struct {
	Items []SavedTrackObject `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PagingSavedTrackObjectAllOf PagingSavedTrackObjectAllOf

// NewPagingSavedTrackObjectAllOf instantiates a new PagingSavedTrackObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingSavedTrackObjectAllOf() *PagingSavedTrackObjectAllOf {
	this := PagingSavedTrackObjectAllOf{}
	return &this
}

// NewPagingSavedTrackObjectAllOfWithDefaults instantiates a new PagingSavedTrackObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingSavedTrackObjectAllOfWithDefaults() *PagingSavedTrackObjectAllOf {
	this := PagingSavedTrackObjectAllOf{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PagingSavedTrackObjectAllOf) GetItems() []SavedTrackObject {
	if o == nil || IsNil(o.Items) {
		var ret []SavedTrackObject
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingSavedTrackObjectAllOf) GetItemsOk() ([]SavedTrackObject, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PagingSavedTrackObjectAllOf) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SavedTrackObject and assigns it to the Items field.
func (o *PagingSavedTrackObjectAllOf) SetItems(v []SavedTrackObject) {
	o.Items = v
}

func (o PagingSavedTrackObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingSavedTrackObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PagingSavedTrackObjectAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPagingSavedTrackObjectAllOf := _PagingSavedTrackObjectAllOf{}

	if err = json.Unmarshal(bytes, &varPagingSavedTrackObjectAllOf); err == nil {
		*o = PagingSavedTrackObjectAllOf(varPagingSavedTrackObjectAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePagingSavedTrackObjectAllOf struct {
	value *PagingSavedTrackObjectAllOf
	isSet bool
}

func (v NullablePagingSavedTrackObjectAllOf) Get() *PagingSavedTrackObjectAllOf {
	return v.value
}

func (v *NullablePagingSavedTrackObjectAllOf) Set(val *PagingSavedTrackObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingSavedTrackObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingSavedTrackObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingSavedTrackObjectAllOf(val *PagingSavedTrackObjectAllOf) *NullablePagingSavedTrackObjectAllOf {
	return &NullablePagingSavedTrackObjectAllOf{value: val, isSet: true}
}

func (v NullablePagingSavedTrackObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingSavedTrackObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


