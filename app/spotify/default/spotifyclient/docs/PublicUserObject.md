# PublicUserObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The name displayed on the user&#39;s profile. &#x60;null&#x60; if not available.  | [optional] 
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known public external URLs for this user.  | [optional] 
**Followers** | Pointer to [**FollowersObject**](FollowersObject.md) | Information about the followers of this user.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint for this user.  | [optional] 
**Id** | Pointer to **string** | The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for this user.  | [optional] 
**Images** | Pointer to [**[]ImageObject**](ImageObject.md) | The user&#39;s profile image.  | [optional] 
**Type** | Pointer to **string** | The object type.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for this user.  | [optional] 

## Methods

### NewPublicUserObject

`func NewPublicUserObject() *PublicUserObject`

NewPublicUserObject instantiates a new PublicUserObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUserObjectWithDefaults

`func NewPublicUserObjectWithDefaults() *PublicUserObject`

NewPublicUserObjectWithDefaults instantiates a new PublicUserObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *PublicUserObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PublicUserObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PublicUserObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PublicUserObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *PublicUserObject) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *PublicUserObject) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetExternalUrls

`func (o *PublicUserObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PublicUserObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PublicUserObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PublicUserObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *PublicUserObject) GetFollowers() FollowersObject`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PublicUserObject) GetFollowersOk() (*FollowersObject, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PublicUserObject) SetFollowers(v FollowersObject)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PublicUserObject) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetHref

`func (o *PublicUserObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PublicUserObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PublicUserObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PublicUserObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PublicUserObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicUserObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicUserObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicUserObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *PublicUserObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PublicUserObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PublicUserObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.

### HasImages

`func (o *PublicUserObject) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetType

`func (o *PublicUserObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicUserObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicUserObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublicUserObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *PublicUserObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PublicUserObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PublicUserObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PublicUserObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


