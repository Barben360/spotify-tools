# SimplifiedAlbumObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlbumGroup** | Pointer to **string** | The field is present when getting an artist&#39;s albums. Compare to album_type this field represents relationship between the artist and the album.  | [optional] 
**Artists** | [**[]SimplifiedArtistObject**](SimplifiedArtistObject.md) | The artists of the album. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | 

## Methods

### NewSimplifiedAlbumObjectAllOf

`func NewSimplifiedAlbumObjectAllOf(artists []SimplifiedArtistObject, ) *SimplifiedAlbumObjectAllOf`

NewSimplifiedAlbumObjectAllOf instantiates a new SimplifiedAlbumObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedAlbumObjectAllOfWithDefaults

`func NewSimplifiedAlbumObjectAllOfWithDefaults() *SimplifiedAlbumObjectAllOf`

NewSimplifiedAlbumObjectAllOfWithDefaults instantiates a new SimplifiedAlbumObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbumGroup

`func (o *SimplifiedAlbumObjectAllOf) GetAlbumGroup() string`

GetAlbumGroup returns the AlbumGroup field if non-nil, zero value otherwise.

### GetAlbumGroupOk

`func (o *SimplifiedAlbumObjectAllOf) GetAlbumGroupOk() (*string, bool)`

GetAlbumGroupOk returns a tuple with the AlbumGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumGroup

`func (o *SimplifiedAlbumObjectAllOf) SetAlbumGroup(v string)`

SetAlbumGroup sets AlbumGroup field to given value.

### HasAlbumGroup

`func (o *SimplifiedAlbumObjectAllOf) HasAlbumGroup() bool`

HasAlbumGroup returns a boolean if a field has been set.

### GetArtists

`func (o *SimplifiedAlbumObjectAllOf) GetArtists() []SimplifiedArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *SimplifiedAlbumObjectAllOf) GetArtistsOk() (*[]SimplifiedArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *SimplifiedAlbumObjectAllOf) SetArtists(v []SimplifiedArtistObject)`

SetArtists sets Artists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


