# PlaylistTracksRefObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | A link to the Web API endpoint where full details of the playlist&#39;s tracks can be retrieved.  | [optional] 
**Total** | Pointer to **int32** | Number of tracks in the playlist.  | [optional] 

## Methods

### NewPlaylistTracksRefObject

`func NewPlaylistTracksRefObject() *PlaylistTracksRefObject`

NewPlaylistTracksRefObject instantiates a new PlaylistTracksRefObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistTracksRefObjectWithDefaults

`func NewPlaylistTracksRefObjectWithDefaults() *PlaylistTracksRefObject`

NewPlaylistTracksRefObjectWithDefaults instantiates a new PlaylistTracksRefObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PlaylistTracksRefObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlaylistTracksRefObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlaylistTracksRefObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlaylistTracksRefObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetTotal

`func (o *PlaylistTracksRefObject) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PlaylistTracksRefObject) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PlaylistTracksRefObject) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PlaylistTracksRefObject) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


