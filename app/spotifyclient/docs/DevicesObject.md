# DevicesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]DeviceObject**](DeviceObject.md) | A list of 0..n Device objects | [optional] 

## Methods

### NewDevicesObject

`func NewDevicesObject() *DevicesObject`

NewDevicesObject instantiates a new DevicesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesObjectWithDefaults

`func NewDevicesObjectWithDefaults() *DevicesObject`

NewDevicesObjectWithDefaults instantiates a new DevicesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *DevicesObject) GetDevices() []DeviceObject`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DevicesObject) GetDevicesOk() (*[]DeviceObject, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DevicesObject) SetDevices(v []DeviceObject)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DevicesObject) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


