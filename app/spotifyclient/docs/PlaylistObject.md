# PlaylistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collaborative** | Pointer to **bool** | &#x60;true&#x60; if the owner allows other users to modify the playlist.  | [optional] 
**Description** | Pointer to **NullableString** | The playlist description. _Only returned for modified, verified playlists, otherwise_ &#x60;null&#x60;.  | [optional] 
**ExternalUrls** | Pointer to [**PlaylistObjectExternalUrls**](PlaylistObjectExternalUrls.md) |  | [optional] 
**Followers** | Pointer to [**PlaylistObjectFollowers**](PlaylistObjectFollowers.md) |  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the playlist.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the playlist.  | [optional] 
**Images** | Pointer to [**[]ImageObject**](ImageObject.md) | Images for the playlist. The array may be empty or contain up to three images. The images are returned by size in descending order. See [Working with Playlists](/documentation/general/guides/working-with-playlists/). _**Note**: If returned, the source URL for the image (&#x60;url&#x60;) is temporary and will expire in less than a day._  | [optional] 
**Name** | Pointer to **string** | The name of the playlist.  | [optional] 
**Owner** | Pointer to [**PlaylistObjectOwner**](PlaylistObjectOwner.md) |  | [optional] 
**Public** | Pointer to **bool** | The playlist&#39;s public/private status: &#x60;true&#x60; the playlist is public, &#x60;false&#x60; the playlist is private, &#x60;null&#x60; the playlist status is not relevant. For more about public/private status, see [Working with Playlists](/documentation/general/guides/working-with-playlists/)  | [optional] 
**SnapshotId** | Pointer to **string** | The version identifier for the current playlist. Can be supplied in other requests to target a specific playlist version  | [optional] 
**Tracks** | Pointer to [**PlaylistObjectTracks**](PlaylistObjectTracks.md) |  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;playlist\&quot;  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the playlist.  | [optional] 

## Methods

### NewPlaylistObject

`func NewPlaylistObject() *PlaylistObject`

NewPlaylistObject instantiates a new PlaylistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistObjectWithDefaults

`func NewPlaylistObjectWithDefaults() *PlaylistObject`

NewPlaylistObjectWithDefaults instantiates a new PlaylistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollaborative

`func (o *PlaylistObject) GetCollaborative() bool`

GetCollaborative returns the Collaborative field if non-nil, zero value otherwise.

### GetCollaborativeOk

`func (o *PlaylistObject) GetCollaborativeOk() (*bool, bool)`

GetCollaborativeOk returns a tuple with the Collaborative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborative

`func (o *PlaylistObject) SetCollaborative(v bool)`

SetCollaborative sets Collaborative field to given value.

### HasCollaborative

`func (o *PlaylistObject) HasCollaborative() bool`

HasCollaborative returns a boolean if a field has been set.

### GetDescription

`func (o *PlaylistObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlaylistObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlaylistObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlaylistObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlaylistObject) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlaylistObject) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternalUrls

`func (o *PlaylistObject) GetExternalUrls() PlaylistObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PlaylistObject) GetExternalUrlsOk() (*PlaylistObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PlaylistObject) SetExternalUrls(v PlaylistObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PlaylistObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *PlaylistObject) GetFollowers() PlaylistObjectFollowers`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PlaylistObject) GetFollowersOk() (*PlaylistObjectFollowers, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PlaylistObject) SetFollowers(v PlaylistObjectFollowers)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PlaylistObject) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetHref

`func (o *PlaylistObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlaylistObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlaylistObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlaylistObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PlaylistObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaylistObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *PlaylistObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PlaylistObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PlaylistObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.

### HasImages

`func (o *PlaylistObject) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetName

`func (o *PlaylistObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlaylistObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlaylistObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlaylistObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *PlaylistObject) GetOwner() PlaylistObjectOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PlaylistObject) GetOwnerOk() (*PlaylistObjectOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PlaylistObject) SetOwner(v PlaylistObjectOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PlaylistObject) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPublic

`func (o *PlaylistObject) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PlaylistObject) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PlaylistObject) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *PlaylistObject) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSnapshotId

`func (o *PlaylistObject) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *PlaylistObject) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *PlaylistObject) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *PlaylistObject) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetTracks

`func (o *PlaylistObject) GetTracks() PlaylistObjectTracks`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *PlaylistObject) GetTracksOk() (*PlaylistObjectTracks, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *PlaylistObject) SetTracks(v PlaylistObjectTracks)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *PlaylistObject) HasTracks() bool`

HasTracks returns a boolean if a field has been set.

### GetType

`func (o *PlaylistObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaylistObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaylistObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaylistObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *PlaylistObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PlaylistObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PlaylistObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PlaylistObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


