# PlayHistoryObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Track** | Pointer to [**PlayHistoryObjectTrack**](PlayHistoryObjectTrack.md) |  | [optional] 
**PlayedAt** | Pointer to **time.Time** | The date and time the track was played. | [optional] 
**Context** | Pointer to [**PlayHistoryObjectContext**](PlayHistoryObjectContext.md) |  | [optional] 

## Methods

### NewPlayHistoryObject

`func NewPlayHistoryObject() *PlayHistoryObject`

NewPlayHistoryObject instantiates a new PlayHistoryObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayHistoryObjectWithDefaults

`func NewPlayHistoryObjectWithDefaults() *PlayHistoryObject`

NewPlayHistoryObjectWithDefaults instantiates a new PlayHistoryObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrack

`func (o *PlayHistoryObject) GetTrack() PlayHistoryObjectTrack`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *PlayHistoryObject) GetTrackOk() (*PlayHistoryObjectTrack, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *PlayHistoryObject) SetTrack(v PlayHistoryObjectTrack)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *PlayHistoryObject) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetPlayedAt

`func (o *PlayHistoryObject) GetPlayedAt() time.Time`

GetPlayedAt returns the PlayedAt field if non-nil, zero value otherwise.

### GetPlayedAtOk

`func (o *PlayHistoryObject) GetPlayedAtOk() (*time.Time, bool)`

GetPlayedAtOk returns a tuple with the PlayedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedAt

`func (o *PlayHistoryObject) SetPlayedAt(v time.Time)`

SetPlayedAt sets PlayedAt field to given value.

### HasPlayedAt

`func (o *PlayHistoryObject) HasPlayedAt() bool`

HasPlayedAt returns a boolean if a field has been set.

### GetContext

`func (o *PlayHistoryObject) GetContext() PlayHistoryObjectContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PlayHistoryObject) GetContextOk() (*PlayHistoryObjectContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PlayHistoryObject) SetContext(v PlayHistoryObjectContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PlayHistoryObject) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


