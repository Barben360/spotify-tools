# SavedAlbumObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAt** | Pointer to **time.Time** | The date and time the album was saved Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.  | [optional] 
**Album** | Pointer to [**AlbumObject**](AlbumObject.md) | Information about the album. | [optional] 

## Methods

### NewSavedAlbumObject

`func NewSavedAlbumObject() *SavedAlbumObject`

NewSavedAlbumObject instantiates a new SavedAlbumObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedAlbumObjectWithDefaults

`func NewSavedAlbumObjectWithDefaults() *SavedAlbumObject`

NewSavedAlbumObjectWithDefaults instantiates a new SavedAlbumObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAt

`func (o *SavedAlbumObject) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *SavedAlbumObject) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *SavedAlbumObject) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *SavedAlbumObject) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetAlbum

`func (o *SavedAlbumObject) GetAlbum() AlbumObject`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *SavedAlbumObject) GetAlbumOk() (*AlbumObject, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *SavedAlbumObject) SetAlbum(v AlbumObject)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *SavedAlbumObject) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


