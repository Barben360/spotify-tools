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

// checks if the EpisodeObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EpisodeObject{}

// EpisodeObject struct for EpisodeObject
type EpisodeObject struct {
	// A URL to a 30 second preview (MP3 format) of the episode. `null` if not available. 
	AudioPreviewUrl string `json:"audio_preview_url"`
	// A description of the episode. HTML tags are stripped away from this field, use `html_description` field in case HTML tags are needed. 
	Description string `json:"description"`
	// A description of the episode. This field may contain HTML tags. 
	HtmlDescription string `json:"html_description"`
	// The episode length in milliseconds. 
	DurationMs int32 `json:"duration_ms"`
	// Whether or not the episode has explicit content (true = yes it does; false = no it does not OR unknown). 
	Explicit bool `json:"explicit"`
	ExternalUrls EpisodeBaseExternalUrls `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the episode. 
	Href string `json:"href"`
	// The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the episode. 
	Id string `json:"id"`
	// The cover art for the episode in various sizes, widest first. 
	Images []ImageObject `json:"images"`
	// True if the episode is hosted outside of Spotify's CDN. 
	IsExternallyHosted bool `json:"is_externally_hosted"`
	// True if the episode is playable in the given market. Otherwise false. 
	IsPlayable bool `json:"is_playable"`
	// The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the `languages` field instead. 
	// Deprecated
	Language *string `json:"language,omitempty"`
	// A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code. 
	Languages []string `json:"languages"`
	// The name of the episode. 
	Name string `json:"name"`
	// The date the episode was first released, for example `\"1981-12-15\"`. Depending on the precision, it might be shown as `\"1981\"` or `\"1981-12\"`. 
	ReleaseDate string `json:"release_date"`
	// The precision with which `release_date` value is known. 
	ReleaseDatePrecision string `json:"release_date_precision"`
	ResumePoint EpisodeBaseResumePoint `json:"resume_point"`
	// The object type. 
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the episode. 
	Uri string `json:"uri"`
	Restrictions *EpisodeBaseRestrictions `json:"restrictions,omitempty"`
	Show EpisodeObjectAllOfShow `json:"show"`
	AdditionalProperties map[string]interface{}
}

type _EpisodeObject EpisodeObject

// NewEpisodeObject instantiates a new EpisodeObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpisodeObject(audioPreviewUrl string, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls EpisodeBaseExternalUrls, href string, id string, images []ImageObject, isExternallyHosted bool, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, resumePoint EpisodeBaseResumePoint, type_ string, uri string, show EpisodeObjectAllOfShow) *EpisodeObject {
	this := EpisodeObject{}
	this.AudioPreviewUrl = audioPreviewUrl
	this.Description = description
	this.HtmlDescription = htmlDescription
	this.DurationMs = durationMs
	this.Explicit = explicit
	this.ExternalUrls = externalUrls
	this.Href = href
	this.Id = id
	this.Images = images
	this.IsExternallyHosted = isExternallyHosted
	this.IsPlayable = isPlayable
	this.Languages = languages
	this.Name = name
	this.ReleaseDate = releaseDate
	this.ReleaseDatePrecision = releaseDatePrecision
	this.ResumePoint = resumePoint
	this.Type = type_
	this.Uri = uri
	this.Show = show
	return &this
}

// NewEpisodeObjectWithDefaults instantiates a new EpisodeObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpisodeObjectWithDefaults() *EpisodeObject {
	this := EpisodeObject{}
	return &this
}

// GetAudioPreviewUrl returns the AudioPreviewUrl field value
func (o *EpisodeObject) GetAudioPreviewUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudioPreviewUrl
}

// GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetAudioPreviewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudioPreviewUrl, true
}

// SetAudioPreviewUrl sets field value
func (o *EpisodeObject) SetAudioPreviewUrl(v string) {
	o.AudioPreviewUrl = v
}

// GetDescription returns the Description field value
func (o *EpisodeObject) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EpisodeObject) SetDescription(v string) {
	o.Description = v
}

// GetHtmlDescription returns the HtmlDescription field value
func (o *EpisodeObject) GetHtmlDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlDescription
}

// GetHtmlDescriptionOk returns a tuple with the HtmlDescription field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetHtmlDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlDescription, true
}

// SetHtmlDescription sets field value
func (o *EpisodeObject) SetHtmlDescription(v string) {
	o.HtmlDescription = v
}

// GetDurationMs returns the DurationMs field value
func (o *EpisodeObject) GetDurationMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetDurationMsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationMs, true
}

// SetDurationMs sets field value
func (o *EpisodeObject) SetDurationMs(v int32) {
	o.DurationMs = v
}

// GetExplicit returns the Explicit field value
func (o *EpisodeObject) GetExplicit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetExplicitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explicit, true
}

// SetExplicit sets field value
func (o *EpisodeObject) SetExplicit(v bool) {
	o.Explicit = v
}

// GetExternalUrls returns the ExternalUrls field value
func (o *EpisodeObject) GetExternalUrls() EpisodeBaseExternalUrls {
	if o == nil {
		var ret EpisodeBaseExternalUrls
		return ret
	}

	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetExternalUrlsOk() (*EpisodeBaseExternalUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrls, true
}

// SetExternalUrls sets field value
func (o *EpisodeObject) SetExternalUrls(v EpisodeBaseExternalUrls) {
	o.ExternalUrls = v
}

// GetHref returns the Href field value
func (o *EpisodeObject) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *EpisodeObject) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *EpisodeObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EpisodeObject) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *EpisodeObject) GetImages() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetImagesOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *EpisodeObject) SetImages(v []ImageObject) {
	o.Images = v
}

// GetIsExternallyHosted returns the IsExternallyHosted field value
func (o *EpisodeObject) GetIsExternallyHosted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExternallyHosted
}

// GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetIsExternallyHostedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExternallyHosted, true
}

// SetIsExternallyHosted sets field value
func (o *EpisodeObject) SetIsExternallyHosted(v bool) {
	o.IsExternallyHosted = v
}

// GetIsPlayable returns the IsPlayable field value
func (o *EpisodeObject) GetIsPlayable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPlayable
}

// GetIsPlayableOk returns a tuple with the IsPlayable field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetIsPlayableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPlayable, true
}

// SetIsPlayable sets field value
func (o *EpisodeObject) SetIsPlayable(v bool) {
	o.IsPlayable = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
// Deprecated
func (o *EpisodeObject) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EpisodeObject) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *EpisodeObject) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
// Deprecated
func (o *EpisodeObject) SetLanguage(v string) {
	o.Language = &v
}

// GetLanguages returns the Languages field value
func (o *EpisodeObject) GetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetLanguagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Languages, true
}

// SetLanguages sets field value
func (o *EpisodeObject) SetLanguages(v []string) {
	o.Languages = v
}

// GetName returns the Name field value
func (o *EpisodeObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EpisodeObject) SetName(v string) {
	o.Name = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *EpisodeObject) GetReleaseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *EpisodeObject) SetReleaseDate(v string) {
	o.ReleaseDate = v
}

// GetReleaseDatePrecision returns the ReleaseDatePrecision field value
func (o *EpisodeObject) GetReleaseDatePrecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDatePrecision
}

// GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetReleaseDatePrecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDatePrecision, true
}

// SetReleaseDatePrecision sets field value
func (o *EpisodeObject) SetReleaseDatePrecision(v string) {
	o.ReleaseDatePrecision = v
}

// GetResumePoint returns the ResumePoint field value
func (o *EpisodeObject) GetResumePoint() EpisodeBaseResumePoint {
	if o == nil {
		var ret EpisodeBaseResumePoint
		return ret
	}

	return o.ResumePoint
}

// GetResumePointOk returns a tuple with the ResumePoint field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetResumePointOk() (*EpisodeBaseResumePoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResumePoint, true
}

// SetResumePoint sets field value
func (o *EpisodeObject) SetResumePoint(v EpisodeBaseResumePoint) {
	o.ResumePoint = v
}

// GetType returns the Type field value
func (o *EpisodeObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EpisodeObject) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *EpisodeObject) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *EpisodeObject) SetUri(v string) {
	o.Uri = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *EpisodeObject) GetRestrictions() EpisodeBaseRestrictions {
	if o == nil || IsNil(o.Restrictions) {
		var ret EpisodeBaseRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetRestrictionsOk() (*EpisodeBaseRestrictions, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *EpisodeObject) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given EpisodeBaseRestrictions and assigns it to the Restrictions field.
func (o *EpisodeObject) SetRestrictions(v EpisodeBaseRestrictions) {
	o.Restrictions = &v
}

// GetShow returns the Show field value
func (o *EpisodeObject) GetShow() EpisodeObjectAllOfShow {
	if o == nil {
		var ret EpisodeObjectAllOfShow
		return ret
	}

	return o.Show
}

// GetShowOk returns a tuple with the Show field value
// and a boolean to check if the value has been set.
func (o *EpisodeObject) GetShowOk() (*EpisodeObjectAllOfShow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Show, true
}

// SetShow sets field value
func (o *EpisodeObject) SetShow(v EpisodeObjectAllOfShow) {
	o.Show = v
}

func (o EpisodeObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpisodeObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audio_preview_url"] = o.AudioPreviewUrl
	toSerialize["description"] = o.Description
	toSerialize["html_description"] = o.HtmlDescription
	toSerialize["duration_ms"] = o.DurationMs
	toSerialize["explicit"] = o.Explicit
	toSerialize["external_urls"] = o.ExternalUrls
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	toSerialize["is_externally_hosted"] = o.IsExternallyHosted
	toSerialize["is_playable"] = o.IsPlayable
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	toSerialize["languages"] = o.Languages
	toSerialize["name"] = o.Name
	toSerialize["release_date"] = o.ReleaseDate
	toSerialize["release_date_precision"] = o.ReleaseDatePrecision
	toSerialize["resume_point"] = o.ResumePoint
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	toSerialize["show"] = o.Show

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EpisodeObject) UnmarshalJSON(bytes []byte) (err error) {
	varEpisodeObject := _EpisodeObject{}

	if err = json.Unmarshal(bytes, &varEpisodeObject); err == nil {
		*o = EpisodeObject(varEpisodeObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "audio_preview_url")
		delete(additionalProperties, "description")
		delete(additionalProperties, "html_description")
		delete(additionalProperties, "duration_ms")
		delete(additionalProperties, "explicit")
		delete(additionalProperties, "external_urls")
		delete(additionalProperties, "href")
		delete(additionalProperties, "id")
		delete(additionalProperties, "images")
		delete(additionalProperties, "is_externally_hosted")
		delete(additionalProperties, "is_playable")
		delete(additionalProperties, "language")
		delete(additionalProperties, "languages")
		delete(additionalProperties, "name")
		delete(additionalProperties, "release_date")
		delete(additionalProperties, "release_date_precision")
		delete(additionalProperties, "resume_point")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uri")
		delete(additionalProperties, "restrictions")
		delete(additionalProperties, "show")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEpisodeObject struct {
	value *EpisodeObject
	isSet bool
}

func (v NullableEpisodeObject) Get() *EpisodeObject {
	return v.value
}

func (v *NullableEpisodeObject) Set(val *EpisodeObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisodeObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisodeObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisodeObject(val *EpisodeObject) *NullableEpisodeObject {
	return &NullableEpisodeObject{value: val, isSet: true}
}

func (v NullableEpisodeObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisodeObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


