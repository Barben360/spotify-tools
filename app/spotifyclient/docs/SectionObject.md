# SectionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **float32** | The starting point (in seconds) of the section. | [optional] 
**Duration** | Pointer to **float32** | The duration (in seconds) of the section. | [optional] 
**Confidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the section&#39;s \&quot;designation\&quot;. | [optional] 
**Loudness** | Pointer to **float32** | The overall loudness of the section in decibels (dB). Loudness values are useful for comparing relative loudness of sections within tracks. | [optional] 
**Tempo** | Pointer to **float32** | The overall estimated tempo of the section in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration. | [optional] 
**TempoConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the tempo. Some tracks contain tempo changes or sounds which don&#39;t contain tempo (like pure speech) which would correspond to a low value in this field. | [optional] 
**Key** | Pointer to **int32** | The estimated overall key of the section. The values in this field ranging from 0 to 11 mapping to pitches using standard Pitch Class notation (E.g. 0 &#x3D; C, 1 &#x3D; C♯/D♭, 2 &#x3D; D, and so on). If no key was detected, the value is -1. | [optional] 
**KeyConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the key. Songs with many key changes may correspond to low values in this field. | [optional] 
**Mode** | Pointer to **float32** | Indicates the modality (major or minor) of a section, the type of scale from which its melodic content is derived. This field will contain a 0 for \&quot;minor\&quot;, a 1 for \&quot;major\&quot;, or a -1 for no result. Note that the major key (e.g. C major) could more likely be confused with the minor key at 3 semitones lower (e.g. A minor) as both keys carry the same pitches. | [optional] 
**ModeConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the &#x60;mode&#x60;. | [optional] 
**TimeSignature** | Pointer to **int32** | An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of \&quot;3/4\&quot;, to \&quot;7/4\&quot;. | [optional] 
**TimeSignatureConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the &#x60;time_signature&#x60;. Sections with time signature changes may correspond to low values in this field. | [optional] 

## Methods

### NewSectionObject

`func NewSectionObject() *SectionObject`

NewSectionObject instantiates a new SectionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionObjectWithDefaults

`func NewSectionObjectWithDefaults() *SectionObject`

NewSectionObjectWithDefaults instantiates a new SectionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *SectionObject) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SectionObject) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SectionObject) SetStart(v float32)`

SetStart sets Start field to given value.

### HasStart

`func (o *SectionObject) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetDuration

`func (o *SectionObject) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SectionObject) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SectionObject) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SectionObject) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetConfidence

`func (o *SectionObject) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *SectionObject) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *SectionObject) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *SectionObject) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetLoudness

`func (o *SectionObject) GetLoudness() float32`

GetLoudness returns the Loudness field if non-nil, zero value otherwise.

### GetLoudnessOk

`func (o *SectionObject) GetLoudnessOk() (*float32, bool)`

GetLoudnessOk returns a tuple with the Loudness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudness

`func (o *SectionObject) SetLoudness(v float32)`

SetLoudness sets Loudness field to given value.

### HasLoudness

`func (o *SectionObject) HasLoudness() bool`

HasLoudness returns a boolean if a field has been set.

### GetTempo

`func (o *SectionObject) GetTempo() float32`

GetTempo returns the Tempo field if non-nil, zero value otherwise.

### GetTempoOk

`func (o *SectionObject) GetTempoOk() (*float32, bool)`

GetTempoOk returns a tuple with the Tempo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempo

`func (o *SectionObject) SetTempo(v float32)`

SetTempo sets Tempo field to given value.

### HasTempo

`func (o *SectionObject) HasTempo() bool`

HasTempo returns a boolean if a field has been set.

### GetTempoConfidence

`func (o *SectionObject) GetTempoConfidence() float32`

GetTempoConfidence returns the TempoConfidence field if non-nil, zero value otherwise.

### GetTempoConfidenceOk

`func (o *SectionObject) GetTempoConfidenceOk() (*float32, bool)`

GetTempoConfidenceOk returns a tuple with the TempoConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempoConfidence

`func (o *SectionObject) SetTempoConfidence(v float32)`

SetTempoConfidence sets TempoConfidence field to given value.

### HasTempoConfidence

`func (o *SectionObject) HasTempoConfidence() bool`

HasTempoConfidence returns a boolean if a field has been set.

### GetKey

`func (o *SectionObject) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SectionObject) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SectionObject) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *SectionObject) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyConfidence

`func (o *SectionObject) GetKeyConfidence() float32`

GetKeyConfidence returns the KeyConfidence field if non-nil, zero value otherwise.

### GetKeyConfidenceOk

`func (o *SectionObject) GetKeyConfidenceOk() (*float32, bool)`

GetKeyConfidenceOk returns a tuple with the KeyConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyConfidence

`func (o *SectionObject) SetKeyConfidence(v float32)`

SetKeyConfidence sets KeyConfidence field to given value.

### HasKeyConfidence

`func (o *SectionObject) HasKeyConfidence() bool`

HasKeyConfidence returns a boolean if a field has been set.

### GetMode

`func (o *SectionObject) GetMode() float32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SectionObject) GetModeOk() (*float32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SectionObject) SetMode(v float32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SectionObject) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModeConfidence

`func (o *SectionObject) GetModeConfidence() float32`

GetModeConfidence returns the ModeConfidence field if non-nil, zero value otherwise.

### GetModeConfidenceOk

`func (o *SectionObject) GetModeConfidenceOk() (*float32, bool)`

GetModeConfidenceOk returns a tuple with the ModeConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeConfidence

`func (o *SectionObject) SetModeConfidence(v float32)`

SetModeConfidence sets ModeConfidence field to given value.

### HasModeConfidence

`func (o *SectionObject) HasModeConfidence() bool`

HasModeConfidence returns a boolean if a field has been set.

### GetTimeSignature

`func (o *SectionObject) GetTimeSignature() int32`

GetTimeSignature returns the TimeSignature field if non-nil, zero value otherwise.

### GetTimeSignatureOk

`func (o *SectionObject) GetTimeSignatureOk() (*int32, bool)`

GetTimeSignatureOk returns a tuple with the TimeSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSignature

`func (o *SectionObject) SetTimeSignature(v int32)`

SetTimeSignature sets TimeSignature field to given value.

### HasTimeSignature

`func (o *SectionObject) HasTimeSignature() bool`

HasTimeSignature returns a boolean if a field has been set.

### GetTimeSignatureConfidence

`func (o *SectionObject) GetTimeSignatureConfidence() float32`

GetTimeSignatureConfidence returns the TimeSignatureConfidence field if non-nil, zero value otherwise.

### GetTimeSignatureConfidenceOk

`func (o *SectionObject) GetTimeSignatureConfidenceOk() (*float32, bool)`

GetTimeSignatureConfidenceOk returns a tuple with the TimeSignatureConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSignatureConfidence

`func (o *SectionObject) SetTimeSignatureConfidence(v float32)`

SetTimeSignatureConfidence sets TimeSignatureConfidence field to given value.

### HasTimeSignatureConfidence

`func (o *SectionObject) HasTimeSignatureConfidence() bool`

HasTimeSignatureConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


