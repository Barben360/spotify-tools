# GetCategories200ResponseCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | A link to the Web API endpoint returning the full result of the request  | 
**Limit** | **int32** | The maximum number of items in the response (as set in the query or by default).  | 
**Next** | **NullableString** | URL to the next page of items. ( &#x60;null&#x60; if none)  | 
**Offset** | **int32** | The offset of the items returned (as set in the query or by default)  | 
**Previous** | **NullableString** | URL to the previous page of items. ( &#x60;null&#x60; if none)  | 
**Total** | **int32** | The total number of items available to return.  | 
**Items** | [**[]CategoryObject**](CategoryObject.md) |  | 

## Methods

### NewGetCategories200ResponseCategories

`func NewGetCategories200ResponseCategories(href string, limit int32, next NullableString, offset int32, previous NullableString, total int32, items []CategoryObject, ) *GetCategories200ResponseCategories`

NewGetCategories200ResponseCategories instantiates a new GetCategories200ResponseCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCategories200ResponseCategoriesWithDefaults

`func NewGetCategories200ResponseCategoriesWithDefaults() *GetCategories200ResponseCategories`

NewGetCategories200ResponseCategoriesWithDefaults instantiates a new GetCategories200ResponseCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *GetCategories200ResponseCategories) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GetCategories200ResponseCategories) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GetCategories200ResponseCategories) SetHref(v string)`

SetHref sets Href field to given value.


### GetLimit

`func (o *GetCategories200ResponseCategories) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetCategories200ResponseCategories) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetCategories200ResponseCategories) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNext

`func (o *GetCategories200ResponseCategories) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetCategories200ResponseCategories) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetCategories200ResponseCategories) SetNext(v string)`

SetNext sets Next field to given value.


### SetNextNil

`func (o *GetCategories200ResponseCategories) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *GetCategories200ResponseCategories) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetOffset

`func (o *GetCategories200ResponseCategories) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetCategories200ResponseCategories) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetCategories200ResponseCategories) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetPrevious

`func (o *GetCategories200ResponseCategories) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *GetCategories200ResponseCategories) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *GetCategories200ResponseCategories) SetPrevious(v string)`

SetPrevious sets Previous field to given value.


### SetPreviousNil

`func (o *GetCategories200ResponseCategories) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *GetCategories200ResponseCategories) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetTotal

`func (o *GetCategories200ResponseCategories) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetCategories200ResponseCategories) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetCategories200ResponseCategories) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *GetCategories200ResponseCategories) GetItems() []CategoryObject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetCategories200ResponseCategories) GetItemsOk() (*[]CategoryObject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetCategories200ResponseCategories) SetItems(v []CategoryObject)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


