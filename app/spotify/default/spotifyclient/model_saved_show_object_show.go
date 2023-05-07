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

// checks if the SavedShowObjectShow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedShowObjectShow{}

// SavedShowObjectShow Information about the show.
type SavedShowObjectShow struct {
	// A list of the countries in which the show can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. 
	AvailableMarkets []string `json:"available_markets"`
	// The copyright statements of the show. 
	Copyrights []CopyrightObject `json:"copyrights"`
	// A description of the show. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed. 
	Description string `json:"description"`
	// A description of the show. This field may contain HTML tags. 
	HtmlDescription string `json:"html_description"`
	// Whether or not the show has explicit content (true = yes it does; false = no it does not OR unknown). 
	Explicit bool `json:"explicit"`
	ExternalUrls ShowBaseExternalUrls `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the show. 
	Href string `json:"href"`
	// The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the show. 
	Id string `json:"id"`
	// The cover art for the show in various sizes, widest first. 
	Images []ImageObject `json:"images"`
	// True if all of the shows episodes are hosted outside of Spotify's CDN. This field might be `null` in some cases. 
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// A list of the languages used in the show, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. 
	Languages []string `json:"languages"`
	// The media type of the show. 
	MediaType string `json:"media_type"`
	// The name of the episode. 
	Name string `json:"name"`
	// The publisher of the show. 
	Publisher string `json:"publisher"`
	// The object type. 
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the show. 
	Uri string `json:"uri"`
	// The total number of episodes in the show. 
	TotalEpisodes int32 `json:"total_episodes"`
}

// NewSavedShowObjectShow instantiates a new SavedShowObjectShow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedShowObjectShow(availableMarkets []string, copyrights []CopyrightObject, description string, htmlDescription string, explicit bool, externalUrls ShowBaseExternalUrls, href string, id string, images []ImageObject, isExternallyHosted bool, languages []string, mediaType string, name string, publisher string, type_ string, uri string, totalEpisodes int32) *SavedShowObjectShow {
	this := SavedShowObjectShow{}
	this.AvailableMarkets = availableMarkets
	this.Copyrights = copyrights
	this.Description = description
	this.HtmlDescription = htmlDescription
	this.Explicit = explicit
	this.ExternalUrls = externalUrls
	this.Href = href
	this.Id = id
	this.Images = images
	this.IsExternallyHosted = isExternallyHosted
	this.Languages = languages
	this.MediaType = mediaType
	this.Name = name
	this.Publisher = publisher
	this.Type = type_
	this.Uri = uri
	this.TotalEpisodes = totalEpisodes
	return &this
}

// NewSavedShowObjectShowWithDefaults instantiates a new SavedShowObjectShow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedShowObjectShowWithDefaults() *SavedShowObjectShow {
	this := SavedShowObjectShow{}
	return &this
}

// GetAvailableMarkets returns the AvailableMarkets field value
func (o *SavedShowObjectShow) GetAvailableMarkets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableMarkets
}

// GetAvailableMarketsOk returns a tuple with the AvailableMarkets field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetAvailableMarketsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableMarkets, true
}

// SetAvailableMarkets sets field value
func (o *SavedShowObjectShow) SetAvailableMarkets(v []string) {
	o.AvailableMarkets = v
}

// GetCopyrights returns the Copyrights field value
func (o *SavedShowObjectShow) GetCopyrights() []CopyrightObject {
	if o == nil {
		var ret []CopyrightObject
		return ret
	}

	return o.Copyrights
}

// GetCopyrightsOk returns a tuple with the Copyrights field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetCopyrightsOk() ([]CopyrightObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Copyrights, true
}

// SetCopyrights sets field value
func (o *SavedShowObjectShow) SetCopyrights(v []CopyrightObject) {
	o.Copyrights = v
}

// GetDescription returns the Description field value
func (o *SavedShowObjectShow) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SavedShowObjectShow) SetDescription(v string) {
	o.Description = v
}

// GetHtmlDescription returns the HtmlDescription field value
func (o *SavedShowObjectShow) GetHtmlDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlDescription
}

// GetHtmlDescriptionOk returns a tuple with the HtmlDescription field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetHtmlDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlDescription, true
}

// SetHtmlDescription sets field value
func (o *SavedShowObjectShow) SetHtmlDescription(v string) {
	o.HtmlDescription = v
}

// GetExplicit returns the Explicit field value
func (o *SavedShowObjectShow) GetExplicit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetExplicitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explicit, true
}

// SetExplicit sets field value
func (o *SavedShowObjectShow) SetExplicit(v bool) {
	o.Explicit = v
}

// GetExternalUrls returns the ExternalUrls field value
func (o *SavedShowObjectShow) GetExternalUrls() ShowBaseExternalUrls {
	if o == nil {
		var ret ShowBaseExternalUrls
		return ret
	}

	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetExternalUrlsOk() (*ShowBaseExternalUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrls, true
}

// SetExternalUrls sets field value
func (o *SavedShowObjectShow) SetExternalUrls(v ShowBaseExternalUrls) {
	o.ExternalUrls = v
}

// GetHref returns the Href field value
func (o *SavedShowObjectShow) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *SavedShowObjectShow) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *SavedShowObjectShow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SavedShowObjectShow) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *SavedShowObjectShow) GetImages() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetImagesOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *SavedShowObjectShow) SetImages(v []ImageObject) {
	o.Images = v
}

// GetIsExternallyHosted returns the IsExternallyHosted field value
func (o *SavedShowObjectShow) GetIsExternallyHosted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExternallyHosted
}

// GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetIsExternallyHostedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExternallyHosted, true
}

// SetIsExternallyHosted sets field value
func (o *SavedShowObjectShow) SetIsExternallyHosted(v bool) {
	o.IsExternallyHosted = v
}

// GetLanguages returns the Languages field value
func (o *SavedShowObjectShow) GetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Languages, true
}

// SetLanguages sets field value
func (o *SavedShowObjectShow) SetLanguages(v []string) {
	o.Languages = v
}

// GetMediaType returns the MediaType field value
func (o *SavedShowObjectShow) GetMediaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *SavedShowObjectShow) SetMediaType(v string) {
	o.MediaType = v
}

// GetName returns the Name field value
func (o *SavedShowObjectShow) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SavedShowObjectShow) SetName(v string) {
	o.Name = v
}

// GetPublisher returns the Publisher field value
func (o *SavedShowObjectShow) GetPublisher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetPublisherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Publisher, true
}

// SetPublisher sets field value
func (o *SavedShowObjectShow) SetPublisher(v string) {
	o.Publisher = v
}

// GetType returns the Type field value
func (o *SavedShowObjectShow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SavedShowObjectShow) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *SavedShowObjectShow) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *SavedShowObjectShow) SetUri(v string) {
	o.Uri = v
}

// GetTotalEpisodes returns the TotalEpisodes field value
func (o *SavedShowObjectShow) GetTotalEpisodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalEpisodes
}

// GetTotalEpisodesOk returns a tuple with the TotalEpisodes field value
// and a boolean to check if the value has been set.
func (o *SavedShowObjectShow) GetTotalEpisodesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalEpisodes, true
}

// SetTotalEpisodes sets field value
func (o *SavedShowObjectShow) SetTotalEpisodes(v int32) {
	o.TotalEpisodes = v
}

func (o SavedShowObjectShow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedShowObjectShow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["available_markets"] = o.AvailableMarkets
	toSerialize["copyrights"] = o.Copyrights
	toSerialize["description"] = o.Description
	toSerialize["html_description"] = o.HtmlDescription
	toSerialize["explicit"] = o.Explicit
	toSerialize["external_urls"] = o.ExternalUrls
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	toSerialize["is_externally_hosted"] = o.IsExternallyHosted
	toSerialize["languages"] = o.Languages
	toSerialize["media_type"] = o.MediaType
	toSerialize["name"] = o.Name
	toSerialize["publisher"] = o.Publisher
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	toSerialize["total_episodes"] = o.TotalEpisodes
	return toSerialize, nil
}

type NullableSavedShowObjectShow struct {
	value *SavedShowObjectShow
	isSet bool
}

func (v NullableSavedShowObjectShow) Get() *SavedShowObjectShow {
	return v.value
}

func (v *NullableSavedShowObjectShow) Set(val *SavedShowObjectShow) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedShowObjectShow) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedShowObjectShow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedShowObjectShow(val *SavedShowObjectShow) *NullableSavedShowObjectShow {
	return &NullableSavedShowObjectShow{value: val, isSet: true}
}

func (v NullableSavedShowObjectShow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedShowObjectShow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


