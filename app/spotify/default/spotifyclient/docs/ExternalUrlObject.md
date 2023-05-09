# ExternalUrlObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spotify** | Pointer to **string** | The [Spotify URL](/documentation/web-api/concepts/spotify-uris-ids) for the object.  | [optional] 

## Methods

### NewExternalUrlObject

`func NewExternalUrlObject() *ExternalUrlObject`

NewExternalUrlObject instantiates a new ExternalUrlObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalUrlObjectWithDefaults

`func NewExternalUrlObjectWithDefaults() *ExternalUrlObject`

NewExternalUrlObjectWithDefaults instantiates a new ExternalUrlObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpotify

`func (o *ExternalUrlObject) GetSpotify() string`

GetSpotify returns the Spotify field if non-nil, zero value otherwise.

### GetSpotifyOk

`func (o *ExternalUrlObject) GetSpotifyOk() (*string, bool)`

GetSpotifyOk returns a tuple with the Spotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotify

`func (o *ExternalUrlObject) SetSpotify(v string)`

SetSpotify sets Spotify field to given value.

### HasSpotify

`func (o *ExternalUrlObject) HasSpotify() bool`

HasSpotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


