# SavedTrackObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedAt** | Pointer to **time.Time** | The date and time the track was saved. Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object.  | [optional] 
**Track** | Pointer to [**TrackObject**](TrackObject.md) | Information about the track. | [optional] 

## Methods

### NewSavedTrackObject

`func NewSavedTrackObject() *SavedTrackObject`

NewSavedTrackObject instantiates a new SavedTrackObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedTrackObjectWithDefaults

`func NewSavedTrackObjectWithDefaults() *SavedTrackObject`

NewSavedTrackObjectWithDefaults instantiates a new SavedTrackObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedAt

`func (o *SavedTrackObject) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *SavedTrackObject) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *SavedTrackObject) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *SavedTrackObject) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetTrack

`func (o *SavedTrackObject) GetTrack() TrackObject`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *SavedTrackObject) GetTrackOk() (*TrackObject, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *SavedTrackObject) SetTrack(v TrackObject)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *SavedTrackObject) HasTrack() bool`

HasTrack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


