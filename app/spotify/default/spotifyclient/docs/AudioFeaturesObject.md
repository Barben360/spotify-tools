# AudioFeaturesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acousticness** | Pointer to **float32** | A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0 represents high confidence the track is acoustic.  | [optional] 
**AnalysisUrl** | Pointer to **string** | A URL to access the full audio analysis of this track. An access token is required to access this data.  | [optional] 
**Danceability** | Pointer to **float32** | Danceability describes how suitable a track is for dancing based on a combination of musical elements including tempo, rhythm stability, beat strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is most danceable.  | [optional] 
**DurationMs** | Pointer to **int32** | The duration of the track in milliseconds.  | [optional] 
**Energy** | Pointer to **float32** | Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of intensity and activity. Typically, energetic tracks feel fast, loud, and noisy. For example, death metal has high energy, while a Bach prelude scores low on the scale. Perceptual features contributing to this attribute include dynamic range, perceived loudness, timbre, onset rate, and general entropy.  | [optional] 
**Id** | Pointer to **string** | The Spotify ID for the track.  | [optional] 
**Instrumentalness** | Pointer to **float32** | Predicts whether a track contains no vocals. \&quot;Ooh\&quot; and \&quot;aah\&quot; sounds are treated as instrumental in this context. Rap or spoken word tracks are clearly \&quot;vocal\&quot;. The closer the instrumentalness value is to 1.0, the greater likelihood the track contains no vocal content. Values above 0.5 are intended to represent instrumental tracks, but confidence is higher as the value approaches 1.0.  | [optional] 
**Key** | Pointer to **int32** | The key the track is in. Integers map to pitches using standard [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 &#x3D; C, 1 &#x3D; C♯/D♭, 2 &#x3D; D, and so on. If no key was detected, the value is -1.  | [optional] 
**Liveness** | Pointer to **float32** | Detects the presence of an audience in the recording. Higher liveness values represent an increased probability that the track was performed live. A value above 0.8 provides strong likelihood that the track is live.  | [optional] 
**Loudness** | Pointer to **float32** | The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db.  | [optional] 
**Mode** | Pointer to **int32** | Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0.  | [optional] 
**Speechiness** | Pointer to **float32** | Speechiness detects the presence of spoken words in a track. The more exclusively speech-like the recording (e.g. talk show, audio book, poetry), the closer to 1.0 the attribute value. Values above 0.66 describe tracks that are probably made entirely of spoken words. Values between 0.33 and 0.66 describe tracks that may contain both music and speech, either in sections or layered, including such cases as rap music. Values below 0.33 most likely represent music and other non-speech-like tracks.  | [optional] 
**Tempo** | Pointer to **float32** | The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration.  | [optional] 
**TimeSignature** | Pointer to **int32** | An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of \&quot;3/4\&quot;, to \&quot;7/4\&quot;. | [optional] 
**TrackHref** | Pointer to **string** | A link to the Web API endpoint providing full details of the track.  | [optional] 
**Type** | Pointer to **string** | The object type.  | [optional] 
**Uri** | Pointer to **string** | The Spotify URI for the track.  | [optional] 
**Valence** | Pointer to **float32** | A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a track. Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric), while tracks with low valence sound more negative (e.g. sad, depressed, angry).  | [optional] 

## Methods

### NewAudioFeaturesObject

`func NewAudioFeaturesObject() *AudioFeaturesObject`

NewAudioFeaturesObject instantiates a new AudioFeaturesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioFeaturesObjectWithDefaults

`func NewAudioFeaturesObjectWithDefaults() *AudioFeaturesObject`

NewAudioFeaturesObjectWithDefaults instantiates a new AudioFeaturesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcousticness

`func (o *AudioFeaturesObject) GetAcousticness() float32`

GetAcousticness returns the Acousticness field if non-nil, zero value otherwise.

### GetAcousticnessOk

`func (o *AudioFeaturesObject) GetAcousticnessOk() (*float32, bool)`

GetAcousticnessOk returns a tuple with the Acousticness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcousticness

`func (o *AudioFeaturesObject) SetAcousticness(v float32)`

SetAcousticness sets Acousticness field to given value.

### HasAcousticness

`func (o *AudioFeaturesObject) HasAcousticness() bool`

HasAcousticness returns a boolean if a field has been set.

### GetAnalysisUrl

`func (o *AudioFeaturesObject) GetAnalysisUrl() string`

GetAnalysisUrl returns the AnalysisUrl field if non-nil, zero value otherwise.

### GetAnalysisUrlOk

`func (o *AudioFeaturesObject) GetAnalysisUrlOk() (*string, bool)`

GetAnalysisUrlOk returns a tuple with the AnalysisUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisUrl

`func (o *AudioFeaturesObject) SetAnalysisUrl(v string)`

SetAnalysisUrl sets AnalysisUrl field to given value.

### HasAnalysisUrl

`func (o *AudioFeaturesObject) HasAnalysisUrl() bool`

HasAnalysisUrl returns a boolean if a field has been set.

### GetDanceability

`func (o *AudioFeaturesObject) GetDanceability() float32`

GetDanceability returns the Danceability field if non-nil, zero value otherwise.

### GetDanceabilityOk

`func (o *AudioFeaturesObject) GetDanceabilityOk() (*float32, bool)`

GetDanceabilityOk returns a tuple with the Danceability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDanceability

`func (o *AudioFeaturesObject) SetDanceability(v float32)`

SetDanceability sets Danceability field to given value.

### HasDanceability

`func (o *AudioFeaturesObject) HasDanceability() bool`

HasDanceability returns a boolean if a field has been set.

### GetDurationMs

`func (o *AudioFeaturesObject) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *AudioFeaturesObject) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *AudioFeaturesObject) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *AudioFeaturesObject) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEnergy

`func (o *AudioFeaturesObject) GetEnergy() float32`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *AudioFeaturesObject) GetEnergyOk() (*float32, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *AudioFeaturesObject) SetEnergy(v float32)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *AudioFeaturesObject) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.

### GetId

`func (o *AudioFeaturesObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudioFeaturesObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudioFeaturesObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AudioFeaturesObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstrumentalness

`func (o *AudioFeaturesObject) GetInstrumentalness() float32`

GetInstrumentalness returns the Instrumentalness field if non-nil, zero value otherwise.

### GetInstrumentalnessOk

`func (o *AudioFeaturesObject) GetInstrumentalnessOk() (*float32, bool)`

GetInstrumentalnessOk returns a tuple with the Instrumentalness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentalness

`func (o *AudioFeaturesObject) SetInstrumentalness(v float32)`

SetInstrumentalness sets Instrumentalness field to given value.

### HasInstrumentalness

`func (o *AudioFeaturesObject) HasInstrumentalness() bool`

HasInstrumentalness returns a boolean if a field has been set.

### GetKey

`func (o *AudioFeaturesObject) GetKey() int32`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AudioFeaturesObject) GetKeyOk() (*int32, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AudioFeaturesObject) SetKey(v int32)`

SetKey sets Key field to given value.

### HasKey

`func (o *AudioFeaturesObject) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLiveness

`func (o *AudioFeaturesObject) GetLiveness() float32`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *AudioFeaturesObject) GetLivenessOk() (*float32, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *AudioFeaturesObject) SetLiveness(v float32)`

SetLiveness sets Liveness field to given value.

### HasLiveness

`func (o *AudioFeaturesObject) HasLiveness() bool`

HasLiveness returns a boolean if a field has been set.

### GetLoudness

`func (o *AudioFeaturesObject) GetLoudness() float32`

GetLoudness returns the Loudness field if non-nil, zero value otherwise.

### GetLoudnessOk

`func (o *AudioFeaturesObject) GetLoudnessOk() (*float32, bool)`

GetLoudnessOk returns a tuple with the Loudness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoudness

`func (o *AudioFeaturesObject) SetLoudness(v float32)`

SetLoudness sets Loudness field to given value.

### HasLoudness

`func (o *AudioFeaturesObject) HasLoudness() bool`

HasLoudness returns a boolean if a field has been set.

### GetMode

`func (o *AudioFeaturesObject) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AudioFeaturesObject) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AudioFeaturesObject) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AudioFeaturesObject) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSpeechiness

`func (o *AudioFeaturesObject) GetSpeechiness() float32`

GetSpeechiness returns the Speechiness field if non-nil, zero value otherwise.

### GetSpeechinessOk

`func (o *AudioFeaturesObject) GetSpeechinessOk() (*float32, bool)`

GetSpeechinessOk returns a tuple with the Speechiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechiness

`func (o *AudioFeaturesObject) SetSpeechiness(v float32)`

SetSpeechiness sets Speechiness field to given value.

### HasSpeechiness

`func (o *AudioFeaturesObject) HasSpeechiness() bool`

HasSpeechiness returns a boolean if a field has been set.

### GetTempo

`func (o *AudioFeaturesObject) GetTempo() float32`

GetTempo returns the Tempo field if non-nil, zero value otherwise.

### GetTempoOk

`func (o *AudioFeaturesObject) GetTempoOk() (*float32, bool)`

GetTempoOk returns a tuple with the Tempo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempo

`func (o *AudioFeaturesObject) SetTempo(v float32)`

SetTempo sets Tempo field to given value.

### HasTempo

`func (o *AudioFeaturesObject) HasTempo() bool`

HasTempo returns a boolean if a field has been set.

### GetTimeSignature

`func (o *AudioFeaturesObject) GetTimeSignature() int32`

GetTimeSignature returns the TimeSignature field if non-nil, zero value otherwise.

### GetTimeSignatureOk

`func (o *AudioFeaturesObject) GetTimeSignatureOk() (*int32, bool)`

GetTimeSignatureOk returns a tuple with the TimeSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSignature

`func (o *AudioFeaturesObject) SetTimeSignature(v int32)`

SetTimeSignature sets TimeSignature field to given value.

### HasTimeSignature

`func (o *AudioFeaturesObject) HasTimeSignature() bool`

HasTimeSignature returns a boolean if a field has been set.

### GetTrackHref

`func (o *AudioFeaturesObject) GetTrackHref() string`

GetTrackHref returns the TrackHref field if non-nil, zero value otherwise.

### GetTrackHrefOk

`func (o *AudioFeaturesObject) GetTrackHrefOk() (*string, bool)`

GetTrackHrefOk returns a tuple with the TrackHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackHref

`func (o *AudioFeaturesObject) SetTrackHref(v string)`

SetTrackHref sets TrackHref field to given value.

### HasTrackHref

`func (o *AudioFeaturesObject) HasTrackHref() bool`

HasTrackHref returns a boolean if a field has been set.

### GetType

`func (o *AudioFeaturesObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudioFeaturesObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudioFeaturesObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AudioFeaturesObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *AudioFeaturesObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AudioFeaturesObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AudioFeaturesObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *AudioFeaturesObject) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetValence

`func (o *AudioFeaturesObject) GetValence() float32`

GetValence returns the Valence field if non-nil, zero value otherwise.

### GetValenceOk

`func (o *AudioFeaturesObject) GetValenceOk() (*float32, bool)`

GetValenceOk returns a tuple with the Valence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValence

`func (o *AudioFeaturesObject) SetValence(v float32)`

SetValence sets Valence field to given value.

### HasValence

`func (o *AudioFeaturesObject) HasValence() bool`

HasValence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


