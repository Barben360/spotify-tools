# FollowArtistsUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | A JSON array of the artist or user [Spotify IDs](/documentation/web-api/concepts/spotify-uris-ids). For example: &#x60;{ids:[\&quot;74ASZWbe4lXaubB36ztrGX\&quot;, \&quot;08td7MxkoHQkXnWAYD8d6Q\&quot;]}&#x60;. A maximum of 50 IDs can be sent in one request. _**Note**: if the &#x60;ids&#x60; parameter is present in the query string, any IDs listed here in the body will be ignored._  | 

## Methods

### NewFollowArtistsUsersRequest

`func NewFollowArtistsUsersRequest(ids []string, ) *FollowArtistsUsersRequest`

NewFollowArtistsUsersRequest instantiates a new FollowArtistsUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowArtistsUsersRequestWithDefaults

`func NewFollowArtistsUsersRequestWithDefaults() *FollowArtistsUsersRequest`

NewFollowArtistsUsersRequestWithDefaults instantiates a new FollowArtistsUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *FollowArtistsUsersRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *FollowArtistsUsersRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *FollowArtistsUsersRequest) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


