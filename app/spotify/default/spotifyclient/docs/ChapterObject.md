# ChapterObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioPreviewUrl** | **string** | A URL to a 30 second preview (MP3 format) of the episode. &#x60;null&#x60; if not available.  | 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the chapter can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**ChapterNumber** | **int32** | The number of the chapter  | 
**Description** | **string** | A description of the episode. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the episode. This field may contain HTML tags.  | 
**DurationMs** | **int32** | The episode length in milliseconds.  | 
**Explicit** | **bool** | Whether or not the episode has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalUrls** | [**EpisodeBaseExternalUrls**](EpisodeBaseExternalUrls.md) |  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the episode.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the episode.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the episode in various sizes, widest first.  | 
**IsPlayable** | **bool** | True if the episode is playable in the given market. Otherwise false.  | 
**Languages** | **[]string** | A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**Name** | **string** | The name of the episode.  | 
**ReleaseDate** | **string** | The date the episode was first released, for example &#x60;\&quot;1981-12-15\&quot;&#x60;. Depending on the precision, it might be shown as &#x60;\&quot;1981\&quot;&#x60; or &#x60;\&quot;1981-12\&quot;&#x60;.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**ResumePoint** | [**EpisodeBaseResumePoint**](EpisodeBaseResumePoint.md) |  | 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the episode.  | 
**Restrictions** | Pointer to [**ChapterBaseRestrictions**](ChapterBaseRestrictions.md) |  | [optional] 
**Audiobook** | [**ChapterObjectAllOfAudiobook**](ChapterObjectAllOfAudiobook.md) |  | 

## Methods

### NewChapterObject

`func NewChapterObject(audioPreviewUrl string, chapterNumber int32, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls EpisodeBaseExternalUrls, href string, id string, images []ImageObject, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, resumePoint EpisodeBaseResumePoint, type_ string, uri string, audiobook ChapterObjectAllOfAudiobook, ) *ChapterObject`

NewChapterObject instantiates a new ChapterObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChapterObjectWithDefaults

`func NewChapterObjectWithDefaults() *ChapterObject`

NewChapterObjectWithDefaults instantiates a new ChapterObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioPreviewUrl

`func (o *ChapterObject) GetAudioPreviewUrl() string`

GetAudioPreviewUrl returns the AudioPreviewUrl field if non-nil, zero value otherwise.

### GetAudioPreviewUrlOk

`func (o *ChapterObject) GetAudioPreviewUrlOk() (*string, bool)`

GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioPreviewUrl

`func (o *ChapterObject) SetAudioPreviewUrl(v string)`

SetAudioPreviewUrl sets AudioPreviewUrl field to given value.


### GetAvailableMarkets

`func (o *ChapterObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *ChapterObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *ChapterObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *ChapterObject) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetChapterNumber

`func (o *ChapterObject) GetChapterNumber() int32`

GetChapterNumber returns the ChapterNumber field if non-nil, zero value otherwise.

### GetChapterNumberOk

`func (o *ChapterObject) GetChapterNumberOk() (*int32, bool)`

GetChapterNumberOk returns a tuple with the ChapterNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapterNumber

`func (o *ChapterObject) SetChapterNumber(v int32)`

SetChapterNumber sets ChapterNumber field to given value.


### GetDescription

`func (o *ChapterObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChapterObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChapterObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *ChapterObject) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *ChapterObject) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *ChapterObject) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetDurationMs

`func (o *ChapterObject) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ChapterObject) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ChapterObject) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExplicit

`func (o *ChapterObject) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *ChapterObject) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *ChapterObject) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *ChapterObject) GetExternalUrls() EpisodeBaseExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *ChapterObject) GetExternalUrlsOk() (*EpisodeBaseExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *ChapterObject) SetExternalUrls(v EpisodeBaseExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *ChapterObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ChapterObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ChapterObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ChapterObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChapterObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChapterObject) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *ChapterObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ChapterObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ChapterObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsPlayable

`func (o *ChapterObject) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *ChapterObject) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *ChapterObject) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.


### GetLanguages

`func (o *ChapterObject) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ChapterObject) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ChapterObject) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetName

`func (o *ChapterObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChapterObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChapterObject) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *ChapterObject) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ChapterObject) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ChapterObject) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *ChapterObject) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *ChapterObject) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *ChapterObject) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetResumePoint

`func (o *ChapterObject) GetResumePoint() EpisodeBaseResumePoint`

GetResumePoint returns the ResumePoint field if non-nil, zero value otherwise.

### GetResumePointOk

`func (o *ChapterObject) GetResumePointOk() (*EpisodeBaseResumePoint, bool)`

GetResumePointOk returns a tuple with the ResumePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePoint

`func (o *ChapterObject) SetResumePoint(v EpisodeBaseResumePoint)`

SetResumePoint sets ResumePoint field to given value.


### GetType

`func (o *ChapterObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChapterObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChapterObject) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *ChapterObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ChapterObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ChapterObject) SetUri(v string)`

SetUri sets Uri field to given value.


### GetRestrictions

`func (o *ChapterObject) GetRestrictions() ChapterBaseRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *ChapterObject) GetRestrictionsOk() (*ChapterBaseRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *ChapterObject) SetRestrictions(v ChapterBaseRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *ChapterObject) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetAudiobook

`func (o *ChapterObject) GetAudiobook() ChapterObjectAllOfAudiobook`

GetAudiobook returns the Audiobook field if non-nil, zero value otherwise.

### GetAudiobookOk

`func (o *ChapterObject) GetAudiobookOk() (*ChapterObjectAllOfAudiobook, bool)`

GetAudiobookOk returns a tuple with the Audiobook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiobook

`func (o *ChapterObject) SetAudiobook(v ChapterObjectAllOfAudiobook)`

SetAudiobook sets Audiobook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


