# PlaylistOwnerObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**PublicUserObjectExternalUrls**](PublicUserObjectExternalUrls.md) |  | [optional] 
**Followers** | Pointer to [**PublicUserObjectFollowers**](PublicUserObjectFollowers.md) |  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint for this user.  | [optional] 
**Id** | Pointer to **string** | The [Spotify user ID](/documentation/web-api/#spotify-uris-and-ids) for this user.  | [optional] 
**Type** | Pointer to **string** | The object type.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for this user.  | [optional] 
**DisplayName** | Pointer to **NullableString** | The name displayed on the user&#39;s profile. &#x60;null&#x60; if not available.  | [optional] 

## Methods

### NewPlaylistOwnerObject

`func NewPlaylistOwnerObject() *PlaylistOwnerObject`

NewPlaylistOwnerObject instantiates a new PlaylistOwnerObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistOwnerObjectWithDefaults

`func NewPlaylistOwnerObjectWithDefaults() *PlaylistOwnerObject`

NewPlaylistOwnerObjectWithDefaults instantiates a new PlaylistOwnerObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *PlaylistOwnerObject) GetExternalUrls() PublicUserObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PlaylistOwnerObject) GetExternalUrlsOk() (*PublicUserObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PlaylistOwnerObject) SetExternalUrls(v PublicUserObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PlaylistOwnerObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *PlaylistOwnerObject) GetFollowers() PublicUserObjectFollowers`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PlaylistOwnerObject) GetFollowersOk() (*PublicUserObjectFollowers, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PlaylistOwnerObject) SetFollowers(v PublicUserObjectFollowers)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PlaylistOwnerObject) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetHref

`func (o *PlaylistOwnerObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlaylistOwnerObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlaylistOwnerObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlaylistOwnerObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PlaylistOwnerObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistOwnerObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistOwnerObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaylistOwnerObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PlaylistOwnerObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaylistOwnerObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaylistOwnerObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaylistOwnerObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *PlaylistOwnerObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PlaylistOwnerObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PlaylistOwnerObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PlaylistOwnerObject) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetDisplayName

`func (o *PlaylistOwnerObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PlaylistOwnerObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PlaylistOwnerObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PlaylistOwnerObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *PlaylistOwnerObject) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *PlaylistOwnerObject) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


