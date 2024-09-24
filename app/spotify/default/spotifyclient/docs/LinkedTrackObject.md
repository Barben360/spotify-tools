# LinkedTrackObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this track.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;track\&quot;.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 

## Methods

### NewLinkedTrackObject

`func NewLinkedTrackObject() *LinkedTrackObject`

NewLinkedTrackObject instantiates a new LinkedTrackObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedTrackObjectWithDefaults

`func NewLinkedTrackObjectWithDefaults() *LinkedTrackObject`

NewLinkedTrackObjectWithDefaults instantiates a new LinkedTrackObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *LinkedTrackObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *LinkedTrackObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *LinkedTrackObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *LinkedTrackObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *LinkedTrackObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LinkedTrackObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LinkedTrackObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LinkedTrackObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *LinkedTrackObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedTrackObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedTrackObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedTrackObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LinkedTrackObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedTrackObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedTrackObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedTrackObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *LinkedTrackObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *LinkedTrackObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *LinkedTrackObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *LinkedTrackObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


