# PlaylistObjectOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**PublicUserObjectExternalUrls**](PublicUserObjectExternalUrls.md) |  | [optional] 
**Followers** | Pointer to [**PublicUserObjectFollowers**](PublicUserObjectFollowers.md) |  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint for this user.  | [optional] 
**Id** | Pointer to **string** | The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user.  | [optional] 
**Type** | Pointer to **string** | The object type.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this user.  | [optional] 
**DisplayName** | Pointer to **NullableString** | The name displayed on the user&#39;s profile. &#x60;null&#x60; if not available.  | [optional] 

## Methods

### NewPlaylistObjectOwner

`func NewPlaylistObjectOwner() *PlaylistObjectOwner`

NewPlaylistObjectOwner instantiates a new PlaylistObjectOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistObjectOwnerWithDefaults

`func NewPlaylistObjectOwnerWithDefaults() *PlaylistObjectOwner`

NewPlaylistObjectOwnerWithDefaults instantiates a new PlaylistObjectOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *PlaylistObjectOwner) GetExternalUrls() PublicUserObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PlaylistObjectOwner) GetExternalUrlsOk() (*PublicUserObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PlaylistObjectOwner) SetExternalUrls(v PublicUserObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PlaylistObjectOwner) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *PlaylistObjectOwner) GetFollowers() PublicUserObjectFollowers`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PlaylistObjectOwner) GetFollowersOk() (*PublicUserObjectFollowers, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PlaylistObjectOwner) SetFollowers(v PublicUserObjectFollowers)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PlaylistObjectOwner) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetHref

`func (o *PlaylistObjectOwner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlaylistObjectOwner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlaylistObjectOwner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlaylistObjectOwner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PlaylistObjectOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaylistObjectOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaylistObjectOwner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaylistObjectOwner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PlaylistObjectOwner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaylistObjectOwner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaylistObjectOwner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaylistObjectOwner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *PlaylistObjectOwner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PlaylistObjectOwner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PlaylistObjectOwner) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PlaylistObjectOwner) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetDisplayName

`func (o *PlaylistObjectOwner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PlaylistObjectOwner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PlaylistObjectOwner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PlaylistObjectOwner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *PlaylistObjectOwner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *PlaylistObjectOwner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


