# SimplifiedArtistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this artist.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the artist.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.  | [optional] 
**Name** | Pointer to **string** | The name of the artist.  | [optional] 
**Type** | Pointer to **string** | The object type.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.  | [optional] 

## Methods

### NewSimplifiedArtistObject

`func NewSimplifiedArtistObject() *SimplifiedArtistObject`

NewSimplifiedArtistObject instantiates a new SimplifiedArtistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedArtistObjectWithDefaults

`func NewSimplifiedArtistObjectWithDefaults() *SimplifiedArtistObject`

NewSimplifiedArtistObjectWithDefaults instantiates a new SimplifiedArtistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *SimplifiedArtistObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SimplifiedArtistObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SimplifiedArtistObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *SimplifiedArtistObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *SimplifiedArtistObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedArtistObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedArtistObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedArtistObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SimplifiedArtistObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedArtistObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedArtistObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimplifiedArtistObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedArtistObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedArtistObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedArtistObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedArtistObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedArtistObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedArtistObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedArtistObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedArtistObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *SimplifiedArtistObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SimplifiedArtistObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SimplifiedArtistObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *SimplifiedArtistObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


