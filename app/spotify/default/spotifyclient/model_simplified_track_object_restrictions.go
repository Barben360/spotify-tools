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

// checks if the SimplifiedTrackObjectRestrictions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedTrackObjectRestrictions{}

// SimplifiedTrackObjectRestrictions Included in the response when a content restriction is applied. 
type SimplifiedTrackObjectRestrictions struct {
	// The reason for the restriction. Supported values: - `market` - The content item is not available in the given market. - `product` - The content item is not available for the user's subscription type. - `explicit` - The content item is explicit and the user's account is set to not play explicit content.  Additional reasons may be added in the future. **Note**: If you use this field, make sure that your application safely handles unknown values. 
	Reason *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedTrackObjectRestrictions SimplifiedTrackObjectRestrictions

// NewSimplifiedTrackObjectRestrictions instantiates a new SimplifiedTrackObjectRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedTrackObjectRestrictions() *SimplifiedTrackObjectRestrictions {
	this := SimplifiedTrackObjectRestrictions{}
	return &this
}

// NewSimplifiedTrackObjectRestrictionsWithDefaults instantiates a new SimplifiedTrackObjectRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedTrackObjectRestrictionsWithDefaults() *SimplifiedTrackObjectRestrictions {
	this := SimplifiedTrackObjectRestrictions{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SimplifiedTrackObjectRestrictions) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedTrackObjectRestrictions) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SimplifiedTrackObjectRestrictions) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SimplifiedTrackObjectRestrictions) SetReason(v string) {
	o.Reason = &v
}

func (o SimplifiedTrackObjectRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedTrackObjectRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedTrackObjectRestrictions) UnmarshalJSON(bytes []byte) (err error) {
	varSimplifiedTrackObjectRestrictions := _SimplifiedTrackObjectRestrictions{}

	if err = json.Unmarshal(bytes, &varSimplifiedTrackObjectRestrictions); err == nil {
		*o = SimplifiedTrackObjectRestrictions(varSimplifiedTrackObjectRestrictions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedTrackObjectRestrictions struct {
	value *SimplifiedTrackObjectRestrictions
	isSet bool
}

func (v NullableSimplifiedTrackObjectRestrictions) Get() *SimplifiedTrackObjectRestrictions {
	return v.value
}

func (v *NullableSimplifiedTrackObjectRestrictions) Set(val *SimplifiedTrackObjectRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedTrackObjectRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedTrackObjectRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedTrackObjectRestrictions(val *SimplifiedTrackObjectRestrictions) *NullableSimplifiedTrackObjectRestrictions {
	return &NullableSimplifiedTrackObjectRestrictions{value: val, isSet: true}
}

func (v NullableSimplifiedTrackObjectRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedTrackObjectRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


