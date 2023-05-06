# StartAUsersPlaybackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextUri** | Pointer to **string** | Optional. Spotify URI of the context to play. Valid contexts are albums, artists &amp; playlists. &#x60;{context_uri:\&quot;spotify:album:1Je1IMUlBXcx1Fz0WE7oPT\&quot;}&#x60;  | [optional] 
**Uris** | Pointer to **[]string** | Optional. A JSON array of the Spotify track URIs to play. For example: &#x60;{\&quot;uris\&quot;: [\&quot;spotify:track:4iV5W9uYEdYUVa79Axb7Rh\&quot;, \&quot;spotify:track:1301WleyT98MSxVHPZCA6M\&quot;]}&#x60;  | [optional] 
**Offset** | Pointer to **map[string]interface{}** | Optional. Indicates from where in the context playback should start. Only available when context_uri corresponds to an album or playlist object \&quot;position\&quot; is zero based and can’t be negative. Example: &#x60;\&quot;offset\&quot;: {\&quot;position\&quot;: 5}&#x60; \&quot;uri\&quot; is a string representing the uri of the item to start at. Example: &#x60;\&quot;offset\&quot;: {\&quot;uri\&quot;: \&quot;spotify:track:1301WleyT98MSxVHPZCA6M\&quot;}&#x60;  | [optional] 
**PositionMs** | Pointer to **int32** | Indicates from what position to start playback. Must be a positive number. Passing in a position that is greater than the length of the track will cause the player to start playing the next song.  | [optional] 

## Methods

### NewStartAUsersPlaybackRequest

`func NewStartAUsersPlaybackRequest() *StartAUsersPlaybackRequest`

NewStartAUsersPlaybackRequest instantiates a new StartAUsersPlaybackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartAUsersPlaybackRequestWithDefaults

`func NewStartAUsersPlaybackRequestWithDefaults() *StartAUsersPlaybackRequest`

NewStartAUsersPlaybackRequestWithDefaults instantiates a new StartAUsersPlaybackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextUri

`func (o *StartAUsersPlaybackRequest) GetContextUri() string`

GetContextUri returns the ContextUri field if non-nil, zero value otherwise.

### GetContextUriOk

`func (o *StartAUsersPlaybackRequest) GetContextUriOk() (*string, bool)`

GetContextUriOk returns a tuple with the ContextUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextUri

`func (o *StartAUsersPlaybackRequest) SetContextUri(v string)`

SetContextUri sets ContextUri field to given value.

### HasContextUri

`func (o *StartAUsersPlaybackRequest) HasContextUri() bool`

HasContextUri returns a boolean if a field has been set.

### GetUris

`func (o *StartAUsersPlaybackRequest) GetUris() []string`

GetUris returns the Uris field if non-nil, zero value otherwise.

### GetUrisOk

`func (o *StartAUsersPlaybackRequest) GetUrisOk() (*[]string, bool)`

GetUrisOk returns a tuple with the Uris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUris

`func (o *StartAUsersPlaybackRequest) SetUris(v []string)`

SetUris sets Uris field to given value.

### HasUris

`func (o *StartAUsersPlaybackRequest) HasUris() bool`

HasUris returns a boolean if a field has been set.

### GetOffset

`func (o *StartAUsersPlaybackRequest) GetOffset() map[string]interface{}`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StartAUsersPlaybackRequest) GetOffsetOk() (*map[string]interface{}, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StartAUsersPlaybackRequest) SetOffset(v map[string]interface{})`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StartAUsersPlaybackRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPositionMs

`func (o *StartAUsersPlaybackRequest) GetPositionMs() int32`

GetPositionMs returns the PositionMs field if non-nil, zero value otherwise.

### GetPositionMsOk

`func (o *StartAUsersPlaybackRequest) GetPositionMsOk() (*int32, bool)`

GetPositionMsOk returns a tuple with the PositionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionMs

`func (o *StartAUsersPlaybackRequest) SetPositionMs(v int32)`

SetPositionMs sets PositionMs field to given value.

### HasPositionMs

`func (o *StartAUsersPlaybackRequest) HasPositionMs() bool`

HasPositionMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


