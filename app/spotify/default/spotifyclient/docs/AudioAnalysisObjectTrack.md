# AudioAnalysisObjectTrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumSamples** | Pointer to **int32** | The exact number of audio samples analyzed from this track. See also &#x60;analysis_sample_rate&#x60;. | [optional] 
**Duration** | Pointer to **float32** | Length of the track in seconds. | [optional] 
**SampleMd5** | Pointer to **string** | This field will always contain the empty string. | [optional] 
**OffsetSeconds** | Pointer to **int32** | An offset to the start of the region of the track that was analyzed. (As the entire track is analyzed, this should always be 0.) | [optional] 
**WindowSeconds** | Pointer to **int32** | The length of the region of the track was analyzed, if a subset of the track was analyzed. (As the entire track is analyzed, this should always be 0.) | [optional] 
**AnalysisSampleRate** | Pointer to **int32** | The sample rate used to decode and analyze this track. May differ from the actual sample rate of this track available on Spotify. | [optional] 
**AnalysisChannels** | Pointer to **int32** | The number of channels used for analysis. If 1, all channels are summed together to mono before analysis. | [optional] 
**EndOfFadeIn** | Pointer to **float32** | The time, in seconds, at which the track&#39;s fade-in period ends. If the track has no fade-in, this will be 0.0. | [optional] 
**StartOfFadeOut** | Pointer to **float32** | The time, in seconds, at which the track&#39;s fade-out period starts. If the track has no fade-out, this should match the track&#39;s length. | [optional] 
**Loudness** | Pointer to **float32** | The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db.  | [optional] 
**Tempo** | Pointer to **float32** | The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.  | [optional] 
**TempoConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the &#x60;tempo&#x60;. | [optional] 
**TimeSignature** | Pointer to **int32** | An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of \&quot;3/4\&quot;, to \&quot;7/4\&quot;. | [optional] 
**TimeSignatureConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the &#x60;time_signature&#x60;. | [optional] 
**Key** | Pointer to **int32** | The key the track is in. Integers map to pitches using standard [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 &#x3D; C, 1 &#x3D; C♯/D♭, 2 &#x3D; D, and so on. If no key was detected, the value is -1.  | [optional] 
**KeyConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the &#x60;key&#x60;. | [optional] 
**Mode** | Pointer to **int32** | Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0.  | [optional] 
**ModeConfidence** | Pointer to **float32** | The confidence, from 0.0 to 1.0, of the reliability of the &#x60;mode&#x60;. | [optional] 
**Codestring** | Pointer to **string** | An [Echo Nest Musical Fingerprint (ENMFP)](https://academiccommons.columbia.edu/doi/10.7916/D8Q248M4) codestring for this track. | [optional] 
**CodeVersion** | Pointer to **float32** | A version number for the Echo Nest Musical Fingerprint format used in the codestring field. | [optional] 
**Echoprintstring** | Pointer to **string** | An [EchoPrint](https://github.com/spotify/echoprint-codegen) codestring for this track. | [optional] 
**EchoprintVersion** | Pointer to **float32** | A version number for the EchoPrint format used in the echoprintstring field. | [optional] 
**Synchstring** | Pointer to **string** | A [Synchstring](https://github.com/echonest/synchdata) for this track. | [optional] 
**SynchVersion** | Pointer to **float32** | A version number for the Synchstring used in the synchstring field. | [optional] 
**Rhythmstring** | Pointer to **string** | A Rhythmstring for this track. The format of this string is similar to the Synchstring. | [optional] 
**RhythmVersion** | Pointer to **float32** | A version number for the Rhythmstring used in the rhythmstring field. | [optional] 

## Methods

### NewAudioAnalysisObjectTrack

`func NewAudioAnalysisObjectTrack() *AudioAnalysisObjectTrack`

NewAudioAnalysisObjectTrack instantiates a new AudioAnalysisObjectTrack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioAnalysisObjectTrackWithDefaults

`func NewAudioAnalysisObjectTrackWithDefaults() *AudioAnalysisObjectTrack`

NewAudioAnalysisObjectTrackWithDefaults instantiates a new AudioAnalysisObjectTrack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumSamples

`func (o *AudioAnalysisObjectTrack) GetNumSamples() int32`

GetNumSamples returns the NumSamples field if non-nil, zero value otherwise.

### GetNumSamplesOk

`func (o *AudioAnalysisObjectTrack) GetNumSamplesOk() (*int32, bool)`

GetNumSamplesOk returns a tuple with the NumSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSamples

`func (o *AudioAnalysisObjectTrack) SetNumSamples(v int32)`

SetNumSamples sets NumSamples field to given value.

### HasNumSamples

`func (o *AudioAnalysisObjectTrack) HasNumSamples() bool`

HasNumSamples returns a boolean if a field has been set.

### GetDuration

`func (o *AudioAnalysisObjectTrack) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AudioAnalysisObjectTrack) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AudioAnalysisObjectTrack) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AudioAnalysisObjectTrack) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetSampleMd5

`func (o *AudioAnalysisObjectTrack) GetSampleMd5() string`

GetSampleMd5 returns the SampleMd5 field if non-nil, zero value otherwise.

### GetSampleMd5Ok

`func (o *AudioAnalysisObjectTrack) GetSampleMd5Ok() (*string, bool)`

GetSampleMd5Ok returns a tuple with the SampleMd5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleMd5

`func (o *AudioAnalysisObjectTrack) SetSampleMd5(v string)`

SetSampleMd5 sets SampleMd5 field to given value.

### HasSampleMd5

`func (o *AudioAnalysisObjectTrack) HasSampleMd5() bool`

HasSampleMd5 returns a boolean if a field has been set.

### GetOffsetSeconds

`func (o *AudioAnalysisObjectTrack) GetOffsetSeconds() int32`

GetOffsetSeconds returns the OffsetSeconds field if non-nil, zero value otherwise.

### GetOffsetSecondsOk

`func (o *AudioAnalysisObjectTrack) GetOffsetSecondsOk() (*int32, bool)`

GetOffsetSecondsOk returns a tuple with the OffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetSeconds

`func (o *AudioAnalysisObjectTrack) SetOffsetSeconds(v int32)`

SetOffsetSeconds sets OffsetSeconds field to given value.

### HasOffsetSeconds

`func (o *AudioAnalysisObjectTrack) HasOffsetSeconds() bool`

HasOffsetSeconds returns a boolean if a field has been set.

### GetWindowSeconds

`func (o *AudioAnalysisObjectTrack) GetWindowSeconds() int32`

GetWindowSeconds returns the WindowSeconds field if non-nil, zero value otherwise.

### GetWindowSecondsOk

`func (o *AudioAnalysisObjectTrack) GetWindowSecondsOk() (*int32, bool)`

GetWindowSecondsOk returns a tuple with the WindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSeconds

`func (o *AudioAnalysisObjectTrack) SetWindowSeconds(v int32)`

SetWindowSeconds sets WindowSeconds field to given value.

### HasWindowSeconds

`func (o *AudioAnalysisObjectTrack) HasWindowSeconds() bool`

HasWindowSeconds returns a boolean if a field has been set.

### GetAnalysisSampleRate

`func (o *AudioAnalysisObjectTrack) GetAnalysisSampleRate() int32`

GetAnalysisSampleRate returns the AnalysisSampleRate field if non-nil, zero value otherwise.

### GetAnalysisSampleRateOk

`func (o *AudioAnalysisObjectTrack) GetAnalysisSampleRateOk() (*int32, bool)`

GetAnalysisSampleRateOk returns a tuple with the AnalysisSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisSampleRate

`func (o *AudioAnalysisObjectTrack) SetAnalysisSampleRate(v int32)`

SetAnalysisSampleRate sets AnalysisSampleRate field to given value.

### HasAnalysisSampleRate

`func (o *AudioAnalysisObjectTrack) HasAnalysisSampleRate() bool`

HasAnalysisSampleRate returns a boolean if a field has been set.

### GetAnalysisChannels

`func (o *AudioAnalysisObjectTrack) GetAnalysisChannels() int32`

GetAnalysisChannels returns the AnalysisChannels field if non-nil, zero value otherwise.

### GetAnalysisChannelsOk

`func (o *AudioAnalysisObjectTrack) GetAnalysisChannelsOk() (*int32, bool)`

GetAnalysisChannelsOk returns a tuple with the AnalysisChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisChannels

`func (o *AudioAnalysisObjectTrack) SetAnalysisChannels(v int32)`

SetAnalysisChannels sets AnalysisChannels field to given value.

### HasAnalysisChannels

`func (o *AudioAnalysisObjectTrack) HasAnalysisChannels() bool`

HasAnalysisChannels returns a boolean if a field has been set.

### GetEndOfFadeIn

`func (o *AudioAnalysisObjectTrack) GetEndOfFadeIn() float32`

GetEndOfFadeIn returns the EndOfFadeIn field if non-nil, zero value otherwise.

### GetEndOfFadeInOk

`func (o *AudioAnalysisObjectTrack) GetEndOfFadeInOk() (*float32, bool)`

GetEndOfFadeInOk returns a tuple with the EndOfFadeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOfFadeIn

`func (o *AudioAnalysisObjectTrack) SetEndOfFadeIn(v float32)`

SetEndOfFadeIn sets EndOfFadeIn field to given value.

### HasEndOfFadeIn

`func (o *AudioAnalysisObjectTrack) HasEndOfFadeIn() bool`

HasEndOfFadeIn returns a boolean if a field has been set.

### GetStartOfFadeOut

`func (o *AudioAnalysisObjectTrack) GetStartOfFadeOut() float32`

GetStartOfFadeOut returns the StartOfFadeOut field if non-nil, zero value otherwise.

### GetStartOfFadeOutOk

`func (o *AudioAnalysisObjectTrack) GetStartOfFadeOutOk() (*float32, bool)`

GetStartOfFadeOutOk returns a tuple with the StartOfFadeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOfFadeOut

`func (o *AudioAnalysisObjectTrack) SetStartOfFadeOut(v float32)`

SetStartOfFadeOut sets StartOfFadeOut field to given value.

### HasStartOfFadeOut

`func (o *AudioAnalysisObjectTrack) HasStartOfFadeOut() bool`

HasStartOfFadeOut returns a boolean if a field has been set.

### GetLoudness

`func (o *AudioAnalysisObjectTrack) GetLoudness() float32`

GetLoudness returns the Loudness field if non-nil, zero value otherwise.

### GetLoudnessOk

`func (o *AudioAnalysisObjectTrack) GetLoudnessOk() (*float32, bool)`

GetLoudnessOk returns a tuple with the Loudness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudness

`func (o *AudioAnalysisObjectTrack) SetLoudness(v float32)`

SetLoudness sets Loudness field to given value.

### HasLoudness

`func (o *AudioAnalysisObjectTrack) HasLoudness() bool`

HasLoudness returns a boolean if a field has been set.

### GetTempo

`func (o *AudioAnalysisObjectTrack) GetTempo() float32`

GetTempo returns the Tempo field if non-nil, zero value otherwise.

### GetTempoOk

`func (o *AudioAnalysisObjectTrack) GetTempoOk() (*float32, bool)`

GetTempoOk returns a tuple with the Tempo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempo

`func (o *AudioAnalysisObjectTrack) SetTempo(v float32)`

SetTempo sets Tempo field to given value.

### HasTempo

`func (o *AudioAnalysisObjectTrack) HasTempo() bool`

HasTempo returns a boolean if a field has been set.

### GetTempoConfidence

`func (o *AudioAnalysisObjectTrack) GetTempoConfidence() float32`

GetTempoConfidence returns the TempoConfidence field if non-nil, zero value otherwise.

### GetTempoConfidenceOk

`func (o *AudioAnalysisObjectTrack) GetTempoConfidenceOk() (*float32, bool)`

GetTempoConfidenceOk returns a tuple with the TempoConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempoConfidence

`func (o *AudioAnalysisObjectTrack) SetTempoConfidence(v float32)`

SetTempoConfidence sets TempoConfidence field to given value.

### HasTempoConfidence

`func (o *AudioAnalysisObjectTrack) HasTempoConfidence() bool`

HasTempoConfidence returns a boolean if a field has been set.

### GetTimeSignature

`func (o *AudioAnalysisObjectTrack) GetTimeSignature() int32`

GetTimeSignature returns the TimeSignature field if non-nil, zero value otherwise.

### GetTimeSignatureOk

`func (o *AudioAnalysisObjectTrack) GetTimeSignatureOk() (*int32, bool)`

GetTimeSignatureOk returns a tuple with the TimeSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSignature

`func (o *AudioAnalysisObjectTrack) SetTimeSignature(v int32)`

SetTimeSignature sets TimeSignature field to given value.

### HasTimeSignature

`func (o *AudioAnalysisObjectTrack) HasTimeSignature() bool`

HasTimeSignature returns a boolean if a field has been set.

### GetTimeSignatureConfidence

`func (o *AudioAnalysisObjectTrack) GetTimeSignatureConfidence() float32`

GetTimeSignatureConfidence returns the TimeSignatureConfidence field if non-nil, zero value otherwise.

### GetTimeSignatureConfidenceOk

`func (o *AudioAnalysisObjectTrack) GetTimeSignatureConfidenceOk() (*float32, bool)`

GetTimeSignatureConfidenceOk returns a tuple with the TimeSignatureConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSignatureConfidence

`func (o *AudioAnalysisObjectTrack) SetTimeSignatureConfidence(v float32)`

SetTimeSignatureConfidence sets TimeSignatureConfidence field to given value.

### HasTimeSignatureConfidence

`func (o *AudioAnalysisObjectTrack) HasTimeSignatureConfidence() bool`

HasTimeSignatureConfidence returns a boolean if a field has been set.

### GetKey

`func (o *AudioAnalysisObjectTrack) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AudioAnalysisObjectTrack) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AudioAnalysisObjectTrack) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *AudioAnalysisObjectTrack) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKeyConfidence

`func (o *AudioAnalysisObjectTrack) GetKeyConfidence() float32`

GetKeyConfidence returns the KeyConfidence field if non-nil, zero value otherwise.

### GetKeyConfidenceOk

`func (o *AudioAnalysisObjectTrack) GetKeyConfidenceOk() (*float32, bool)`

GetKeyConfidenceOk returns a tuple with the KeyConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyConfidence

`func (o *AudioAnalysisObjectTrack) SetKeyConfidence(v float32)`

SetKeyConfidence sets KeyConfidence field to given value.

### HasKeyConfidence

`func (o *AudioAnalysisObjectTrack) HasKeyConfidence() bool`

HasKeyConfidence returns a boolean if a field has been set.

### GetMode

`func (o *AudioAnalysisObjectTrack) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AudioAnalysisObjectTrack) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AudioAnalysisObjectTrack) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AudioAnalysisObjectTrack) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModeConfidence

`func (o *AudioAnalysisObjectTrack) GetModeConfidence() float32`

GetModeConfidence returns the ModeConfidence field if non-nil, zero value otherwise.

### GetModeConfidenceOk

`func (o *AudioAnalysisObjectTrack) GetModeConfidenceOk() (*float32, bool)`

GetModeConfidenceOk returns a tuple with the ModeConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeConfidence

`func (o *AudioAnalysisObjectTrack) SetModeConfidence(v float32)`

SetModeConfidence sets ModeConfidence field to given value.

### HasModeConfidence

`func (o *AudioAnalysisObjectTrack) HasModeConfidence() bool`

HasModeConfidence returns a boolean if a field has been set.

### GetCodestring

`func (o *AudioAnalysisObjectTrack) GetCodestring() string`

GetCodestring returns the Codestring field if non-nil, zero value otherwise.

### GetCodestringOk

`func (o *AudioAnalysisObjectTrack) GetCodestringOk() (*string, bool)`

GetCodestringOk returns a tuple with the Codestring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodestring

`func (o *AudioAnalysisObjectTrack) SetCodestring(v string)`

SetCodestring sets Codestring field to given value.

### HasCodestring

`func (o *AudioAnalysisObjectTrack) HasCodestring() bool`

HasCodestring returns a boolean if a field has been set.

### GetCodeVersion

`func (o *AudioAnalysisObjectTrack) GetCodeVersion() float32`

GetCodeVersion returns the CodeVersion field if non-nil, zero value otherwise.

### GetCodeVersionOk

`func (o *AudioAnalysisObjectTrack) GetCodeVersionOk() (*float32, bool)`

GetCodeVersionOk returns a tuple with the CodeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVersion

`func (o *AudioAnalysisObjectTrack) SetCodeVersion(v float32)`

SetCodeVersion sets CodeVersion field to given value.

### HasCodeVersion

`func (o *AudioAnalysisObjectTrack) HasCodeVersion() bool`

HasCodeVersion returns a boolean if a field has been set.

### GetEchoprintstring

`func (o *AudioAnalysisObjectTrack) GetEchoprintstring() string`

GetEchoprintstring returns the Echoprintstring field if non-nil, zero value otherwise.

### GetEchoprintstringOk

`func (o *AudioAnalysisObjectTrack) GetEchoprintstringOk() (*string, bool)`

GetEchoprintstringOk returns a tuple with the Echoprintstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEchoprintstring

`func (o *AudioAnalysisObjectTrack) SetEchoprintstring(v string)`

SetEchoprintstring sets Echoprintstring field to given value.

### HasEchoprintstring

`func (o *AudioAnalysisObjectTrack) HasEchoprintstring() bool`

HasEchoprintstring returns a boolean if a field has been set.

### GetEchoprintVersion

`func (o *AudioAnalysisObjectTrack) GetEchoprintVersion() float32`

GetEchoprintVersion returns the EchoprintVersion field if non-nil, zero value otherwise.

### GetEchoprintVersionOk

`func (o *AudioAnalysisObjectTrack) GetEchoprintVersionOk() (*float32, bool)`

GetEchoprintVersionOk returns a tuple with the EchoprintVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEchoprintVersion

`func (o *AudioAnalysisObjectTrack) SetEchoprintVersion(v float32)`

SetEchoprintVersion sets EchoprintVersion field to given value.

### HasEchoprintVersion

`func (o *AudioAnalysisObjectTrack) HasEchoprintVersion() bool`

HasEchoprintVersion returns a boolean if a field has been set.

### GetSynchstring

`func (o *AudioAnalysisObjectTrack) GetSynchstring() string`

GetSynchstring returns the Synchstring field if non-nil, zero value otherwise.

### GetSynchstringOk

`func (o *AudioAnalysisObjectTrack) GetSynchstringOk() (*string, bool)`

GetSynchstringOk returns a tuple with the Synchstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchstring

`func (o *AudioAnalysisObjectTrack) SetSynchstring(v string)`

SetSynchstring sets Synchstring field to given value.

### HasSynchstring

`func (o *AudioAnalysisObjectTrack) HasSynchstring() bool`

HasSynchstring returns a boolean if a field has been set.

### GetSynchVersion

`func (o *AudioAnalysisObjectTrack) GetSynchVersion() float32`

GetSynchVersion returns the SynchVersion field if non-nil, zero value otherwise.

### GetSynchVersionOk

`func (o *AudioAnalysisObjectTrack) GetSynchVersionOk() (*float32, bool)`

GetSynchVersionOk returns a tuple with the SynchVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchVersion

`func (o *AudioAnalysisObjectTrack) SetSynchVersion(v float32)`

SetSynchVersion sets SynchVersion field to given value.

### HasSynchVersion

`func (o *AudioAnalysisObjectTrack) HasSynchVersion() bool`

HasSynchVersion returns a boolean if a field has been set.

### GetRhythmstring

`func (o *AudioAnalysisObjectTrack) GetRhythmstring() string`

GetRhythmstring returns the Rhythmstring field if non-nil, zero value otherwise.

### GetRhythmstringOk

`func (o *AudioAnalysisObjectTrack) GetRhythmstringOk() (*string, bool)`

GetRhythmstringOk returns a tuple with the Rhythmstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhythmstring

`func (o *AudioAnalysisObjectTrack) SetRhythmstring(v string)`

SetRhythmstring sets Rhythmstring field to given value.

### HasRhythmstring

`func (o *AudioAnalysisObjectTrack) HasRhythmstring() bool`

HasRhythmstring returns a boolean if a field has been set.

### GetRhythmVersion

`func (o *AudioAnalysisObjectTrack) GetRhythmVersion() float32`

GetRhythmVersion returns the RhythmVersion field if non-nil, zero value otherwise.

### GetRhythmVersionOk

`func (o *AudioAnalysisObjectTrack) GetRhythmVersionOk() (*float32, bool)`

GetRhythmVersionOk returns a tuple with the RhythmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhythmVersion

`func (o *AudioAnalysisObjectTrack) SetRhythmVersion(v float32)`

SetRhythmVersion sets RhythmVersion field to given value.

### HasRhythmVersion

`func (o *AudioAnalysisObjectTrack) HasRhythmVersion() bool`

HasRhythmVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


