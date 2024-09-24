# ArtistObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUrls** | Pointer to [**ExternalUrlObject**](ExternalUrlObject.md) | Known external URLs for this artist.  | [optional] 
**Followers** | Pointer to [**FollowersObject**](FollowersObject.md) | Information about the followers of the artist.  | [optional] 
**Genres** | Pointer to **[]string** | A list of the genres the artist is associated with. If not yet classified, the array is empty.  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the artist.  | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the artist.  | [optional] 
**Images** | Pointer to [**[]ImageObject**](ImageObject.md) | Images of the artist in various sizes, widest first.  | [optional] 
**Name** | Pointer to **string** | The name of the artist.  | [optional] 
**Popularity** | Pointer to **int32** | The popularity of the artist. The value will be between 0 and 100, with 100 being the most popular. The artist&#39;s popularity is calculated from the popularity of all the artist&#39;s tracks.  | [optional] 
**Type** | Pointer to **string** | The object type.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the artist.  | [optional] 

## Methods

### NewArtistObject

`func NewArtistObject() *ArtistObject`

NewArtistObject instantiates a new ArtistObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistObjectWithDefaults

`func NewArtistObjectWithDefaults() *ArtistObject`

NewArtistObjectWithDefaults instantiates a new ArtistObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUrls

`func (o *ArtistObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *ArtistObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *ArtistObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *ArtistObject) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetFollowers

`func (o *ArtistObject) GetFollowers() FollowersObject`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *ArtistObject) GetFollowersOk() (*FollowersObject, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *ArtistObject) SetFollowers(v FollowersObject)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *ArtistObject) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.

### GetGenres

`func (o *ArtistObject) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *ArtistObject) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *ArtistObject) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *ArtistObject) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetHref

`func (o *ArtistObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ArtistObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ArtistObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ArtistObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ArtistObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtistObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtistObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ArtistObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *ArtistObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ArtistObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ArtistObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.

### HasImages

`func (o *ArtistObject) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetName

`func (o *ArtistObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtistObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtistObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtistObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPopularity

`func (o *ArtistObject) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *ArtistObject) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *ArtistObject) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *ArtistObject) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetType

`func (o *ArtistObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtistObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtistObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ArtistObject) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *ArtistObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ArtistObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ArtistObject) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ArtistObject) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


