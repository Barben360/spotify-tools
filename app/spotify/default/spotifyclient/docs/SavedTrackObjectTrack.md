# SavedTrackObjectTrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Album** | Pointer to [**TrackObjectAlbum**](TrackObjectAlbum.md) |  | [optional] 
**Artists** | Pointer to [**[]ArtistObject**](ArtistObject.md) | The artists who performed the track. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | [optional] 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**DiscNumber** | Pointer to **int32** | The disc number (usually &#x60;1&#x60; unless the album consists of more than one disc).  | [optional] 
**DurationMs** | Pointer to **int32** | The track length in milliseconds.  | [optional] 
**Explicit** | Pointer to **bool** | Whether or not the track has explicit lyrics ( &#x60;true&#x60; &#x3D; yes it does; &#x60;false&#x60; &#x3D; no it does not OR unknown).  | [optional] 
**ExternalIds** | Pointer to [**TrackObjectExternalIds**](TrackObjectExternalIds.md) |  | [optional] 
**ExternalUrls** | Pointer to [**LinkedTrackObjectExternalUrls**](LinkedTrackObjectExternalUrls.md) |  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**IsPlayable** | Pointer to **bool** | Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied. If &#x60;true&#x60;, the track is playable in the given market. Otherwise &#x60;false&#x60;.  | [optional] 
**LinkedFrom** | Pointer to [**TrackObjectLinkedFrom**](TrackObjectLinkedFrom.md) |  | [optional] 
**Restrictions** | Pointer to [**SimplifiedTrackObjectRestrictions**](SimplifiedTrackObjectRestrictions.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the track.  | [optional] 
**Popularity** | Pointer to **int32** | The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.&lt;br/&gt;The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.&lt;br/&gt;Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._  | [optional] 
**PreviewUrl** | Pointer to **string** | A link to a 30 second preview (MP3 format) of the track. Can be &#x60;null&#x60;  | [optional] 
**TrackNumber** | Pointer to **int32** | The number of the track. If an album has several discs, the track number is the number on the specified disc.  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;track\&quot;.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**IsLocal** | Pointer to **bool** | Whether or not the track is from a local file.  | [optional] 

## Methods

### NewSavedTrackObjectTrack

`func NewSavedTrackObjectTrack() *SavedTrackObjectTrack`

NewSavedTrackObjectTrack instantiates a new SavedTrackObjectTrack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedTrackObjectTrackWithDefaults

`func NewSavedTrackObjectTrackWithDefaults() *SavedTrackObjectTrack`

NewSavedTrackObjectTrackWithDefaults instantiates a new SavedTrackObjectTrack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbum

`func (o *SavedTrackObjectTrack) GetAlbum() TrackObjectAlbum`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *SavedTrackObjectTrack) GetAlbumOk() (*TrackObjectAlbum, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *SavedTrackObjectTrack) SetAlbum(v TrackObjectAlbum)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *SavedTrackObjectTrack) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetArtists

`func (o *SavedTrackObjectTrack) GetArtists() []ArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *SavedTrackObjectTrack) GetArtistsOk() (*[]ArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *SavedTrackObjectTrack) SetArtists(v []ArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *SavedTrackObjectTrack) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetAvailableMarkets

`func (o *SavedTrackObjectTrack) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *SavedTrackObjectTrack) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *SavedTrackObjectTrack) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *SavedTrackObjectTrack) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetDiscNumber

`func (o *SavedTrackObjectTrack) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *SavedTrackObjectTrack) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *SavedTrackObjectTrack) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *SavedTrackObjectTrack) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDurationMs

`func (o *SavedTrackObjectTrack) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SavedTrackObjectTrack) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SavedTrackObjectTrack) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *SavedTrackObjectTrack) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetExplicit

`func (o *SavedTrackObjectTrack) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *SavedTrackObjectTrack) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *SavedTrackObjectTrack) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *SavedTrackObjectTrack) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### GetExternalIds

`func (o *SavedTrackObjectTrack) GetExternalIds() TrackObjectExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *SavedTrackObjectTrack) GetExternalIdsOk() (*TrackObjectExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *SavedTrackObjectTrack) SetExternalIds(v TrackObjectExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *SavedTrackObjectTrack) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetExternalUrls

`func (o *SavedTrackObjectTrack) GetExternalUrls() LinkedTrackObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SavedTrackObjectTrack) GetExternalUrlsOk() (*LinkedTrackObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SavedTrackObjectTrack) SetExternalUrls(v LinkedTrackObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *SavedTrackObjectTrack) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *SavedTrackObjectTrack) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SavedTrackObjectTrack) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SavedTrackObjectTrack) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SavedTrackObjectTrack) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SavedTrackObjectTrack) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedTrackObjectTrack) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedTrackObjectTrack) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SavedTrackObjectTrack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPlayable

`func (o *SavedTrackObjectTrack) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *SavedTrackObjectTrack) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *SavedTrackObjectTrack) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.

### HasIsPlayable

`func (o *SavedTrackObjectTrack) HasIsPlayable() bool`

HasIsPlayable returns a boolean if a field has been set.

### GetLinkedFrom

`func (o *SavedTrackObjectTrack) GetLinkedFrom() TrackObjectLinkedFrom`

GetLinkedFrom returns the LinkedFrom field if non-nil, zero value otherwise.

### GetLinkedFromOk

`func (o *SavedTrackObjectTrack) GetLinkedFromOk() (*TrackObjectLinkedFrom, bool)`

GetLinkedFromOk returns a tuple with the LinkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFrom

`func (o *SavedTrackObjectTrack) SetLinkedFrom(v TrackObjectLinkedFrom)`

SetLinkedFrom sets LinkedFrom field to given value.

### HasLinkedFrom

`func (o *SavedTrackObjectTrack) HasLinkedFrom() bool`

HasLinkedFrom returns a boolean if a field has been set.

### GetRestrictions

`func (o *SavedTrackObjectTrack) GetRestrictions() SimplifiedTrackObjectRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *SavedTrackObjectTrack) GetRestrictionsOk() (*SimplifiedTrackObjectRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *SavedTrackObjectTrack) SetRestrictions(v SimplifiedTrackObjectRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *SavedTrackObjectTrack) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetName

`func (o *SavedTrackObjectTrack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedTrackObjectTrack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedTrackObjectTrack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SavedTrackObjectTrack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPopularity

`func (o *SavedTrackObjectTrack) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *SavedTrackObjectTrack) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *SavedTrackObjectTrack) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *SavedTrackObjectTrack) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *SavedTrackObjectTrack) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *SavedTrackObjectTrack) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *SavedTrackObjectTrack) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *SavedTrackObjectTrack) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### GetTrackNumber

`func (o *SavedTrackObjectTrack) GetTrackNumber() int32`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *SavedTrackObjectTrack) GetTrackNumberOk() (*int32, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *SavedTrackObjectTrack) SetTrackNumber(v int32)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *SavedTrackObjectTrack) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### GetType

`func (o *SavedTrackObjectTrack) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedTrackObjectTrack) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedTrackObjectTrack) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SavedTrackObjectTrack) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *SavedTrackObjectTrack) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SavedTrackObjectTrack) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SavedTrackObjectTrack) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *SavedTrackObjectTrack) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetIsLocal

`func (o *SavedTrackObjectTrack) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *SavedTrackObjectTrack) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *SavedTrackObjectTrack) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *SavedTrackObjectTrack) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


