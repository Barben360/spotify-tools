# RecommendationsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seeds** | [**[]RecommendationSeedObject**](RecommendationSeedObject.md) | An array of recommendation seed objects.  | 
**Tracks** | [**[]TrackObject**](TrackObject.md) | An array of track objects ordered according to the parameters supplied.  | 

## Methods

### NewRecommendationsObject

`func NewRecommendationsObject(seeds []RecommendationSeedObject, tracks []TrackObject, ) *RecommendationsObject`

NewRecommendationsObject instantiates a new RecommendationsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationsObjectWithDefaults

`func NewRecommendationsObjectWithDefaults() *RecommendationsObject`

NewRecommendationsObjectWithDefaults instantiates a new RecommendationsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeeds

`func (o *RecommendationsObject) GetSeeds() []RecommendationSeedObject`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *RecommendationsObject) GetSeedsOk() (*[]RecommendationSeedObject, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *RecommendationsObject) SetSeeds(v []RecommendationSeedObject)`

SetSeeds sets Seeds field to given value.


### GetTracks

`func (o *RecommendationsObject) GetTracks() []TrackObject`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *RecommendationsObject) GetTracksOk() (*[]TrackObject, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *RecommendationsObject) SetTracks(v []TrackObject)`

SetTracks sets Tracks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


