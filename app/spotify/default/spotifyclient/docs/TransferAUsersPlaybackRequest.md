# TransferAUsersPlaybackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | **[]string** | A JSON array containing the ID of the device on which playback should be started/transferred.&lt;br/&gt;For example:&#x60;{device_ids:[\&quot;74ASZWbe4lXaubB36ztrGX\&quot;]}&#x60;&lt;br/&gt;_**Note**: Although an array is accepted, only a single device_id is currently supported. Supplying more than one will return &#x60;400 Bad Request&#x60;_  | 
**Play** | Pointer to **bool** | **true**: ensure playback happens on new device.&lt;br/&gt;**false** or not provided: keep the current playback state.  | [optional] 

## Methods

### NewTransferAUsersPlaybackRequest

`func NewTransferAUsersPlaybackRequest(deviceIds []string, ) *TransferAUsersPlaybackRequest`

NewTransferAUsersPlaybackRequest instantiates a new TransferAUsersPlaybackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAUsersPlaybackRequestWithDefaults

`func NewTransferAUsersPlaybackRequestWithDefaults() *TransferAUsersPlaybackRequest`

NewTransferAUsersPlaybackRequestWithDefaults instantiates a new TransferAUsersPlaybackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *TransferAUsersPlaybackRequest) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *TransferAUsersPlaybackRequest) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *TransferAUsersPlaybackRequest) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.


### GetPlay

`func (o *TransferAUsersPlaybackRequest) GetPlay() bool`

GetPlay returns the Play field if non-nil, zero value otherwise.

### GetPlayOk

`func (o *TransferAUsersPlaybackRequest) GetPlayOk() (*bool, bool)`

GetPlayOk returns a tuple with the Play field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlay

`func (o *TransferAUsersPlaybackRequest) SetPlay(v bool)`

SetPlay sets Play field to given value.

### HasPlay

`func (o *TransferAUsersPlaybackRequest) HasPlay() bool`

HasPlay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


