# AlbumBaseRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The reason for the restriction. Albums may be restricted if the content is not available in a given market, to the user&#39;s subscription type, or when the user&#39;s account is set to not play explicit content. Additional reasons may be added in the future.  | [optional] 

## Methods

### NewAlbumBaseRestrictions

`func NewAlbumBaseRestrictions() *AlbumBaseRestrictions`

NewAlbumBaseRestrictions instantiates a new AlbumBaseRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumBaseRestrictionsWithDefaults

`func NewAlbumBaseRestrictionsWithDefaults() *AlbumBaseRestrictions`

NewAlbumBaseRestrictionsWithDefaults instantiates a new AlbumBaseRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *AlbumBaseRestrictions) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AlbumBaseRestrictions) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AlbumBaseRestrictions) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AlbumBaseRestrictions) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


