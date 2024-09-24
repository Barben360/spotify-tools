# SimplifiedPlaylistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collaborative** | Pointer to **bool** | &#x60;true&#x60; if the owner allows other users to modify the playlist.  | [optional] 
**Description** | Pointer to **string** | The playlist description. _Only returned for modified, verified playlists, otherwise_ &#x60;null&#x60;.  | [optional] 
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this playlist.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the playlist.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.  | [optional] 
**Images** | Pointer to [**[]ImageObject**](ImageObject.md) | Images for the playlist. The array may be empty or contain up to three images. The images are returned by size in descending order. See [Working with Playlists](/documentation/web-api/concepts/playlists). _**Note**: If returned, the source URL for the image (&#x60;url&#x60;) is temporary and will expire in less than a day._  | [optional] 
**Name** | Pointer to **string** | The name of the playlist.  | [optional] 
**Owner** | Pointer to [**PlaylistOwnerObject**](PlaylistOwnerObject.md) | The user who owns the playlist  | [optional] 
**Public** | Pointer to **bool** | The playlist&#39;s public/private status (if it is added to the user&#39;s profile): &#x60;true&#x60; the playlist is public, &#x60;false&#x60; the playlist is private, &#x60;null&#x60; the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/web-api/concepts/playlists)  | [optional] 
**SnapshotId** | Pointer to **string** | The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version  | [optional] 
**Tracks** | Pointer to [**PlaylistTracksRefObject**](PlaylistTracksRefObject.md) | A collection containing a link ( &#x60;href&#x60; ) to the Web API endpoint where full details of the playlist&#39;s tracks can be retrieved, along with the &#x60;total&#x60; number of tracks in the playlist. Note, a track object may be &#x60;null&#x60;. This can happen if a track is no longer available.  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;playlist\&quot;  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the playlist.  | [optional] 

## Methods

### NewSimplifiedPlaylistObject

`func NewSimplifiedPlaylistObject() *SimplifiedPlaylistObject`

NewSimplifiedPlaylistObject instantiates a new SimplifiedPlaylistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedPlaylistObjectWithDefaults

`func NewSimplifiedPlaylistObjectWithDefaults() *SimplifiedPlaylistObject`

NewSimplifiedPlaylistObjectWithDefaults instantiates a new SimplifiedPlaylistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollaborative

`func (o *SimplifiedPlaylistObject) GetCollaborative() bool`

GetCollaborative returns the Collaborative field if non-nil, zero value otherwise.

### GetCollaborativeOk

`func (o *SimplifiedPlaylistObject) GetCollaborativeOk() (*bool, bool)`

GetCollaborativeOk returns a tuple with the Collaborative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborative

`func (o *SimplifiedPlaylistObject) SetCollaborative(v bool)`

SetCollaborative sets Collaborative field to given value.

### HasCollaborative

`func (o *SimplifiedPlaylistObject) HasCollaborative() bool`

HasCollaborative returns a boolean if a field has been set.

### GetDescription

`func (o *SimplifiedPlaylistObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimplifiedPlaylistObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimplifiedPlaylistObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimplifiedPlaylistObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalUrls

`func (o *SimplifiedPlaylistObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SimplifiedPlaylistObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SimplifiedPlaylistObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *SimplifiedPlaylistObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *SimplifiedPlaylistObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedPlaylistObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedPlaylistObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedPlaylistObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SimplifiedPlaylistObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedPlaylistObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedPlaylistObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimplifiedPlaylistObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *SimplifiedPlaylistObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SimplifiedPlaylistObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SimplifiedPlaylistObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.

### HasImages

`func (o *SimplifiedPlaylistObject) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedPlaylistObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedPlaylistObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedPlaylistObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedPlaylistObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *SimplifiedPlaylistObject) GetOwner() PlaylistOwnerObject`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SimplifiedPlaylistObject) GetOwnerOk() (*PlaylistOwnerObject, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SimplifiedPlaylistObject) SetOwner(v PlaylistOwnerObject)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SimplifiedPlaylistObject) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublic

`func (o *SimplifiedPlaylistObject) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *SimplifiedPlaylistObject) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *SimplifiedPlaylistObject) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *SimplifiedPlaylistObject) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSnapshotId

`func (o *SimplifiedPlaylistObject) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SimplifiedPlaylistObject) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SimplifiedPlaylistObject) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *SimplifiedPlaylistObject) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetTracks

`func (o *SimplifiedPlaylistObject) GetTracks() PlaylistTracksRefObject`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *SimplifiedPlaylistObject) GetTracksOk() (*PlaylistTracksRefObject, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *SimplifiedPlaylistObject) SetTracks(v PlaylistTracksRefObject)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *SimplifiedPlaylistObject) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedPlaylistObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedPlaylistObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedPlaylistObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedPlaylistObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *SimplifiedPlaylistObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SimplifiedPlaylistObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SimplifiedPlaylistObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *SimplifiedPlaylistObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


