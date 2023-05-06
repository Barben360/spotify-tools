/*
Spotify Web API with fixes and improvements from sonallux

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 2023.2.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spotifyclient

import (
	"encoding/json"
)

// checks if the TuneableTrackObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TuneableTrackObject{}

// TuneableTrackObject struct for TuneableTrackObject
type TuneableTrackObject struct {
	// A confidence measure from 0.0 to 1.0 of whether the track is acoustic. 1.0 represents high confidence the track is acoustic. 
	Acousticness *float32 `json:"acousticness,omitempty"`
	// Danceability describes how suitable a track is for dancing based on a combination of musical elements including tempo, rhythm stability, beat strength, and overall regularity. A value of 0.0 is least danceable and 1.0 is most danceable. 
	Danceability *float32 `json:"danceability,omitempty"`
	// The duration of the track in milliseconds. 
	DurationMs *int32 `json:"duration_ms,omitempty"`
	// Energy is a measure from 0.0 to 1.0 and represents a perceptual measure of intensity and activity. Typically, energetic tracks feel fast, loud, and noisy. For example, death metal has high energy, while a Bach prelude scores low on the scale. Perceptual features contributing to this attribute include dynamic range, perceived loudness, timbre, onset rate, and general entropy. 
	Energy *float32 `json:"energy,omitempty"`
	// Predicts whether a track contains no vocals. \"Ooh\" and \"aah\" sounds are treated as instrumental in this context. Rap or spoken word tracks are clearly \"vocal\". The closer the instrumentalness value is to 1.0, the greater likelihood the track contains no vocal content. Values above 0.5 are intended to represent instrumental tracks, but confidence is higher as the value approaches 1.0. 
	Instrumentalness *float32 `json:"instrumentalness,omitempty"`
	// The key the track is in. Integers map to pitches using standard [Pitch Class notation](https://en.wikipedia.org/wiki/Pitch_class). E.g. 0 = C, 1 = C♯/D♭, 2 = D, and so on. If no key was detected, the value is -1. 
	Key *int32 `json:"key,omitempty"`
	// Detects the presence of an audience in the recording. Higher liveness values represent an increased probability that the track was performed live. A value above 0.8 provides strong likelihood that the track is live. 
	Liveness *float32 `json:"liveness,omitempty"`
	// The overall loudness of a track in decibels (dB). Loudness values are averaged across the entire track and are useful for comparing relative loudness of tracks. Loudness is the quality of a sound that is the primary psychological correlate of physical strength (amplitude). Values typically range between -60 and 0 db. 
	Loudness *float32 `json:"loudness,omitempty"`
	// Mode indicates the modality (major or minor) of a track, the type of scale from which its melodic content is derived. Major is represented by 1 and minor is 0. 
	Mode *int32 `json:"mode,omitempty"`
	// The popularity of the track. The value will be between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are. _**Note**: When applying track relinking via the `market` parameter, it is expected to find relinked tracks with popularities that do not match `min_*`, `max_*`and `target_*` popularities. These relinked tracks are accurate replacements for unplayable tracks with the expected popularity scores. Original, non-relinked tracks are available via the `linked_from` attribute of the [relinked track response](/documentation/general/guides/track-relinking-guide)._ 
	Popularity *float32 `json:"popularity,omitempty"`
	// Speechiness detects the presence of spoken words in a track. The more exclusively speech-like the recording (e.g. talk show, audio book, poetry), the closer to 1.0 the attribute value. Values above 0.66 describe tracks that are probably made entirely of spoken words. Values between 0.33 and 0.66 describe tracks that may contain both music and speech, either in sections or layered, including such cases as rap music. Values below 0.33 most likely represent music and other non-speech-like tracks. 
	Speechiness *float32 `json:"speechiness,omitempty"`
	// The overall estimated tempo of a track in beats per minute (BPM). In musical terminology, tempo is the speed or pace of a given piece and derives directly from the average beat duration. 
	Tempo *float32 `json:"tempo,omitempty"`
	// An estimated time signature. The time signature (meter) is a notational convention to specify how many beats are in each bar (or measure). The time signature ranges from 3 to 7 indicating time signatures of \"3/4\", to \"7/4\".
	TimeSignature *int32 `json:"time_signature,omitempty"`
	// A measure from 0.0 to 1.0 describing the musical positiveness conveyed by a track. Tracks with high valence sound more positive (e.g. happy, cheerful, euphoric), while tracks with low valence sound more negative (e.g. sad, depressed, angry). 
	Valence *float32 `json:"valence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TuneableTrackObject TuneableTrackObject

// NewTuneableTrackObject instantiates a new TuneableTrackObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTuneableTrackObject() *TuneableTrackObject {
	this := TuneableTrackObject{}
	return &this
}

// NewTuneableTrackObjectWithDefaults instantiates a new TuneableTrackObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTuneableTrackObjectWithDefaults() *TuneableTrackObject {
	this := TuneableTrackObject{}
	return &this
}

// GetAcousticness returns the Acousticness field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetAcousticness() float32 {
	if o == nil || IsNil(o.Acousticness) {
		var ret float32
		return ret
	}
	return *o.Acousticness
}

// GetAcousticnessOk returns a tuple with the Acousticness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetAcousticnessOk() (*float32, bool) {
	if o == nil || IsNil(o.Acousticness) {
		return nil, false
	}
	return o.Acousticness, true
}

// HasAcousticness returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasAcousticness() bool {
	if o != nil && !IsNil(o.Acousticness) {
		return true
	}

	return false
}

// SetAcousticness gets a reference to the given float32 and assigns it to the Acousticness field.
func (o *TuneableTrackObject) SetAcousticness(v float32) {
	o.Acousticness = &v
}

// GetDanceability returns the Danceability field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetDanceability() float32 {
	if o == nil || IsNil(o.Danceability) {
		var ret float32
		return ret
	}
	return *o.Danceability
}

// GetDanceabilityOk returns a tuple with the Danceability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetDanceabilityOk() (*float32, bool) {
	if o == nil || IsNil(o.Danceability) {
		return nil, false
	}
	return o.Danceability, true
}

// HasDanceability returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasDanceability() bool {
	if o != nil && !IsNil(o.Danceability) {
		return true
	}

	return false
}

// SetDanceability gets a reference to the given float32 and assigns it to the Danceability field.
func (o *TuneableTrackObject) SetDanceability(v float32) {
	o.Danceability = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetDurationMs() int32 {
	if o == nil || IsNil(o.DurationMs) {
		var ret int32
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetDurationMsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationMs) {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasDurationMs() bool {
	if o != nil && !IsNil(o.DurationMs) {
		return true
	}

	return false
}

// SetDurationMs gets a reference to the given int32 and assigns it to the DurationMs field.
func (o *TuneableTrackObject) SetDurationMs(v int32) {
	o.DurationMs = &v
}

// GetEnergy returns the Energy field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetEnergy() float32 {
	if o == nil || IsNil(o.Energy) {
		var ret float32
		return ret
	}
	return *o.Energy
}

// GetEnergyOk returns a tuple with the Energy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetEnergyOk() (*float32, bool) {
	if o == nil || IsNil(o.Energy) {
		return nil, false
	}
	return o.Energy, true
}

// HasEnergy returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasEnergy() bool {
	if o != nil && !IsNil(o.Energy) {
		return true
	}

	return false
}

// SetEnergy gets a reference to the given float32 and assigns it to the Energy field.
func (o *TuneableTrackObject) SetEnergy(v float32) {
	o.Energy = &v
}

// GetInstrumentalness returns the Instrumentalness field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetInstrumentalness() float32 {
	if o == nil || IsNil(o.Instrumentalness) {
		var ret float32
		return ret
	}
	return *o.Instrumentalness
}

// GetInstrumentalnessOk returns a tuple with the Instrumentalness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetInstrumentalnessOk() (*float32, bool) {
	if o == nil || IsNil(o.Instrumentalness) {
		return nil, false
	}
	return o.Instrumentalness, true
}

// HasInstrumentalness returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasInstrumentalness() bool {
	if o != nil && !IsNil(o.Instrumentalness) {
		return true
	}

	return false
}

// SetInstrumentalness gets a reference to the given float32 and assigns it to the Instrumentalness field.
func (o *TuneableTrackObject) SetInstrumentalness(v float32) {
	o.Instrumentalness = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetKey() int32 {
	if o == nil || IsNil(o.Key) {
		var ret int32
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetKeyOk() (*int32, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given int32 and assigns it to the Key field.
func (o *TuneableTrackObject) SetKey(v int32) {
	o.Key = &v
}

// GetLiveness returns the Liveness field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetLiveness() float32 {
	if o == nil || IsNil(o.Liveness) {
		var ret float32
		return ret
	}
	return *o.Liveness
}

// GetLivenessOk returns a tuple with the Liveness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetLivenessOk() (*float32, bool) {
	if o == nil || IsNil(o.Liveness) {
		return nil, false
	}
	return o.Liveness, true
}

// HasLiveness returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasLiveness() bool {
	if o != nil && !IsNil(o.Liveness) {
		return true
	}

	return false
}

// SetLiveness gets a reference to the given float32 and assigns it to the Liveness field.
func (o *TuneableTrackObject) SetLiveness(v float32) {
	o.Liveness = &v
}

// GetLoudness returns the Loudness field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetLoudness() float32 {
	if o == nil || IsNil(o.Loudness) {
		var ret float32
		return ret
	}
	return *o.Loudness
}

// GetLoudnessOk returns a tuple with the Loudness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetLoudnessOk() (*float32, bool) {
	if o == nil || IsNil(o.Loudness) {
		return nil, false
	}
	return o.Loudness, true
}

// HasLoudness returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasLoudness() bool {
	if o != nil && !IsNil(o.Loudness) {
		return true
	}

	return false
}

// SetLoudness gets a reference to the given float32 and assigns it to the Loudness field.
func (o *TuneableTrackObject) SetLoudness(v float32) {
	o.Loudness = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetMode() int32 {
	if o == nil || IsNil(o.Mode) {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetModeOk() (*int32, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *TuneableTrackObject) SetMode(v int32) {
	o.Mode = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetPopularity() float32 {
	if o == nil || IsNil(o.Popularity) {
		var ret float32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetPopularityOk() (*float32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given float32 and assigns it to the Popularity field.
func (o *TuneableTrackObject) SetPopularity(v float32) {
	o.Popularity = &v
}

// GetSpeechiness returns the Speechiness field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetSpeechiness() float32 {
	if o == nil || IsNil(o.Speechiness) {
		var ret float32
		return ret
	}
	return *o.Speechiness
}

// GetSpeechinessOk returns a tuple with the Speechiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetSpeechinessOk() (*float32, bool) {
	if o == nil || IsNil(o.Speechiness) {
		return nil, false
	}
	return o.Speechiness, true
}

// HasSpeechiness returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasSpeechiness() bool {
	if o != nil && !IsNil(o.Speechiness) {
		return true
	}

	return false
}

// SetSpeechiness gets a reference to the given float32 and assigns it to the Speechiness field.
func (o *TuneableTrackObject) SetSpeechiness(v float32) {
	o.Speechiness = &v
}

// GetTempo returns the Tempo field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetTempo() float32 {
	if o == nil || IsNil(o.Tempo) {
		var ret float32
		return ret
	}
	return *o.Tempo
}

// GetTempoOk returns a tuple with the Tempo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetTempoOk() (*float32, bool) {
	if o == nil || IsNil(o.Tempo) {
		return nil, false
	}
	return o.Tempo, true
}

// HasTempo returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasTempo() bool {
	if o != nil && !IsNil(o.Tempo) {
		return true
	}

	return false
}

// SetTempo gets a reference to the given float32 and assigns it to the Tempo field.
func (o *TuneableTrackObject) SetTempo(v float32) {
	o.Tempo = &v
}

// GetTimeSignature returns the TimeSignature field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetTimeSignature() int32 {
	if o == nil || IsNil(o.TimeSignature) {
		var ret int32
		return ret
	}
	return *o.TimeSignature
}

// GetTimeSignatureOk returns a tuple with the TimeSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetTimeSignatureOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeSignature) {
		return nil, false
	}
	return o.TimeSignature, true
}

// HasTimeSignature returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasTimeSignature() bool {
	if o != nil && !IsNil(o.TimeSignature) {
		return true
	}

	return false
}

// SetTimeSignature gets a reference to the given int32 and assigns it to the TimeSignature field.
func (o *TuneableTrackObject) SetTimeSignature(v int32) {
	o.TimeSignature = &v
}

// GetValence returns the Valence field value if set, zero value otherwise.
func (o *TuneableTrackObject) GetValence() float32 {
	if o == nil || IsNil(o.Valence) {
		var ret float32
		return ret
	}
	return *o.Valence
}

// GetValenceOk returns a tuple with the Valence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TuneableTrackObject) GetValenceOk() (*float32, bool) {
	if o == nil || IsNil(o.Valence) {
		return nil, false
	}
	return o.Valence, true
}

// HasValence returns a boolean if a field has been set.
func (o *TuneableTrackObject) HasValence() bool {
	if o != nil && !IsNil(o.Valence) {
		return true
	}

	return false
}

// SetValence gets a reference to the given float32 and assigns it to the Valence field.
func (o *TuneableTrackObject) SetValence(v float32) {
	o.Valence = &v
}

func (o TuneableTrackObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TuneableTrackObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acousticness) {
		toSerialize["acousticness"] = o.Acousticness
	}
	if !IsNil(o.Danceability) {
		toSerialize["danceability"] = o.Danceability
	}
	if !IsNil(o.DurationMs) {
		toSerialize["duration_ms"] = o.DurationMs
	}
	if !IsNil(o.Energy) {
		toSerialize["energy"] = o.Energy
	}
	if !IsNil(o.Instrumentalness) {
		toSerialize["instrumentalness"] = o.Instrumentalness
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Liveness) {
		toSerialize["liveness"] = o.Liveness
	}
	if !IsNil(o.Loudness) {
		toSerialize["loudness"] = o.Loudness
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !IsNil(o.Speechiness) {
		toSerialize["speechiness"] = o.Speechiness
	}
	if !IsNil(o.Tempo) {
		toSerialize["tempo"] = o.Tempo
	}
	if !IsNil(o.TimeSignature) {
		toSerialize["time_signature"] = o.TimeSignature
	}
	if !IsNil(o.Valence) {
		toSerialize["valence"] = o.Valence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TuneableTrackObject) UnmarshalJSON(bytes []byte) (err error) {
	varTuneableTrackObject := _TuneableTrackObject{}

	if err = json.Unmarshal(bytes, &varTuneableTrackObject); err == nil {
		*o = TuneableTrackObject(varTuneableTrackObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "acousticness")
		delete(additionalProperties, "danceability")
		delete(additionalProperties, "duration_ms")
		delete(additionalProperties, "energy")
		delete(additionalProperties, "instrumentalness")
		delete(additionalProperties, "key")
		delete(additionalProperties, "liveness")
		delete(additionalProperties, "loudness")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "popularity")
		delete(additionalProperties, "speechiness")
		delete(additionalProperties, "tempo")
		delete(additionalProperties, "time_signature")
		delete(additionalProperties, "valence")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTuneableTrackObject struct {
	value *TuneableTrackObject
	isSet bool
}

func (v NullableTuneableTrackObject) Get() *TuneableTrackObject {
	return v.value
}

func (v *NullableTuneableTrackObject) Set(val *TuneableTrackObject) {
	v.value = val
	v.isSet = true
}

func (v NullableTuneableTrackObject) IsSet() bool {
	return v.isSet
}

func (v *NullableTuneableTrackObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTuneableTrackObject(val *TuneableTrackObject) *NullableTuneableTrackObject {
	return &NullableTuneableTrackObject{value: val, isSet: true}
}

func (v NullableTuneableTrackObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTuneableTrackObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


