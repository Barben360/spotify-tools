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

// checks if the ChapterRestrictionObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChapterRestrictionObject{}

// ChapterRestrictionObject struct for ChapterRestrictionObject
type ChapterRestrictionObject struct {
	// The reason for the restriction. Supported values: - `market` - The content item is not available in the given market. - `product` - The content item is not available for the user's subscription type. - `explicit` - The content item is explicit and the user's account is set to not play explicit content. - `payment_required` - Payment is required to play the content item.  Additional reasons may be added in the future. **Note**: If you use this field, make sure that your application safely handles unknown values. 
	Reason *string `json:"reason,omitempty"`
}

// NewChapterRestrictionObject instantiates a new ChapterRestrictionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChapterRestrictionObject() *ChapterRestrictionObject {
	this := ChapterRestrictionObject{}
	return &this
}

// NewChapterRestrictionObjectWithDefaults instantiates a new ChapterRestrictionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChapterRestrictionObjectWithDefaults() *ChapterRestrictionObject {
	this := ChapterRestrictionObject{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ChapterRestrictionObject) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterRestrictionObject) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ChapterRestrictionObject) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ChapterRestrictionObject) SetReason(v string) {
	o.Reason = &v
}

func (o ChapterRestrictionObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChapterRestrictionObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableChapterRestrictionObject struct {
	value *ChapterRestrictionObject
	isSet bool
}

func (v NullableChapterRestrictionObject) Get() *ChapterRestrictionObject {
	return v.value
}

func (v *NullableChapterRestrictionObject) Set(val *ChapterRestrictionObject) {
	v.value = val
	v.isSet = true
}

func (v NullableChapterRestrictionObject) IsSet() bool {
	return v.isSet
}

func (v *NullableChapterRestrictionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChapterRestrictionObject(val *ChapterRestrictionObject) *NullableChapterRestrictionObject {
	return &NullableChapterRestrictionObject{value: val, isSet: true}
}

func (v NullableChapterRestrictionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChapterRestrictionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


