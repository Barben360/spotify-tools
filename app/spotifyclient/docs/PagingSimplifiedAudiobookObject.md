# PagingSimplifiedAudiobookObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | A link to the Web API endpoint returning the full result of the request  | 
**Limit** | **int32** | The maximum number of items in the response (as set in the query or by default).  | 
**Next** | **NullableString** | URL to the next page of items. ( &#x60;null&#x60; if none)  | 
**Offset** | **int32** | The offset of the items returned (as set in the query or by default)  | 
**Previous** | **NullableString** | URL to the previous page of items. ( &#x60;null&#x60; if none)  | 
**Total** | **int32** | The total number of items available to return.  | 
**Items** | [**[]SimplifiedAudiobookObject**](SimplifiedAudiobookObject.md) |  | 

## Methods

### NewPagingSimplifiedAudiobookObject

`func NewPagingSimplifiedAudiobookObject(href string, limit int32, next NullableString, offset int32, previous NullableString, total int32, items []SimplifiedAudiobookObject, ) *PagingSimplifiedAudiobookObject`

NewPagingSimplifiedAudiobookObject instantiates a new PagingSimplifiedAudiobookObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagingSimplifiedAudiobookObjectWithDefaults

`func NewPagingSimplifiedAudiobookObjectWithDefaults() *PagingSimplifiedAudiobookObject`

NewPagingSimplifiedAudiobookObjectWithDefaults instantiates a new PagingSimplifiedAudiobookObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PagingSimplifiedAudiobookObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PagingSimplifiedAudiobookObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PagingSimplifiedAudiobookObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetLimit

`func (o *PagingSimplifiedAudiobookObject) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PagingSimplifiedAudiobookObject) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PagingSimplifiedAudiobookObject) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *PagingSimplifiedAudiobookObject) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PagingSimplifiedAudiobookObject) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PagingSimplifiedAudiobookObject) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *PagingSimplifiedAudiobookObject) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PagingSimplifiedAudiobookObject) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetOffset

`func (o *PagingSimplifiedAudiobookObject) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PagingSimplifiedAudiobookObject) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PagingSimplifiedAudiobookObject) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetPrevious

`func (o *PagingSimplifiedAudiobookObject) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PagingSimplifiedAudiobookObject) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PagingSimplifiedAudiobookObject) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *PagingSimplifiedAudiobookObject) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PagingSimplifiedAudiobookObject) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotal

`func (o *PagingSimplifiedAudiobookObject) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PagingSimplifiedAudiobookObject) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PagingSimplifiedAudiobookObject) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *PagingSimplifiedAudiobookObject) GetItems() []SimplifiedAudiobookObject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagingSimplifiedAudiobookObject) GetItemsOk() (*[]SimplifiedAudiobookObject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagingSimplifiedAudiobookObject) SetItems(v []SimplifiedAudiobookObject)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


