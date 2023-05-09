# QueueObjectQueueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Album** | Pointer to [**TrackObjectAlbum**](TrackObjectAlbum.md) |  | [optional] 
**Artists** | Pointer to [**[]ArtistObject**](ArtistObject.md) | The artists who performed the track. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | [optional] 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**DiscNumber** | Pointer to **int32** | The disc number (usually &#x60;1&#x60; unless the album consists of more than one disc).  | [optional] 
**DurationMs** | **int32** | The episode length in milliseconds.  | 
**Explicit** | **bool** | Whether or not the episode has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalIds** | Pointer to [**TrackObjectExternalIds**](TrackObjectExternalIds.md) |  | [optional] 
**ExternalUrls** | [**EpisodeBaseExternalUrls**](EpisodeBaseExternalUrls.md) |  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the episode.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the episode.  | 
**IsPlayable** | **bool** | True if the episode is playable in the given market. Otherwise false.  | 
**LinkedFrom** | Pointer to [**TrackObjectLinkedFrom**](TrackObjectLinkedFrom.md) |  | [optional] 
**Restrictions** | Pointer to [**EpisodeBaseRestrictions**](EpisodeBaseRestrictions.md) |  | [optional] 
**Name** | **string** | The name of the episode.  | 
**Popularity** | Pointer to **int32** | The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.&lt;br/&gt;The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.&lt;br/&gt;Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._  | [optional] 
**PreviewUrl** | Pointer to **string** | A link to a 30 second preview (MP3 format) of the track. Can be &#x60;null&#x60;  | [optional] 
**TrackNumber** | Pointer to **int32** | The number of the track. If an album has several discs, the track number is the number on the specified disc.  | [optional] 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the episode.  | 
**IsLocal** | Pointer to **bool** | Whether or not the track is from a local file.  | [optional] 
**AudioPreviewUrl** | **string** | A URL to a 30 second preview (MP3 format) of the episode. &#x60;null&#x60; if not available.  | 
**Description** | **string** | A description of the episode. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the episode. This field may contain HTML tags.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the episode in various sizes, widest first.  | 
**IsExternallyHosted** | **bool** | True if the episode is hosted outside of Spotify&#39;s CDN.  | 
**Language** | Pointer to **string** | The language used in the episode, identified by a [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code. This field is deprecated and might be removed in the future. Please use the &#x60;languages&#x60; field instead.  | [optional] 
**Languages** | **[]string** | A list of the languages used in the episode, identified by their [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**ReleaseDate** | **string** | The date the episode was first released, for example &#x60;\&quot;1981-12-15\&quot;&#x60;. Depending on the precision, it might be shown as &#x60;\&quot;1981\&quot;&#x60; or &#x60;\&quot;1981-12\&quot;&#x60;.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**ResumePoint** | [**EpisodeBaseResumePoint**](EpisodeBaseResumePoint.md) |  | 
**Show** | [**EpisodeObjectAllOfShow**](EpisodeObjectAllOfShow.md) |  | 

## Methods

### NewQueueObjectQueueInner

`func NewQueueObjectQueueInner(durationMs int32, explicit bool, externalUrls EpisodeBaseExternalUrls, href string, id string, isPlayable bool, name string, type_ string, uri string, audioPreviewUrl string, description string, htmlDescription string, images []ImageObject, isExternallyHosted bool, languages []string, releaseDate string, releaseDatePrecision string, resumePoint EpisodeBaseResumePoint, show EpisodeObjectAllOfShow, ) *QueueObjectQueueInner`

NewQueueObjectQueueInner instantiates a new QueueObjectQueueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueObjectQueueInnerWithDefaults

`func NewQueueObjectQueueInnerWithDefaults() *QueueObjectQueueInner`

NewQueueObjectQueueInnerWithDefaults instantiates a new QueueObjectQueueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbum

`func (o *QueueObjectQueueInner) GetAlbum() TrackObjectAlbum`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *QueueObjectQueueInner) GetAlbumOk() (*TrackObjectAlbum, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *QueueObjectQueueInner) SetAlbum(v TrackObjectAlbum)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *QueueObjectQueueInner) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetArtists

`func (o *QueueObjectQueueInner) GetArtists() []ArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *QueueObjectQueueInner) GetArtistsOk() (*[]ArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *QueueObjectQueueInner) SetArtists(v []ArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *QueueObjectQueueInner) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetAvailableMarkets

`func (o *QueueObjectQueueInner) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *QueueObjectQueueInner) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *QueueObjectQueueInner) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *QueueObjectQueueInner) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetDiscNumber

`func (o *QueueObjectQueueInner) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *QueueObjectQueueInner) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *QueueObjectQueueInner) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *QueueObjectQueueInner) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDurationMs

`func (o *QueueObjectQueueInner) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *QueueObjectQueueInner) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *QueueObjectQueueInner) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExplicit

`func (o *QueueObjectQueueInner) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *QueueObjectQueueInner) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *QueueObjectQueueInner) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalIds

`func (o *QueueObjectQueueInner) GetExternalIds() TrackObjectExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *QueueObjectQueueInner) GetExternalIdsOk() (*TrackObjectExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *QueueObjectQueueInner) SetExternalIds(v TrackObjectExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *QueueObjectQueueInner) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetExternalUrls

`func (o *QueueObjectQueueInner) GetExternalUrls() EpisodeBaseExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *QueueObjectQueueInner) GetExternalUrlsOk() (*EpisodeBaseExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *QueueObjectQueueInner) SetExternalUrls(v EpisodeBaseExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *QueueObjectQueueInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *QueueObjectQueueInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *QueueObjectQueueInner) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *QueueObjectQueueInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueObjectQueueInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueObjectQueueInner) SetId(v string)`

SetId sets Id field to given value.


### GetIsPlayable

`func (o *QueueObjectQueueInner) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *QueueObjectQueueInner) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *QueueObjectQueueInner) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.


### GetLinkedFrom

`func (o *QueueObjectQueueInner) GetLinkedFrom() TrackObjectLinkedFrom`

GetLinkedFrom returns the LinkedFrom field if non-nil, zero value otherwise.

### GetLinkedFromOk

`func (o *QueueObjectQueueInner) GetLinkedFromOk() (*TrackObjectLinkedFrom, bool)`

GetLinkedFromOk returns a tuple with the LinkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFrom

`func (o *QueueObjectQueueInner) SetLinkedFrom(v TrackObjectLinkedFrom)`

SetLinkedFrom sets LinkedFrom field to given value.

### HasLinkedFrom

`func (o *QueueObjectQueueInner) HasLinkedFrom() bool`

HasLinkedFrom returns a boolean if a field has been set.

### GetRestrictions

`func (o *QueueObjectQueueInner) GetRestrictions() EpisodeBaseRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *QueueObjectQueueInner) GetRestrictionsOk() (*EpisodeBaseRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *QueueObjectQueueInner) SetRestrictions(v EpisodeBaseRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *QueueObjectQueueInner) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetName

`func (o *QueueObjectQueueInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueueObjectQueueInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueueObjectQueueInner) SetName(v string)`

SetName sets Name field to given value.


### GetPopularity

`func (o *QueueObjectQueueInner) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *QueueObjectQueueInner) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *QueueObjectQueueInner) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *QueueObjectQueueInner) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *QueueObjectQueueInner) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *QueueObjectQueueInner) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *QueueObjectQueueInner) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *QueueObjectQueueInner) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### GetTrackNumber

`func (o *QueueObjectQueueInner) GetTrackNumber() int32`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *QueueObjectQueueInner) GetTrackNumberOk() (*int32, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *QueueObjectQueueInner) SetTrackNumber(v int32)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *QueueObjectQueueInner) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### GetType

`func (o *QueueObjectQueueInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueueObjectQueueInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueueObjectQueueInner) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *QueueObjectQueueInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *QueueObjectQueueInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *QueueObjectQueueInner) SetUri(v string)`

SetUri sets Uri field to given value.


### GetIsLocal

`func (o *QueueObjectQueueInner) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *QueueObjectQueueInner) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *QueueObjectQueueInner) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *QueueObjectQueueInner) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetAudioPreviewUrl

`func (o *QueueObjectQueueInner) GetAudioPreviewUrl() string`

GetAudioPreviewUrl returns the AudioPreviewUrl field if non-nil, zero value otherwise.

### GetAudioPreviewUrlOk

`func (o *QueueObjectQueueInner) GetAudioPreviewUrlOk() (*string, bool)`

GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioPreviewUrl

`func (o *QueueObjectQueueInner) SetAudioPreviewUrl(v string)`

SetAudioPreviewUrl sets AudioPreviewUrl field to given value.


### GetDescription

`func (o *QueueObjectQueueInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueueObjectQueueInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueueObjectQueueInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *QueueObjectQueueInner) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *QueueObjectQueueInner) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *QueueObjectQueueInner) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetImages

`func (o *QueueObjectQueueInner) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *QueueObjectQueueInner) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *QueueObjectQueueInner) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsExternallyHosted

`func (o *QueueObjectQueueInner) GetIsExternallyHosted() bool`

GetIsExternallyHosted returns the IsExternallyHosted field if non-nil, zero value otherwise.

### GetIsExternallyHostedOk

`func (o *QueueObjectQueueInner) GetIsExternallyHostedOk() (*bool, bool)`

GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyHosted

`func (o *QueueObjectQueueInner) SetIsExternallyHosted(v bool)`

SetIsExternallyHosted sets IsExternallyHosted field to given value.


### GetLanguage

`func (o *QueueObjectQueueInner) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *QueueObjectQueueInner) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *QueueObjectQueueInner) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *QueueObjectQueueInner) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLanguages

`func (o *QueueObjectQueueInner) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *QueueObjectQueueInner) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *QueueObjectQueueInner) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetReleaseDate

`func (o *QueueObjectQueueInner) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *QueueObjectQueueInner) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *QueueObjectQueueInner) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *QueueObjectQueueInner) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *QueueObjectQueueInner) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *QueueObjectQueueInner) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetResumePoint

`func (o *QueueObjectQueueInner) GetResumePoint() EpisodeBaseResumePoint`

GetResumePoint returns the ResumePoint field if non-nil, zero value otherwise.

### GetResumePointOk

`func (o *QueueObjectQueueInner) GetResumePointOk() (*EpisodeBaseResumePoint, bool)`

GetResumePointOk returns a tuple with the ResumePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePoint

`func (o *QueueObjectQueueInner) SetResumePoint(v EpisodeBaseResumePoint)`

SetResumePoint sets ResumePoint field to given value.


### GetShow

`func (o *QueueObjectQueueInner) GetShow() EpisodeObjectAllOfShow`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *QueueObjectQueueInner) GetShowOk() (*EpisodeObjectAllOfShow, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *QueueObjectQueueInner) SetShow(v EpisodeObjectAllOfShow)`

SetShow sets Show field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


