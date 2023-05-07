# TimeIntervalObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **float32** | The starting point (in seconds) of the time interval. | [optional] 
**Duration** | Pointer to **float32** | The duration (in seconds) of the time interval. | [optional] 
**Confidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the interval. | [optional] 

## Methods

### NewTimeIntervalObject

`func NewTimeIntervalObject() *TimeIntervalObject`

NewTimeIntervalObject instantiates a new TimeIntervalObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeIntervalObjectWithDefaults

`func NewTimeIntervalObjectWithDefaults() *TimeIntervalObject`

NewTimeIntervalObjectWithDefaults instantiates a new TimeIntervalObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *TimeIntervalObject) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TimeIntervalObject) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TimeIntervalObject) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *TimeIntervalObject) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetDuration

`func (o *TimeIntervalObject) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TimeIntervalObject) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TimeIntervalObject) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TimeIntervalObject) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetConfidence

`func (o *TimeIntervalObject) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *TimeIntervalObject) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *TimeIntervalObject) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *TimeIntervalObject) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


