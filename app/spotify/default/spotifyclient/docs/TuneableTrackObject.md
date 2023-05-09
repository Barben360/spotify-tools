# TuneableTrackObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acousticness** | Pointer to **float32** | A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0 represents high confidence the track is acoustic.  | [optional] 
**Danceability** | Pointer to **float32** | Danceability describes how suitable a track is for dancing based on a combination of musical elements including tempo, rhythm stability, beat strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is most danceable.  | [optional] 
**DurationMs** | Pointer to **int32** | The duration of the track in milliseconds.  | [optional] 
**Energy** | Pointer to **float32** | Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of intensity and activity. Typically, energetic tracks feel fast, loud, and noisy. For example, death metal has high energy, while a Bach prelude scores low on the scale. Perceptual features contributing to this attribute include dynamic range, perceived loudness, timbre, onset rate, and general entropy.  | [optional] 
**Instrumentalness** | Pointer to **float32** | Predicts whether a track contains no vocals. \&quot;Ooh\&quot; and \&quot;aah\&quot; sounds are treated as instrumental in this context. Rap or spoken word tracks are clearly \&quot;vocal\&quot;. The closer the instrumentalness value is to 1.0, the greater likelihood the track contains no vocal content. Values above 0.5 are intended to represent instrumental tracks, but confidence is higher as the value approaches 1.0.  | [optional] 
**Key** | Pointer to **int32** | The key the track is in. Integers map to pitches using standard [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 &#x3D; C, 1 &#x3D; C♯/D♭, 2 &#x3D; D, and so on. If no key was detected, the value is -1.  | [optional] 
**Liveness** | Pointer to **float32** | Detects the presence of an audience in the recording. Higher liveness values represent an increased probability that the track was performed live. A value above 0.8 provides strong likelihood that the track is live.  | [optional] 
**Loudness** | Pointer to **float32** | The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db.  | [optional] 
**Mode** | Pointer to **int32** | Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0.  | [optional] 
**Popularity** | Pointer to **float32** | The popularity of the track. The value will be between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are. _**Note**: When applying track relinking via the &#x60;market&#x60; parameter, it is expected to find relinked tracks with popularities that do not match &#x60;min_*&#x60;, &#x60;max_*&#x60;and &#x60;target_*&#x60; popularities. These relinked tracks are accurate replacements for unplayable tracks with the expected popularity scores. Original, non-relinked tracks are available via the &#x60;linked_from&#x60; attribute of the [relinked track response](/documentation/web-api/concepts/track-relinking)._  | [optional] 
**Speechiness** | Pointer to **float32** | Speechiness detects the presence of spoken words in a track. The more exclusively speech-like the recording (e.g. talk show, audio book, poetry), the closer to 1.0 the attribute value. Values above 0.66 describe tracks that are probably made entirely of spoken words. Values between 0.33 and 0.66 describe tracks that may contain both music and speech, either in sections or layered, including such cases as rap music. Values below 0.33 most likely represent music and other non-speech-like tracks.  | [optional] 
**Tempo** | Pointer to **float32** | The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.  | [optional] 
**TimeSignature** | Pointer to **int32** | An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of \&quot;3/4\&quot;, to \&quot;7/4\&quot;. | [optional] 
**Valence** | Pointer to **float32** | A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a track. Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric), while tracks with low valence sound more negative (e.g. sad, depressed, angry).  | [optional] 

## Methods

### NewTuneableTrackObject

`func NewTuneableTrackObject() *TuneableTrackObject`

NewTuneableTrackObject instantiates a new TuneableTrackObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTuneableTrackObjectWithDefaults

`func NewTuneableTrackObjectWithDefaults() *TuneableTrackObject`

NewTuneableTrackObjectWithDefaults instantiates a new TuneableTrackObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcousticness

`func (o *TuneableTrackObject) GetAcousticness() float32`

GetAcousticness returns the Acousticness field if non-nil, zero value otherwise.

### GetAcousticnessOk

`func (o *TuneableTrackObject) GetAcousticnessOk() (*float32, bool)`

GetAcousticnessOk returns a tuple with the Acousticness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcousticness

`func (o *TuneableTrackObject) SetAcousticness(v float32)`

SetAcousticness sets Acousticness field to given value.

### HasAcousticness

`func (o *TuneableTrackObject) HasAcousticness() bool`

HasAcousticness returns a boolean if a field has been set.

### GetDanceability

`func (o *TuneableTrackObject) GetDanceability() float32`

GetDanceability returns the Danceability field if non-nil, zero value otherwise.

### GetDanceabilityOk

`func (o *TuneableTrackObject) GetDanceabilityOk() (*float32, bool)`

GetDanceabilityOk returns a tuple with the Danceability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDanceability

`func (o *TuneableTrackObject) SetDanceability(v float32)`

SetDanceability sets Danceability field to given value.

### HasDanceability

`func (o *TuneableTrackObject) HasDanceability() bool`

HasDanceability returns a boolean if a field has been set.

### GetDurationMs

`func (o *TuneableTrackObject) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *TuneableTrackObject) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *TuneableTrackObject) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *TuneableTrackObject) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEnergy

`func (o *TuneableTrackObject) GetEnergy() float32`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *TuneableTrackObject) GetEnergyOk() (*float32, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *TuneableTrackObject) SetEnergy(v float32)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *TuneableTrackObject) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.

### GetInstrumentalness

`func (o *TuneableTrackObject) GetInstrumentalness() float32`

GetInstrumentalness returns the Instrumentalness field if non-nil, zero value otherwise.

### GetInstrumentalnessOk

`func (o *TuneableTrackObject) GetInstrumentalnessOk() (*float32, bool)`

GetInstrumentalnessOk returns a tuple with the Instrumentalness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentalness

`func (o *TuneableTrackObject) SetInstrumentalness(v float32)`

SetInstrumentalness sets Instrumentalness field to given value.

### HasInstrumentalness

`func (o *TuneableTrackObject) HasInstrumentalness() bool`

HasInstrumentalness returns a boolean if a field has been set.

### GetKey

`func (o *TuneableTrackObject) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TuneableTrackObject) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TuneableTrackObject) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *TuneableTrackObject) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLiveness

`func (o *TuneableTrackObject) GetLiveness() float32`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *TuneableTrackObject) GetLivenessOk() (*float32, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *TuneableTrackObject) SetLiveness(v float32)`

SetLiveness sets Liveness field to given value.

### HasLiveness

`func (o *TuneableTrackObject) HasLiveness() bool`

HasLiveness returns a boolean if a field has been set.

### GetLoudness

`func (o *TuneableTrackObject) GetLoudness() float32`

GetLoudness returns the Loudness field if non-nil, zero value otherwise.

### GetLoudnessOk

`func (o *TuneableTrackObject) GetLoudnessOk() (*float32, bool)`

GetLoudnessOk returns a tuple with the Loudness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudness

`func (o *TuneableTrackObject) SetLoudness(v float32)`

SetLoudness sets Loudness field to given value.

### HasLoudness

`func (o *TuneableTrackObject) HasLoudness() bool`

HasLoudness returns a boolean if a field has been set.

### GetMode

`func (o *TuneableTrackObject) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TuneableTrackObject) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TuneableTrackObject) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TuneableTrackObject) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPopularity

`func (o *TuneableTrackObject) GetPopularity() float32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *TuneableTrackObject) GetPopularityOk() (*float32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *TuneableTrackObject) SetPopularity(v float32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *TuneableTrackObject) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetSpeechiness

`func (o *TuneableTrackObject) GetSpeechiness() float32`

GetSpeechiness returns the Speechiness field if non-nil, zero value otherwise.

### GetSpeechinessOk

`func (o *TuneableTrackObject) GetSpeechinessOk() (*float32, bool)`

GetSpeechinessOk returns a tuple with the Speechiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechiness

`func (o *TuneableTrackObject) SetSpeechiness(v float32)`

SetSpeechiness sets Speechiness field to given value.

### HasSpeechiness

`func (o *TuneableTrackObject) HasSpeechiness() bool`

HasSpeechiness returns a boolean if a field has been set.

### GetTempo

`func (o *TuneableTrackObject) GetTempo() float32`

GetTempo returns the Tempo field if non-nil, zero value otherwise.

### GetTempoOk

`func (o *TuneableTrackObject) GetTempoOk() (*float32, bool)`

GetTempoOk returns a tuple with the Tempo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempo

`func (o *TuneableTrackObject) SetTempo(v float32)`

SetTempo sets Tempo field to given value.

### HasTempo

`func (o *TuneableTrackObject) HasTempo() bool`

HasTempo returns a boolean if a field has been set.

### GetTimeSignature

`func (o *TuneableTrackObject) GetTimeSignature() int32`

GetTimeSignature returns the TimeSignature field if non-nil, zero value otherwise.

### GetTimeSignatureOk

`func (o *TuneableTrackObject) GetTimeSignatureOk() (*int32, bool)`

GetTimeSignatureOk returns a tuple with the TimeSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSignature

`func (o *TuneableTrackObject) SetTimeSignature(v int32)`

SetTimeSignature sets TimeSignature field to given value.

### HasTimeSignature

`func (o *TuneableTrackObject) HasTimeSignature() bool`

HasTimeSignature returns a boolean if a field has been set.

### GetValence

`func (o *TuneableTrackObject) GetValence() float32`

GetValence returns the Valence field if non-nil, zero value otherwise.

### GetValenceOk

`func (o *TuneableTrackObject) GetValenceOk() (*float32, bool)`

GetValenceOk returns a tuple with the Valence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValence

`func (o *TuneableTrackObject) SetValence(v float32)`

SetValence sets Valence field to given value.

### HasValence

`func (o *TuneableTrackObject) HasValence() bool`

HasValence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


