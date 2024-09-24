# SavedEpisodeObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAt** | Pointer to **time.Time** | The date and time the episode was saved. Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ.  | [optional] 
**Episode** | Pointer to [**EpisodeObject**](EpisodeObject.md) | Information about the episode. | [optional] 

## Methods

### NewSavedEpisodeObject

`func NewSavedEpisodeObject() *SavedEpisodeObject`

NewSavedEpisodeObject instantiates a new SavedEpisodeObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedEpisodeObjectWithDefaults

`func NewSavedEpisodeObjectWithDefaults() *SavedEpisodeObject`

NewSavedEpisodeObjectWithDefaults instantiates a new SavedEpisodeObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAt

`func (o *SavedEpisodeObject) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *SavedEpisodeObject) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *SavedEpisodeObject) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *SavedEpisodeObject) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetEpisode

`func (o *SavedEpisodeObject) GetEpisode() EpisodeObject`

GetEpisode returns the Episode field if non-nil, zero value otherwise.

### GetEpisodeOk

`func (o *SavedEpisodeObject) GetEpisodeOk() (*EpisodeObject, bool)`

GetEpisodeOk returns a tuple with the Episode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisode

`func (o *SavedEpisodeObject) SetEpisode(v EpisodeObject)`

SetEpisode sets Episode field to given value.

### HasEpisode

`func (o *SavedEpisodeObject) HasEpisode() bool`

HasEpisode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


