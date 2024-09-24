# AudiobookObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authors** | [**[]AuthorObject**](AuthorObject.md) | The author(s) for the audiobook.  | 
**AvailableMarkets** | **[]string** | A list of the countries in which the audiobook can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | 
**Copyrights** | [**[]CopyrightObject**](CopyrightObject.md) | The copyright statements of the audiobook.  | 
**Description** | **string** | A description of the audiobook. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the audiobook. This field may contain HTML tags.  | 
**Edition** | Pointer to **string** | The edition of the audiobook.  | [optional] 
**Explicit** | **bool** | Whether or not the audiobook has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalUrls** | [**ExternalUrlObject**](ExternalUrlObject.md) | External URLs for this audiobook.  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the audiobook.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the audiobook in various sizes, widest first.  | 
**Languages** | **[]string** | A list of the languages used in the audiobook, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**MediaType** | **string** | The media type of the audiobook.  | 
**Name** | **string** | The name of the audiobook.  | 
**Narrators** | [**[]NarratorObject**](NarratorObject.md) | The narrator(s) for the audiobook.  | 
**Publisher** | **string** | The publisher of the audiobook.  | 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/concepts/spotify-uris-ids) for the audiobook.  | 
**TotalChapters** | **int32** | The number of chapters in this audiobook.  | 
**Chapters** | **map[string]interface{}** | The chapters of the audiobook.  | 

## Methods

### NewAudiobookObject

`func NewAudiobookObject(authors []AuthorObject, availableMarkets []string, copyrights []CopyrightObject, description string, htmlDescription string, explicit bool, externalUrls ExternalUrlObject, href string, id string, images []ImageObject, languages []string, mediaType string, name string, narrators []NarratorObject, publisher string, type_ string, uri string, totalChapters int32, chapters map[string]interface{}, ) *AudiobookObject`

NewAudiobookObject instantiates a new AudiobookObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudiobookObjectWithDefaults

`func NewAudiobookObjectWithDefaults() *AudiobookObject`

NewAudiobookObjectWithDefaults instantiates a new AudiobookObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthors

`func (o *AudiobookObject) GetAuthors() []AuthorObject`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *AudiobookObject) GetAuthorsOk() (*[]AuthorObject, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *AudiobookObject) SetAuthors(v []AuthorObject)`

SetAuthors sets Authors field to given value.


### GetAvailableMarkets

`func (o *AudiobookObject) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *AudiobookObject) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *AudiobookObject) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.


### GetCopyrights

`func (o *AudiobookObject) GetCopyrights() []CopyrightObject`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *AudiobookObject) GetCopyrightsOk() (*[]CopyrightObject, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *AudiobookObject) SetCopyrights(v []CopyrightObject)`

SetCopyrights sets Copyrights field to given value.


### GetDescription

`func (o *AudiobookObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AudiobookObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AudiobookObject) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *AudiobookObject) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *AudiobookObject) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *AudiobookObject) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetEdition

`func (o *AudiobookObject) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *AudiobookObject) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *AudiobookObject) SetEdition(v string)`

SetEdition sets Edition field to given value.

### HasEdition

`func (o *AudiobookObject) HasEdition() bool`

HasEdition returns a boolean if a field has been set.

### GetExplicit

`func (o *AudiobookObject) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *AudiobookObject) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *AudiobookObject) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *AudiobookObject) GetExternalUrls() ExternalUrlObject`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *AudiobookObject) GetExternalUrlsOk() (*ExternalUrlObject, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *AudiobookObject) SetExternalUrls(v ExternalUrlObject)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *AudiobookObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AudiobookObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AudiobookObject) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *AudiobookObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudiobookObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudiobookObject) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *AudiobookObject) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *AudiobookObject) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *AudiobookObject) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetLanguages

`func (o *AudiobookObject) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *AudiobookObject) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *AudiobookObject) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetMediaType

`func (o *AudiobookObject) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AudiobookObject) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AudiobookObject) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetName

`func (o *AudiobookObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudiobookObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudiobookObject) SetName(v string)`

SetName sets Name field to given value.


### GetNarrators

`func (o *AudiobookObject) GetNarrators() []NarratorObject`

GetNarrators returns the Narrators field if non-nil, zero value otherwise.

### GetNarratorsOk

`func (o *AudiobookObject) GetNarratorsOk() (*[]NarratorObject, bool)`

GetNarratorsOk returns a tuple with the Narrators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNarrators

`func (o *AudiobookObject) SetNarrators(v []NarratorObject)`

SetNarrators sets Narrators field to given value.


### GetPublisher

`func (o *AudiobookObject) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AudiobookObject) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AudiobookObject) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetType

`func (o *AudiobookObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudiobookObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudiobookObject) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *AudiobookObject) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AudiobookObject) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AudiobookObject) SetUri(v string)`

SetUri sets Uri field to given value.


### GetTotalChapters

`func (o *AudiobookObject) GetTotalChapters() int32`

GetTotalChapters returns the TotalChapters field if non-nil, zero value otherwise.

### GetTotalChaptersOk

`func (o *AudiobookObject) GetTotalChaptersOk() (*int32, bool)`

GetTotalChaptersOk returns a tuple with the TotalChapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChapters

`func (o *AudiobookObject) SetTotalChapters(v int32)`

SetTotalChapters sets TotalChapters field to given value.


### GetChapters

`func (o *AudiobookObject) GetChapters() map[string]interface{}`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *AudiobookObject) GetChaptersOk() (*map[string]interface{}, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *AudiobookObject) SetChapters(v map[string]interface{})`

SetChapters sets Chapters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


