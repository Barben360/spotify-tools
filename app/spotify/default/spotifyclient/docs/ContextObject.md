# ContextObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The object type, e.g. \&quot;artist\&quot;, \&quot;playlist\&quot;, \&quot;album\&quot;, \&quot;show\&quot;.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track. | [optional] 
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | External URLs for this context. | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the context.  | [optional] 

## Methods

### NewContextObject

`func NewContextObject() *ContextObject`

NewContextObject instantiates a new ContextObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextObjectWithDefaults

`func NewContextObjectWithDefaults() *ContextObject`

NewContextObjectWithDefaults instantiates a new ContextObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContextObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContextObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContextObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContextObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ContextObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ContextObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ContextObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ContextObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetExternalUrls

`func (o *ContextObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *ContextObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *ContextObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *ContextObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetUri

`func (o *ContextObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ContextObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ContextObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ContextObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


