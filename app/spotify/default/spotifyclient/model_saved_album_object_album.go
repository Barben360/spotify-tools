/*
Spotify Web API with fixes and improvements from sonallux

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 2023.2.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spotifyclient

import (
	"encoding/json"
)

// checks if the SavedAlbumObjectAlbum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedAlbumObjectAlbum{}

// SavedAlbumObjectAlbum Information about the album.
type SavedAlbumObjectAlbum struct {
	// The type of the album. 
	AlbumType string `json:"album_type"`
	// The number of tracks in the album.
	TotalTracks int32 `json:"total_tracks"`
	// The markets in which the album is available: [ISO 3166-1 alpha-2 country codes](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). _**NOTE**: an album is considered available in a market when at least 1 of its tracks is available in that market._ 
	AvailableMarkets []string `json:"available_markets"`
	ExternalUrls AlbumBaseExternalUrls `json:"external_urls"`
	// A link to the Web API endpoint providing full details of the album. 
	Href string `json:"href"`
	// The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the album. 
	Id string `json:"id"`
	// The cover art for the album in various sizes, widest first. 
	Images []ImageObject `json:"images"`
	// The name of the album. In case of an album takedown, the value may be an empty string. 
	Name string `json:"name"`
	// The date the album was first released. 
	ReleaseDate string `json:"release_date"`
	// The precision with which `release_date` value is known. 
	ReleaseDatePrecision string `json:"release_date_precision"`
	Restrictions *AlbumBaseRestrictions `json:"restrictions,omitempty"`
	// The object type. 
	Type string `json:"type"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the album. 
	Uri string `json:"uri"`
	// The artists of the album. Each artist object includes a link in `href` to more detailed information about the artist. 
	Artists []SimplifiedArtistObject `json:"artists,omitempty"`
	Tracks *AlbumObjectAllOfTracks `json:"tracks,omitempty"`
	// The popularity of the album, with 100 being the most popular. The popularity is calculated from the popularity of the album's individual tracks.
	Popularity *int32 `json:"popularity,omitempty"`
	// The label for the album.
	Label *string `json:"label,omitempty"`
	ExternalIds *AlbumObjectAllOfExternalIds `json:"external_ids,omitempty"`
	// A list of the genres used to classify the album. (If not yet classified, the array is empty.)
	Genres []string `json:"genres,omitempty"`
	// The copyright statements of the album.
	Copyrights []CopyrightObject `json:"copyrights,omitempty"`
}

// NewSavedAlbumObjectAlbum instantiates a new SavedAlbumObjectAlbum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedAlbumObjectAlbum(albumType string, totalTracks int32, availableMarkets []string, externalUrls AlbumBaseExternalUrls, href string, id string, images []ImageObject, name string, releaseDate string, releaseDatePrecision string, type_ string, uri string) *SavedAlbumObjectAlbum {
	this := SavedAlbumObjectAlbum{}
	this.AlbumType = albumType
	this.TotalTracks = totalTracks
	this.AvailableMarkets = availableMarkets
	this.ExternalUrls = externalUrls
	this.Href = href
	this.Id = id
	this.Images = images
	this.Name = name
	this.ReleaseDate = releaseDate
	this.ReleaseDatePrecision = releaseDatePrecision
	this.Type = type_
	this.Uri = uri
	return &this
}

// NewSavedAlbumObjectAlbumWithDefaults instantiates a new SavedAlbumObjectAlbum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedAlbumObjectAlbumWithDefaults() *SavedAlbumObjectAlbum {
	this := SavedAlbumObjectAlbum{}
	return &this
}

// GetAlbumType returns the AlbumType field value
func (o *SavedAlbumObjectAlbum) GetAlbumType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlbumType
}

// GetAlbumTypeOk returns a tuple with the AlbumType field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetAlbumTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlbumType, true
}

// SetAlbumType sets field value
func (o *SavedAlbumObjectAlbum) SetAlbumType(v string) {
	o.AlbumType = v
}

// GetTotalTracks returns the TotalTracks field value
func (o *SavedAlbumObjectAlbum) GetTotalTracks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalTracks
}

// GetTotalTracksOk returns a tuple with the TotalTracks field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetTotalTracksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalTracks, true
}

// SetTotalTracks sets field value
func (o *SavedAlbumObjectAlbum) SetTotalTracks(v int32) {
	o.TotalTracks = v
}

// GetAvailableMarkets returns the AvailableMarkets field value
func (o *SavedAlbumObjectAlbum) GetAvailableMarkets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableMarkets
}

// GetAvailableMarketsOk returns a tuple with the AvailableMarkets field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetAvailableMarketsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableMarkets, true
}

// SetAvailableMarkets sets field value
func (o *SavedAlbumObjectAlbum) SetAvailableMarkets(v []string) {
	o.AvailableMarkets = v
}

// GetExternalUrls returns the ExternalUrls field value
func (o *SavedAlbumObjectAlbum) GetExternalUrls() AlbumBaseExternalUrls {
	if o == nil {
		var ret AlbumBaseExternalUrls
		return ret
	}

	return o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetExternalUrlsOk() (*AlbumBaseExternalUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrls, true
}

// SetExternalUrls sets field value
func (o *SavedAlbumObjectAlbum) SetExternalUrls(v AlbumBaseExternalUrls) {
	o.ExternalUrls = v
}

// GetHref returns the Href field value
func (o *SavedAlbumObjectAlbum) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *SavedAlbumObjectAlbum) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *SavedAlbumObjectAlbum) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SavedAlbumObjectAlbum) SetId(v string) {
	o.Id = v
}

// GetImages returns the Images field value
func (o *SavedAlbumObjectAlbum) GetImages() []ImageObject {
	if o == nil {
		var ret []ImageObject
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetImagesOk() ([]ImageObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *SavedAlbumObjectAlbum) SetImages(v []ImageObject) {
	o.Images = v
}

// GetName returns the Name field value
func (o *SavedAlbumObjectAlbum) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SavedAlbumObjectAlbum) SetName(v string) {
	o.Name = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *SavedAlbumObjectAlbum) GetReleaseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *SavedAlbumObjectAlbum) SetReleaseDate(v string) {
	o.ReleaseDate = v
}

// GetReleaseDatePrecision returns the ReleaseDatePrecision field value
func (o *SavedAlbumObjectAlbum) GetReleaseDatePrecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDatePrecision
}

// GetReleaseDatePrecisionOk returns a tuple with the ReleaseDatePrecision field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetReleaseDatePrecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDatePrecision, true
}

// SetReleaseDatePrecision sets field value
func (o *SavedAlbumObjectAlbum) SetReleaseDatePrecision(v string) {
	o.ReleaseDatePrecision = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetRestrictions() AlbumBaseRestrictions {
	if o == nil || IsNil(o.Restrictions) {
		var ret AlbumBaseRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetRestrictionsOk() (*AlbumBaseRestrictions, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given AlbumBaseRestrictions and assigns it to the Restrictions field.
func (o *SavedAlbumObjectAlbum) SetRestrictions(v AlbumBaseRestrictions) {
	o.Restrictions = &v
}

// GetType returns the Type field value
func (o *SavedAlbumObjectAlbum) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SavedAlbumObjectAlbum) SetType(v string) {
	o.Type = v
}

// GetUri returns the Uri field value
func (o *SavedAlbumObjectAlbum) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *SavedAlbumObjectAlbum) SetUri(v string) {
	o.Uri = v
}

// GetArtists returns the Artists field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetArtists() []SimplifiedArtistObject {
	if o == nil || IsNil(o.Artists) {
		var ret []SimplifiedArtistObject
		return ret
	}
	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetArtistsOk() ([]SimplifiedArtistObject, bool) {
	if o == nil || IsNil(o.Artists) {
		return nil, false
	}
	return o.Artists, true
}

// HasArtists returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasArtists() bool {
	if o != nil && !IsNil(o.Artists) {
		return true
	}

	return false
}

// SetArtists gets a reference to the given []SimplifiedArtistObject and assigns it to the Artists field.
func (o *SavedAlbumObjectAlbum) SetArtists(v []SimplifiedArtistObject) {
	o.Artists = v
}

// GetTracks returns the Tracks field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetTracks() AlbumObjectAllOfTracks {
	if o == nil || IsNil(o.Tracks) {
		var ret AlbumObjectAllOfTracks
		return ret
	}
	return *o.Tracks
}

// GetTracksOk returns a tuple with the Tracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetTracksOk() (*AlbumObjectAllOfTracks, bool) {
	if o == nil || IsNil(o.Tracks) {
		return nil, false
	}
	return o.Tracks, true
}

// HasTracks returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasTracks() bool {
	if o != nil && !IsNil(o.Tracks) {
		return true
	}

	return false
}

// SetTracks gets a reference to the given AlbumObjectAllOfTracks and assigns it to the Tracks field.
func (o *SavedAlbumObjectAlbum) SetTracks(v AlbumObjectAllOfTracks) {
	o.Tracks = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetPopularity() int32 {
	if o == nil || IsNil(o.Popularity) {
		var ret int32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetPopularityOk() (*int32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given int32 and assigns it to the Popularity field.
func (o *SavedAlbumObjectAlbum) SetPopularity(v int32) {
	o.Popularity = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SavedAlbumObjectAlbum) SetLabel(v string) {
	o.Label = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetExternalIds() AlbumObjectAllOfExternalIds {
	if o == nil || IsNil(o.ExternalIds) {
		var ret AlbumObjectAllOfExternalIds
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetExternalIdsOk() (*AlbumObjectAllOfExternalIds, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given AlbumObjectAllOfExternalIds and assigns it to the ExternalIds field.
func (o *SavedAlbumObjectAlbum) SetExternalIds(v AlbumObjectAllOfExternalIds) {
	o.ExternalIds = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetGenres() []string {
	if o == nil || IsNil(o.Genres) {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *SavedAlbumObjectAlbum) SetGenres(v []string) {
	o.Genres = v
}

// GetCopyrights returns the Copyrights field value if set, zero value otherwise.
func (o *SavedAlbumObjectAlbum) GetCopyrights() []CopyrightObject {
	if o == nil || IsNil(o.Copyrights) {
		var ret []CopyrightObject
		return ret
	}
	return o.Copyrights
}

// GetCopyrightsOk returns a tuple with the Copyrights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedAlbumObjectAlbum) GetCopyrightsOk() ([]CopyrightObject, bool) {
	if o == nil || IsNil(o.Copyrights) {
		return nil, false
	}
	return o.Copyrights, true
}

// HasCopyrights returns a boolean if a field has been set.
func (o *SavedAlbumObjectAlbum) HasCopyrights() bool {
	if o != nil && !IsNil(o.Copyrights) {
		return true
	}

	return false
}

// SetCopyrights gets a reference to the given []CopyrightObject and assigns it to the Copyrights field.
func (o *SavedAlbumObjectAlbum) SetCopyrights(v []CopyrightObject) {
	o.Copyrights = v
}

func (o SavedAlbumObjectAlbum) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedAlbumObjectAlbum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["album_type"] = o.AlbumType
	toSerialize["total_tracks"] = o.TotalTracks
	toSerialize["available_markets"] = o.AvailableMarkets
	toSerialize["external_urls"] = o.ExternalUrls
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	toSerialize["images"] = o.Images
	toSerialize["name"] = o.Name
	toSerialize["release_date"] = o.ReleaseDate
	toSerialize["release_date_precision"] = o.ReleaseDatePrecision
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	toSerialize["type"] = o.Type
	toSerialize["uri"] = o.Uri
	if !IsNil(o.Artists) {
		toSerialize["artists"] = o.Artists
	}
	if !IsNil(o.Tracks) {
		toSerialize["tracks"] = o.Tracks
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["external_ids"] = o.ExternalIds
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.Copyrights) {
		toSerialize["copyrights"] = o.Copyrights
	}
	return toSerialize, nil
}

type NullableSavedAlbumObjectAlbum struct {
	value *SavedAlbumObjectAlbum
	isSet bool
}

func (v NullableSavedAlbumObjectAlbum) Get() *SavedAlbumObjectAlbum {
	return v.value
}

func (v *NullableSavedAlbumObjectAlbum) Set(val *SavedAlbumObjectAlbum) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedAlbumObjectAlbum) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedAlbumObjectAlbum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedAlbumObjectAlbum(val *SavedAlbumObjectAlbum) *NullableSavedAlbumObjectAlbum {
	return &NullableSavedAlbumObjectAlbum{value: val, isSet: true}
}

func (v NullableSavedAlbumObjectAlbum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedAlbumObjectAlbum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


