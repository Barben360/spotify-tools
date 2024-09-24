# SimplifiedChapterObject

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

### NewSimplifiedChapterObject

`func NewSimplifiedChapterObject(audioPreviewUrl NullableString, chapterNumber int32, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, ) *SimplifiedChapterObject`

NewSimplifiedChapterObject instantiates a new SimplifiedChapterObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedChapterObjectWithDefaults

`func NewSimplifiedChapterObjectWithDefaults() *SimplifiedChapterObject`

NewSimplifiedChapterObjectWithDefaults instantiates a new SimplifiedChapterObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioPreviewUrl

`func (o *SimplifiedChapterObject) GetAudioPreviewUrl() string`

GetAudioPreviewUrl returns the AudioPreviewUrl field if non-nil, zero value otherwise.

### GetAudioPreviewUrlOk

`func (o *SimplifiedChapterObject) GetAudioPreviewUrlOk() (*string, bool)`

GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioPreviewUrl

`func (o *SimplifiedChapterObject) SetAudioPreviewUrl(v string)`

SetAudioPreviewUrl sets AudioPreviewUrl field to given value.


### SetAudioPreviewUrlNil

`func (o *SimplifiedChapterObject) SetAudioPreviewUrlNil(b bool)`

 SetAudioPreviewUrlNil sets the value for AudioPreviewUrl to be an explicit nil

### UnsetAudioPreviewUrl
`func (o *SimplifiedChapterObject) UnsetAudioPreviewUrl()`

UnsetAudioPreviewUrl ensures that no value is present for AudioPreviewUrl, not even an explicit nil
### GetAvailableMarkets

`func (o *SimplifiedChapterObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *SimplifiedChapterObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *SimplifiedChapterObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *SimplifiedChapterObject) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetChapterNumber

`func (o *SimplifiedChapterObject) GetChapterNumber() int32`

GetChapterNumber returns the ChapterNumber field if non-nil, zero value otherwise.

### GetChapterNumberOk

`func (o *SimplifiedChapterObject) GetChapterNumberOk() (*int32, bool)`

GetChapterNumberOk returns a tuple with the ChapterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapterNumber

`func (o *SimplifiedChapterObject) SetChapterNumber(v int32)`

SetChapterNumber sets ChapterNumber field to given value.


### GetDescription

`func (o *SimplifiedChapterObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimplifiedChapterObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimplifiedChapterObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *SimplifiedChapterObject) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *SimplifiedChapterObject) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *SimplifiedChapterObject) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetDurationMs

`func (o *SimplifiedChapterObject) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SimplifiedChapterObject) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SimplifiedChapterObject) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExplicit

`func (o *SimplifiedChapterObject) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *SimplifiedChapterObject) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *SimplifiedChapterObject) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *SimplifiedChapterObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SimplifiedChapterObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SimplifiedChapterObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *SimplifiedChapterObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedChapterObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedChapterObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *SimplifiedChapterObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedChapterObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedChapterObject) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *SimplifiedChapterObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SimplifiedChapterObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SimplifiedChapterObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsPlayable

`func (o *SimplifiedChapterObject) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *SimplifiedChapterObject) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *SimplifiedChapterObject) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.


### GetLanguages

`func (o *SimplifiedChapterObject) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *SimplifiedChapterObject) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *SimplifiedChapterObject) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetName

`func (o *SimplifiedChapterObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedChapterObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedChapterObject) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *SimplifiedChapterObject) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SimplifiedChapterObject) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SimplifiedChapterObject) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *SimplifiedChapterObject) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *SimplifiedChapterObject) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *SimplifiedChapterObject) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetResumePoint

`func (o *SimplifiedChapterObject) GetResumePoint() ResumePointObject`

GetResumePoint returns the ResumePoint field if non-nil, zero value otherwise.

### GetResumePointOk

`func (o *SimplifiedChapterObject) GetResumePointOk() (*ResumePointObject, bool)`

GetResumePointOk returns a tuple with the ResumePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePoint

`func (o *SimplifiedChapterObject) SetResumePoint(v ResumePointObject)`

SetResumePoint sets ResumePoint field to given value.

### HasResumePoint

`func (o *SimplifiedChapterObject) HasResumePoint() bool`

HasResumePoint returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedChapterObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedChapterObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedChapterObject) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *SimplifiedChapterObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SimplifiedChapterObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SimplifiedChapterObject) SetUri(v string)`

SetUri sets Uri field to given value.


### GetRestrictions

`func (o *SimplifiedChapterObject) GetRestrictions() ChapterRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *SimplifiedChapterObject) GetRestrictionsOk() (*ChapterRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *SimplifiedChapterObject) SetRestrictions(v ChapterRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *SimplifiedChapterObject) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


