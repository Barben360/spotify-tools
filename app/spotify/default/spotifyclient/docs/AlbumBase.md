# AlbumBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlbumType** | **string** | The type of the album.  | 
**TotalTracks** | **int32** | The number of tracks in the album. | 
**AvailableMarkets** | **[]string** | The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._  | 
**ExternalUrls** | [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this album.  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the album.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the album.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the album in various sizes, widest first.  | 
**Name** | **string** | The name of the album. In case of an album takedown, the value may be an empty string.  | 
**ReleaseDate** | **string** | The date the album was first released.  | 
**ReleaseDatePrecision** | **string** | The precision with which &#x60;release_date&#x60; value is known.  | 
**Restrictions** | Pointer to [**AlbumRestrictionObject**](AlbumRestrictionObject.md) | Included in the response when a content restriction is applied.  | [optional] 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the album.  | 

## Methods

### NewAlbumBase

`func NewAlbumBase(albumType string, totalTracks int32, availableMarkets []string, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string, ) *AlbumBase`

NewAlbumBase instantiates a new AlbumBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumBaseWithDefaults

`func NewAlbumBaseWithDefaults() *AlbumBase`

NewAlbumBaseWithDefaults instantiates a new AlbumBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbumType

`func (o *AlbumBase) GetAlbumType() string`

GetAlbumType returns the AlbumType field if non-nil, zero value otherwise.

### GetAlbumTypeOk

`func (o *AlbumBase) GetAlbumTypeOk() (*string, bool)`

GetAlbumTypeOk returns a tuple with the AlbumType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumType

`func (o *AlbumBase) SetAlbumType(v string)`

SetAlbumType sets AlbumType field to given value.


### GetTotalTracks

`func (o *AlbumBase) GetTotalTracks() int32`

GetTotalTracks returns the TotalTracks field if non-nil, zero value otherwise.

### GetTotalTracksOk

`func (o *AlbumBase) GetTotalTracksOk() (*int32, bool)`

GetTotalTracksOk returns a tuple with the TotalTracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTracks

`func (o *AlbumBase) SetTotalTracks(v int32)`

SetTotalTracks sets TotalTracks field to given value.


### GetAvailableMarkets

`func (o *AlbumBase) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *AlbumBase) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *AlbumBase) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.


### GetExternalUrls

`func (o *AlbumBase) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *AlbumBase) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *AlbumBase) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *AlbumBase) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AlbumBase) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AlbumBase) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *AlbumBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlbumBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlbumBase) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *AlbumBase) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AlbumBase) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AlbumBase) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetName

`func (o *AlbumBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlbumBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlbumBase) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseDate

`func (o *AlbumBase) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AlbumBase) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AlbumBase) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.


### GetReleaseDatePrecision

`func (o *AlbumBase) GetReleaseDatePrecision() string`

GetReleaseDatePrecision returns the ReleaseDatePrecision field if non-nil, zero value otherwise.

### GetReleaseDatePrecisionOk

`func (o *AlbumBase) GetReleaseDatePrecisionOk() (*string, bool)`

GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDatePrecision

`func (o *AlbumBase) SetReleaseDatePrecision(v string)`

SetReleaseDatePrecision sets ReleaseDatePrecision field to given value.


### GetRestrictions

`func (o *AlbumBase) GetRestrictions() AlbumRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *AlbumBase) GetRestrictionsOk() (*AlbumRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *AlbumBase) SetRestrictions(v AlbumRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *AlbumBase) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetType

`func (o *AlbumBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlbumBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlbumBase) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *AlbumBase) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AlbumBase) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AlbumBase) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


