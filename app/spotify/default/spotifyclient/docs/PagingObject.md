# PagingObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | A link to the Web API endpoint returning the full result of the request  | 
**Limit** | **int32** | The maximum number of items in the response (as set in the query or by default).  | 
**Next** | **NullableString** | URL to the next page of items. ( &#x60;null&#x60; if none)  | 
**Offset** | **int32** | The offset of the items returned (as set in the query or by default)  | 
**Previous** | **NullableString** | URL to the previous page of items. ( &#x60;null&#x60; if none)  | 
**Total** | **int32** | The total number of items available to return.  | 

## Methods

### NewPagingObject

`func NewPagingObject(href string, limit int32, next NullableString, offset int32, previous NullableString, total int32, ) *PagingObject`

NewPagingObject instantiates a new PagingObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingObjectWithDefaults

`func NewPagingObjectWithDefaults() *PagingObject`

NewPagingObjectWithDefaults instantiates a new PagingObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PagingObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PagingObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PagingObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetLimit

`func (o *PagingObject) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagingObject) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagingObject) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *PagingObject) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PagingObject) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PagingObject) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *PagingObject) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PagingObject) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetOffset

`func (o *PagingObject) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PagingObject) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PagingObject) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetPrevious

`func (o *PagingObject) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PagingObject) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PagingObject) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *PagingObject) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PagingObject) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotal

`func (o *PagingObject) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PagingObject) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PagingObject) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


