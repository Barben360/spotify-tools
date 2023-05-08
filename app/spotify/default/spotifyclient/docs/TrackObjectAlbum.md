# TrackObjectAlbum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlbumType** | **string** | The type of the album.  | 
**TotalTracks** | **int32** | The number of tracks in the album. | 
**AvailableMarkets** | **[]string** | The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._  | 
**ExternalUrls** | [**AlbumBaseExternalUrls**](AlbumBaseExternalUrls.md) |  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the album.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the album.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the album in various sizes, widest first.  | 
**Name** | **string** | The name of the album. In case of an album takedown, the value may be an empty string.  | 
**ReleaseDate** | **string** | The date the album was first released.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**Restrictions** | Pointer to [**AlbumBaseRestrictions**](AlbumBaseRestrictions.md) |  | [optional] 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the album.  | 
**IsPlayable** | Pointer to **bool** | Fix by Barben360  | [optional] 
**AlbumGroup** | Pointer to **string** | The field is present when getting an artist&#39;s albums. Compare to album_type this field represents relationship between the artist and the album.  | [optional] 
**Artists** | [**[]SimplifiedArtistObject**](SimplifiedArtistObject.md) | The artists of the album. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | 

## Methods

### NewTrackObjectAlbum

`func NewTrackObjectAlbum(albumType string, totalTracks int32, availableMarkets []string, externalUrls AlbumBaseExternalUrls, href string, id string, images []ImageObject, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, artists []SimplifiedArtistObject, ) *TrackObjectAlbum`

NewTrackObjectAlbum instantiates a new TrackObjectAlbum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackObjectAlbumWithDefaults

`func NewTrackObjectAlbumWithDefaults() *TrackObjectAlbum`

NewTrackObjectAlbumWithDefaults instantiates a new TrackObjectAlbum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbumType

`func (o *TrackObjectAlbum) GetAlbumType() string`

GetAlbumType returns the AlbumType field if non-nil, zero value otherwise.

### GetAlbumTypeOk

`func (o *TrackObjectAlbum) GetAlbumTypeOk() (*string, bool)`

GetAlbumTypeOk returns a tuple with the AlbumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumType

`func (o *TrackObjectAlbum) SetAlbumType(v string)`

SetAlbumType sets AlbumType field to given value.


### GetTotalTracks

`func (o *TrackObjectAlbum) GetTotalTracks() int32`

GetTotalTracks returns the TotalTracks field if non-nil, zero value otherwise.

### GetTotalTracksOk

`func (o *TrackObjectAlbum) GetTotalTracksOk() (*int32, bool)`

GetTotalTracksOk returns a tuple with the TotalTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTracks

`func (o *TrackObjectAlbum) SetTotalTracks(v int32)`

SetTotalTracks sets TotalTracks field to given value.


### GetAvailableMarkets

`func (o *TrackObjectAlbum) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *TrackObjectAlbum) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *TrackObjectAlbum) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.


### GetExternalUrls

`func (o *TrackObjectAlbum) GetExternalUrls() AlbumBaseExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *TrackObjectAlbum) GetExternalUrlsOk() (*AlbumBaseExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *TrackObjectAlbum) SetExternalUrls(v AlbumBaseExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *TrackObjectAlbum) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TrackObjectAlbum) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TrackObjectAlbum) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *TrackObjectAlbum) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackObjectAlbum) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackObjectAlbum) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *TrackObjectAlbum) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *TrackObjectAlbum) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *TrackObjectAlbum) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetName

`func (o *TrackObjectAlbum) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackObjectAlbum) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackObjectAlbum) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *TrackObjectAlbum) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *TrackObjectAlbum) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *TrackObjectAlbum) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *TrackObjectAlbum) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *TrackObjectAlbum) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *TrackObjectAlbum) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetRestrictions

`func (o *TrackObjectAlbum) GetRestrictions() AlbumBaseRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *TrackObjectAlbum) GetRestrictionsOk() (*AlbumBaseRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *TrackObjectAlbum) SetRestrictions(v AlbumBaseRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *TrackObjectAlbum) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetType

`func (o *TrackObjectAlbum) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrackObjectAlbum) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrackObjectAlbum) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *TrackObjectAlbum) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TrackObjectAlbum) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TrackObjectAlbum) SetUri(v string)`

SetUri sets Uri field to given value.


### GetIsPlayable

`func (o *TrackObjectAlbum) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *TrackObjectAlbum) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *TrackObjectAlbum) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.

### HasIsPlayable

`func (o *TrackObjectAlbum) HasIsPlayable() bool`

HasIsPlayable returns a boolean if a field has been set.

### GetAlbumGroup

`func (o *TrackObjectAlbum) GetAlbumGroup() string`

GetAlbumGroup returns the AlbumGroup field if non-nil, zero value otherwise.

### GetAlbumGroupOk

`func (o *TrackObjectAlbum) GetAlbumGroupOk() (*string, bool)`

GetAlbumGroupOk returns a tuple with the AlbumGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumGroup

`func (o *TrackObjectAlbum) SetAlbumGroup(v string)`

SetAlbumGroup sets AlbumGroup field to given value.

### HasAlbumGroup

`func (o *TrackObjectAlbum) HasAlbumGroup() bool`

HasAlbumGroup returns a boolean if a field has been set.

### GetArtists

`func (o *TrackObjectAlbum) GetArtists() []SimplifiedArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *TrackObjectAlbum) GetArtistsOk() (*[]SimplifiedArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *TrackObjectAlbum) SetArtists(v []SimplifiedArtistObject)`

SetArtists sets Artists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


