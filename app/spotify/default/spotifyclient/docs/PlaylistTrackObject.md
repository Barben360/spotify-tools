# PlaylistTrackObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAt** | Pointer to **time.Time** | The date and time the track or episode was added. _**Note**: some very old playlists may return &#x60;null&#x60; in this field._  | [optional] 
**AddedBy** | Pointer to [**PlaylistUserObject**](PlaylistUserObject.md) | The Spotify user who added the track or episode. _**Note**: some very old playlists may return &#x60;null&#x60; in this field._  | [optional] 
**IsLocal** | Pointer to **bool** | Whether this track or episode is a [local file](/documentation/web-api/concepts/playlists/#local-files) or not.  | [optional] 
**Track** | Pointer to [**PlaylistTrackObjectTrack**](PlaylistTrackObjectTrack.md) |  | [optional] 

## Methods

### NewPlaylistTrackObject

`func NewPlaylistTrackObject() *PlaylistTrackObject`

NewPlaylistTrackObject instantiates a new PlaylistTrackObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistTrackObjectWithDefaults

`func NewPlaylistTrackObjectWithDefaults() *PlaylistTrackObject`

NewPlaylistTrackObjectWithDefaults instantiates a new PlaylistTrackObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAt

`func (o *PlaylistTrackObject) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *PlaylistTrackObject) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *PlaylistTrackObject) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *PlaylistTrackObject) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetAddedBy

`func (o *PlaylistTrackObject) GetAddedBy() PlaylistUserObject`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *PlaylistTrackObject) GetAddedByOk() (*PlaylistUserObject, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *PlaylistTrackObject) SetAddedBy(v PlaylistUserObject)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *PlaylistTrackObject) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetIsLocal

`func (o *PlaylistTrackObject) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *PlaylistTrackObject) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *PlaylistTrackObject) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *PlaylistTrackObject) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetTrack

`func (o *PlaylistTrackObject) GetTrack() PlaylistTrackObjectTrack`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *PlaylistTrackObject) GetTrackOk() (*PlaylistTrackObjectTrack, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *PlaylistTrackObject) SetTrack(v PlaylistTrackObjectTrack)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *PlaylistTrackObject) HasTrack() bool`

HasTrack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


