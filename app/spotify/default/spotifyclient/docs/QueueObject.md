# QueueObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentlyPlaying** | Pointer to [**QueueObjectCurrentlyPlaying**](QueueObjectCurrentlyPlaying.md) |  | [optional] 
**Queue** | Pointer to [**[]QueueObjectQueueInner**](QueueObjectQueueInner.md) | The tracks or episodes in the queue. Can be empty. | [optional] 

## Methods

### NewQueueObject

`func NewQueueObject() *QueueObject`

NewQueueObject instantiates a new QueueObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueObjectWithDefaults

`func NewQueueObjectWithDefaults() *QueueObject`

NewQueueObjectWithDefaults instantiates a new QueueObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentlyPlaying

`func (o *QueueObject) GetCurrentlyPlaying() QueueObjectCurrentlyPlaying`

GetCurrentlyPlaying returns the CurrentlyPlaying field if non-nil, zero value otherwise.

### GetCurrentlyPlayingOk

`func (o *QueueObject) GetCurrentlyPlayingOk() (*QueueObjectCurrentlyPlaying, bool)`

GetCurrentlyPlayingOk returns a tuple with the CurrentlyPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentlyPlaying

`func (o *QueueObject) SetCurrentlyPlaying(v QueueObjectCurrentlyPlaying)`

SetCurrentlyPlaying sets CurrentlyPlaying field to given value.

### HasCurrentlyPlaying

`func (o *QueueObject) HasCurrentlyPlaying() bool`

HasCurrentlyPlaying returns a boolean if a field has been set.

### GetQueue

`func (o *QueueObject) GetQueue() []QueueObjectQueueInner`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *QueueObject) GetQueueOk() (*[]QueueObjectQueueInner, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *QueueObject) SetQueue(v []QueueObjectQueueInner)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *QueueObject) HasQueue() bool`

HasQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


