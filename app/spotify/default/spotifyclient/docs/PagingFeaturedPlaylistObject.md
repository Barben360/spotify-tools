# PagingFeaturedPlaylistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The localized message of a playlist.  | [optional] 
**Playlists** | Pointer to [**PagingPlaylistObject**](PagingPlaylistObject.md) |  | [optional] 

## Methods

### NewPagingFeaturedPlaylistObject

`func NewPagingFeaturedPlaylistObject() *PagingFeaturedPlaylistObject`

NewPagingFeaturedPlaylistObject instantiates a new PagingFeaturedPlaylistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingFeaturedPlaylistObjectWithDefaults

`func NewPagingFeaturedPlaylistObjectWithDefaults() *PagingFeaturedPlaylistObject`

NewPagingFeaturedPlaylistObjectWithDefaults instantiates a new PagingFeaturedPlaylistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PagingFeaturedPlaylistObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PagingFeaturedPlaylistObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PagingFeaturedPlaylistObject) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PagingFeaturedPlaylistObject) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPlaylists

`func (o *PagingFeaturedPlaylistObject) GetPlaylists() PagingPlaylistObject`

GetPlaylists returns the Playlists field if non-nil, zero value otherwise.

### GetPlaylistsOk

`func (o *PagingFeaturedPlaylistObject) GetPlaylistsOk() (*PagingPlaylistObject, bool)`

GetPlaylistsOk returns a tuple with the Playlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylists

`func (o *PagingFeaturedPlaylistObject) SetPlaylists(v PagingPlaylistObject)`

SetPlaylists sets Playlists field to given value.

### HasPlaylists

`func (o *PagingFeaturedPlaylistObject) HasPlaylists() bool`

HasPlaylists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


