# SavedEpisodeObjectEpisode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Episode** | Pointer to **bool** | Fix by Barben360  | [optional] 
**Track** | Pointer to **bool** | Fix by Barben360  | [optional] 
**AudioPreviewUrl** | **string** | A URL to a 30 second preview (MP3 format) of the episode. &#x60;null&#x60; if not available.  | 
**Description** | **string** | A description of the episode. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the episode. This field may contain HTML tags.  | 
**DurationMs** | **int32** | The episode length in milliseconds.  | 
**Explicit** | **bool** | Whether or not the episode has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalUrls** | [**EpisodeBaseExternalUrls**](EpisodeBaseExternalUrls.md) |  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the episode.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the episode.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the episode in various sizes, widest first.  | 
**IsExternallyHosted** | **bool** | True if the episode is hosted outside of Spotify&#39;s CDN.  | 
**IsPlayable** | **bool** | True if the episode is playable in the given market. Otherwise false.  | 
**Language** | Pointer to **string** | The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the &#x60;languages&#x60; field instead.  | [optional] 
**Languages** | **[]string** | A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**Name** | **string** | The name of the episode.  | 
**ReleaseDate** | **string** | The date the episode was first released, for example &#x60;\&quot;1981-12-15\&quot;&#x60;. Depending on the precision, it might be shown as &#x60;\&quot;1981\&quot;&#x60; or &#x60;\&quot;1981-12\&quot;&#x60;.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**ResumePoint** | [**EpisodeBaseResumePoint**](EpisodeBaseResumePoint.md) |  | 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the episode.  | 
**Restrictions** | Pointer to [**EpisodeBaseRestrictions**](EpisodeBaseRestrictions.md) |  | [optional] 
**Show** | [**EpisodeObjectAllOfShow**](EpisodeObjectAllOfShow.md) |  | 

## Methods

### NewSavedEpisodeObjectEpisode

`func NewSavedEpisodeObjectEpisode(audioPreviewUrl string, description string, htmlDescription string, durationMs int32, explicit bool, externalUrls EpisodeBaseExternalUrls, href string, id string, images []ImageObject, isExternallyHosted bool, isPlayable bool, languages []string, name string, releaseDate string, releaseDatePrecision string, resumePoint EpisodeBaseResumePoint, type_ string, uri string, show EpisodeObjectAllOfShow, ) *SavedEpisodeObjectEpisode`

NewSavedEpisodeObjectEpisode instantiates a new SavedEpisodeObjectEpisode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedEpisodeObjectEpisodeWithDefaults

`func NewSavedEpisodeObjectEpisodeWithDefaults() *SavedEpisodeObjectEpisode`

NewSavedEpisodeObjectEpisodeWithDefaults instantiates a new SavedEpisodeObjectEpisode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpisode

`func (o *SavedEpisodeObjectEpisode) GetEpisode() bool`

GetEpisode returns the Episode field if non-nil, zero value otherwise.

### GetEpisodeOk

`func (o *SavedEpisodeObjectEpisode) GetEpisodeOk() (*bool, bool)`

GetEpisodeOk returns a tuple with the Episode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisode

`func (o *SavedEpisodeObjectEpisode) SetEpisode(v bool)`

SetEpisode sets Episode field to given value.

### HasEpisode

`func (o *SavedEpisodeObjectEpisode) HasEpisode() bool`

HasEpisode returns a boolean if a field has been set.

### GetTrack

`func (o *SavedEpisodeObjectEpisode) GetTrack() bool`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *SavedEpisodeObjectEpisode) GetTrackOk() (*bool, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *SavedEpisodeObjectEpisode) SetTrack(v bool)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *SavedEpisodeObjectEpisode) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetAudioPreviewUrl

`func (o *SavedEpisodeObjectEpisode) GetAudioPreviewUrl() string`

GetAudioPreviewUrl returns the AudioPreviewUrl field if non-nil, zero value otherwise.

### GetAudioPreviewUrlOk

`func (o *SavedEpisodeObjectEpisode) GetAudioPreviewUrlOk() (*string, bool)`

GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioPreviewUrl

`func (o *SavedEpisodeObjectEpisode) SetAudioPreviewUrl(v string)`

SetAudioPreviewUrl sets AudioPreviewUrl field to given value.


### GetDescription

`func (o *SavedEpisodeObjectEpisode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedEpisodeObjectEpisode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedEpisodeObjectEpisode) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *SavedEpisodeObjectEpisode) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *SavedEpisodeObjectEpisode) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *SavedEpisodeObjectEpisode) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetDurationMs

`func (o *SavedEpisodeObjectEpisode) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SavedEpisodeObjectEpisode) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SavedEpisodeObjectEpisode) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExplicit

`func (o *SavedEpisodeObjectEpisode) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *SavedEpisodeObjectEpisode) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *SavedEpisodeObjectEpisode) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *SavedEpisodeObjectEpisode) GetExternalUrls() EpisodeBaseExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SavedEpisodeObjectEpisode) GetExternalUrlsOk() (*EpisodeBaseExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SavedEpisodeObjectEpisode) SetExternalUrls(v EpisodeBaseExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *SavedEpisodeObjectEpisode) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SavedEpisodeObjectEpisode) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SavedEpisodeObjectEpisode) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *SavedEpisodeObjectEpisode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedEpisodeObjectEpisode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedEpisodeObjectEpisode) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *SavedEpisodeObjectEpisode) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SavedEpisodeObjectEpisode) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SavedEpisodeObjectEpisode) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsExternallyHosted

`func (o *SavedEpisodeObjectEpisode) GetIsExternallyHosted() bool`

GetIsExternallyHosted returns the IsExternallyHosted field if non-nil, zero value otherwise.

### GetIsExternallyHostedOk

`func (o *SavedEpisodeObjectEpisode) GetIsExternallyHostedOk() (*bool, bool)`

GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyHosted

`func (o *SavedEpisodeObjectEpisode) SetIsExternallyHosted(v bool)`

SetIsExternallyHosted sets IsExternallyHosted field to given value.


### GetIsPlayable

`func (o *SavedEpisodeObjectEpisode) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *SavedEpisodeObjectEpisode) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *SavedEpisodeObjectEpisode) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.


### GetLanguage

`func (o *SavedEpisodeObjectEpisode) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SavedEpisodeObjectEpisode) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SavedEpisodeObjectEpisode) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SavedEpisodeObjectEpisode) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLanguages

`func (o *SavedEpisodeObjectEpisode) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *SavedEpisodeObjectEpisode) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *SavedEpisodeObjectEpisode) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetName

`func (o *SavedEpisodeObjectEpisode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedEpisodeObjectEpisode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedEpisodeObjectEpisode) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *SavedEpisodeObjectEpisode) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SavedEpisodeObjectEpisode) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SavedEpisodeObjectEpisode) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *SavedEpisodeObjectEpisode) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *SavedEpisodeObjectEpisode) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *SavedEpisodeObjectEpisode) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetResumePoint

`func (o *SavedEpisodeObjectEpisode) GetResumePoint() EpisodeBaseResumePoint`

GetResumePoint returns the ResumePoint field if non-nil, zero value otherwise.

### GetResumePointOk

`func (o *SavedEpisodeObjectEpisode) GetResumePointOk() (*EpisodeBaseResumePoint, bool)`

GetResumePointOk returns a tuple with the ResumePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePoint

`func (o *SavedEpisodeObjectEpisode) SetResumePoint(v EpisodeBaseResumePoint)`

SetResumePoint sets ResumePoint field to given value.


### GetType

`func (o *SavedEpisodeObjectEpisode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedEpisodeObjectEpisode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedEpisodeObjectEpisode) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *SavedEpisodeObjectEpisode) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SavedEpisodeObjectEpisode) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SavedEpisodeObjectEpisode) SetUri(v string)`

SetUri sets Uri field to given value.


### GetRestrictions

`func (o *SavedEpisodeObjectEpisode) GetRestrictions() EpisodeBaseRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *SavedEpisodeObjectEpisode) GetRestrictionsOk() (*EpisodeBaseRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *SavedEpisodeObjectEpisode) SetRestrictions(v EpisodeBaseRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *SavedEpisodeObjectEpisode) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetShow

`func (o *SavedEpisodeObjectEpisode) GetShow() EpisodeObjectAllOfShow`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *SavedEpisodeObjectEpisode) GetShowOk() (*EpisodeObjectAllOfShow, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *SavedEpisodeObjectEpisode) SetShow(v EpisodeObjectAllOfShow)`

SetShow sets Show field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


