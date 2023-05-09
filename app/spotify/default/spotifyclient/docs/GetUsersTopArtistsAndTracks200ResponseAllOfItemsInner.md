# GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**LinkedTrackObjectExternalUrls**](LinkedTrackObjectExternalUrls.md) |  | [optional] 
**Followers** | Pointer to [**ArtistObjectFollowers**](ArtistObjectFollowers.md) |  | [optional] 
**Genres** | Pointer to **[]string** | A list of the genres the artist is associated with. If not yet classified, the array is empty.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**Images** | Pointer to [**[]ImageObject**](ImageObject.md) | Images of the artist in various sizes, widest first.  | [optional] 
**Name** | Pointer to **string** | The name of the track.  | [optional] 
**Popularity** | Pointer to **int32** | The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.&lt;br/&gt;The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.&lt;br/&gt;Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;track\&quot;.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**Album** | Pointer to [**TrackObjectAlbum**](TrackObjectAlbum.md) |  | [optional] 
**Artists** | Pointer to [**[]ArtistObject**](ArtistObject.md) | The artists who performed the track. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | [optional] 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**DiscNumber** | Pointer to **int32** | The disc number (usually &#x60;1&#x60; unless the album consists of more than one disc).  | [optional] 
**DurationMs** | Pointer to **int32** | The track length in milliseconds.  | [optional] 
**Explicit** | Pointer to **bool** | Whether or not the track has explicit lyrics ( &#x60;true&#x60; &#x3D; yes it does; &#x60;false&#x60; &#x3D; no it does not OR unknown).  | [optional] 
**ExternalIds** | Pointer to [**TrackObjectExternalIds**](TrackObjectExternalIds.md) |  | [optional] 
**IsPlayable** | Pointer to **bool** | Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking) is applied. If &#x60;true&#x60;, the track is playable in the given market. Otherwise &#x60;false&#x60;.  | [optional] 
**LinkedFrom** | Pointer to [**TrackObjectLinkedFrom**](TrackObjectLinkedFrom.md) |  | [optional] 
**Restrictions** | Pointer to [**SimplifiedTrackObjectRestrictions**](SimplifiedTrackObjectRestrictions.md) |  | [optional] 
**PreviewUrl** | Pointer to **string** | A link to a 30 second preview (MP3 format) of the track. Can be &#x60;null&#x60;  | [optional] 
**TrackNumber** | Pointer to **int32** | The number of the track. If an album has several discs, the track number is the number on the specified disc.  | [optional] 
**IsLocal** | Pointer to **bool** | Whether or not the track is from a local file.  | [optional] 

## Methods

### NewGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner

`func NewGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner() *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner`

NewGetUsersTopArtistsAndTracks200ResponseAllOfItemsInner instantiates a new GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsersTopArtistsAndTracks200ResponseAllOfItemsInnerWithDefaults

`func NewGetUsersTopArtistsAndTracks200ResponseAllOfItemsInnerWithDefaults() *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner`

NewGetUsersTopArtistsAndTracks200ResponseAllOfItemsInnerWithDefaults instantiates a new GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetExternalUrls() LinkedTrackObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetExternalUrlsOk() (*LinkedTrackObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetExternalUrls(v LinkedTrackObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetFollowers() ArtistObjectFollowers`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetFollowersOk() (*ArtistObjectFollowers, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetFollowers(v ArtistObjectFollowers)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetGenres

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetHref

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetImages(v []ImageObject)`

SetImages sets Images field to given value.

### HasImages

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetName

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPopularity

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetType

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetAlbum

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetAlbum() TrackObjectAlbum`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetAlbumOk() (*TrackObjectAlbum, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetAlbum(v TrackObjectAlbum)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetArtists

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetArtists() []ArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetArtistsOk() (*[]ArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetArtists(v []ArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetAvailableMarkets

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetDiscNumber

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDurationMs

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetExplicit

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### GetExternalIds

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetExternalIds() TrackObjectExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetExternalIdsOk() (*TrackObjectExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetExternalIds(v TrackObjectExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetIsPlayable

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.

### HasIsPlayable

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasIsPlayable() bool`

HasIsPlayable returns a boolean if a field has been set.

### GetLinkedFrom

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetLinkedFrom() TrackObjectLinkedFrom`

GetLinkedFrom returns the LinkedFrom field if non-nil, zero value otherwise.

### GetLinkedFromOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetLinkedFromOk() (*TrackObjectLinkedFrom, bool)`

GetLinkedFromOk returns a tuple with the LinkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFrom

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetLinkedFrom(v TrackObjectLinkedFrom)`

SetLinkedFrom sets LinkedFrom field to given value.

### HasLinkedFrom

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasLinkedFrom() bool`

HasLinkedFrom returns a boolean if a field has been set.

### GetRestrictions

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetRestrictions() SimplifiedTrackObjectRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetRestrictionsOk() (*SimplifiedTrackObjectRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetRestrictions(v SimplifiedTrackObjectRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### GetTrackNumber

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetTrackNumber() int32`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetTrackNumberOk() (*int32, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetTrackNumber(v int32)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### GetIsLocal

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *GetUsersTopArtistsAndTracks200ResponseAllOfItemsInner) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


