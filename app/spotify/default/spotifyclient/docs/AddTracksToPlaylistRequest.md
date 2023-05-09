# AddTracksToPlaylistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uris** | Pointer to **[]string** | A JSON array of the [Spotify URIs](/documentation/web-api/concepts/spotify-uris-ids) to add. For example: &#x60;{\&quot;uris\&quot;: [\&quot;spotify:track:4iV5W9uYEdYUVa79Axb7Rh\&quot;,\&quot;spotify:track:1301WleyT98MSxVHPZCA6M\&quot;, \&quot;spotify:episode:512ojhOuo1ktJprKbVcKyQ\&quot;]}&#x60;&lt;br/&gt;A maximum of 100 items can be added in one request. _**Note**: if the &#x60;uris&#x60; parameter is present in the query string, any URIs listed here in the body will be ignored._  | [optional] 
**Position** | Pointer to **int32** | The position to insert the items, a zero-based index. For example, to insert the items in the first position: &#x60;position&#x3D;0&#x60; ; to insert the items in the third position: &#x60;position&#x3D;2&#x60;. If omitted, the items will be appended to the playlist. Items are added in the order they appear in the uris array. For example: &#x60;{\&quot;uris\&quot;: [\&quot;spotify:track:4iV5W9uYEdYUVa79Axb7Rh\&quot;,\&quot;spotify:track:1301WleyT98MSxVHPZCA6M\&quot;], \&quot;position\&quot;: 3}&#x60;  | [optional] 

## Methods

### NewAddTracksToPlaylistRequest

`func NewAddTracksToPlaylistRequest() *AddTracksToPlaylistRequest`

NewAddTracksToPlaylistRequest instantiates a new AddTracksToPlaylistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTracksToPlaylistRequestWithDefaults

`func NewAddTracksToPlaylistRequestWithDefaults() *AddTracksToPlaylistRequest`

NewAddTracksToPlaylistRequestWithDefaults instantiates a new AddTracksToPlaylistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUris

`func (o *AddTracksToPlaylistRequest) GetUris() []string`

GetUris returns the Uris field if non-nil, zero value otherwise.

### GetUrisOk

`func (o *AddTracksToPlaylistRequest) GetUrisOk() (*[]string, bool)`

GetUrisOk returns a tuple with the Uris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUris

`func (o *AddTracksToPlaylistRequest) SetUris(v []string)`

SetUris sets Uris field to given value.

### HasUris

`func (o *AddTracksToPlaylistRequest) HasUris() bool`

HasUris returns a boolean if a field has been set.

### GetPosition

`func (o *AddTracksToPlaylistRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AddTracksToPlaylistRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AddTracksToPlaylistRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AddTracksToPlaylistRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


