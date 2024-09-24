# CursorPagingObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | A link to the Web API endpoint returning the full result of the request. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of items in the response (as set in the query or by default). | [optional] 
**Next** | Pointer to **string** | URL to the next page of items. ( &#x60;null&#x60; if none) | [optional] 
**Cursors** | Pointer to [**CursorObject**](CursorObject.md) | The cursors used to find the next set of items. | [optional] 
**Total** | Pointer to **int32** | The total number of items available to return. | [optional] 

## Methods

### NewCursorPagingObject

`func NewCursorPagingObject() *CursorPagingObject`

NewCursorPagingObject instantiates a new CursorPagingObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPagingObjectWithDefaults

`func NewCursorPagingObjectWithDefaults() *CursorPagingObject`

NewCursorPagingObjectWithDefaults instantiates a new CursorPagingObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CursorPagingObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CursorPagingObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CursorPagingObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CursorPagingObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *CursorPagingObject) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CursorPagingObject) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CursorPagingObject) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CursorPagingObject) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *CursorPagingObject) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CursorPagingObject) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CursorPagingObject) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *CursorPagingObject) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCursors

`func (o *CursorPagingObject) GetCursors() CursorObject`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *CursorPagingObject) GetCursorsOk() (*CursorObject, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *CursorPagingObject) SetCursors(v CursorObject)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *CursorPagingObject) HasCursors() bool`

HasCursors returns a boolean if a field has been set.

### GetTotal

`func (o *CursorPagingObject) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CursorPagingObject) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CursorPagingObject) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CursorPagingObject) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


