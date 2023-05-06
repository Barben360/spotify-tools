# AudioAnalysisObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyzerVersion** | Pointer to **string** | The version of the Analyzer used to analyze this track. | [optional] 
**Platform** | Pointer to **string** | The platform used to read the track&#39;s audio data. | [optional] 
**DetailedStatus** | Pointer to **string** | A detailed status code for this track. If analysis data is missing, this code may explain why. | [optional] 
**StatusCode** | Pointer to **int32** | The return code of the analyzer process. 0 if successful, 1 if any errors occurred. | [optional] 
**Timestamp** | Pointer to **int32** | The Unix timestamp (in seconds) at which this track was analyzed. | [optional] 
**AnalysisTime** | Pointer to **float32** | The amount of time taken to analyze this track. | [optional] 
**InputProcess** | Pointer to **string** | The method used to read the track&#39;s audio data. | [optional] 

## Methods

### NewAudioAnalysisObjectMeta

`func NewAudioAnalysisObjectMeta() *AudioAnalysisObjectMeta`

NewAudioAnalysisObjectMeta instantiates a new AudioAnalysisObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioAnalysisObjectMetaWithDefaults

`func NewAudioAnalysisObjectMetaWithDefaults() *AudioAnalysisObjectMeta`

NewAudioAnalysisObjectMetaWithDefaults instantiates a new AudioAnalysisObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyzerVersion

`func (o *AudioAnalysisObjectMeta) GetAnalyzerVersion() string`

GetAnalyzerVersion returns the AnalyzerVersion field if non-nil, zero value otherwise.

### GetAnalyzerVersionOk

`func (o *AudioAnalysisObjectMeta) GetAnalyzerVersionOk() (*string, bool)`

GetAnalyzerVersionOk returns a tuple with the AnalyzerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzerVersion

`func (o *AudioAnalysisObjectMeta) SetAnalyzerVersion(v string)`

SetAnalyzerVersion sets AnalyzerVersion field to given value.

### HasAnalyzerVersion

`func (o *AudioAnalysisObjectMeta) HasAnalyzerVersion() bool`

HasAnalyzerVersion returns a boolean if a field has been set.

### GetPlatform

`func (o *AudioAnalysisObjectMeta) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AudioAnalysisObjectMeta) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AudioAnalysisObjectMeta) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AudioAnalysisObjectMeta) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDetailedStatus

`func (o *AudioAnalysisObjectMeta) GetDetailedStatus() string`

GetDetailedStatus returns the DetailedStatus field if non-nil, zero value otherwise.

### GetDetailedStatusOk

`func (o *AudioAnalysisObjectMeta) GetDetailedStatusOk() (*string, bool)`

GetDetailedStatusOk returns a tuple with the DetailedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedStatus

`func (o *AudioAnalysisObjectMeta) SetDetailedStatus(v string)`

SetDetailedStatus sets DetailedStatus field to given value.

### HasDetailedStatus

`func (o *AudioAnalysisObjectMeta) HasDetailedStatus() bool`

HasDetailedStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *AudioAnalysisObjectMeta) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AudioAnalysisObjectMeta) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AudioAnalysisObjectMeta) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AudioAnalysisObjectMeta) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *AudioAnalysisObjectMeta) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AudioAnalysisObjectMeta) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AudioAnalysisObjectMeta) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AudioAnalysisObjectMeta) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAnalysisTime

`func (o *AudioAnalysisObjectMeta) GetAnalysisTime() float32`

GetAnalysisTime returns the AnalysisTime field if non-nil, zero value otherwise.

### GetAnalysisTimeOk

`func (o *AudioAnalysisObjectMeta) GetAnalysisTimeOk() (*float32, bool)`

GetAnalysisTimeOk returns a tuple with the AnalysisTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisTime

`func (o *AudioAnalysisObjectMeta) SetAnalysisTime(v float32)`

SetAnalysisTime sets AnalysisTime field to given value.

### HasAnalysisTime

`func (o *AudioAnalysisObjectMeta) HasAnalysisTime() bool`

HasAnalysisTime returns a boolean if a field has been set.

### GetInputProcess

`func (o *AudioAnalysisObjectMeta) GetInputProcess() string`

GetInputProcess returns the InputProcess field if non-nil, zero value otherwise.

### GetInputProcessOk

`func (o *AudioAnalysisObjectMeta) GetInputProcessOk() (*string, bool)`

GetInputProcessOk returns a tuple with the InputProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputProcess

`func (o *AudioAnalysisObjectMeta) SetInputProcess(v string)`

SetInputProcess sets InputProcess field to given value.

### HasInputProcess

`func (o *AudioAnalysisObjectMeta) HasInputProcess() bool`

HasInputProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


