# AudioAnalysisObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**AudioAnalysisObjectMeta**](AudioAnalysisObjectMeta.md) |  | [optional] 
**Track** | Pointer to [**AudioAnalysisObjectTrack**](AudioAnalysisObjectTrack.md) |  | [optional] 
**Bars** | Pointer to [**[]TimeIntervalObject**](TimeIntervalObject.md) | The time intervals of the bars throughout the track. A bar (or measure) is a segment of time defined as a given number of beats. | [optional] 
**Beats** | Pointer to [**[]TimeIntervalObject**](TimeIntervalObject.md) | The time intervals of beats throughout the track. A beat is the basic time unit of a piece of music; for example, each tick of a metronome. Beats are typically multiples of tatums. | [optional] 
**Sections** | Pointer to [**[]SectionObject**](SectionObject.md) | Sections are defined by large variations in rhythm or timbre, e.g. chorus, verse, bridge, guitar solo, etc. Each section contains its own descriptions of tempo, key, mode, time_signature, and loudness. | [optional] 
**Segments** | Pointer to [**[]SegmentObject**](SegmentObject.md) | Each segment contains a roughly conisistent sound throughout its duration. | [optional] 
**Tatums** | Pointer to [**[]TimeIntervalObject**](TimeIntervalObject.md) | A tatum represents the lowest regular pulse train that a listener intuitively infers from the timing of perceived musical events (segments). | [optional] 

## Methods

### NewAudioAnalysisObject

`func NewAudioAnalysisObject() *AudioAnalysisObject`

NewAudioAnalysisObject instantiates a new AudioAnalysisObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioAnalysisObjectWithDefaults

`func NewAudioAnalysisObjectWithDefaults() *AudioAnalysisObject`

NewAudioAnalysisObjectWithDefaults instantiates a new AudioAnalysisObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AudioAnalysisObject) GetMeta() AudioAnalysisObjectMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AudioAnalysisObject) GetMetaOk() (*AudioAnalysisObjectMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AudioAnalysisObject) SetMeta(v AudioAnalysisObjectMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AudioAnalysisObject) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTrack

`func (o *AudioAnalysisObject) GetTrack() AudioAnalysisObjectTrack`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *AudioAnalysisObject) GetTrackOk() (*AudioAnalysisObjectTrack, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *AudioAnalysisObject) SetTrack(v AudioAnalysisObjectTrack)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *AudioAnalysisObject) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetBars

`func (o *AudioAnalysisObject) GetBars() []TimeIntervalObject`

GetBars returns the Bars field if non-nil, zero value otherwise.

### GetBarsOk

`func (o *AudioAnalysisObject) GetBarsOk() (*[]TimeIntervalObject, bool)`

GetBarsOk returns a tuple with the Bars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBars

`func (o *AudioAnalysisObject) SetBars(v []TimeIntervalObject)`

SetBars sets Bars field to given value.

### HasBars

`func (o *AudioAnalysisObject) HasBars() bool`

HasBars returns a boolean if a field has been set.

### GetBeats

`func (o *AudioAnalysisObject) GetBeats() []TimeIntervalObject`

GetBeats returns the Beats field if non-nil, zero value otherwise.

### GetBeatsOk

`func (o *AudioAnalysisObject) GetBeatsOk() (*[]TimeIntervalObject, bool)`

GetBeatsOk returns a tuple with the Beats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeats

`func (o *AudioAnalysisObject) SetBeats(v []TimeIntervalObject)`

SetBeats sets Beats field to given value.

### HasBeats

`func (o *AudioAnalysisObject) HasBeats() bool`

HasBeats returns a boolean if a field has been set.

### GetSections

`func (o *AudioAnalysisObject) GetSections() []SectionObject`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *AudioAnalysisObject) GetSectionsOk() (*[]SectionObject, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *AudioAnalysisObject) SetSections(v []SectionObject)`

SetSections sets Sections field to given value.

### HasSections

`func (o *AudioAnalysisObject) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSegments

`func (o *AudioAnalysisObject) GetSegments() []SegmentObject`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *AudioAnalysisObject) GetSegmentsOk() (*[]SegmentObject, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *AudioAnalysisObject) SetSegments(v []SegmentObject)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *AudioAnalysisObject) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetTatums

`func (o *AudioAnalysisObject) GetTatums() []TimeIntervalObject`

GetTatums returns the Tatums field if non-nil, zero value otherwise.

### GetTatumsOk

`func (o *AudioAnalysisObject) GetTatumsOk() (*[]TimeIntervalObject, bool)`

GetTatumsOk returns a tuple with the Tatums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTatums

`func (o *AudioAnalysisObject) SetTatums(v []TimeIntervalObject)`

SetTatums sets Tatums field to given value.

### HasTatums

`func (o *AudioAnalysisObject) HasTatums() bool`

HasTatums returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


