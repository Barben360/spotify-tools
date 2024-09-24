# PlaylistUserObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known public external URLs for this user.  | [optional] 
**Followers** | Pointer to [**FollowersObject**](FollowersObject.md) | Information about the followers of this user.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint for this user.  | [optional] 
**Id** | Pointer to **string** | The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user.  | [optional] 
**Type** | Pointer to **string** | The object type.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this user.  | [optional] 

## Methods

### NewPlaylistUserObject

`func NewPlaylistUserObject() *PlaylistUserObject`

NewPlaylistUserObject instantiates a new PlaylistUserObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistUserObjectWithDefaults

`func NewPlaylistUserObjectWithDefaults() *PlaylistUserObject`

NewPlaylistUserObjectWithDefaults instantiates a new PlaylistUserObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *PlaylistUserObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PlaylistUserObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PlaylistUserObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PlaylistUserObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *PlaylistUserObject) GetFollowers() FollowersObject`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PlaylistUserObject) GetFollowersOk() (*FollowersObject, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PlaylistUserObject) SetFollowers(v FollowersObject)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PlaylistUserObject) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetHref

`func (o *PlaylistUserObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlaylistUserObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlaylistUserObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlaylistUserObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PlaylistUserObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistUserObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistUserObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaylistUserObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PlaylistUserObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaylistUserObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaylistUserObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaylistUserObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *PlaylistUserObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PlaylistUserObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PlaylistUserObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PlaylistUserObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


