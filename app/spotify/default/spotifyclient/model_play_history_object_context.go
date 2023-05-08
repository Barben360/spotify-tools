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

// checks if the PlayHistoryObjectContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayHistoryObjectContext{}

// PlayHistoryObjectContext The context the track was played from.
type PlayHistoryObjectContext struct {
	// The object type, e.g. \"artist\", \"playlist\", \"album\", \"show\". 
	Type *string `json:"type,omitempty"`
	// A link to the Web API endpoint providing full details of the track.
	Href *string `json:"href,omitempty"`
	ExternalUrls *ContextObjectExternalUrls `json:"external_urls,omitempty"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the context. 
	Uri *string `json:"uri,omitempty"`
}

// NewPlayHistoryObjectContext instantiates a new PlayHistoryObjectContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayHistoryObjectContext() *PlayHistoryObjectContext {
	this := PlayHistoryObjectContext{}
	return &this
}

// NewPlayHistoryObjectContextWithDefaults instantiates a new PlayHistoryObjectContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayHistoryObjectContextWithDefaults() *PlayHistoryObjectContext {
	this := PlayHistoryObjectContext{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlayHistoryObjectContext) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayHistoryObjectContext) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlayHistoryObjectContext) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlayHistoryObjectContext) SetType(v string) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PlayHistoryObjectContext) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayHistoryObjectContext) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PlayHistoryObjectContext) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PlayHistoryObjectContext) SetHref(v string) {
	o.Href = &v
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *PlayHistoryObjectContext) GetExternalUrls() ContextObjectExternalUrls {
	if o == nil || IsNil(o.ExternalUrls) {
		var ret ContextObjectExternalUrls
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayHistoryObjectContext) GetExternalUrlsOk() (*ContextObjectExternalUrls, bool) {
	if o == nil || IsNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *PlayHistoryObjectContext) HasExternalUrls() bool {
	if o != nil && !IsNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given ContextObjectExternalUrls and assigns it to the ExternalUrls field.
func (o *PlayHistoryObjectContext) SetExternalUrls(v ContextObjectExternalUrls) {
	o.ExternalUrls = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *PlayHistoryObjectContext) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayHistoryObjectContext) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PlayHistoryObjectContext) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *PlayHistoryObjectContext) SetUri(v string) {
	o.Uri = &v
}

func (o PlayHistoryObjectContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayHistoryObjectContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullablePlayHistoryObjectContext struct {
	value *PlayHistoryObjectContext
	isSet bool
}

func (v NullablePlayHistoryObjectContext) Get() *PlayHistoryObjectContext {
	return v.value
}

func (v *NullablePlayHistoryObjectContext) Set(val *PlayHistoryObjectContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayHistoryObjectContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayHistoryObjectContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayHistoryObjectContext(val *PlayHistoryObjectContext) *NullablePlayHistoryObjectContext {
	return &NullablePlayHistoryObjectContext{value: val, isSet: true}
}

func (v NullablePlayHistoryObjectContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayHistoryObjectContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

