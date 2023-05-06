# RecommendationSeedObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterFilteringSize** | Pointer to **int32** | The number of tracks available after min\\_\\* and max\\_\\* filters have been applied.  | [optional] 
**AfterRelinkingSize** | Pointer to **int32** | The number of tracks available after relinking for regional availability.  | [optional] 
**Href** | Pointer to **string** | A link to the full track or artist data for this seed. For tracks this will be a link to a Track Object. For artists a link to an Artist Object. For genre seeds, this value will be &#x60;null&#x60;.  | [optional] 
**Id** | Pointer to **string** | The id used to select this seed. This will be the same as the string used in the &#x60;seed_artists&#x60;, &#x60;seed_tracks&#x60; or &#x60;seed_genres&#x60; parameter.  | [optional] 
**InitialPoolSize** | Pointer to **int32** | The number of recommended tracks available for this seed.  | [optional] 
**Type** | Pointer to **string** | The entity type of this seed. One of &#x60;artist&#x60;, &#x60;track&#x60; or &#x60;genre&#x60;.  | [optional] 

## Methods

### NewRecommendationSeedObject

`func NewRecommendationSeedObject() *RecommendationSeedObject`

NewRecommendationSeedObject instantiates a new RecommendationSeedObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationSeedObjectWithDefaults

`func NewRecommendationSeedObjectWithDefaults() *RecommendationSeedObject`

NewRecommendationSeedObjectWithDefaults instantiates a new RecommendationSeedObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterFilteringSize

`func (o *RecommendationSeedObject) GetAfterFilteringSize() int32`

GetAfterFilteringSize returns the AfterFilteringSize field if non-nil, zero value otherwise.

### GetAfterFilteringSizeOk

`func (o *RecommendationSeedObject) GetAfterFilteringSizeOk() (*int32, bool)`

GetAfterFilteringSizeOk returns a tuple with the AfterFilteringSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterFilteringSize

`func (o *RecommendationSeedObject) SetAfterFilteringSize(v int32)`

SetAfterFilteringSize sets AfterFilteringSize field to given value.

### HasAfterFilteringSize

`func (o *RecommendationSeedObject) HasAfterFilteringSize() bool`

HasAfterFilteringSize returns a boolean if a field has been set.

### GetAfterRelinkingSize

`func (o *RecommendationSeedObject) GetAfterRelinkingSize() int32`

GetAfterRelinkingSize returns the AfterRelinkingSize field if non-nil, zero value otherwise.

### GetAfterRelinkingSizeOk

`func (o *RecommendationSeedObject) GetAfterRelinkingSizeOk() (*int32, bool)`

GetAfterRelinkingSizeOk returns a tuple with the AfterRelinkingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterRelinkingSize

`func (o *RecommendationSeedObject) SetAfterRelinkingSize(v int32)`

SetAfterRelinkingSize sets AfterRelinkingSize field to given value.

### HasAfterRelinkingSize

`func (o *RecommendationSeedObject) HasAfterRelinkingSize() bool`

HasAfterRelinkingSize returns a boolean if a field has been set.

### GetHref

`func (o *RecommendationSeedObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RecommendationSeedObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RecommendationSeedObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RecommendationSeedObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *RecommendationSeedObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecommendationSeedObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecommendationSeedObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecommendationSeedObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *RecommendationSeedObject) GetInitialPoolSize() int32`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *RecommendationSeedObject) GetInitialPoolSizeOk() (*int32, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *RecommendationSeedObject) SetInitialPoolSize(v int32)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *RecommendationSeedObject) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetType

`func (o *RecommendationSeedObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecommendationSeedObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecommendationSeedObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecommendationSeedObject) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


