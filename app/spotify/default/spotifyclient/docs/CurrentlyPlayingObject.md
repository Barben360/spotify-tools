# CurrentlyPlayingObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**ContextObject**](ContextObject.md) | A Context Object. Can be &#x60;null&#x60;. | [optional] 
**Timestamp** | Pointer to **int64** | Unix Millisecond Timestamp when data was fetched | [optional] 
**ProgressMs** | Pointer to **int32** | Progress into the currently playing track or episode. Can be &#x60;null&#x60;. | [optional] 
**IsPlaying** | Pointer to **bool** | If something is currently playing, return &#x60;true&#x60;. | [optional] 
**Item** | Pointer to [**QueueObjectCurrentlyPlaying**](QueueObjectCurrentlyPlaying.md) |  | [optional] 
**CurrentlyPlayingType** | Pointer to **string** | The object type of the currently playing item. Can be one of &#x60;track&#x60;, &#x60;episode&#x60;, &#x60;ad&#x60; or &#x60;unknown&#x60;.  | [optional] 
**Actions** | Pointer to [**DisallowsObject**](DisallowsObject.md) | Allows to update the user interface based on which playback actions are available within the current context.  | [optional] 

## Methods

### NewCurrentlyPlayingObject

`func NewCurrentlyPlayingObject() *CurrentlyPlayingObject`

NewCurrentlyPlayingObject instantiates a new CurrentlyPlayingObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentlyPlayingObjectWithDefaults

`func NewCurrentlyPlayingObjectWithDefaults() *CurrentlyPlayingObject`

NewCurrentlyPlayingObjectWithDefaults instantiates a new CurrentlyPlayingObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *CurrentlyPlayingObject) GetContext() ContextObject`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CurrentlyPlayingObject) GetContextOk() (*ContextObject, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CurrentlyPlayingObject) SetContext(v ContextObject)`

SetContext sets Context field to given value.

### HasContext

`func (o *CurrentlyPlayingObject) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetTimestamp

`func (o *CurrentlyPlayingObject) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CurrentlyPlayingObject) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CurrentlyPlayingObject) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CurrentlyPlayingObject) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetProgressMs

`func (o *CurrentlyPlayingObject) GetProgressMs() int32`

GetProgressMs returns the ProgressMs field if non-nil, zero value otherwise.

### GetProgressMsOk

`func (o *CurrentlyPlayingObject) GetProgressMsOk() (*int32, bool)`

GetProgressMsOk returns a tuple with the ProgressMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMs

`func (o *CurrentlyPlayingObject) SetProgressMs(v int32)`

SetProgressMs sets ProgressMs field to given value.

### HasProgressMs

`func (o *CurrentlyPlayingObject) HasProgressMs() bool`

HasProgressMs returns a boolean if a field has been set.

### GetIsPlaying

`func (o *CurrentlyPlayingObject) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *CurrentlyPlayingObject) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *CurrentlyPlayingObject) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *CurrentlyPlayingObject) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetItem

`func (o *CurrentlyPlayingObject) GetItem() QueueObjectCurrentlyPlaying`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CurrentlyPlayingObject) GetItemOk() (*QueueObjectCurrentlyPlaying, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CurrentlyPlayingObject) SetItem(v QueueObjectCurrentlyPlaying)`

SetItem sets Item field to given value.

### HasItem

`func (o *CurrentlyPlayingObject) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetCurrentlyPlayingType

`func (o *CurrentlyPlayingObject) GetCurrentlyPlayingType() string`

GetCurrentlyPlayingType returns the CurrentlyPlayingType field if non-nil, zero value otherwise.

### GetCurrentlyPlayingTypeOk

`func (o *CurrentlyPlayingObject) GetCurrentlyPlayingTypeOk() (*string, bool)`

GetCurrentlyPlayingTypeOk returns a tuple with the CurrentlyPlayingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyPlayingType

`func (o *CurrentlyPlayingObject) SetCurrentlyPlayingType(v string)`

SetCurrentlyPlayingType sets CurrentlyPlayingType field to given value.

### HasCurrentlyPlayingType

`func (o *CurrentlyPlayingObject) HasCurrentlyPlayingType() bool`

HasCurrentlyPlayingType returns a boolean if a field has been set.

### GetActions

`func (o *CurrentlyPlayingObject) GetActions() DisallowsObject`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CurrentlyPlayingObject) GetActionsOk() (*DisallowsObject, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CurrentlyPlayingObject) SetActions(v DisallowsObject)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CurrentlyPlayingObject) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


