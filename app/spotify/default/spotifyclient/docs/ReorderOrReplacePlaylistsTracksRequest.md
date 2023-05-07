# ReorderOrReplacePlaylistsTracksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uris** | Pointer to **[]string** |  | [optional] 
**RangeStart** | Pointer to **int32** | The position of the first item to be reordered.  | [optional] 
**InsertBefore** | Pointer to **int32** | The position where the items should be inserted.&lt;br/&gt;To reorder the items to the end of the playlist, simply set _insert_before_ to the position after the last item.&lt;br/&gt;Examples:&lt;br/&gt;To reorder the first item to the last position in a playlist with 10 items, set _range_start_ to 0, and _insert_before_ to 10.&lt;br/&gt;To reorder the last item in a playlist with 10 items to the start of the playlist, set _range_start_ to 9, and _insert_before_ to 0.  | [optional] 
**RangeLength** | Pointer to **int32** | The amount of items to be reordered. Defaults to 1 if not set.&lt;br/&gt;The range of items to be reordered begins from the _range_start_ position, and includes the _range_length_ subsequent items.&lt;br/&gt;Example:&lt;br/&gt;To move the items at index 9-10 to the start of the playlist, _range_start_ is set to 9, and _range_length_ is set to 2.  | [optional] 
**SnapshotId** | Pointer to **string** | The playlist&#39;s snapshot ID against which you want to make the changes.  | [optional] 

## Methods

### NewReorderOrReplacePlaylistsTracksRequest

`func NewReorderOrReplacePlaylistsTracksRequest() *ReorderOrReplacePlaylistsTracksRequest`

NewReorderOrReplacePlaylistsTracksRequest instantiates a new ReorderOrReplacePlaylistsTracksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReorderOrReplacePlaylistsTracksRequestWithDefaults

`func NewReorderOrReplacePlaylistsTracksRequestWithDefaults() *ReorderOrReplacePlaylistsTracksRequest`

NewReorderOrReplacePlaylistsTracksRequestWithDefaults instantiates a new ReorderOrReplacePlaylistsTracksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUris

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetUris() []string`

GetUris returns the Uris field if non-nil, zero value otherwise.

### GetUrisOk

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetUrisOk() (*[]string, bool)`

GetUrisOk returns a tuple with the Uris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUris

`func (o *ReorderOrReplacePlaylistsTracksRequest) SetUris(v []string)`

SetUris sets Uris field to given value.

### HasUris

`func (o *ReorderOrReplacePlaylistsTracksRequest) HasUris() bool`

HasUris returns a boolean if a field has been set.

### GetRangeStart

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetRangeStart() int32`

GetRangeStart returns the RangeStart field if non-nil, zero value otherwise.

### GetRangeStartOk

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetRangeStartOk() (*int32, bool)`

GetRangeStartOk returns a tuple with the RangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStart

`func (o *ReorderOrReplacePlaylistsTracksRequest) SetRangeStart(v int32)`

SetRangeStart sets RangeStart field to given value.

### HasRangeStart

`func (o *ReorderOrReplacePlaylistsTracksRequest) HasRangeStart() bool`

HasRangeStart returns a boolean if a field has been set.

### GetInsertBefore

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetInsertBefore() int32`

GetInsertBefore returns the InsertBefore field if non-nil, zero value otherwise.

### GetInsertBeforeOk

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetInsertBeforeOk() (*int32, bool)`

GetInsertBeforeOk returns a tuple with the InsertBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertBefore

`func (o *ReorderOrReplacePlaylistsTracksRequest) SetInsertBefore(v int32)`

SetInsertBefore sets InsertBefore field to given value.

### HasInsertBefore

`func (o *ReorderOrReplacePlaylistsTracksRequest) HasInsertBefore() bool`

HasInsertBefore returns a boolean if a field has been set.

### GetRangeLength

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetRangeLength() int32`

GetRangeLength returns the RangeLength field if non-nil, zero value otherwise.

### GetRangeLengthOk

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetRangeLengthOk() (*int32, bool)`

GetRangeLengthOk returns a tuple with the RangeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeLength

`func (o *ReorderOrReplacePlaylistsTracksRequest) SetRangeLength(v int32)`

SetRangeLength sets RangeLength field to given value.

### HasRangeLength

`func (o *ReorderOrReplacePlaylistsTracksRequest) HasRangeLength() bool`

HasRangeLength returns a boolean if a field has been set.

### GetSnapshotId

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ReorderOrReplacePlaylistsTracksRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ReorderOrReplacePlaylistsTracksRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ReorderOrReplacePlaylistsTracksRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


