# SaveAlbumsUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: &#x60;[\&quot;4iV5W9uYEdYUVa79Axb7Rh\&quot;, \&quot;1301WleyT98MSxVHPZCA6M\&quot;]&#x60;&lt;br/&gt;A maximum of 50 items can be specified in one request. _**Note**: if the &#x60;ids&#x60; parameter is present in the query string, any IDs listed here in the body will be ignored._  | [optional] 

## Methods

### NewSaveAlbumsUserRequest

`func NewSaveAlbumsUserRequest() *SaveAlbumsUserRequest`

NewSaveAlbumsUserRequest instantiates a new SaveAlbumsUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveAlbumsUserRequestWithDefaults

`func NewSaveAlbumsUserRequestWithDefaults() *SaveAlbumsUserRequest`

NewSaveAlbumsUserRequestWithDefaults instantiates a new SaveAlbumsUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *SaveAlbumsUserRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SaveAlbumsUserRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SaveAlbumsUserRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *SaveAlbumsUserRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


