# TrackObjectLinkedFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**LinkedTrackObjectExternalUrls**](LinkedTrackObjectExternalUrls.md) |  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;track\&quot;.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 

## Methods

### NewTrackObjectLinkedFrom

`func NewTrackObjectLinkedFrom() *TrackObjectLinkedFrom`

NewTrackObjectLinkedFrom instantiates a new TrackObjectLinkedFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackObjectLinkedFromWithDefaults

`func NewTrackObjectLinkedFromWithDefaults() *TrackObjectLinkedFrom`

NewTrackObjectLinkedFromWithDefaults instantiates a new TrackObjectLinkedFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *TrackObjectLinkedFrom) GetExternalUrls() LinkedTrackObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *TrackObjectLinkedFrom) GetExternalUrlsOk() (*LinkedTrackObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *TrackObjectLinkedFrom) SetExternalUrls(v LinkedTrackObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *TrackObjectLinkedFrom) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *TrackObjectLinkedFrom) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TrackObjectLinkedFrom) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TrackObjectLinkedFrom) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TrackObjectLinkedFrom) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *TrackObjectLinkedFrom) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackObjectLinkedFrom) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackObjectLinkedFrom) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackObjectLinkedFrom) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TrackObjectLinkedFrom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrackObjectLinkedFrom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrackObjectLinkedFrom) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrackObjectLinkedFrom) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *TrackObjectLinkedFrom) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TrackObjectLinkedFrom) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TrackObjectLinkedFrom) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TrackObjectLinkedFrom) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


