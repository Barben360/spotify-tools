# SavedShowObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAt** | Pointer to **time.Time** | The date and time the show was saved. Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.  | [optional] 
**Show** | Pointer to [**SimplifiedShowObject**](SimplifiedShowObject.md) | Information about the show. | [optional] 

## Methods

### NewSavedShowObject

`func NewSavedShowObject() *SavedShowObject`

NewSavedShowObject instantiates a new SavedShowObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedShowObjectWithDefaults

`func NewSavedShowObjectWithDefaults() *SavedShowObject`

NewSavedShowObjectWithDefaults instantiates a new SavedShowObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAt

`func (o *SavedShowObject) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *SavedShowObject) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *SavedShowObject) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *SavedShowObject) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetShow

`func (o *SavedShowObject) GetShow() SimplifiedShowObject`

GetShow returns the Show field if non-nil, zero value otherwise.

### GetShowOk

`func (o *SavedShowObject) GetShowOk() (*SimplifiedShowObject, bool)`

GetShowOk returns a tuple with the Show field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShow

`func (o *SavedShowObject) SetShow(v SimplifiedShowObject)`

SetShow sets Show field to given value.

### HasShow

`func (o *SavedShowObject) HasShow() bool`

HasShow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


