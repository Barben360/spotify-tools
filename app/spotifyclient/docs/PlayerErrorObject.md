# PlayerErrorObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | The HTTP status code. Either &#x60;404 NOT FOUND&#x60; or &#x60;403 FORBIDDEN&#x60;.  Also returned in the response header.  | [optional] 
**Message** | Pointer to **string** | A short description of the cause of the error.  | [optional] 
**Reason** | Pointer to [**PlayerErrorReasons**](PlayerErrorReasons.md) |  | [optional] 

## Methods

### NewPlayerErrorObject

`func NewPlayerErrorObject() *PlayerErrorObject`

NewPlayerErrorObject instantiates a new PlayerErrorObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerErrorObjectWithDefaults

`func NewPlayerErrorObjectWithDefaults() *PlayerErrorObject`

NewPlayerErrorObjectWithDefaults instantiates a new PlayerErrorObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PlayerErrorObject) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlayerErrorObject) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlayerErrorObject) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlayerErrorObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *PlayerErrorObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlayerErrorObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlayerErrorObject) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PlayerErrorObject) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *PlayerErrorObject) GetReason() PlayerErrorReasons`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PlayerErrorObject) GetReasonOk() (*PlayerErrorReasons, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PlayerErrorObject) SetReason(v PlayerErrorReasons)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PlayerErrorObject) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


