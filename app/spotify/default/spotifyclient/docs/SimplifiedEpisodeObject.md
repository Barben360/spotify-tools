# SimplifiedEpisodeObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioPreviewUrl** | **NullableString** | A URL to a 30 second preview (MP3 format) of the episode. &#x60;null&#x60; if not available.  | 
**Description** | **string** | A description of the episode. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the episode. This field may contain HTML tags.  | 
**DurationMs** | **int32** | The episode length in milliseconds.  | 
**Explicit** | **bool** | Whether or not the episode has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalUrls** | [**ExternalUrlObject**](ExternalUrlObject.md) | External URLs for this episode.  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the episode.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the episode.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the episode in various sizes, widest first.  | 
**IsExternallyHosted** | **bool** | True if the episode is hosted outside of Spotify&#39;s CDN.  | 
**IsPlayable** | **bool** | True if the episode is playable in the given market. Otherwise false.  | 
**Language** | Pointer to **string** | The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the &#x60;languages&#x60; field instead.  | [optional] 
**Languages** | **[]string** | A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**Name** | **string** | The name of the episode.  | 
**ReleaseDate** | **string** | The date the episode was first released, for example &#x60;\&quot;1981-12-15\&quot;&#x60;. Depending on the precision, it might be shown as &#x60;\&quot;1981\&quot;&#x60; or &#x60;\&quot;1981-12\&quot;&#x60;.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**ResumePoint** | Pointer to [**ResumePointObject**](ResumePointObject.md) | The user&#39;s most recent position in the episode. Set if the supplied access token is a user token and has the scope &#39;user-read-playback-position&#39;.  | [optional] 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the episode.  | 
**Restrictions** | Pointer to [**EpisodeRestrictionObject**](EpisodeRestrictionObject.md) | Included in the response when a content restriction is applied.  | [optional] 

## Methods

### NewSimplifiedEpisodeObject

`func NewSimplifiedEpisodeObject(audioPreviewUrl NullableString, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, isExternallyHosted bool, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, ) *SimplifiedEpisodeObject`

NewSimplifiedEpisodeObject instantiates a new SimplifiedEpisodeObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedEpisodeObjectWithDefaults

`func NewSimplifiedEpisodeObjectWithDefaults() *SimplifiedEpisodeObject`

NewSimplifiedEpisodeObjectWithDefaults instantiates a new SimplifiedEpisodeObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioPreviewUrl

`func (o *SimplifiedEpisodeObject) GetAudioPreviewUrl() string`

GetAudioPreviewUrl returns the AudioPreviewUrl field if non-nil, zero value otherwise.

### GetAudioPreviewUrlOk

`func (o *SimplifiedEpisodeObject) GetAudioPreviewUrlOk() (*string, bool)`

GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioPreviewUrl

`func (o *SimplifiedEpisodeObject) SetAudioPreviewUrl(v string)`

SetAudioPreviewUrl sets AudioPreviewUrl field to given value.


### SetAudioPreviewUrlNil

`func (o *SimplifiedEpisodeObject) SetAudioPreviewUrlNil(b bool)`

 SetAudioPreviewUrlNil sets the value for AudioPreviewUrl to be an explicit nil

### UnsetAudioPreviewUrl
`func (o *SimplifiedEpisodeObject) UnsetAudioPreviewUrl()`

UnsetAudioPreviewUrl ensures that no value is present for AudioPreviewUrl, not even an explicit nil
### GetDescription

`func (o *SimplifiedEpisodeObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimplifiedEpisodeObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimplifiedEpisodeObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *SimplifiedEpisodeObject) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *SimplifiedEpisodeObject) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *SimplifiedEpisodeObject) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetDurationMs

`func (o *SimplifiedEpisodeObject) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SimplifiedEpisodeObject) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SimplifiedEpisodeObject) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExplicit

`func (o *SimplifiedEpisodeObject) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *SimplifiedEpisodeObject) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *SimplifiedEpisodeObject) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *SimplifiedEpisodeObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SimplifiedEpisodeObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SimplifiedEpisodeObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *SimplifiedEpisodeObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedEpisodeObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedEpisodeObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *SimplifiedEpisodeObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedEpisodeObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedEpisodeObject) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *SimplifiedEpisodeObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SimplifiedEpisodeObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SimplifiedEpisodeObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsExternallyHosted

`func (o *SimplifiedEpisodeObject) GetIsExternallyHosted() bool`

GetIsExternallyHosted returns the IsExternallyHosted field if non-nil, zero value otherwise.

### GetIsExternallyHostedOk

`func (o *SimplifiedEpisodeObject) GetIsExternallyHostedOk() (*bool, bool)`

GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyHosted

`func (o *SimplifiedEpisodeObject) SetIsExternallyHosted(v bool)`

SetIsExternallyHosted sets IsExternallyHosted field to given value.


### GetIsPlayable

`func (o *SimplifiedEpisodeObject) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *SimplifiedEpisodeObject) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *SimplifiedEpisodeObject) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.


### GetLanguage

`func (o *SimplifiedEpisodeObject) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SimplifiedEpisodeObject) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SimplifiedEpisodeObject) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SimplifiedEpisodeObject) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLanguages

`func (o *SimplifiedEpisodeObject) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *SimplifiedEpisodeObject) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *SimplifiedEpisodeObject) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetName

`func (o *SimplifiedEpisodeObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedEpisodeObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedEpisodeObject) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *SimplifiedEpisodeObject) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SimplifiedEpisodeObject) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SimplifiedEpisodeObject) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *SimplifiedEpisodeObject) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *SimplifiedEpisodeObject) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *SimplifiedEpisodeObject) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetResumePoint

`func (o *SimplifiedEpisodeObject) GetResumePoint() ResumePointObject`

GetResumePoint returns the ResumePoint field if non-nil, zero value otherwise.

### GetResumePointOk

`func (o *SimplifiedEpisodeObject) GetResumePointOk() (*ResumePointObject, bool)`

GetResumePointOk returns a tuple with the ResumePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePoint

`func (o *SimplifiedEpisodeObject) SetResumePoint(v ResumePointObject)`

SetResumePoint sets ResumePoint field to given value.

### HasResumePoint

`func (o *SimplifiedEpisodeObject) HasResumePoint() bool`

HasResumePoint returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedEpisodeObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedEpisodeObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedEpisodeObject) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *SimplifiedEpisodeObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SimplifiedEpisodeObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SimplifiedEpisodeObject) SetUri(v string)`

SetUri sets Uri field to given value.


### GetRestrictions

`func (o *SimplifiedEpisodeObject) GetRestrictions() EpisodeRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *SimplifiedEpisodeObject) GetRestrictionsOk() (*EpisodeRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *SimplifiedEpisodeObject) SetRestrictions(v EpisodeRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *SimplifiedEpisodeObject) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


