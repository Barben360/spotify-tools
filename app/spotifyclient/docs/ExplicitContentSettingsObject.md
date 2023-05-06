# ExplicitContentSettingsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterEnabled** | Pointer to **bool** | When &#x60;true&#x60;, indicates that explicit content should not be played.  | [optional] 
**FilterLocked** | Pointer to **bool** | When &#x60;true&#x60;, indicates that the explicit content setting is locked and can&#39;t be changed by the user.  | [optional] 

## Methods

### NewExplicitContentSettingsObject

`func NewExplicitContentSettingsObject() *ExplicitContentSettingsObject`

NewExplicitContentSettingsObject instantiates a new ExplicitContentSettingsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplicitContentSettingsObjectWithDefaults

`func NewExplicitContentSettingsObjectWithDefaults() *ExplicitContentSettingsObject`

NewExplicitContentSettingsObjectWithDefaults instantiates a new ExplicitContentSettingsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterEnabled

`func (o *ExplicitContentSettingsObject) GetFilterEnabled() bool`

GetFilterEnabled returns the FilterEnabled field if non-nil, zero value otherwise.

### GetFilterEnabledOk

`func (o *ExplicitContentSettingsObject) GetFilterEnabledOk() (*bool, bool)`

GetFilterEnabledOk returns a tuple with the FilterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterEnabled

`func (o *ExplicitContentSettingsObject) SetFilterEnabled(v bool)`

SetFilterEnabled sets FilterEnabled field to given value.

### HasFilterEnabled

`func (o *ExplicitContentSettingsObject) HasFilterEnabled() bool`

HasFilterEnabled returns a boolean if a field has been set.

### GetFilterLocked

`func (o *ExplicitContentSettingsObject) GetFilterLocked() bool`

GetFilterLocked returns the FilterLocked field if non-nil, zero value otherwise.

### GetFilterLockedOk

`func (o *ExplicitContentSettingsObject) GetFilterLockedOk() (*bool, bool)`

GetFilterLockedOk returns a tuple with the FilterLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLocked

`func (o *ExplicitContentSettingsObject) SetFilterLocked(v bool)`

SetFilterLocked sets FilterLocked field to given value.

### HasFilterLocked

`func (o *ExplicitContentSettingsObject) HasFilterLocked() bool`

HasFilterLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


