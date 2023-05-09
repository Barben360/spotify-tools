# RemoveTracksPlaylistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tracks** | [**[]RemoveTracksPlaylistRequestTracksInner**](RemoveTracksPlaylistRequestTracksInner.md) | An array of objects containing [Spotify URIs](/documentation/web-api/concepts/spotify-uris-ids) of the tracks or episodes to remove. For example: &#x60;{ \&quot;tracks\&quot;: [{ \&quot;uri\&quot;: \&quot;spotify:track:4iV5W9uYEdYUVa79Axb7Rh\&quot; },{ \&quot;uri\&quot;: \&quot;spotify:track:1301WleyT98MSxVHPZCA6M\&quot; }] }&#x60;. A maximum of 100 objects can be sent at once.  | 
**SnapshotId** | Pointer to **string** | The playlist&#39;s snapshot ID against which you want to make the changes. The API will validate that the specified items exist and in the specified positions and make the changes, even if more recent changes have been made to the playlist.  | [optional] 

## Methods

### NewRemoveTracksPlaylistRequest

`func NewRemoveTracksPlaylistRequest(tracks []RemoveTracksPlaylistRequestTracksInner, ) *RemoveTracksPlaylistRequest`

NewRemoveTracksPlaylistRequest instantiates a new RemoveTracksPlaylistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveTracksPlaylistRequestWithDefaults

`func NewRemoveTracksPlaylistRequestWithDefaults() *RemoveTracksPlaylistRequest`

NewRemoveTracksPlaylistRequestWithDefaults instantiates a new RemoveTracksPlaylistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTracks

`func (o *RemoveTracksPlaylistRequest) GetTracks() []RemoveTracksPlaylistRequestTracksInner`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *RemoveTracksPlaylistRequest) GetTracksOk() (*[]RemoveTracksPlaylistRequestTracksInner, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *RemoveTracksPlaylistRequest) SetTracks(v []RemoveTracksPlaylistRequestTracksInner)`

SetTracks sets Tracks field to given value.


### GetSnapshotId

`func (o *RemoveTracksPlaylistRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RemoveTracksPlaylistRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RemoveTracksPlaylistRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *RemoveTracksPlaylistRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


