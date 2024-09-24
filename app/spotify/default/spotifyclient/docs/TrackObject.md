# TrackObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Album** | Pointer to [**SimplifiedAlbumObject**](SimplifiedAlbumObject.md) | The album on which the track appears. The album object includes a link in &#x60;href&#x60; to full information about the album.  | [optional] 
**Artists** | Pointer to [**[]ArtistObject**](ArtistObject.md) | The artists who performed the track. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | [optional] 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**DiscNumber** | Pointer to **int32** | The disc number (usually &#x60;1&#x60; unless the album consists of more than one disc).  | [optional] 
**DurationMs** | Pointer to **int32** | The track length in milliseconds.  | [optional] 
**Explicit** | Pointer to **bool** | Whether or not the track has explicit lyrics ( &#x60;true&#x60; &#x3D; yes it does; &#x60;false&#x60; &#x3D; no it does not OR unknown).  | [optional] 
**ExternalIds** | Pointer to [**ExternalIdObject**](ExternalIdObject.md) | Known external IDs for the track.  | [optional] 
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this track.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**IsPlayable** | Pointer to **bool** | Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied. If &#x60;true&#x60;, the track is playable in the given market. Otherwise &#x60;false&#x60;.  | [optional] 
**LinkedFrom** | Pointer to [**LinkedTrackObject**](LinkedTrackObject.md) | Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied, and the requested track has been replaced with different track. The track in the &#x60;linked_from&#x60; object contains information about the originally requested track. | [optional] 
**Restrictions** | Pointer to [**TrackRestrictionObject**](TrackRestrictionObject.md) | Included in the response when a content restriction is applied.  | [optional] 
**Name** | Pointer to **string** | The name of the track.  | [optional] 
**Popularity** | Pointer to **int32** | The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.&lt;br/&gt;The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.&lt;br/&gt;Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._  | [optional] 
**PreviewUrl** | Pointer to **NullableString** | A link to a 30 second preview (MP3 format) of the track. Can be &#x60;null&#x60;  | [optional] 
**TrackNumber** | Pointer to **int32** | The number of the track. If an album has several discs, the track number is the number on the specified disc.  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;track\&quot;.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**IsLocal** | Pointer to **bool** | Whether or not the track is from a local file.  | [optional] 

## Methods

### NewTrackObject

`func NewTrackObject() *TrackObject`

NewTrackObject instantiates a new TrackObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackObjectWithDefaults

`func NewTrackObjectWithDefaults() *TrackObject`

NewTrackObjectWithDefaults instantiates a new TrackObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbum

`func (o *TrackObject) GetAlbum() SimplifiedAlbumObject`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *TrackObject) GetAlbumOk() (*SimplifiedAlbumObject, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *TrackObject) SetAlbum(v SimplifiedAlbumObject)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *TrackObject) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetArtists

`func (o *TrackObject) GetArtists() []ArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *TrackObject) GetArtistsOk() (*[]ArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *TrackObject) SetArtists(v []ArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *TrackObject) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetAvailableMarkets

`func (o *TrackObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *TrackObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *TrackObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *TrackObject) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetDiscNumber

`func (o *TrackObject) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *TrackObject) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *TrackObject) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *TrackObject) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDurationMs

`func (o *TrackObject) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *TrackObject) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *TrackObject) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *TrackObject) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetExplicit

`func (o *TrackObject) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *TrackObject) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *TrackObject) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *TrackObject) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### GetExternalIds

`func (o *TrackObject) GetExternalIds() ExternalIdObject`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *TrackObject) GetExternalIdsOk() (*ExternalIdObject, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *TrackObject) SetExternalIds(v ExternalIdObject)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *TrackObject) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetExternalUrls

`func (o *TrackObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *TrackObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *TrackObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *TrackObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *TrackObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TrackObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TrackObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TrackObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *TrackObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPlayable

`func (o *TrackObject) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *TrackObject) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *TrackObject) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.

### HasIsPlayable

`func (o *TrackObject) HasIsPlayable() bool`

HasIsPlayable returns a boolean if a field has been set.

### GetLinkedFrom

`func (o *TrackObject) GetLinkedFrom() LinkedTrackObject`

GetLinkedFrom returns the LinkedFrom field if non-nil, zero value otherwise.

### GetLinkedFromOk

`func (o *TrackObject) GetLinkedFromOk() (*LinkedTrackObject, bool)`

GetLinkedFromOk returns a tuple with the LinkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFrom

`func (o *TrackObject) SetLinkedFrom(v LinkedTrackObject)`

SetLinkedFrom sets LinkedFrom field to given value.

### HasLinkedFrom

`func (o *TrackObject) HasLinkedFrom() bool`

HasLinkedFrom returns a boolean if a field has been set.

### GetRestrictions

`func (o *TrackObject) GetRestrictions() TrackRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *TrackObject) GetRestrictionsOk() (*TrackRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *TrackObject) SetRestrictions(v TrackRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *TrackObject) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetName

`func (o *TrackObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrackObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPopularity

`func (o *TrackObject) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *TrackObject) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *TrackObject) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *TrackObject) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *TrackObject) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *TrackObject) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *TrackObject) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *TrackObject) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### SetPreviewUrlNil

`func (o *TrackObject) SetPreviewUrlNil(b bool)`

 SetPreviewUrlNil sets the value for PreviewUrl to be an explicit nil

### UnsetPreviewUrl
`func (o *TrackObject) UnsetPreviewUrl()`

UnsetPreviewUrl ensures that no value is present for PreviewUrl, not even an explicit nil
### GetTrackNumber

`func (o *TrackObject) GetTrackNumber() int32`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *TrackObject) GetTrackNumberOk() (*int32, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *TrackObject) SetTrackNumber(v int32)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *TrackObject) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### GetType

`func (o *TrackObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrackObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrackObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrackObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *TrackObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TrackObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TrackObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TrackObject) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetIsLocal

`func (o *TrackObject) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *TrackObject) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *TrackObject) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *TrackObject) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


