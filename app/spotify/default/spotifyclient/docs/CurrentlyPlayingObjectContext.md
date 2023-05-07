# CurrentlyPlayingObjectContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The object type, e.g. \&quot;artist\&quot;, \&quot;playlist\&quot;, \&quot;album\&quot;, \&quot;show\&quot;.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track. | [optional] 
**ExternalUrls** | Pointer to [**ContextObjectExternalUrls**](ContextObjectExternalUrls.md) |  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the context.  | [optional] 

## Methods

### NewCurrentlyPlayingObjectContext

`func NewCurrentlyPlayingObjectContext() *CurrentlyPlayingObjectContext`

NewCurrentlyPlayingObjectContext instantiates a new CurrentlyPlayingObjectContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentlyPlayingObjectContextWithDefaults

`func NewCurrentlyPlayingObjectContextWithDefaults() *CurrentlyPlayingObjectContext`

NewCurrentlyPlayingObjectContextWithDefaults instantiates a new CurrentlyPlayingObjectContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CurrentlyPlayingObjectContext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CurrentlyPlayingObjectContext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CurrentlyPlayingObjectContext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CurrentlyPlayingObjectContext) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *CurrentlyPlayingObjectContext) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CurrentlyPlayingObjectContext) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CurrentlyPlayingObjectContext) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CurrentlyPlayingObjectContext) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetExternalUrls

`func (o *CurrentlyPlayingObjectContext) GetExternalUrls() ContextObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *CurrentlyPlayingObjectContext) GetExternalUrlsOk() (*ContextObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *CurrentlyPlayingObjectContext) SetExternalUrls(v ContextObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *CurrentlyPlayingObjectContext) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetUri

`func (o *CurrentlyPlayingObjectContext) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CurrentlyPlayingObjectContext) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CurrentlyPlayingObjectContext) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CurrentlyPlayingObjectContext) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


