# SimplifiedTrackObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artists** | Pointer to [**[]SimplifiedArtistObject**](SimplifiedArtistObject.md) | The artists who performed the track. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist. | [optional] 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**DiscNumber** | Pointer to **int32** | The disc number (usually &#x60;1&#x60; unless the album consists of more than one disc). | [optional] 
**DurationMs** | Pointer to **int32** | The track length in milliseconds. | [optional] 
**Explicit** | Pointer to **bool** | Whether or not the track has explicit lyrics ( &#x60;true&#x60; &#x3D; yes it does; &#x60;false&#x60; &#x3D; no it does not OR unknown). | [optional] 
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | External URLs for this track.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track. | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**IsPlayable** | Pointer to **bool** | Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking/) is applied. If &#x60;true&#x60;, the track is playable in the given market. Otherwise &#x60;false&#x60;.  | [optional] 
**LinkedFrom** | Pointer to [**LinkedTrackObject**](LinkedTrackObject.md) | Part of the response when [Track Relinking](/documentation/web-api/concepts/track-relinking/) is applied and is only part of the response if the track linking, in fact, exists. The requested track has been replaced with a different track. The track in the &#x60;linked_from&#x60; object contains information about the originally requested track. | [optional] 
**Restrictions** | Pointer to [**TrackRestrictionObject**](TrackRestrictionObject.md) | Included in the response when a content restriction is applied.  | [optional] 
**Name** | Pointer to **string** | The name of the track. | [optional] 
**PreviewUrl** | Pointer to **NullableString** | A URL to a 30 second preview (MP3 format) of the track.  | [optional] 
**TrackNumber** | Pointer to **int32** | The number of the track. If an album has several discs, the track number is the number on the specified disc.  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;track\&quot;.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the track.  | [optional] 
**IsLocal** | Pointer to **bool** | Whether or not the track is from a local file.  | [optional] 

## Methods

### NewSimplifiedTrackObject

`func NewSimplifiedTrackObject() *SimplifiedTrackObject`

NewSimplifiedTrackObject instantiates a new SimplifiedTrackObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedTrackObjectWithDefaults

`func NewSimplifiedTrackObjectWithDefaults() *SimplifiedTrackObject`

NewSimplifiedTrackObjectWithDefaults instantiates a new SimplifiedTrackObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtists

`func (o *SimplifiedTrackObject) GetArtists() []SimplifiedArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *SimplifiedTrackObject) GetArtistsOk() (*[]SimplifiedArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *SimplifiedTrackObject) SetArtists(v []SimplifiedArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *SimplifiedTrackObject) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetAvailableMarkets

`func (o *SimplifiedTrackObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *SimplifiedTrackObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *SimplifiedTrackObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *SimplifiedTrackObject) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetDiscNumber

`func (o *SimplifiedTrackObject) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *SimplifiedTrackObject) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *SimplifiedTrackObject) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *SimplifiedTrackObject) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDurationMs

`func (o *SimplifiedTrackObject) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *SimplifiedTrackObject) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *SimplifiedTrackObject) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *SimplifiedTrackObject) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetExplicit

`func (o *SimplifiedTrackObject) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *SimplifiedTrackObject) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *SimplifiedTrackObject) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *SimplifiedTrackObject) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### GetExternalUrls

`func (o *SimplifiedTrackObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SimplifiedTrackObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SimplifiedTrackObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *SimplifiedTrackObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *SimplifiedTrackObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedTrackObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedTrackObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedTrackObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *SimplifiedTrackObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedTrackObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedTrackObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimplifiedTrackObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPlayable

`func (o *SimplifiedTrackObject) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *SimplifiedTrackObject) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *SimplifiedTrackObject) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.

### HasIsPlayable

`func (o *SimplifiedTrackObject) HasIsPlayable() bool`

HasIsPlayable returns a boolean if a field has been set.

### GetLinkedFrom

`func (o *SimplifiedTrackObject) GetLinkedFrom() LinkedTrackObject`

GetLinkedFrom returns the LinkedFrom field if non-nil, zero value otherwise.

### GetLinkedFromOk

`func (o *SimplifiedTrackObject) GetLinkedFromOk() (*LinkedTrackObject, bool)`

GetLinkedFromOk returns a tuple with the LinkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFrom

`func (o *SimplifiedTrackObject) SetLinkedFrom(v LinkedTrackObject)`

SetLinkedFrom sets LinkedFrom field to given value.

### HasLinkedFrom

`func (o *SimplifiedTrackObject) HasLinkedFrom() bool`

HasLinkedFrom returns a boolean if a field has been set.

### GetRestrictions

`func (o *SimplifiedTrackObject) GetRestrictions() TrackRestrictionObject`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *SimplifiedTrackObject) GetRestrictionsOk() (*TrackRestrictionObject, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *SimplifiedTrackObject) SetRestrictions(v TrackRestrictionObject)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *SimplifiedTrackObject) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedTrackObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedTrackObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedTrackObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedTrackObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *SimplifiedTrackObject) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *SimplifiedTrackObject) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *SimplifiedTrackObject) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *SimplifiedTrackObject) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### SetPreviewUrlNil

`func (o *SimplifiedTrackObject) SetPreviewUrlNil(b bool)`

 SetPreviewUrlNil sets the value for PreviewUrl to be an explicit nil

### UnsetPreviewUrl
`func (o *SimplifiedTrackObject) UnsetPreviewUrl()`

UnsetPreviewUrl ensures that no value is present for PreviewUrl, not even an explicit nil
### GetTrackNumber

`func (o *SimplifiedTrackObject) GetTrackNumber() int32`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *SimplifiedTrackObject) GetTrackNumberOk() (*int32, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *SimplifiedTrackObject) SetTrackNumber(v int32)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *SimplifiedTrackObject) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedTrackObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedTrackObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedTrackObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedTrackObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *SimplifiedTrackObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SimplifiedTrackObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SimplifiedTrackObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *SimplifiedTrackObject) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetIsLocal

`func (o *SimplifiedTrackObject) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *SimplifiedTrackObject) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *SimplifiedTrackObject) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *SimplifiedTrackObject) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


