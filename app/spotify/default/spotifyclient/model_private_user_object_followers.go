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

// checks if the PrivateUserObjectFollowers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateUserObjectFollowers{}

// PrivateUserObjectFollowers Information about the followers of the user.
type PrivateUserObjectFollowers struct {
	// This will always be set to null, as the Web API does not support it at the moment. 
	Href NullableString `json:"href,omitempty"`
	// The total number of followers. 
	Total *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivateUserObjectFollowers PrivateUserObjectFollowers

// NewPrivateUserObjectFollowers instantiates a new PrivateUserObjectFollowers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateUserObjectFollowers() *PrivateUserObjectFollowers {
	this := PrivateUserObjectFollowers{}
	return &this
}

// NewPrivateUserObjectFollowersWithDefaults instantiates a new PrivateUserObjectFollowers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateUserObjectFollowersWithDefaults() *PrivateUserObjectFollowers {
	this := PrivateUserObjectFollowers{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateUserObjectFollowers) GetHref() string {
	if o == nil || IsNil(o.Href.Get()) {
		var ret string
		return ret
	}
	return *o.Href.Get()
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUserObjectFollowers) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Href.Get(), o.Href.IsSet()
}

// HasHref returns a boolean if a field has been set.
func (o *PrivateUserObjectFollowers) HasHref() bool {
	if o != nil && o.Href.IsSet() {
		return true
	}

	return false
}

// SetHref gets a reference to the given NullableString and assigns it to the Href field.
func (o *PrivateUserObjectFollowers) SetHref(v string) {
	o.Href.Set(&v)
}
// SetHrefNil sets the value for Href to be an explicit nil
func (o *PrivateUserObjectFollowers) SetHrefNil() {
	o.Href.Set(nil)
}

// UnsetHref ensures that no value is present for Href, not even an explicit nil
func (o *PrivateUserObjectFollowers) UnsetHref() {
	o.Href.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *PrivateUserObjectFollowers) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateUserObjectFollowers) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *PrivateUserObjectFollowers) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *PrivateUserObjectFollowers) SetTotal(v int32) {
	o.Total = &v
}

func (o PrivateUserObjectFollowers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateUserObjectFollowers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Href.IsSet() {
		toSerialize["href"] = o.Href.Get()
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrivateUserObjectFollowers) UnmarshalJSON(bytes []byte) (err error) {
	varPrivateUserObjectFollowers := _PrivateUserObjectFollowers{}

	if err = json.Unmarshal(bytes, &varPrivateUserObjectFollowers); err == nil {
		*o = PrivateUserObjectFollowers(varPrivateUserObjectFollowers)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrivateUserObjectFollowers struct {
	value *PrivateUserObjectFollowers
	isSet bool
}

func (v NullablePrivateUserObjectFollowers) Get() *PrivateUserObjectFollowers {
	return v.value
}

func (v *NullablePrivateUserObjectFollowers) Set(val *PrivateUserObjectFollowers) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateUserObjectFollowers) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateUserObjectFollowers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateUserObjectFollowers(val *PrivateUserObjectFollowers) *NullablePrivateUserObjectFollowers {
	return &NullablePrivateUserObjectFollowers{value: val, isSet: true}
}

func (v NullablePrivateUserObjectFollowers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateUserObjectFollowers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


