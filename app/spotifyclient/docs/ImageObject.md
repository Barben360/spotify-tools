# ImageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The source URL of the image.  | 
**Height** | **NullableInt32** | The image height in pixels.  | 
**Width** | **NullableInt32** | The image width in pixels.  | 

## Methods

### NewImageObject

`func NewImageObject(url string, height NullableInt32, width NullableInt32, ) *ImageObject`

NewImageObject instantiates a new ImageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageObjectWithDefaults

`func NewImageObjectWithDefaults() *ImageObject`

NewImageObjectWithDefaults instantiates a new ImageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ImageObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageObject) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeight

`func (o *ImageObject) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ImageObject) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ImageObject) SetHeight(v int32)`

SetHeight sets Height field to given value.


### SetHeightNil

`func (o *ImageObject) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ImageObject) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *ImageObject) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ImageObject) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ImageObject) SetWidth(v int32)`

SetWidth sets Width field to given value.


### SetWidthNil

`func (o *ImageObject) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ImageObject) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


