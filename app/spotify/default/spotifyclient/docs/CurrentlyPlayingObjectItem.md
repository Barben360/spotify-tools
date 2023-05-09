# CurrentlyPlayingObjectItem

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

### NewCurrentlyPlayingObjectItem

`func NewCurrentlyPlayingObjectItem(durationMs int32, explicit bool, externalUrls EpisodeBaseExternalUrls, href string, id string, isPlayable bool, name string, type_ string, uri string, audioPreviewUrl string, description string, htmlDescription string, images []ImageObject, isExternallyHosted bool, languages []string, releaseDate string, releaseDatePrecision string, resumePoint EpisodeBaseResumePoint, show EpisodeObjectAllOfShow, ) *CurrentlyPlayingObjectItem`

NewCurrentlyPlayingObjectItem instantiates a new CurrentlyPlayingObjectItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentlyPlayingObjectItemWithDefaults

`func NewCurrentlyPlayingObjectItemWithDefaults() *CurrentlyPlayingObjectItem`

NewCurrentlyPlayingObjectItemWithDefaults instantiates a new CurrentlyPlayingObjectItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbum

`func (o *CurrentlyPlayingObjectItem) GetAlbum() TrackObjectAlbum`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *CurrentlyPlayingObjectItem) GetAlbumOk() (*TrackObjectAlbum, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *CurrentlyPlayingObjectItem) SetAlbum(v TrackObjectAlbum)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *CurrentlyPlayingObjectItem) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetArtists

`func (o *CurrentlyPlayingObjectItem) GetArtists() []ArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *CurrentlyPlayingObjectItem) GetArtistsOk() (*[]ArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *CurrentlyPlayingObjectItem) SetArtists(v []ArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *CurrentlyPlayingObjectItem) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetAvailableMarkets

`func (o *CurrentlyPlayingObjectItem) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *CurrentlyPlayingObjectItem) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *CurrentlyPlayingObjectItem) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *CurrentlyPlayingObjectItem) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetDiscNumber

`func (o *CurrentlyPlayingObjectItem) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *CurrentlyPlayingObjectItem) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *CurrentlyPlayingObjectItem) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *CurrentlyPlayingObjectItem) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDurationMs

`func (o *CurrentlyPlayingObjectItem) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *CurrentlyPlayingObjectItem) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *CurrentlyPlayingObjectItem) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### GetExplicit

`func (o *CurrentlyPlayingObjectItem) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *CurrentlyPlayingObjectItem) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *CurrentlyPlayingObjectItem) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalIds

`func (o *CurrentlyPlayingObjectItem) GetExternalIds() TrackObjectExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *CurrentlyPlayingObjectItem) GetExternalIdsOk() (*TrackObjectExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *CurrentlyPlayingObjectItem) SetExternalIds(v TrackObjectExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *CurrentlyPlayingObjectItem) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetExternalUrls

`func (o *CurrentlyPlayingObjectItem) GetExternalUrls() EpisodeBaseExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *CurrentlyPlayingObjectItem) GetExternalUrlsOk() (*EpisodeBaseExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *CurrentlyPlayingObjectItem) SetExternalUrls(v EpisodeBaseExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *CurrentlyPlayingObjectItem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CurrentlyPlayingObjectItem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CurrentlyPlayingObjectItem) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *CurrentlyPlayingObjectItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentlyPlayingObjectItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentlyPlayingObjectItem) SetId(v string)`

SetId sets Id field to given value.


### GetIsPlayable

`func (o *CurrentlyPlayingObjectItem) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *CurrentlyPlayingObjectItem) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *CurrentlyPlayingObjectItem) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.


### GetLinkedFrom

`func (o *CurrentlyPlayingObjectItem) GetLinkedFrom() TrackObjectLinkedFrom`

GetLinkedFrom returns the LinkedFrom field if non-nil, zero value otherwise.

### GetLinkedFromOk

`func (o *CurrentlyPlayingObjectItem) GetLinkedFromOk() (*TrackObjectLinkedFrom, bool)`

GetLinkedFromOk returns a tuple with the LinkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFrom

`func (o *CurrentlyPlayingObjectItem) SetLinkedFrom(v TrackObjectLinkedFrom)`

SetLinkedFrom sets LinkedFrom field to given value.

### HasLinkedFrom

`func (o *CurrentlyPlayingObjectItem) HasLinkedFrom() bool`

HasLinkedFrom returns a boolean if a field has been set.

### GetRestrictions

`func (o *CurrentlyPlayingObjectItem) GetRestrictions() EpisodeBaseRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *CurrentlyPlayingObjectItem) GetRestrictionsOk() (*EpisodeBaseRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *CurrentlyPlayingObjectItem) SetRestrictions(v EpisodeBaseRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *CurrentlyPlayingObjectItem) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetName

`func (o *CurrentlyPlayingObjectItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentlyPlayingObjectItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentlyPlayingObjectItem) SetName(v string)`

SetName sets Name field to given value.


### GetPopularity

`func (o *CurrentlyPlayingObjectItem) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *CurrentlyPlayingObjectItem) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *CurrentlyPlayingObjectItem) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *CurrentlyPlayingObjectItem) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *CurrentlyPlayingObjectItem) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *CurrentlyPlayingObjectItem) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *CurrentlyPlayingObjectItem) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *CurrentlyPlayingObjectItem) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### GetTrackNumber

`func (o *CurrentlyPlayingObjectItem) GetTrackNumber() int32`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *CurrentlyPlayingObjectItem) GetTrackNumberOk() (*int32, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *CurrentlyPlayingObjectItem) SetTrackNumber(v int32)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *CurrentlyPlayingObjectItem) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### GetType

`func (o *CurrentlyPlayingObjectItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CurrentlyPlayingObjectItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CurrentlyPlayingObjectItem) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *CurrentlyPlayingObjectItem) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CurrentlyPlayingObjectItem) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CurrentlyPlayingObjectItem) SetUri(v string)`

SetUri sets Uri field to given value.


### GetIsLocal

`func (o *CurrentlyPlayingObjectItem) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *CurrentlyPlayingObjectItem) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *CurrentlyPlayingObjectItem) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *CurrentlyPlayingObjectItem) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetAudioPreviewUrl

`func (o *CurrentlyPlayingObjectItem) GetAudioPreviewUrl() string`

GetAudioPreviewUrl returns the AudioPreviewUrl field if non-nil, zero value otherwise.

### GetAudioPreviewUrlOk

`func (o *CurrentlyPlayingObjectItem) GetAudioPreviewUrlOk() (*string, bool)`

GetAudioPreviewUrlOk returns a tuple with the AudioPreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioPreviewUrl

`func (o *CurrentlyPlayingObjectItem) SetAudioPreviewUrl(v string)`

SetAudioPreviewUrl sets AudioPreviewUrl field to given value.


### GetDescription

`func (o *CurrentlyPlayingObjectItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurrentlyPlayingObjectItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurrentlyPlayingObjectItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *CurrentlyPlayingObjectItem) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *CurrentlyPlayingObjectItem) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *CurrentlyPlayingObjectItem) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetImages

`func (o *CurrentlyPlayingObjectItem) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CurrentlyPlayingObjectItem) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CurrentlyPlayingObjectItem) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsExternallyHosted

`func (o *CurrentlyPlayingObjectItem) GetIsExternallyHosted() bool`

GetIsExternallyHosted returns the IsExternallyHosted field if non-nil, zero value otherwise.

### GetIsExternallyHostedOk

`func (o *CurrentlyPlayingObjectItem) GetIsExternallyHostedOk() (*bool, bool)`

GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyHosted

`func (o *CurrentlyPlayingObjectItem) SetIsExternallyHosted(v bool)`

SetIsExternallyHosted sets IsExternallyHosted field to given value.


### GetLanguage

`func (o *CurrentlyPlayingObjectItem) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CurrentlyPlayingObjectItem) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CurrentlyPlayingObjectItem) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CurrentlyPlayingObjectItem) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLanguages

`func (o *CurrentlyPlayingObjectItem) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CurrentlyPlayingObjectItem) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CurrentlyPlayingObjectItem) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetReleaseDate

`func (o *CurrentlyPlayingObjectItem) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *CurrentlyPlayingObjectItem) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *CurrentlyPlayingObjectItem) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *CurrentlyPlayingObjectItem) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *CurrentlyPlayingObjectItem) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *CurrentlyPlayingObjectItem) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetResumePoint

`func (o *CurrentlyPlayingObjectItem) GetResumePoint() EpisodeBaseResumePoint`

GetResumePoint returns the ResumePoint field if non-nil, zero value otherwise.

### GetResumePointOk

`func (o *CurrentlyPlayingObjectItem) GetResumePointOk() (*EpisodeBaseResumePoint, bool)`

GetResumePointOk returns a tuple with the ResumePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePoint

`func (o *CurrentlyPlayingObjectItem) SetResumePoint(v EpisodeBaseResumePoint)`

SetResumePoint sets ResumePoint field to given value.


### GetShow

`func (o *CurrentlyPlayingObjectItem) GetShow() EpisodeObjectAllOfShow`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *CurrentlyPlayingObjectItem) GetShowOk() (*EpisodeObjectAllOfShow, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *CurrentlyPlayingObjectItem) SetShow(v EpisodeObjectAllOfShow)`

SetShow sets Show field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


