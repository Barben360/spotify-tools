# ExternalIdObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isrc** | Pointer to **string** | [International Standard Recording Code](http://en.wikipedia.org/wiki/International_Standard_Recording_Code)  | [optional] 
**Ean** | Pointer to **string** | [International Article Number](http://en.wikipedia.org/wiki/International_Article_Number_%28EAN%29)  | [optional] 
**Upc** | Pointer to **string** | [Universal Product Code](http://en.wikipedia.org/wiki/Universal_Product_Code)  | [optional] 

## Methods

### NewExternalIdObject

`func NewExternalIdObject() *ExternalIdObject`

NewExternalIdObject instantiates a new ExternalIdObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIdObjectWithDefaults

`func NewExternalIdObjectWithDefaults() *ExternalIdObject`

NewExternalIdObjectWithDefaults instantiates a new ExternalIdObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsrc

`func (o *ExternalIdObject) GetIsrc() string`

GetIsrc returns the Isrc field if non-nil, zero value otherwise.

### GetIsrcOk

`func (o *ExternalIdObject) GetIsrcOk() (*string, bool)`

GetIsrcOk returns a tuple with the Isrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsrc

`func (o *ExternalIdObject) SetIsrc(v string)`

SetIsrc sets Isrc field to given value.

### HasIsrc

`func (o *ExternalIdObject) HasIsrc() bool`

HasIsrc returns a boolean if a field has been set.

### GetEan

`func (o *ExternalIdObject) GetEan() string`

GetEan returns the Ean field if non-nil, zero value otherwise.

### GetEanOk

`func (o *ExternalIdObject) GetEanOk() (*string, bool)`

GetEanOk returns a tuple with the Ean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEan

`func (o *ExternalIdObject) SetEan(v string)`

SetEan sets Ean field to given value.

### HasEan

`func (o *ExternalIdObject) HasEan() bool`

HasEan returns a boolean if a field has been set.

### GetUpc

`func (o *ExternalIdObject) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *ExternalIdObject) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *ExternalIdObject) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *ExternalIdObject) HasUpc() bool`

HasUpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


