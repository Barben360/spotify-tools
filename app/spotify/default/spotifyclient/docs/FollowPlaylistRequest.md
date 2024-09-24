# FollowPlaylistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | Pointer to **bool** | Defaults to &#x60;true&#x60;. If &#x60;true&#x60; the playlist will be included in user&#39;s public playlists (added to profile), if &#x60;false&#x60; it will remain private. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)  | [optional] 

## Methods

### NewFollowPlaylistRequest

`func NewFollowPlaylistRequest() *FollowPlaylistRequest`

NewFollowPlaylistRequest instantiates a new FollowPlaylistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowPlaylistRequestWithDefaults

`func NewFollowPlaylistRequestWithDefaults() *FollowPlaylistRequest`

NewFollowPlaylistRequestWithDefaults instantiates a new FollowPlaylistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *FollowPlaylistRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FollowPlaylistRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FollowPlaylistRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *FollowPlaylistRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


