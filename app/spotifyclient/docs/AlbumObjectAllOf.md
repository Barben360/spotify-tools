# AlbumObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artists** | Pointer to [**[]SimplifiedArtistObject**](SimplifiedArtistObject.md) | The artists of the album. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | [optional] 
**Tracks** | Pointer to [**AlbumObjectAllOfTracks**](AlbumObjectAllOfTracks.md) |  | [optional] 
**Popularity** | Pointer to **int32** | The popularity of the album, with 100 being the most popular. The popularity is calculated from the popularity of the album&#39;s individual tracks. | [optional] 
**Label** | Pointer to **string** | The label for the album. | [optional] 
**ExternalIds** | Pointer to [**AlbumObjectAllOfExternalIds**](AlbumObjectAllOfExternalIds.md) |  | [optional] 
**Genres** | Pointer to **[]string** | A list of the genres used to classify the album. (If not yet classified, the array is empty.) | [optional] 
**Copyrights** | Pointer to [**[]CopyrightObject**](CopyrightObject.md) | The copyright statements of the album. | [optional] 

## Methods

### NewAlbumObjectAllOf

`func NewAlbumObjectAllOf() *AlbumObjectAllOf`

NewAlbumObjectAllOf instantiates a new AlbumObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumObjectAllOfWithDefaults

`func NewAlbumObjectAllOfWithDefaults() *AlbumObjectAllOf`

NewAlbumObjectAllOfWithDefaults instantiates a new AlbumObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtists

`func (o *AlbumObjectAllOf) GetArtists() []SimplifiedArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *AlbumObjectAllOf) GetArtistsOk() (*[]SimplifiedArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *AlbumObjectAllOf) SetArtists(v []SimplifiedArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *AlbumObjectAllOf) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetTracks

`func (o *AlbumObjectAllOf) GetTracks() AlbumObjectAllOfTracks`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *AlbumObjectAllOf) GetTracksOk() (*AlbumObjectAllOfTracks, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *AlbumObjectAllOf) SetTracks(v AlbumObjectAllOfTracks)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *AlbumObjectAllOf) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### GetPopularity

`func (o *AlbumObjectAllOf) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *AlbumObjectAllOf) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *AlbumObjectAllOf) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *AlbumObjectAllOf) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetLabel

`func (o *AlbumObjectAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AlbumObjectAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AlbumObjectAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AlbumObjectAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetExternalIds

`func (o *AlbumObjectAllOf) GetExternalIds() AlbumObjectAllOfExternalIds`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *AlbumObjectAllOf) GetExternalIdsOk() (*AlbumObjectAllOfExternalIds, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *AlbumObjectAllOf) SetExternalIds(v AlbumObjectAllOfExternalIds)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *AlbumObjectAllOf) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetGenres

`func (o *AlbumObjectAllOf) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *AlbumObjectAllOf) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *AlbumObjectAllOf) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *AlbumObjectAllOf) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetCopyrights

`func (o *AlbumObjectAllOf) GetCopyrights() []CopyrightObject`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *AlbumObjectAllOf) GetCopyrightsOk() (*[]CopyrightObject, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *AlbumObjectAllOf) SetCopyrights(v []CopyrightObject)`

SetCopyrights sets Copyrights field to given value.

### HasCopyrights

`func (o *AlbumObjectAllOf) HasCopyrights() bool`

HasCopyrights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


