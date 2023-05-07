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

// checks if the SimplifiedArtistObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedArtistObject{}

// SimplifiedArtistObject struct for SimplifiedArtistObject
type SimplifiedArtistObject struct {
	ExternalUrls *ArtistObjectExternalUrls `json:"external_urls,omitempty"`
	// A link to the Web API endpoint providing full details of the artist. 
	Href *string `json:"href,omitempty"`
	// The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the artist. 
	Id *string `json:"id,omitempty"`
	// The name of the artist. 
	Name *string `json:"name,omitempty"`
	// The object type. 
	Type *string `json:"type,omitempty"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the artist. 
	Uri *string `json:"uri,omitempty"`
}

// NewSimplifiedArtistObject instantiates a new SimplifiedArtistObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedArtistObject() *SimplifiedArtistObject {
	this := SimplifiedArtistObject{}
	return &this
}

// NewSimplifiedArtistObjectWithDefaults instantiates a new SimplifiedArtistObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedArtistObjectWithDefaults() *SimplifiedArtistObject {
	this := SimplifiedArtistObject{}
	return &this
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *SimplifiedArtistObject) GetExternalUrls() ArtistObjectExternalUrls {
	if o == nil || IsNil(o.ExternalUrls) {
		var ret ArtistObjectExternalUrls
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedArtistObject) GetExternalUrlsOk() (*ArtistObjectExternalUrls, bool) {
	if o == nil || IsNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *SimplifiedArtistObject) HasExternalUrls() bool {
	if o != nil && !IsNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given ArtistObjectExternalUrls and assigns it to the ExternalUrls field.
func (o *SimplifiedArtistObject) SetExternalUrls(v ArtistObjectExternalUrls) {
	o.ExternalUrls = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SimplifiedArtistObject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedArtistObject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SimplifiedArtistObject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SimplifiedArtistObject) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SimplifiedArtistObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedArtistObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SimplifiedArtistObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SimplifiedArtistObject) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimplifiedArtistObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedArtistObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimplifiedArtistObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimplifiedArtistObject) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimplifiedArtistObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedArtistObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimplifiedArtistObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SimplifiedArtistObject) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *SimplifiedArtistObject) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedArtistObject) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *SimplifiedArtistObject) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *SimplifiedArtistObject) SetUri(v string) {
	o.Uri = &v
}

func (o SimplifiedArtistObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedArtistObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableSimplifiedArtistObject struct {
	value *SimplifiedArtistObject
	isSet bool
}

func (v NullableSimplifiedArtistObject) Get() *SimplifiedArtistObject {
	return v.value
}

func (v *NullableSimplifiedArtistObject) Set(val *SimplifiedArtistObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedArtistObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedArtistObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedArtistObject(val *SimplifiedArtistObject) *NullableSimplifiedArtistObject {
	return &NullableSimplifiedArtistObject{value: val, isSet: true}
}

func (v NullableSimplifiedArtistObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedArtistObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


