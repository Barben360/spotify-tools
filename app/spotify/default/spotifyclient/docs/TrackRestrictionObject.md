# TrackRestrictionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The reason for the restriction. Supported values: - &#x60;market&#x60; - The content item is not available in the given market. - &#x60;product&#x60; - The content item is not available for the user&#39;s subscription type. - &#x60;explicit&#x60; - The content item is explicit and the user&#39;s account is set to not play explicit content.  Additional reasons may be added in the future. **Note**: If you use this field, make sure that your application safely handles unknown values.  | [optional] 

## Methods

### NewTrackRestrictionObject

`func NewTrackRestrictionObject() *TrackRestrictionObject`

NewTrackRestrictionObject instantiates a new TrackRestrictionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackRestrictionObjectWithDefaults

`func NewTrackRestrictionObjectWithDefaults() *TrackRestrictionObject`

NewTrackRestrictionObjectWithDefaults instantiates a new TrackRestrictionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *TrackRestrictionObject) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TrackRestrictionObject) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TrackRestrictionObject) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TrackRestrictionObject) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


