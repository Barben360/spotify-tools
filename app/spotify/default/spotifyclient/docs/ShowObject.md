# ShowObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableMarkets** | **[]string** | A list of the countries in which the show can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | 
**Copyrights** | [**[]CopyrightObject**](CopyrightObject.md) | The copyright statements of the show.  | 
**Description** | **string** | A description of the show. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the show. This field may contain HTML tags.  | 
**Explicit** | **bool** | Whether or not the show has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalUrls** | [**ExternalUrlObject**](ExternalUrlObject.md) | External URLs for this show.  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the show.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the show.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the show in various sizes, widest first.  | 
**IsExternallyHosted** | **bool** | True if all of the shows episodes are hosted outside of Spotify&#39;s CDN. This field might be &#x60;null&#x60; in some cases.  | 
**Languages** | **[]string** | A list of the languages used in the show, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**MediaType** | **string** | The media type of the show.  | 
**Name** | **string** | The name of the episode.  | 
**Publisher** | **string** | The publisher of the show.  | 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the show.  | 
**TotalEpisodes** | **int32** | The total number of episodes in the show.  | 
**Episodes** | **map[string]interface{}** | The episodes of the show.  | 

## Methods

### NewShowObject

`func NewShowObject(availableMarkets []string, copyrights []CopyrightObject, description string, htmlDescription string, explicit bool, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, isExternallyHosted bool, languages []string, mediaType string, name string, publisher string, type_ string, uri string, totalEpisodes int32, episodes map[string]interface{}, ) *ShowObject`

NewShowObject instantiates a new ShowObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowObjectWithDefaults

`func NewShowObjectWithDefaults() *ShowObject`

NewShowObjectWithDefaults instantiates a new ShowObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableMarkets

`func (o *ShowObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *ShowObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *ShowObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.


### GetCopyrights

`func (o *ShowObject) GetCopyrights() []CopyrightObject`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *ShowObject) GetCopyrightsOk() (*[]CopyrightObject, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *ShowObject) SetCopyrights(v []CopyrightObject)`

SetCopyrights sets Copyrights field to given value.


### GetDescription

`func (o *ShowObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShowObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShowObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *ShowObject) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *ShowObject) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *ShowObject) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetExplicit

`func (o *ShowObject) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *ShowObject) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *ShowObject) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *ShowObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *ShowObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *ShowObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *ShowObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ShowObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ShowObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *ShowObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShowObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShowObject) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *ShowObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ShowObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ShowObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsExternallyHosted

`func (o *ShowObject) GetIsExternallyHosted() bool`

GetIsExternallyHosted returns the IsExternallyHosted field if non-nil, zero value otherwise.

### GetIsExternallyHostedOk

`func (o *ShowObject) GetIsExternallyHostedOk() (*bool, bool)`

GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyHosted

`func (o *ShowObject) SetIsExternallyHosted(v bool)`

SetIsExternallyHosted sets IsExternallyHosted field to given value.


### GetLanguages

`func (o *ShowObject) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ShowObject) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ShowObject) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetMediaType

`func (o *ShowObject) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ShowObject) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ShowObject) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetName

`func (o *ShowObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShowObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShowObject) SetName(v string)`

SetName sets Name field to given value.


### GetPublisher

`func (o *ShowObject) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ShowObject) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ShowObject) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetType

`func (o *ShowObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShowObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShowObject) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *ShowObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ShowObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ShowObject) SetUri(v string)`

SetUri sets Uri field to given value.


### GetTotalEpisodes

`func (o *ShowObject) GetTotalEpisodes() int32`

GetTotalEpisodes returns the TotalEpisodes field if non-nil, zero value otherwise.

### GetTotalEpisodesOk

`func (o *ShowObject) GetTotalEpisodesOk() (*int32, bool)`

GetTotalEpisodesOk returns a tuple with the TotalEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEpisodes

`func (o *ShowObject) SetTotalEpisodes(v int32)`

SetTotalEpisodes sets TotalEpisodes field to given value.


### GetEpisodes

`func (o *ShowObject) GetEpisodes() map[string]interface{}`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *ShowObject) GetEpisodesOk() (*map[string]interface{}, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *ShowObject) SetEpisodes(v map[string]interface{})`

SetEpisodes sets Episodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


