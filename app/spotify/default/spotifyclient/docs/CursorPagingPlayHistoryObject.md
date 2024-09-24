# CursorPagingPlayHistoryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | A link to the Web API endpoint returning the full result of the request. | [optional] 
**Limit** | Pointer to **int32** | The maximum number of items in the response (as set in the query or by default). | [optional] 
**Next** | Pointer to **string** | URL to the next page of items. ( &#x60;null&#x60; if none) | [optional] 
**Cursors** | Pointer to [**CursorObject**](CursorObject.md) | The cursors used to find the next set of items. | [optional] 
**Total** | Pointer to **int32** | The total number of items available to return. | [optional] 
**Items** | Pointer to [**[]PlayHistoryObject**](PlayHistoryObject.md) |  | [optional] 

## Methods

### NewCursorPagingPlayHistoryObject

`func NewCursorPagingPlayHistoryObject() *CursorPagingPlayHistoryObject`

NewCursorPagingPlayHistoryObject instantiates a new CursorPagingPlayHistoryObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorPagingPlayHistoryObjectWithDefaults

`func NewCursorPagingPlayHistoryObjectWithDefaults() *CursorPagingPlayHistoryObject`

NewCursorPagingPlayHistoryObjectWithDefaults instantiates a new CursorPagingPlayHistoryObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CursorPagingPlayHistoryObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CursorPagingPlayHistoryObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CursorPagingPlayHistoryObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CursorPagingPlayHistoryObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetLimit

`func (o *CursorPagingPlayHistoryObject) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CursorPagingPlayHistoryObject) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CursorPagingPlayHistoryObject) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CursorPagingPlayHistoryObject) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *CursorPagingPlayHistoryObject) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CursorPagingPlayHistoryObject) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CursorPagingPlayHistoryObject) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *CursorPagingPlayHistoryObject) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetCursors

`func (o *CursorPagingPlayHistoryObject) GetCursors() CursorObject`

GetCursors returns the Cursors field if non-nil, zero value otherwise.

### GetCursorsOk

`func (o *CursorPagingPlayHistoryObject) GetCursorsOk() (*CursorObject, bool)`

GetCursorsOk returns a tuple with the Cursors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursors

`func (o *CursorPagingPlayHistoryObject) SetCursors(v CursorObject)`

SetCursors sets Cursors field to given value.

### HasCursors

`func (o *CursorPagingPlayHistoryObject) HasCursors() bool`

HasCursors returns a boolean if a field has been set.

### GetTotal

`func (o *CursorPagingPlayHistoryObject) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CursorPagingPlayHistoryObject) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CursorPagingPlayHistoryObject) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CursorPagingPlayHistoryObject) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *CursorPagingPlayHistoryObject) GetItems() []PlayHistoryObject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CursorPagingPlayHistoryObject) GetItemsOk() (*[]PlayHistoryObject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CursorPagingPlayHistoryObject) SetItems(v []PlayHistoryObject)`

SetItems sets Items field to given value.

### HasItems

`func (o *CursorPagingPlayHistoryObject) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


