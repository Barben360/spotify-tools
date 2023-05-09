# SaveEpisodesUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | A JSON array of the [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). &lt;br/&gt;A maximum of 50 items can be specified in one request. _**Note**: if the &#x60;ids&#x60; parameter is present in the query string, any IDs listed here in the body will be ignored._  | 

## Methods

### NewSaveEpisodesUserRequest

`func NewSaveEpisodesUserRequest(ids []string, ) *SaveEpisodesUserRequest`

NewSaveEpisodesUserRequest instantiates a new SaveEpisodesUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveEpisodesUserRequestWithDefaults

`func NewSaveEpisodesUserRequestWithDefaults() *SaveEpisodesUserRequest`

NewSaveEpisodesUserRequestWithDefaults instantiates a new SaveEpisodesUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *SaveEpisodesUserRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SaveEpisodesUserRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SaveEpisodesUserRequest) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


