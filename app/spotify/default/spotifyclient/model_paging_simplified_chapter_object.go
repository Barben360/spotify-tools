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

// checks if the PagingSimplifiedChapterObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingSimplifiedChapterObject{}

// PagingSimplifiedChapterObject struct for PagingSimplifiedChapterObject
type PagingSimplifiedChapterObject struct {
	// A link to the Web API endpoint returning the full result of the request 
	Href string `json:"href"`
	// The maximum number of items in the response (as set in the query or by default). 
	Limit int32 `json:"limit"`
	// URL to the next page of items. ( `null` if none) 
	Next NullableString `json:"next"`
	// The offset of the items returned (as set in the query or by default) 
	Offset int32 `json:"offset"`
	// URL to the previous page of items. ( `null` if none) 
	Previous NullableString `json:"previous"`
	// The total number of items available to return. 
	Total int32 `json:"total"`
	Items []SimplifiedChapterObject `json:"items"`
	AdditionalProperties map[string]interface{}
}

type _PagingSimplifiedChapterObject PagingSimplifiedChapterObject

// NewPagingSimplifiedChapterObject instantiates a new PagingSimplifiedChapterObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingSimplifiedChapterObject(href string, limit int32, next NullableString, offset int32, previous NullableString, total int32, items []SimplifiedChapterObject) *PagingSimplifiedChapterObject {
	this := PagingSimplifiedChapterObject{}
	this.Href = href
	this.Limit = limit
	this.Next = next
	this.Offset = offset
	this.Previous = previous
	this.Total = total
	this.Items = items
	return &this
}

// NewPagingSimplifiedChapterObjectWithDefaults instantiates a new PagingSimplifiedChapterObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingSimplifiedChapterObjectWithDefaults() *PagingSimplifiedChapterObject {
	this := PagingSimplifiedChapterObject{}
	return &this
}

// GetHref returns the Href field value
func (o *PagingSimplifiedChapterObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *PagingSimplifiedChapterObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *PagingSimplifiedChapterObject) SetHref(v string) {
	o.Href = v
}

// GetLimit returns the Limit field value
func (o *PagingSimplifiedChapterObject) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PagingSimplifiedChapterObject) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PagingSimplifiedChapterObject) SetLimit(v int32) {
	o.Limit = v
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PagingSimplifiedChapterObject) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PagingSimplifiedChapterObject) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *PagingSimplifiedChapterObject) SetNext(v string) {
	o.Next.Set(&v)
}

// GetOffset returns the Offset field value
func (o *PagingSimplifiedChapterObject) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *PagingSimplifiedChapterObject) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *PagingSimplifiedChapterObject) SetOffset(v int32) {
	o.Offset = v
}

// GetPrevious returns the Previous field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PagingSimplifiedChapterObject) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}

	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PagingSimplifiedChapterObject) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// SetPrevious sets field value
func (o *PagingSimplifiedChapterObject) SetPrevious(v string) {
	o.Previous.Set(&v)
}

// GetTotal returns the Total field value
func (o *PagingSimplifiedChapterObject) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PagingSimplifiedChapterObject) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PagingSimplifiedChapterObject) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *PagingSimplifiedChapterObject) GetItems() []SimplifiedChapterObject {
	if o == nil {
		var ret []SimplifiedChapterObject
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PagingSimplifiedChapterObject) GetItemsOk() ([]SimplifiedChapterObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PagingSimplifiedChapterObject) SetItems(v []SimplifiedChapterObject) {
	o.Items = v
}

func (o PagingSimplifiedChapterObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingSimplifiedChapterObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["limit"] = o.Limit
	toSerialize["next"] = o.Next.Get()
	toSerialize["offset"] = o.Offset
	toSerialize["previous"] = o.Previous.Get()
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PagingSimplifiedChapterObject) UnmarshalJSON(bytes []byte) (err error) {
	varPagingSimplifiedChapterObject := _PagingSimplifiedChapterObject{}

	if err = json.Unmarshal(bytes, &varPagingSimplifiedChapterObject); err == nil {
		*o = PagingSimplifiedChapterObject(varPagingSimplifiedChapterObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next")
		delete(additionalProperties, "offset")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "total")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePagingSimplifiedChapterObject struct {
	value *PagingSimplifiedChapterObject
	isSet bool
}

func (v NullablePagingSimplifiedChapterObject) Get() *PagingSimplifiedChapterObject {
	return v.value
}

func (v *NullablePagingSimplifiedChapterObject) Set(val *PagingSimplifiedChapterObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingSimplifiedChapterObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingSimplifiedChapterObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingSimplifiedChapterObject(val *PagingSimplifiedChapterObject) *NullablePagingSimplifiedChapterObject {
	return &NullablePagingSimplifiedChapterObject{value: val, isSet: true}
}

func (v NullablePagingSimplifiedChapterObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingSimplifiedChapterObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


