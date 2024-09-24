# CurrentlyPlayingContextObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**DeviceObject**](DeviceObject.md) | The device that is currently active.  | [optional] 
**RepeatState** | Pointer to **string** | off, track, context | [optional] 
**ShuffleState** | Pointer to **bool** | If shuffle is on or off. | [optional] 
**Context** | Pointer to [**ContextObject**](ContextObject.md) | A Context Object. Can be &#x60;null&#x60;. | [optional] 
**Timestamp** | Pointer to **int64** | Unix Millisecond Timestamp when playback state was last changed (play, pause, skip, scrub, new song, etc.). | [optional] 
**ProgressMs** | Pointer to **int32** | Progress into the currently playing track or episode. Can be &#x60;null&#x60;. | [optional] 
**IsPlaying** | Pointer to **bool** | If something is currently playing, return &#x60;true&#x60;. | [optional] 
**Item** | Pointer to [**QueueObjectCurrentlyPlaying**](QueueObjectCurrentlyPlaying.md) |  | [optional] 
**CurrentlyPlayingType** | Pointer to **string** | The object type of the currently playing item. Can be one of &#x60;track&#x60;, &#x60;episode&#x60;, &#x60;ad&#x60; or &#x60;unknown&#x60;.  | [optional] 
**Actions** | Pointer to [**DisallowsObject**](DisallowsObject.md) | Allows to update the user interface based on which playback actions are available within the current context.  | [optional] 

## Methods

### NewCurrentlyPlayingContextObject

`func NewCurrentlyPlayingContextObject() *CurrentlyPlayingContextObject`

NewCurrentlyPlayingContextObject instantiates a new CurrentlyPlayingContextObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentlyPlayingContextObjectWithDefaults

`func NewCurrentlyPlayingContextObjectWithDefaults() *CurrentlyPlayingContextObject`

NewCurrentlyPlayingContextObjectWithDefaults instantiates a new CurrentlyPlayingContextObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *CurrentlyPlayingContextObject) GetDevice() DeviceObject`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CurrentlyPlayingContextObject) GetDeviceOk() (*DeviceObject, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CurrentlyPlayingContextObject) SetDevice(v DeviceObject)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CurrentlyPlayingContextObject) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetRepeatState

`func (o *CurrentlyPlayingContextObject) GetRepeatState() string`

GetRepeatState returns the RepeatState field if non-nil, zero value otherwise.

### GetRepeatStateOk

`func (o *CurrentlyPlayingContextObject) GetRepeatStateOk() (*string, bool)`

GetRepeatStateOk returns a tuple with the RepeatState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatState

`func (o *CurrentlyPlayingContextObject) SetRepeatState(v string)`

SetRepeatState sets RepeatState field to given value.

### HasRepeatState

`func (o *CurrentlyPlayingContextObject) HasRepeatState() bool`

HasRepeatState returns a boolean if a field has been set.

### GetShuffleState

`func (o *CurrentlyPlayingContextObject) GetShuffleState() bool`

GetShuffleState returns the ShuffleState field if non-nil, zero value otherwise.

### GetShuffleStateOk

`func (o *CurrentlyPlayingContextObject) GetShuffleStateOk() (*bool, bool)`

GetShuffleStateOk returns a tuple with the ShuffleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShuffleState

`func (o *CurrentlyPlayingContextObject) SetShuffleState(v bool)`

SetShuffleState sets ShuffleState field to given value.

### HasShuffleState

`func (o *CurrentlyPlayingContextObject) HasShuffleState() bool`

HasShuffleState returns a boolean if a field has been set.

### GetContext

`func (o *CurrentlyPlayingContextObject) GetContext() ContextObject`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CurrentlyPlayingContextObject) GetContextOk() (*ContextObject, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CurrentlyPlayingContextObject) SetContext(v ContextObject)`

SetContext sets Context field to given value.

### HasContext

`func (o *CurrentlyPlayingContextObject) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetTimestamp

`func (o *CurrentlyPlayingContextObject) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CurrentlyPlayingContextObject) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CurrentlyPlayingContextObject) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CurrentlyPlayingContextObject) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetProgressMs

`func (o *CurrentlyPlayingContextObject) GetProgressMs() int32`

GetProgressMs returns the ProgressMs field if non-nil, zero value otherwise.

### GetProgressMsOk

`func (o *CurrentlyPlayingContextObject) GetProgressMsOk() (*int32, bool)`

GetProgressMsOk returns a tuple with the ProgressMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressMs

`func (o *CurrentlyPlayingContextObject) SetProgressMs(v int32)`

SetProgressMs sets ProgressMs field to given value.

### HasProgressMs

`func (o *CurrentlyPlayingContextObject) HasProgressMs() bool`

HasProgressMs returns a boolean if a field has been set.

### GetIsPlaying

`func (o *CurrentlyPlayingContextObject) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *CurrentlyPlayingContextObject) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *CurrentlyPlayingContextObject) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *CurrentlyPlayingContextObject) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetItem

`func (o *CurrentlyPlayingContextObject) GetItem() QueueObjectCurrentlyPlaying`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CurrentlyPlayingContextObject) GetItemOk() (*QueueObjectCurrentlyPlaying, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CurrentlyPlayingContextObject) SetItem(v QueueObjectCurrentlyPlaying)`

SetItem sets Item field to given value.

### HasItem

`func (o *CurrentlyPlayingContextObject) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetCurrentlyPlayingType

`func (o *CurrentlyPlayingContextObject) GetCurrentlyPlayingType() string`

GetCurrentlyPlayingType returns the CurrentlyPlayingType field if non-nil, zero value otherwise.

### GetCurrentlyPlayingTypeOk

`func (o *CurrentlyPlayingContextObject) GetCurrentlyPlayingTypeOk() (*string, bool)`

GetCurrentlyPlayingTypeOk returns a tuple with the CurrentlyPlayingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyPlayingType

`func (o *CurrentlyPlayingContextObject) SetCurrentlyPlayingType(v string)`

SetCurrentlyPlayingType sets CurrentlyPlayingType field to given value.

### HasCurrentlyPlayingType

`func (o *CurrentlyPlayingContextObject) HasCurrentlyPlayingType() bool`

HasCurrentlyPlayingType returns a boolean if a field has been set.

### GetActions

`func (o *CurrentlyPlayingContextObject) GetActions() DisallowsObject`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CurrentlyPlayingContextObject) GetActionsOk() (*DisallowsObject, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CurrentlyPlayingContextObject) SetActions(v DisallowsObject)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CurrentlyPlayingContextObject) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


