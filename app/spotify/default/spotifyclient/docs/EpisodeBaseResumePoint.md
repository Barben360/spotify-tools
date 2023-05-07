# EpisodeBaseResumePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyPlayed** | Pointer to **bool** | Whether or not the episode has been fully played by the user.  | [optional] 
**ResumePositionMs** | Pointer to **int32** | The user&#39;s most recent position in the episode in milliseconds.  | [optional] 

## Methods

### NewEpisodeBaseResumePoint

`func NewEpisodeBaseResumePoint() *EpisodeBaseResumePoint`

NewEpisodeBaseResumePoint instantiates a new EpisodeBaseResumePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpisodeBaseResumePointWithDefaults

`func NewEpisodeBaseResumePointWithDefaults() *EpisodeBaseResumePoint`

NewEpisodeBaseResumePointWithDefaults instantiates a new EpisodeBaseResumePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullyPlayed

`func (o *EpisodeBaseResumePoint) GetFullyPlayed() bool`

GetFullyPlayed returns the FullyPlayed field if non-nil, zero value otherwise.

### GetFullyPlayedOk

`func (o *EpisodeBaseResumePoint) GetFullyPlayedOk() (*bool, bool)`

GetFullyPlayedOk returns a tuple with the FullyPlayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyPlayed

`func (o *EpisodeBaseResumePoint) SetFullyPlayed(v bool)`

SetFullyPlayed sets FullyPlayed field to given value.

### HasFullyPlayed

`func (o *EpisodeBaseResumePoint) HasFullyPlayed() bool`

HasFullyPlayed returns a boolean if a field has been set.

### GetResumePositionMs

`func (o *EpisodeBaseResumePoint) GetResumePositionMs() int32`

GetResumePositionMs returns the ResumePositionMs field if non-nil, zero value otherwise.

### GetResumePositionMsOk

`func (o *EpisodeBaseResumePoint) GetResumePositionMsOk() (*int32, bool)`

GetResumePositionMsOk returns a tuple with the ResumePositionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumePositionMs

`func (o *EpisodeBaseResumePoint) SetResumePositionMs(v int32)`

SetResumePositionMs sets ResumePositionMs field to given value.

### HasResumePositionMs

`func (o *EpisodeBaseResumePoint) HasResumePositionMs() bool`

HasResumePositionMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


