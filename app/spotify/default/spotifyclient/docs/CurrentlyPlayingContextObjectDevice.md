# CurrentlyPlayingContextObjectDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The device ID. | [optional] 
**IsActive** | Pointer to **bool** | If this device is the currently active device. | [optional] 
**IsPrivateSession** | Pointer to **bool** | If this device is currently in a private session. | [optional] 
**IsRestricted** | Pointer to **bool** | Whether controlling this device is restricted. At present if this is \&quot;true\&quot; then no Web API commands will be accepted by this device. | [optional] 
**Name** | Pointer to **string** | A human-readable name for the device. Some devices have a name that the user can configure (e.g. \\\&quot;Loudest speaker\\\&quot;) and some devices have a generic name associated with the manufacturer or device model. | [optional] 
**Type** | Pointer to **string** | Device type, such as \&quot;computer\&quot;, \&quot;smartphone\&quot; or \&quot;speaker\&quot;. | [optional] 
**VolumePercent** | Pointer to **NullableInt32** | The current volume in percent. | [optional] 

## Methods

### NewCurrentlyPlayingContextObjectDevice

`func NewCurrentlyPlayingContextObjectDevice() *CurrentlyPlayingContextObjectDevice`

NewCurrentlyPlayingContextObjectDevice instantiates a new CurrentlyPlayingContextObjectDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentlyPlayingContextObjectDeviceWithDefaults

`func NewCurrentlyPlayingContextObjectDeviceWithDefaults() *CurrentlyPlayingContextObjectDevice`

NewCurrentlyPlayingContextObjectDeviceWithDefaults instantiates a new CurrentlyPlayingContextObjectDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrentlyPlayingContextObjectDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentlyPlayingContextObjectDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentlyPlayingContextObjectDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurrentlyPlayingContextObjectDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CurrentlyPlayingContextObjectDevice) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CurrentlyPlayingContextObjectDevice) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsActive

`func (o *CurrentlyPlayingContextObjectDevice) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CurrentlyPlayingContextObjectDevice) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CurrentlyPlayingContextObjectDevice) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CurrentlyPlayingContextObjectDevice) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsPrivateSession

`func (o *CurrentlyPlayingContextObjectDevice) GetIsPrivateSession() bool`

GetIsPrivateSession returns the IsPrivateSession field if non-nil, zero value otherwise.

### GetIsPrivateSessionOk

`func (o *CurrentlyPlayingContextObjectDevice) GetIsPrivateSessionOk() (*bool, bool)`

GetIsPrivateSessionOk returns a tuple with the IsPrivateSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateSession

`func (o *CurrentlyPlayingContextObjectDevice) SetIsPrivateSession(v bool)`

SetIsPrivateSession sets IsPrivateSession field to given value.

### HasIsPrivateSession

`func (o *CurrentlyPlayingContextObjectDevice) HasIsPrivateSession() bool`

HasIsPrivateSession returns a boolean if a field has been set.

### GetIsRestricted

`func (o *CurrentlyPlayingContextObjectDevice) GetIsRestricted() bool`

GetIsRestricted returns the IsRestricted field if non-nil, zero value otherwise.

### GetIsRestrictedOk

`func (o *CurrentlyPlayingContextObjectDevice) GetIsRestrictedOk() (*bool, bool)`

GetIsRestrictedOk returns a tuple with the IsRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestricted

`func (o *CurrentlyPlayingContextObjectDevice) SetIsRestricted(v bool)`

SetIsRestricted sets IsRestricted field to given value.

### HasIsRestricted

`func (o *CurrentlyPlayingContextObjectDevice) HasIsRestricted() bool`

HasIsRestricted returns a boolean if a field has been set.

### GetName

`func (o *CurrentlyPlayingContextObjectDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentlyPlayingContextObjectDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentlyPlayingContextObjectDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrentlyPlayingContextObjectDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CurrentlyPlayingContextObjectDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CurrentlyPlayingContextObjectDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CurrentlyPlayingContextObjectDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CurrentlyPlayingContextObjectDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolumePercent

`func (o *CurrentlyPlayingContextObjectDevice) GetVolumePercent() int32`

GetVolumePercent returns the VolumePercent field if non-nil, zero value otherwise.

### GetVolumePercentOk

`func (o *CurrentlyPlayingContextObjectDevice) GetVolumePercentOk() (*int32, bool)`

GetVolumePercentOk returns a tuple with the VolumePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePercent

`func (o *CurrentlyPlayingContextObjectDevice) SetVolumePercent(v int32)`

SetVolumePercent sets VolumePercent field to given value.

### HasVolumePercent

`func (o *CurrentlyPlayingContextObjectDevice) HasVolumePercent() bool`

HasVolumePercent returns a boolean if a field has been set.

### SetVolumePercentNil

`func (o *CurrentlyPlayingContextObjectDevice) SetVolumePercentNil(b bool)`

 SetVolumePercentNil sets the value for VolumePercent to be an explicit nil

### UnsetVolumePercent
`func (o *CurrentlyPlayingContextObjectDevice) UnsetVolumePercent()`

UnsetVolumePercent ensures that no value is present for VolumePercent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


