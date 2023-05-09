# PlayHistoryObjectContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The object type, e.g. \&quot;artist\&quot;, \&quot;playlist\&quot;, \&quot;album\&quot;, \&quot;show\&quot;.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track. | [optional] 
**ExternalUrls** | Pointer to [**ContextObjectExternalUrls**](ContextObjectExternalUrls.md) |  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the context.  | [optional] 

## Methods

### NewPlayHistoryObjectContext

`func NewPlayHistoryObjectContext() *PlayHistoryObjectContext`

NewPlayHistoryObjectContext instantiates a new PlayHistoryObjectContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayHistoryObjectContextWithDefaults

`func NewPlayHistoryObjectContextWithDefaults() *PlayHistoryObjectContext`

NewPlayHistoryObjectContextWithDefaults instantiates a new PlayHistoryObjectContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PlayHistoryObjectContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayHistoryObjectContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayHistoryObjectContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlayHistoryObjectContext) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *PlayHistoryObjectContext) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlayHistoryObjectContext) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlayHistoryObjectContext) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlayHistoryObjectContext) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetExternalUrls

`func (o *PlayHistoryObjectContext) GetExternalUrls() ContextObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PlayHistoryObjectContext) GetExternalUrlsOk() (*ContextObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PlayHistoryObjectContext) SetExternalUrls(v ContextObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PlayHistoryObjectContext) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetUri

`func (o *PlayHistoryObjectContext) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PlayHistoryObjectContext) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PlayHistoryObjectContext) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PlayHistoryObjectContext) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


