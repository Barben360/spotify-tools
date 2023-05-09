# CategoryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | A link to the Web API endpoint returning full details of the category.  | 
**Icons** | [**[]ImageObject**](ImageObject.md) | The category icon, in various sizes.  | 
**Id** | **string** | The [Spotify category ID](/documentation/web-api/concepts/spotify-uris-ids) of the category.  | 
**Name** | **string** | The name of the category.  | 

## Methods

### NewCategoryObject

`func NewCategoryObject(href string, icons []ImageObject, id string, name string, ) *CategoryObject`

NewCategoryObject instantiates a new CategoryObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryObjectWithDefaults

`func NewCategoryObjectWithDefaults() *CategoryObject`

NewCategoryObjectWithDefaults instantiates a new CategoryObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CategoryObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CategoryObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CategoryObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetIcons

`func (o *CategoryObject) GetIcons() []ImageObject`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *CategoryObject) GetIconsOk() (*[]ImageObject, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *CategoryObject) SetIcons(v []ImageObject)`

SetIcons sets Icons field to given value.


### GetId

`func (o *CategoryObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CategoryObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CategoryObject) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CategoryObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryObject) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


