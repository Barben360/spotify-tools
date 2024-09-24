# ChangePlaylistDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name for the playlist, for example &#x60;\&quot;My New Playlist Title\&quot;&#x60;  | [optional] 
**Public** | Pointer to **bool** | The playlist&#39;s public/private status (if it should be added to the user&#39;s profile or not): &#x60;true&#x60; the playlist will be public, &#x60;false&#x60; the playlist will be private, &#x60;null&#x60; the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)  | [optional] 
**Collaborative** | Pointer to **bool** | If &#x60;true&#x60;, the playlist will become collaborative and other users will be able to modify the playlist in their Spotify client. &lt;br/&gt; _**Note**: You can only set &#x60;collaborative&#x60; to &#x60;true&#x60; on non-public playlists._  | [optional] 
**Description** | Pointer to **string** | Value for playlist description as displayed in Spotify Clients and in the Web API.  | [optional] 

## Methods

### NewChangePlaylistDetailsRequest

`func NewChangePlaylistDetailsRequest() *ChangePlaylistDetailsRequest`

NewChangePlaylistDetailsRequest instantiates a new ChangePlaylistDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePlaylistDetailsRequestWithDefaults

`func NewChangePlaylistDetailsRequestWithDefaults() *ChangePlaylistDetailsRequest`

NewChangePlaylistDetailsRequestWithDefaults instantiates a new ChangePlaylistDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChangePlaylistDetailsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChangePlaylistDetailsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChangePlaylistDetailsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChangePlaylistDetailsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *ChangePlaylistDetailsRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ChangePlaylistDetailsRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ChangePlaylistDetailsRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ChangePlaylistDetailsRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCollaborative

`func (o *ChangePlaylistDetailsRequest) GetCollaborative() bool`

GetCollaborative returns the Collaborative field if non-nil, zero value otherwise.

### GetCollaborativeOk

`func (o *ChangePlaylistDetailsRequest) GetCollaborativeOk() (*bool, bool)`

GetCollaborativeOk returns a tuple with the Collaborative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborative

`func (o *ChangePlaylistDetailsRequest) SetCollaborative(v bool)`

SetCollaborative sets Collaborative field to given value.

### HasCollaborative

`func (o *ChangePlaylistDetailsRequest) HasCollaborative() bool`

HasCollaborative returns a boolean if a field has been set.

### GetDescription

`func (o *ChangePlaylistDetailsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChangePlaylistDetailsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChangePlaylistDetailsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChangePlaylistDetailsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


