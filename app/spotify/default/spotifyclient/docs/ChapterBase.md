# ChapterBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioPreviewUrl** | **NullableString** | A URL to a 30 second preview (MP3 format) of the chapter. &#x60;null&#x60; if not available.  | 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the chapter can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**ChapterNumber** | **int32** | The number of the chapter  | 
**Description** | **string** | A description of the chapter. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the chapter. This field may contain HTML tags.  | 
**DurationMs** | **int32** | The chapter length in milliseconds.  | 
**Explicit** | **bool** | Whether or not the chapter has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalUrls** | [**ExternalUrlObject**](ExternalUrlObject.md) | External URLs for this chapter.  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the chapter.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the chapter.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the chapter in various sizes, widest first.  | 
**IsPlayable** | **bool** | True if the chapter is playable in the given market. Otherwise false.  | 
**Languages** | **[]string** | A list of the languages used in the chapter, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**Name** | **string** | The name of the chapter.  | 
**ReleaseDate** | **string** | The date the chapter was first released, for example &#x60;\&quot;1981-12-15\&quot;&#x60;. Depending on the precision, it might be shown as &#x60;\&quot;1981\&quot;&#x60; or &#x60;\&quot;1981-12\&quot;&#x60;.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**ResumePoint** | Pointer to [**ResumePointObject**](ResumePointObject.md) | The user&#39;s most recent position in the chapter. Set if the supplied access token is a user token and has the scope &#39;user-read-playback-position&#39;.  | [optional] 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the chapter.  | 
**Restrictions** | Pointer to [**ChapterRestrictionObject**](ChapterRestrictionObject.md) | Included in the response when a content restriction is applied.  | [optional] 

## Methods

### NewChapterBase

`func NewChapterBase(audioPreviewUrl NullableString, chapterNumber int32, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, ) *ChapterBase`

NewChapterBase instantiates a new ChapterBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChapterBaseWithDefaults

`func NewChapterBaseWithDefaults() *ChapterBase`

NewChapterBaseWithDefaults instantiates a new ChapterBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioPreviewUrl

`func (o *ChapterBase) GetAudioPreviewUrl() string`

GetAudioPreviewUrl returns the AudioPreviewUrl field if non-nil, zero value otherwise.

### GetAudioPreviewUrlOk

`func (o *ChapterBase) GetAudioPreviewUrlOk() (*string, bool)`

GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioPreviewUrl

`func (o *ChapterBase) SetAudioPreviewUrl(v string)`

SetAudioPreviewUrl sets AudioPreviewUrl field to given value.


### SetAudioPreviewUrlNil

`func (o *ChapterBase) SetAudioPreviewUrlNil(b bool)`

 SetAudioPreviewUrlNil sets the value for AudioPreviewUrl to be an explicit nil

### UnsetAudioPreviewUrl
`func (o *ChapterBase) UnsetAudioPreviewUrl()`

UnsetAudioPreviewUrl ensures that no value is present for AudioPreviewUrl, not even an explicit nil
### GetAvailableMarkets

`func (o *ChapterBase) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *ChapterBase) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *ChapterBase) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *ChapterBase) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetChapterNumber

`func (o *ChapterBase) GetChapterNumber() int32`

GetChapterNumber returns the ChapterNumber field if non-nil, zero value otherwise.

### GetChapterNumberOk

`func (o *ChapterBase) GetChapterNumberOk() (*int32, bool)`

GetChapterNumberOk returns a tuple with the ChapterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapterNumber

`func (o *ChapterBase) SetChapterNumber(v int32)`

SetChapterNumber sets ChapterNumber field to given value.


### GetDescription

`func (o *ChapterBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChapterBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChapterBase) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *ChapterBase) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *ChapterBase) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *ChapterBase) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetDurationMs

`func (o *ChapterBase) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ChapterBase) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ChapterBase) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExplicit

`func (o *ChapterBase) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *ChapterBase) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *ChapterBase) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *ChapterBase) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *ChapterBase) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *ChapterBase) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *ChapterBase) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ChapterBase) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ChapterBase) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ChapterBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChapterBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChapterBase) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *ChapterBase) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ChapterBase) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ChapterBase) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsPlayable

`func (o *ChapterBase) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *ChapterBase) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *ChapterBase) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.


### GetLanguages

`func (o *ChapterBase) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ChapterBase) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ChapterBase) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetName

`func (o *ChapterBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChapterBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChapterBase) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *ChapterBase) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ChapterBase) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ChapterBase) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *ChapterBase) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *ChapterBase) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *ChapterBase) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetResumePoint

`func (o *ChapterBase) GetResumePoint() ResumePointObject`

GetResumePoint returns the ResumePoint field if non-nil, zero value otherwise.

### GetResumePointOk

`func (o *ChapterBase) GetResumePointOk() (*ResumePointObject, bool)`

GetResumePointOk returns a tuple with the ResumePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePoint

`func (o *ChapterBase) SetResumePoint(v ResumePointObject)`

SetResumePoint sets ResumePoint field to given value.

### HasResumePoint

`func (o *ChapterBase) HasResumePoint() bool`

HasResumePoint returns a boolean if a field has been set.

### GetType

`func (o *ChapterBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChapterBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChapterBase) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *ChapterBase) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ChapterBase) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ChapterBase) SetUri(v string)`

SetUri sets Uri field to given value.


### GetRestrictions

`func (o *ChapterBase) GetRestrictions() ChapterRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *ChapterBase) GetRestrictionsOk() (*ChapterRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *ChapterBase) SetRestrictions(v ChapterRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *ChapterBase) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


