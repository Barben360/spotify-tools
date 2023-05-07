# DisallowsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterruptingPlayback** | Pointer to **bool** | Interrupting playback. Optional field. | [optional] 
**Pausing** | Pointer to **bool** | Pausing. Optional field. | [optional] 
**Resuming** | Pointer to **bool** | Resuming. Optional field. | [optional] 
**Seeking** | Pointer to **bool** | Seeking playback location. Optional field. | [optional] 
**SkippingNext** | Pointer to **bool** | Skipping to the next context. Optional field. | [optional] 
**SkippingPrev** | Pointer to **bool** | Skipping to the previous context. Optional field. | [optional] 
**TogglingRepeatContext** | Pointer to **bool** | Toggling repeat context flag. Optional field. | [optional] 
**TogglingShuffle** | Pointer to **bool** | Toggling shuffle flag. Optional field. | [optional] 
**TogglingRepeatTrack** | Pointer to **bool** | Toggling repeat track flag. Optional field. | [optional] 
**TransferringPlayback** | Pointer to **bool** | Transfering playback between devices. Optional field. | [optional] 

## Methods

### NewDisallowsObject

`func NewDisallowsObject() *DisallowsObject`

NewDisallowsObject instantiates a new DisallowsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisallowsObjectWithDefaults

`func NewDisallowsObjectWithDefaults() *DisallowsObject`

NewDisallowsObjectWithDefaults instantiates a new DisallowsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterruptingPlayback

`func (o *DisallowsObject) GetInterruptingPlayback() bool`

GetInterruptingPlayback returns the InterruptingPlayback field if non-nil, zero value otherwise.

### GetInterruptingPlaybackOk

`func (o *DisallowsObject) GetInterruptingPlaybackOk() (*bool, bool)`

GetInterruptingPlaybackOk returns a tuple with the InterruptingPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptingPlayback

`func (o *DisallowsObject) SetInterruptingPlayback(v bool)`

SetInterruptingPlayback sets InterruptingPlayback field to given value.

### HasInterruptingPlayback

`func (o *DisallowsObject) HasInterruptingPlayback() bool`

HasInterruptingPlayback returns a boolean if a field has been set.

### GetPausing

`func (o *DisallowsObject) GetPausing() bool`

GetPausing returns the Pausing field if non-nil, zero value otherwise.

### GetPausingOk

`func (o *DisallowsObject) GetPausingOk() (*bool, bool)`

GetPausingOk returns a tuple with the Pausing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausing

`func (o *DisallowsObject) SetPausing(v bool)`

SetPausing sets Pausing field to given value.

### HasPausing

`func (o *DisallowsObject) HasPausing() bool`

HasPausing returns a boolean if a field has been set.

### GetResuming

`func (o *DisallowsObject) GetResuming() bool`

GetResuming returns the Resuming field if non-nil, zero value otherwise.

### GetResumingOk

`func (o *DisallowsObject) GetResumingOk() (*bool, bool)`

GetResumingOk returns a tuple with the Resuming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResuming

`func (o *DisallowsObject) SetResuming(v bool)`

SetResuming sets Resuming field to given value.

### HasResuming

`func (o *DisallowsObject) HasResuming() bool`

HasResuming returns a boolean if a field has been set.

### GetSeeking

`func (o *DisallowsObject) GetSeeking() bool`

GetSeeking returns the Seeking field if non-nil, zero value otherwise.

### GetSeekingOk

`func (o *DisallowsObject) GetSeekingOk() (*bool, bool)`

GetSeekingOk returns a tuple with the Seeking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeking

`func (o *DisallowsObject) SetSeeking(v bool)`

SetSeeking sets Seeking field to given value.

### HasSeeking

`func (o *DisallowsObject) HasSeeking() bool`

HasSeeking returns a boolean if a field has been set.

### GetSkippingNext

`func (o *DisallowsObject) GetSkippingNext() bool`

GetSkippingNext returns the SkippingNext field if non-nil, zero value otherwise.

### GetSkippingNextOk

`func (o *DisallowsObject) GetSkippingNextOk() (*bool, bool)`

GetSkippingNextOk returns a tuple with the SkippingNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippingNext

`func (o *DisallowsObject) SetSkippingNext(v bool)`

SetSkippingNext sets SkippingNext field to given value.

### HasSkippingNext

`func (o *DisallowsObject) HasSkippingNext() bool`

HasSkippingNext returns a boolean if a field has been set.

### GetSkippingPrev

`func (o *DisallowsObject) GetSkippingPrev() bool`

GetSkippingPrev returns the SkippingPrev field if non-nil, zero value otherwise.

### GetSkippingPrevOk

`func (o *DisallowsObject) GetSkippingPrevOk() (*bool, bool)`

GetSkippingPrevOk returns a tuple with the SkippingPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippingPrev

`func (o *DisallowsObject) SetSkippingPrev(v bool)`

SetSkippingPrev sets SkippingPrev field to given value.

### HasSkippingPrev

`func (o *DisallowsObject) HasSkippingPrev() bool`

HasSkippingPrev returns a boolean if a field has been set.

### GetTogglingRepeatContext

`func (o *DisallowsObject) GetTogglingRepeatContext() bool`

GetTogglingRepeatContext returns the TogglingRepeatContext field if non-nil, zero value otherwise.

### GetTogglingRepeatContextOk

`func (o *DisallowsObject) GetTogglingRepeatContextOk() (*bool, bool)`

GetTogglingRepeatContextOk returns a tuple with the TogglingRepeatContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglingRepeatContext

`func (o *DisallowsObject) SetTogglingRepeatContext(v bool)`

SetTogglingRepeatContext sets TogglingRepeatContext field to given value.

### HasTogglingRepeatContext

`func (o *DisallowsObject) HasTogglingRepeatContext() bool`

HasTogglingRepeatContext returns a boolean if a field has been set.

### GetTogglingShuffle

`func (o *DisallowsObject) GetTogglingShuffle() bool`

GetTogglingShuffle returns the TogglingShuffle field if non-nil, zero value otherwise.

### GetTogglingShuffleOk

`func (o *DisallowsObject) GetTogglingShuffleOk() (*bool, bool)`

GetTogglingShuffleOk returns a tuple with the TogglingShuffle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglingShuffle

`func (o *DisallowsObject) SetTogglingShuffle(v bool)`

SetTogglingShuffle sets TogglingShuffle field to given value.

### HasTogglingShuffle

`func (o *DisallowsObject) HasTogglingShuffle() bool`

HasTogglingShuffle returns a boolean if a field has been set.

### GetTogglingRepeatTrack

`func (o *DisallowsObject) GetTogglingRepeatTrack() bool`

GetTogglingRepeatTrack returns the TogglingRepeatTrack field if non-nil, zero value otherwise.

### GetTogglingRepeatTrackOk

`func (o *DisallowsObject) GetTogglingRepeatTrackOk() (*bool, bool)`

GetTogglingRepeatTrackOk returns a tuple with the TogglingRepeatTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglingRepeatTrack

`func (o *DisallowsObject) SetTogglingRepeatTrack(v bool)`

SetTogglingRepeatTrack sets TogglingRepeatTrack field to given value.

### HasTogglingRepeatTrack

`func (o *DisallowsObject) HasTogglingRepeatTrack() bool`

HasTogglingRepeatTrack returns a boolean if a field has been set.

### GetTransferringPlayback

`func (o *DisallowsObject) GetTransferringPlayback() bool`

GetTransferringPlayback returns the TransferringPlayback field if non-nil, zero value otherwise.

### GetTransferringPlaybackOk

`func (o *DisallowsObject) GetTransferringPlaybackOk() (*bool, bool)`

GetTransferringPlaybackOk returns a tuple with the TransferringPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferringPlayback

`func (o *DisallowsObject) SetTransferringPlayback(v bool)`

SetTransferringPlayback sets TransferringPlayback field to given value.

### HasTransferringPlayback

`func (o *DisallowsObject) HasTransferringPlayback() bool`

HasTransferringPlayback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


