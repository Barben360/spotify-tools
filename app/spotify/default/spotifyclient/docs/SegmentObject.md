# SegmentObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **float32** | The starting point (in seconds) of the segment. | [optional] 
**Duration** | Pointer to **float32** | The duration (in seconds) of the segment. | [optional] 
**Confidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the segmentation. Segments of the song which are difficult to logically segment (e.g: noise) may correspond to low values in this field.  | [optional] 
**LoudnessStart** | Pointer to **float32** | The onset loudness of the segment in decibels (dB). Combined with &#x60;loudness_max&#x60; and &#x60;loudness_max_time&#x60;, these components can be used to describe the \&quot;attack\&quot; of the segment. | [optional] 
**LoudnessMax** | Pointer to **float32** | The peak loudness of the segment in decibels (dB). Combined with &#x60;loudness_start&#x60; and &#x60;loudness_max_time&#x60;, these components can be used to describe the \&quot;attack\&quot; of the segment. | [optional] 
**LoudnessMaxTime** | Pointer to **float32** | The segment-relative offset of the segment peak loudness in seconds. Combined with &#x60;loudness_start&#x60; and &#x60;loudness_max&#x60;, these components can be used to desctibe the \&quot;attack\&quot; of the segment. | [optional] 
**LoudnessEnd** | Pointer to **float32** | The offset loudness of the segment in decibels (dB). This value should be equivalent to the loudness_start of the following segment. | [optional] 
**Pitches** | Pointer to **[]float32** | Pitch content is given by a “chroma” vector, corresponding to the 12 pitch classes C, C#, D to B, with values ranging from 0 to 1 that describe the relative dominance of every pitch in the chromatic scale. For example a C Major chord would likely be represented by large values of C, E and G (i.e. classes 0, 4, and 7).  Vectors are normalized to 1 by their strongest dimension, therefore noisy sounds are likely represented by values that are all close to 1, while pure tones are described by one value at 1 (the pitch) and others near 0. As can be seen below, the 12 vector indices are a combination of low-power spectrum values at their respective pitch frequencies. ![pitch vector](https://developer.spotify.com/assets/audio/Pitch_vector.png)  | [optional] 
**Timbre** | Pointer to **[]float32** | Timbre is the quality of a musical note or sound that distinguishes different types of musical instruments, or voices. It is a complex notion also referred to as sound color, texture, or tone quality, and is derived from the shape of a segment’s spectro-temporal surface, independently of pitch and loudness. The timbre feature is a vector that includes 12 unbounded values roughly centered around 0. Those values are high level abstractions of the spectral surface, ordered by degree of importance.  For completeness however, the first dimension represents the average loudness of the segment; second emphasizes brightness; third is more closely correlated to the flatness of a sound; fourth to sounds with a stronger attack; etc. See an image below representing the 12 basis functions (i.e. template segments). ![timbre basis functions](https://developer.spotify.com/assets/audio/Timbre_basis_functions.png)  The actual timbre of the segment is best described as a linear combination of these 12 basis functions weighted by the coefficient values: timbre &#x3D; c1 x b1 + c2 x b2 + ... + c12 x b12, where c1 to c12 represent the 12 coefficients and b1 to b12 the 12 basis functions as displayed below. Timbre vectors are best used in comparison with each other.  | [optional] 

## Methods

### NewSegmentObject

`func NewSegmentObject() *SegmentObject`

NewSegmentObject instantiates a new SegmentObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentObjectWithDefaults

`func NewSegmentObjectWithDefaults() *SegmentObject`

NewSegmentObjectWithDefaults instantiates a new SegmentObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *SegmentObject) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SegmentObject) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SegmentObject) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SegmentObject) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetDuration

`func (o *SegmentObject) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SegmentObject) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SegmentObject) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SegmentObject) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetConfidence

`func (o *SegmentObject) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SegmentObject) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SegmentObject) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SegmentObject) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetLoudnessStart

`func (o *SegmentObject) GetLoudnessStart() float32`

GetLoudnessStart returns the LoudnessStart field if non-nil, zero value otherwise.

### GetLoudnessStartOk

`func (o *SegmentObject) GetLoudnessStartOk() (*float32, bool)`

GetLoudnessStartOk returns a tuple with the LoudnessStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudnessStart

`func (o *SegmentObject) SetLoudnessStart(v float32)`

SetLoudnessStart sets LoudnessStart field to given value.

### HasLoudnessStart

`func (o *SegmentObject) HasLoudnessStart() bool`

HasLoudnessStart returns a boolean if a field has been set.

### GetLoudnessMax

`func (o *SegmentObject) GetLoudnessMax() float32`

GetLoudnessMax returns the LoudnessMax field if non-nil, zero value otherwise.

### GetLoudnessMaxOk

`func (o *SegmentObject) GetLoudnessMaxOk() (*float32, bool)`

GetLoudnessMaxOk returns a tuple with the LoudnessMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudnessMax

`func (o *SegmentObject) SetLoudnessMax(v float32)`

SetLoudnessMax sets LoudnessMax field to given value.

### HasLoudnessMax

`func (o *SegmentObject) HasLoudnessMax() bool`

HasLoudnessMax returns a boolean if a field has been set.

### GetLoudnessMaxTime

`func (o *SegmentObject) GetLoudnessMaxTime() float32`

GetLoudnessMaxTime returns the LoudnessMaxTime field if non-nil, zero value otherwise.

### GetLoudnessMaxTimeOk

`func (o *SegmentObject) GetLoudnessMaxTimeOk() (*float32, bool)`

GetLoudnessMaxTimeOk returns a tuple with the LoudnessMaxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudnessMaxTime

`func (o *SegmentObject) SetLoudnessMaxTime(v float32)`

SetLoudnessMaxTime sets LoudnessMaxTime field to given value.

### HasLoudnessMaxTime

`func (o *SegmentObject) HasLoudnessMaxTime() bool`

HasLoudnessMaxTime returns a boolean if a field has been set.

### GetLoudnessEnd

`func (o *SegmentObject) GetLoudnessEnd() float32`

GetLoudnessEnd returns the LoudnessEnd field if non-nil, zero value otherwise.

### GetLoudnessEndOk

`func (o *SegmentObject) GetLoudnessEndOk() (*float32, bool)`

GetLoudnessEndOk returns a tuple with the LoudnessEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudnessEnd

`func (o *SegmentObject) SetLoudnessEnd(v float32)`

SetLoudnessEnd sets LoudnessEnd field to given value.

### HasLoudnessEnd

`func (o *SegmentObject) HasLoudnessEnd() bool`

HasLoudnessEnd returns a boolean if a field has been set.

### GetPitches

`func (o *SegmentObject) GetPitches() []float32`

GetPitches returns the Pitches field if non-nil, zero value otherwise.

### GetPitchesOk

`func (o *SegmentObject) GetPitchesOk() (*[]float32, bool)`

GetPitchesOk returns a tuple with the Pitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPitches

`func (o *SegmentObject) SetPitches(v []float32)`

SetPitches sets Pitches field to given value.

### HasPitches

`func (o *SegmentObject) HasPitches() bool`

HasPitches returns a boolean if a field has been set.

### GetTimbre

`func (o *SegmentObject) GetTimbre() []float32`

GetTimbre returns the Timbre field if non-nil, zero value otherwise.

### GetTimbreOk

`func (o *SegmentObject) GetTimbreOk() (*[]float32, bool)`

GetTimbreOk returns a tuple with the Timbre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimbre

`func (o *SegmentObject) SetTimbre(v []float32)`

SetTimbre sets Timbre field to given value.

### HasTimbre

`func (o *SegmentObject) HasTimbre() bool`

HasTimbre returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


