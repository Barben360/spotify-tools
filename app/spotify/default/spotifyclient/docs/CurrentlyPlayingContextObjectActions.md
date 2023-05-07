# CurrentlyPlayingContextObjectActions

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

### NewCurrentlyPlayingContextObjectActions

`func NewCurrentlyPlayingContextObjectActions() *CurrentlyPlayingContextObjectActions`

NewCurrentlyPlayingContextObjectActions instantiates a new CurrentlyPlayingContextObjectActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentlyPlayingContextObjectActionsWithDefaults

`func NewCurrentlyPlayingContextObjectActionsWithDefaults() *CurrentlyPlayingContextObjectActions`

NewCurrentlyPlayingContextObjectActionsWithDefaults instantiates a new CurrentlyPlayingContextObjectActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterruptingPlayback

`func (o *CurrentlyPlayingContextObjectActions) GetInterruptingPlayback() bool`

GetInterruptingPlayback returns the InterruptingPlayback field if non-nil, zero value otherwise.

### GetInterruptingPlaybackOk

`func (o *CurrentlyPlayingContextObjectActions) GetInterruptingPlaybackOk() (*bool, bool)`

GetInterruptingPlaybackOk returns a tuple with the InterruptingPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterruptingPlayback

`func (o *CurrentlyPlayingContextObjectActions) SetInterruptingPlayback(v bool)`

SetInterruptingPlayback sets InterruptingPlayback field to given value.

### HasInterruptingPlayback

`func (o *CurrentlyPlayingContextObjectActions) HasInterruptingPlayback() bool`

HasInterruptingPlayback returns a boolean if a field has been set.

### GetPausing

`func (o *CurrentlyPlayingContextObjectActions) GetPausing() bool`

GetPausing returns the Pausing field if non-nil, zero value otherwise.

### GetPausingOk

`func (o *CurrentlyPlayingContextObjectActions) GetPausingOk() (*bool, bool)`

GetPausingOk returns a tuple with the Pausing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausing

`func (o *CurrentlyPlayingContextObjectActions) SetPausing(v bool)`

SetPausing sets Pausing field to given value.

### HasPausing

`func (o *CurrentlyPlayingContextObjectActions) HasPausing() bool`

HasPausing returns a boolean if a field has been set.

### GetResuming

`func (o *CurrentlyPlayingContextObjectActions) GetResuming() bool`

GetResuming returns the Resuming field if non-nil, zero value otherwise.

### GetResumingOk

`func (o *CurrentlyPlayingContextObjectActions) GetResumingOk() (*bool, bool)`

GetResumingOk returns a tuple with the Resuming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResuming

`func (o *CurrentlyPlayingContextObjectActions) SetResuming(v bool)`

SetResuming sets Resuming field to given value.

### HasResuming

`func (o *CurrentlyPlayingContextObjectActions) HasResuming() bool`

HasResuming returns a boolean if a field has been set.

### GetSeeking

`func (o *CurrentlyPlayingContextObjectActions) GetSeeking() bool`

GetSeeking returns the Seeking field if non-nil, zero value otherwise.

### GetSeekingOk

`func (o *CurrentlyPlayingContextObjectActions) GetSeekingOk() (*bool, bool)`

GetSeekingOk returns a tuple with the Seeking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeking

`func (o *CurrentlyPlayingContextObjectActions) SetSeeking(v bool)`

SetSeeking sets Seeking field to given value.

### HasSeeking

`func (o *CurrentlyPlayingContextObjectActions) HasSeeking() bool`

HasSeeking returns a boolean if a field has been set.

### GetSkippingNext

`func (o *CurrentlyPlayingContextObjectActions) GetSkippingNext() bool`

GetSkippingNext returns the SkippingNext field if non-nil, zero value otherwise.

### GetSkippingNextOk

`func (o *CurrentlyPlayingContextObjectActions) GetSkippingNextOk() (*bool, bool)`

GetSkippingNextOk returns a tuple with the SkippingNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippingNext

`func (o *CurrentlyPlayingContextObjectActions) SetSkippingNext(v bool)`

SetSkippingNext sets SkippingNext field to given value.

### HasSkippingNext

`func (o *CurrentlyPlayingContextObjectActions) HasSkippingNext() bool`

HasSkippingNext returns a boolean if a field has been set.

### GetSkippingPrev

`func (o *CurrentlyPlayingContextObjectActions) GetSkippingPrev() bool`

GetSkippingPrev returns the SkippingPrev field if non-nil, zero value otherwise.

### GetSkippingPrevOk

`func (o *CurrentlyPlayingContextObjectActions) GetSkippingPrevOk() (*bool, bool)`

GetSkippingPrevOk returns a tuple with the SkippingPrev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippingPrev

`func (o *CurrentlyPlayingContextObjectActions) SetSkippingPrev(v bool)`

SetSkippingPrev sets SkippingPrev field to given value.

### HasSkippingPrev

`func (o *CurrentlyPlayingContextObjectActions) HasSkippingPrev() bool`

HasSkippingPrev returns a boolean if a field has been set.

### GetTogglingRepeatContext

`func (o *CurrentlyPlayingContextObjectActions) GetTogglingRepeatContext() bool`

GetTogglingRepeatContext returns the TogglingRepeatContext field if non-nil, zero value otherwise.

### GetTogglingRepeatContextOk

`func (o *CurrentlyPlayingContextObjectActions) GetTogglingRepeatContextOk() (*bool, bool)`

GetTogglingRepeatContextOk returns a tuple with the TogglingRepeatContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglingRepeatContext

`func (o *CurrentlyPlayingContextObjectActions) SetTogglingRepeatContext(v bool)`

SetTogglingRepeatContext sets TogglingRepeatContext field to given value.

### HasTogglingRepeatContext

`func (o *CurrentlyPlayingContextObjectActions) HasTogglingRepeatContext() bool`

HasTogglingRepeatContext returns a boolean if a field has been set.

### GetTogglingShuffle

`func (o *CurrentlyPlayingContextObjectActions) GetTogglingShuffle() bool`

GetTogglingShuffle returns the TogglingShuffle field if non-nil, zero value otherwise.

### GetTogglingShuffleOk

`func (o *CurrentlyPlayingContextObjectActions) GetTogglingShuffleOk() (*bool, bool)`

GetTogglingShuffleOk returns a tuple with the TogglingShuffle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglingShuffle

`func (o *CurrentlyPlayingContextObjectActions) SetTogglingShuffle(v bool)`

SetTogglingShuffle sets TogglingShuffle field to given value.

### HasTogglingShuffle

`func (o *CurrentlyPlayingContextObjectActions) HasTogglingShuffle() bool`

HasTogglingShuffle returns a boolean if a field has been set.

### GetTogglingRepeatTrack

`func (o *CurrentlyPlayingContextObjectActions) GetTogglingRepeatTrack() bool`

GetTogglingRepeatTrack returns the TogglingRepeatTrack field if non-nil, zero value otherwise.

### GetTogglingRepeatTrackOk

`func (o *CurrentlyPlayingContextObjectActions) GetTogglingRepeatTrackOk() (*bool, bool)`

GetTogglingRepeatTrackOk returns a tuple with the TogglingRepeatTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTogglingRepeatTrack

`func (o *CurrentlyPlayingContextObjectActions) SetTogglingRepeatTrack(v bool)`

SetTogglingRepeatTrack sets TogglingRepeatTrack field to given value.

### HasTogglingRepeatTrack

`func (o *CurrentlyPlayingContextObjectActions) HasTogglingRepeatTrack() bool`

HasTogglingRepeatTrack returns a boolean if a field has been set.

### GetTransferringPlayback

`func (o *CurrentlyPlayingContextObjectActions) GetTransferringPlayback() bool`

GetTransferringPlayback returns the TransferringPlayback field if non-nil, zero value otherwise.

### GetTransferringPlaybackOk

`func (o *CurrentlyPlayingContextObjectActions) GetTransferringPlaybackOk() (*bool, bool)`

GetTransferringPlaybackOk returns a tuple with the TransferringPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferringPlayback

`func (o *CurrentlyPlayingContextObjectActions) SetTransferringPlayback(v bool)`

SetTransferringPlayback sets TransferringPlayback field to given value.

### HasTransferringPlayback

`func (o *CurrentlyPlayingContextObjectActions) HasTransferringPlayback() bool`

HasTransferringPlayback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


