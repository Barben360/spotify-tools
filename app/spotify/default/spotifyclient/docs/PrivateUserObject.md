# PrivateUserObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The country of the user, as set in the user&#39;s account profile. An [ISO 3166-1 alpha-2 country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _This field is only available when the current user has granted access to the [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._  | [optional] 
**DisplayName** | Pointer to **string** | The name displayed on the user&#39;s profile. &#x60;null&#x60; if not available.  | [optional] 
**Email** | Pointer to **string** | The user&#39;s email address, as entered by the user when creating their account. _**Important!** This email address is unverified; there is no proof that it actually belongs to the user._ _This field is only available when the current user has granted access to the [user-read-email](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._  | [optional] 
**ExplicitContent** | Pointer to [**ExplicitContentSettingsObject**](ExplicitContentSettingsObject.md) | The user&#39;s explicit content settings. _This field is only available when the current user has granted access to the [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._  | [optional] 
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this user. | [optional] 
**Followers** | Pointer to [**FollowersObject**](FollowersObject.md) | Information about the followers of the user. | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint for this user.  | [optional] 
**Id** | Pointer to **string** | The [Spotify user ID](/documentation/web-api/concepts/spotify-uris-ids) for the user.  | [optional] 
**Images** | Pointer to [**[]ImageObject**](ImageObject.md) | The user&#39;s profile image. | [optional] 
**Product** | Pointer to **string** | The user&#39;s Spotify subscription level: \&quot;premium\&quot;, \&quot;free\&quot;, etc. (The subscription level \&quot;open\&quot; can be considered the same as \&quot;free\&quot;.) _This field is only available when the current user has granted access to the [user-read-private](/documentation/web-api/concepts/scopes/#list-of-scopes) scope._  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;user\&quot;  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the user.  | [optional] 

## Methods

### NewPrivateUserObject

`func NewPrivateUserObject() *PrivateUserObject`

NewPrivateUserObject instantiates a new PrivateUserObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateUserObjectWithDefaults

`func NewPrivateUserObjectWithDefaults() *PrivateUserObject`

NewPrivateUserObjectWithDefaults instantiates a new PrivateUserObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *PrivateUserObject) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PrivateUserObject) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PrivateUserObject) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PrivateUserObject) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayName

`func (o *PrivateUserObject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PrivateUserObject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PrivateUserObject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PrivateUserObject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *PrivateUserObject) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PrivateUserObject) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PrivateUserObject) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PrivateUserObject) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExplicitContent

`func (o *PrivateUserObject) GetExplicitContent() ExplicitContentSettingsObject`

GetExplicitContent returns the ExplicitContent field if non-nil, zero value otherwise.

### GetExplicitContentOk

`func (o *PrivateUserObject) GetExplicitContentOk() (*ExplicitContentSettingsObject, bool)`

GetExplicitContentOk returns a tuple with the ExplicitContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitContent

`func (o *PrivateUserObject) SetExplicitContent(v ExplicitContentSettingsObject)`

SetExplicitContent sets ExplicitContent field to given value.

### HasExplicitContent

`func (o *PrivateUserObject) HasExplicitContent() bool`

HasExplicitContent returns a boolean if a field has been set.

### GetExternalUrls

`func (o *PrivateUserObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PrivateUserObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PrivateUserObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PrivateUserObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *PrivateUserObject) GetFollowers() FollowersObject`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *PrivateUserObject) GetFollowersOk() (*FollowersObject, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *PrivateUserObject) SetFollowers(v FollowersObject)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *PrivateUserObject) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetHref

`func (o *PrivateUserObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PrivateUserObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PrivateUserObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PrivateUserObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PrivateUserObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateUserObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateUserObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateUserObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *PrivateUserObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PrivateUserObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PrivateUserObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.

### HasImages

`func (o *PrivateUserObject) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetProduct

`func (o *PrivateUserObject) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PrivateUserObject) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PrivateUserObject) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PrivateUserObject) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetType

`func (o *PrivateUserObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateUserObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateUserObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateUserObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *PrivateUserObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PrivateUserObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PrivateUserObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PrivateUserObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


