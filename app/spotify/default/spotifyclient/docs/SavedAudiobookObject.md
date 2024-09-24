# SavedAudiobookObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAt** | Pointer to **time.Time** | The date and time the audiobook was saved Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.  | [optional] 
**Audiobook** | Pointer to [**AudiobookObject**](AudiobookObject.md) | Information about the audiobook. | [optional] 

## Methods

### NewSavedAudiobookObject

`func NewSavedAudiobookObject() *SavedAudiobookObject`

NewSavedAudiobookObject instantiates a new SavedAudiobookObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedAudiobookObjectWithDefaults

`func NewSavedAudiobookObjectWithDefaults() *SavedAudiobookObject`

NewSavedAudiobookObjectWithDefaults instantiates a new SavedAudiobookObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAt

`func (o *SavedAudiobookObject) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *SavedAudiobookObject) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *SavedAudiobookObject) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *SavedAudiobookObject) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetAudiobook

`func (o *SavedAudiobookObject) GetAudiobook() AudiobookObject`

GetAudiobook returns the Audiobook field if non-nil, zero value otherwise.

### GetAudiobookOk

`func (o *SavedAudiobookObject) GetAudiobookOk() (*AudiobookObject, bool)`

GetAudiobookOk returns a tuple with the Audiobook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiobook

`func (o *SavedAudiobookObject) SetAudiobook(v AudiobookObject)`

SetAudiobook sets Audiobook field to given value.

### HasAudiobook

`func (o *SavedAudiobookObject) HasAudiobook() bool`

HasAudiobook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


