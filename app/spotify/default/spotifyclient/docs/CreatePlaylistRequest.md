# CreatePlaylistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for the new playlist, for example &#x60;\&quot;Your Coolest Playlist\&quot;&#x60;. This name does not need to be unique; a user may have several playlists with the same name.  | 
**Public** | Pointer to **bool** | Defaults to &#x60;true&#x60;. The playlist&#39;s public/private status (if it should be added to the user&#39;s profile or not): &#x60;true&#x60; the playlist will be public, &#x60;false&#x60; the playlist will be private. To be able to create private playlists, the user must have granted the &#x60;playlist-modify-private&#x60; [scope](/documentation/web-api/concepts/scopes/#list-of-scopes). For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)  | [optional] 
**Collaborative** | Pointer to **bool** | Defaults to &#x60;false&#x60;. If &#x60;true&#x60; the playlist will be collaborative. _**Note**: to create a collaborative playlist you must also set &#x60;public&#x60; to &#x60;false&#x60;. To create collaborative playlists you must have granted &#x60;playlist-modify-private&#x60; and &#x60;playlist-modify-public&#x60; [scopes](/documentation/web-api/concepts/scopes/#list-of-scopes)._  | [optional] 
**Description** | Pointer to **string** | value for playlist description as displayed in Spotify Clients and in the Web API.  | [optional] 

## Methods

### NewCreatePlaylistRequest

`func NewCreatePlaylistRequest(name string, ) *CreatePlaylistRequest`

NewCreatePlaylistRequest instantiates a new CreatePlaylistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlaylistRequestWithDefaults

`func NewCreatePlaylistRequestWithDefaults() *CreatePlaylistRequest`

NewCreatePlaylistRequestWithDefaults instantiates a new CreatePlaylistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePlaylistRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlaylistRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlaylistRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPublic

`func (o *CreatePlaylistRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CreatePlaylistRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CreatePlaylistRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CreatePlaylistRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCollaborative

`func (o *CreatePlaylistRequest) GetCollaborative() bool`

GetCollaborative returns the Collaborative field if non-nil, zero value otherwise.

### GetCollaborativeOk

`func (o *CreatePlaylistRequest) GetCollaborativeOk() (*bool, bool)`

GetCollaborativeOk returns a tuple with the Collaborative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborative

`func (o *CreatePlaylistRequest) SetCollaborative(v bool)`

SetCollaborative sets Collaborative field to given value.

### HasCollaborative

`func (o *CreatePlaylistRequest) HasCollaborative() bool`

HasCollaborative returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePlaylistRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePlaylistRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePlaylistRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePlaylistRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


