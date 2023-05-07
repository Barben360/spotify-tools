# \TracksApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTracksToPlaylist**](TracksApi.md#AddTracksToPlaylist) | **Post** /playlists/{playlist_id}/tracks | Add Items to Playlist 
[**CheckUsersSavedTracks**](TracksApi.md#CheckUsersSavedTracks) | **Get** /me/tracks/contains | Check User&#39;s Saved Tracks 
[**GetAnAlbumsTracks**](TracksApi.md#GetAnAlbumsTracks) | **Get** /albums/{id}/tracks | Get Album Tracks 
[**GetAnArtistsTopTracks**](TracksApi.md#GetAnArtistsTopTracks) | **Get** /artists/{id}/top-tracks | Get Artist&#39;s Top Tracks 
[**GetAudioAnalysis**](TracksApi.md#GetAudioAnalysis) | **Get** /audio-analysis/{id} | Get Track&#39;s Audio Analysis 
[**GetAudioFeatures**](TracksApi.md#GetAudioFeatures) | **Get** /audio-features/{id} | Get Track&#39;s Audio Features 
[**GetPlaylistsTracks**](TracksApi.md#GetPlaylistsTracks) | **Get** /playlists/{playlist_id}/tracks | Get Playlist Items 
[**GetRecommendations**](TracksApi.md#GetRecommendations) | **Get** /recommendations | Get Recommendations 
[**GetSeveralAudioFeatures**](TracksApi.md#GetSeveralAudioFeatures) | **Get** /audio-features | Get Tracks&#39; Audio Features 
[**GetSeveralTracks**](TracksApi.md#GetSeveralTracks) | **Get** /tracks | Get Several Tracks 
[**GetTrack**](TracksApi.md#GetTrack) | **Get** /tracks/{id} | Get Track 
[**GetUsersSavedTracks**](TracksApi.md#GetUsersSavedTracks) | **Get** /me/tracks | Get User&#39;s Saved Tracks 
[**GetUsersTopTracks**](TracksApi.md#GetUsersTopTracks) | **Get** /me/top/tracks | Get User&#39;s Top Tracks 
[**RemoveTracksPlaylist**](TracksApi.md#RemoveTracksPlaylist) | **Delete** /playlists/{playlist_id}/tracks | Remove Playlist Items 
[**RemoveTracksUser**](TracksApi.md#RemoveTracksUser) | **Delete** /me/tracks | Remove User&#39;s Saved Tracks 
[**ReorderOrReplacePlaylistsTracks**](TracksApi.md#ReorderOrReplacePlaylistsTracks) | **Put** /playlists/{playlist_id}/tracks | Update Playlist Items 
[**SaveTracksUser**](TracksApi.md#SaveTracksUser) | **Put** /me/tracks | Save Tracks for Current User 



## AddTracksToPlaylist

> ReorderOrReplacePlaylistsTracks200Response AddTracksToPlaylist(ctx, playlistId).Position(position).Uris(uris).RequestBody(requestBody).Execute()

Add Items to Playlist 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
    position := int32(0) // int32 |  (optional)
    uris := "spotify:track:4iV5W9uYEdYUVa79Axb7Rh,spotify:track:1301WleyT98MSxVHPZCA6M" // string |  (optional)
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.AddTracksToPlaylist(context.Background(), playlistId).Position(position).Uris(uris).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.AddTracksToPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTracksToPlaylist`: ReorderOrReplacePlaylistsTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.AddTracksToPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTracksToPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **position** | **int32** |  | 
 **uris** | **string** |  | 
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**ReorderOrReplacePlaylistsTracks200Response**](ReorderOrReplacePlaylistsTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckUsersSavedTracks

> []bool CheckUsersSavedTracks(ctx).Ids(ids).Execute()

Check User's Saved Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.CheckUsersSavedTracks(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.CheckUsersSavedTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckUsersSavedTracks`: []bool
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.CheckUsersSavedTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

**[]bool**

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnAlbumsTracks

> PagingSimplifiedTrackObject GetAnAlbumsTracks(ctx, id).Market(market).Limit(limit).Offset(offset).Execute()

Get Album Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    id := "4aawyAB9vmqN3uQ7FjRGTy" // string | 
    market := "ES" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetAnAlbumsTracks(context.Background(), id).Market(market).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetAnAlbumsTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnAlbumsTracks`: PagingSimplifiedTrackObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetAnAlbumsTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnAlbumsTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSimplifiedTrackObject**](PagingSimplifiedTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnArtistsTopTracks

> GetAnArtistsTopTracks200Response GetAnArtistsTopTracks(ctx, id).Market(market).Execute()

Get Artist's Top Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    id := "0TnOYISbd1XYRBk9myaseg" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetAnArtistsTopTracks(context.Background(), id).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetAnArtistsTopTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnArtistsTopTracks`: GetAnArtistsTopTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetAnArtistsTopTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnArtistsTopTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**GetAnArtistsTopTracks200Response**](GetAnArtistsTopTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudioAnalysis

> AudioAnalysisObject GetAudioAnalysis(ctx, id).Execute()

Get Track's Audio Analysis 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    id := "11dFghVXANMlKmJXsNCbNl" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetAudioAnalysis(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetAudioAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudioAnalysis`: AudioAnalysisObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetAudioAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudioAnalysisObject**](AudioAnalysisObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudioFeatures

> AudioFeaturesObject GetAudioFeatures(ctx, id).Execute()

Get Track's Audio Features 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    id := "11dFghVXANMlKmJXsNCbNl" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetAudioFeatures(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetAudioFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudioFeatures`: AudioFeaturesObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetAudioFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudioFeaturesObject**](AudioFeaturesObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistsTracks

> PagingPlaylistTrackObject GetPlaylistsTracks(ctx, playlistId).Market(market).Fields(fields).Limit(limit).Offset(offset).AdditionalTypes(additionalTypes).Execute()

Get Playlist Items 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
    market := "ES" // string |  (optional)
    fields := "items(added_by.id,track(name,href,album(name,href)))" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)
    additionalTypes := "additionalTypes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetPlaylistsTracks(context.Background(), playlistId).Market(market).Fields(fields).Limit(limit).Offset(offset).AdditionalTypes(additionalTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetPlaylistsTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlaylistsTracks`: PagingPlaylistTrackObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetPlaylistsTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistsTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **fields** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **additionalTypes** | **string** |  | 

### Return type

[**PagingPlaylistTrackObject**](PagingPlaylistTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendations

> RecommendationsObject GetRecommendations(ctx).Limit(limit).Market(market).SeedArtists(seedArtists).SeedGenres(seedGenres).SeedTracks(seedTracks).MinAcousticness(minAcousticness).MaxAcousticness(maxAcousticness).TargetAcousticness(targetAcousticness).MinDanceability(minDanceability).MaxDanceability(maxDanceability).TargetDanceability(targetDanceability).MinDurationMs(minDurationMs).MaxDurationMs(maxDurationMs).TargetDurationMs(targetDurationMs).MinEnergy(minEnergy).MaxEnergy(maxEnergy).TargetEnergy(targetEnergy).MinInstrumentalness(minInstrumentalness).MaxInstrumentalness(maxInstrumentalness).TargetInstrumentalness(targetInstrumentalness).MinKey(minKey).MaxKey(maxKey).TargetKey(targetKey).MinLiveness(minLiveness).MaxLiveness(maxLiveness).TargetLiveness(targetLiveness).MinLoudness(minLoudness).MaxLoudness(maxLoudness).TargetLoudness(targetLoudness).MinMode(minMode).MaxMode(maxMode).TargetMode(targetMode).MinPopularity(minPopularity).MaxPopularity(maxPopularity).TargetPopularity(targetPopularity).MinSpeechiness(minSpeechiness).MaxSpeechiness(maxSpeechiness).TargetSpeechiness(targetSpeechiness).MinTempo(minTempo).MaxTempo(maxTempo).TargetTempo(targetTempo).MinTimeSignature(minTimeSignature).MaxTimeSignature(maxTimeSignature).TargetTimeSignature(targetTimeSignature).MinValence(minValence).MaxValence(maxValence).TargetValence(targetValence).Execute()

Get Recommendations 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    limit := int32(10) // int32 |  (optional) (default to 20)
    market := "ES" // string |  (optional)
    seedArtists := "4NHQUGzhtTLFvgF5SZesLK" // string |  (optional)
    seedGenres := "classical,country" // string |  (optional)
    seedTracks := "0c6xIDDpzE81m2q797ordA" // string |  (optional)
    minAcousticness := float32(8.14) // float32 |  (optional)
    maxAcousticness := float32(8.14) // float32 |  (optional)
    targetAcousticness := float32(8.14) // float32 |  (optional)
    minDanceability := float32(8.14) // float32 |  (optional)
    maxDanceability := float32(8.14) // float32 |  (optional)
    targetDanceability := float32(8.14) // float32 |  (optional)
    minDurationMs := int32(56) // int32 |  (optional)
    maxDurationMs := int32(56) // int32 |  (optional)
    targetDurationMs := int32(56) // int32 |  (optional)
    minEnergy := float32(8.14) // float32 |  (optional)
    maxEnergy := float32(8.14) // float32 |  (optional)
    targetEnergy := float32(8.14) // float32 |  (optional)
    minInstrumentalness := float32(8.14) // float32 |  (optional)
    maxInstrumentalness := float32(8.14) // float32 |  (optional)
    targetInstrumentalness := float32(8.14) // float32 |  (optional)
    minKey := int32(56) // int32 |  (optional)
    maxKey := int32(56) // int32 |  (optional)
    targetKey := int32(56) // int32 |  (optional)
    minLiveness := float32(8.14) // float32 |  (optional)
    maxLiveness := float32(8.14) // float32 |  (optional)
    targetLiveness := float32(8.14) // float32 |  (optional)
    minLoudness := float32(8.14) // float32 |  (optional)
    maxLoudness := float32(8.14) // float32 |  (optional)
    targetLoudness := float32(8.14) // float32 |  (optional)
    minMode := int32(56) // int32 |  (optional)
    maxMode := int32(56) // int32 |  (optional)
    targetMode := int32(56) // int32 |  (optional)
    minPopularity := int32(56) // int32 |  (optional)
    maxPopularity := int32(56) // int32 |  (optional)
    targetPopularity := int32(56) // int32 |  (optional)
    minSpeechiness := float32(8.14) // float32 |  (optional)
    maxSpeechiness := float32(8.14) // float32 |  (optional)
    targetSpeechiness := float32(8.14) // float32 |  (optional)
    minTempo := float32(8.14) // float32 |  (optional)
    maxTempo := float32(8.14) // float32 |  (optional)
    targetTempo := float32(8.14) // float32 |  (optional)
    minTimeSignature := int32(56) // int32 |  (optional)
    maxTimeSignature := int32(56) // int32 |  (optional)
    targetTimeSignature := int32(56) // int32 |  (optional)
    minValence := float32(8.14) // float32 |  (optional)
    maxValence := float32(8.14) // float32 |  (optional)
    targetValence := float32(8.14) // float32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetRecommendations(context.Background()).Limit(limit).Market(market).SeedArtists(seedArtists).SeedGenres(seedGenres).SeedTracks(seedTracks).MinAcousticness(minAcousticness).MaxAcousticness(maxAcousticness).TargetAcousticness(targetAcousticness).MinDanceability(minDanceability).MaxDanceability(maxDanceability).TargetDanceability(targetDanceability).MinDurationMs(minDurationMs).MaxDurationMs(maxDurationMs).TargetDurationMs(targetDurationMs).MinEnergy(minEnergy).MaxEnergy(maxEnergy).TargetEnergy(targetEnergy).MinInstrumentalness(minInstrumentalness).MaxInstrumentalness(maxInstrumentalness).TargetInstrumentalness(targetInstrumentalness).MinKey(minKey).MaxKey(maxKey).TargetKey(targetKey).MinLiveness(minLiveness).MaxLiveness(maxLiveness).TargetLiveness(targetLiveness).MinLoudness(minLoudness).MaxLoudness(maxLoudness).TargetLoudness(targetLoudness).MinMode(minMode).MaxMode(maxMode).TargetMode(targetMode).MinPopularity(minPopularity).MaxPopularity(maxPopularity).TargetPopularity(targetPopularity).MinSpeechiness(minSpeechiness).MaxSpeechiness(maxSpeechiness).TargetSpeechiness(targetSpeechiness).MinTempo(minTempo).MaxTempo(maxTempo).TargetTempo(targetTempo).MinTimeSignature(minTimeSignature).MaxTimeSignature(maxTimeSignature).TargetTimeSignature(targetTimeSignature).MinValence(minValence).MaxValence(maxValence).TargetValence(targetValence).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendations`: RecommendationsObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **market** | **string** |  | 
 **seedArtists** | **string** |  | 
 **seedGenres** | **string** |  | 
 **seedTracks** | **string** |  | 
 **minAcousticness** | **float32** |  | 
 **maxAcousticness** | **float32** |  | 
 **targetAcousticness** | **float32** |  | 
 **minDanceability** | **float32** |  | 
 **maxDanceability** | **float32** |  | 
 **targetDanceability** | **float32** |  | 
 **minDurationMs** | **int32** |  | 
 **maxDurationMs** | **int32** |  | 
 **targetDurationMs** | **int32** |  | 
 **minEnergy** | **float32** |  | 
 **maxEnergy** | **float32** |  | 
 **targetEnergy** | **float32** |  | 
 **minInstrumentalness** | **float32** |  | 
 **maxInstrumentalness** | **float32** |  | 
 **targetInstrumentalness** | **float32** |  | 
 **minKey** | **int32** |  | 
 **maxKey** | **int32** |  | 
 **targetKey** | **int32** |  | 
 **minLiveness** | **float32** |  | 
 **maxLiveness** | **float32** |  | 
 **targetLiveness** | **float32** |  | 
 **minLoudness** | **float32** |  | 
 **maxLoudness** | **float32** |  | 
 **targetLoudness** | **float32** |  | 
 **minMode** | **int32** |  | 
 **maxMode** | **int32** |  | 
 **targetMode** | **int32** |  | 
 **minPopularity** | **int32** |  | 
 **maxPopularity** | **int32** |  | 
 **targetPopularity** | **int32** |  | 
 **minSpeechiness** | **float32** |  | 
 **maxSpeechiness** | **float32** |  | 
 **targetSpeechiness** | **float32** |  | 
 **minTempo** | **float32** |  | 
 **maxTempo** | **float32** |  | 
 **targetTempo** | **float32** |  | 
 **minTimeSignature** | **int32** |  | 
 **maxTimeSignature** | **int32** |  | 
 **targetTimeSignature** | **int32** |  | 
 **minValence** | **float32** |  | 
 **maxValence** | **float32** |  | 
 **targetValence** | **float32** |  | 

### Return type

[**RecommendationsObject**](RecommendationsObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeveralAudioFeatures

> GetSeveralAudioFeatures200Response GetSeveralAudioFeatures(ctx).Ids(ids).Execute()

Get Tracks' Audio Features 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetSeveralAudioFeatures(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetSeveralAudioFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeveralAudioFeatures`: GetSeveralAudioFeatures200Response
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetSeveralAudioFeatures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeveralAudioFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

[**GetSeveralAudioFeatures200Response**](GetSeveralAudioFeatures200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeveralTracks

> GetAnArtistsTopTracks200Response GetSeveralTracks(ctx).Ids(ids).Market(market).Execute()

Get Several Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetSeveralTracks(context.Background()).Ids(ids).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetSeveralTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeveralTracks`: GetAnArtistsTopTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetSeveralTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeveralTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 

### Return type

[**GetAnArtistsTopTracks200Response**](GetAnArtistsTopTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrack

> TrackObject GetTrack(ctx, id).Market(market).Execute()

Get Track 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    id := "11dFghVXANMlKmJXsNCbNl" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetTrack(context.Background(), id).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrack`: TrackObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetTrack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**TrackObject**](TrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSavedTracks

> PagingSavedTrackObject GetUsersSavedTracks(ctx).Market(market).Limit(limit).Offset(offset).Execute()

Get User's Saved Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    market := "ES" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetUsersSavedTracks(context.Background()).Market(market).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetUsersSavedTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersSavedTracks`: PagingSavedTrackObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetUsersSavedTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSavedTrackObject**](PagingSavedTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersTopTracks

> PagingTrackObject GetUsersTopTracks(ctx).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()

Get User's Top Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    timeRange := "medium_term" // string |  (optional) (default to "medium_term")
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.GetUsersTopTracks(context.Background()).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.GetUsersTopTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersTopTracks`: PagingTrackObject
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.GetUsersTopTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersTopTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeRange** | **string** |  | [default to &quot;medium_term&quot;]
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingTrackObject**](PagingTrackObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTracksPlaylist

> ReorderOrReplacePlaylistsTracks200Response RemoveTracksPlaylist(ctx, playlistId).RemoveTracksPlaylistRequest(removeTracksPlaylistRequest).Execute()

Remove Playlist Items 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
    removeTracksPlaylistRequest := *openapiclient.NewRemoveTracksPlaylistRequest([]openapiclient.RemoveTracksPlaylistRequestTracksInner{*openapiclient.NewRemoveTracksPlaylistRequestTracksInner()}) // RemoveTracksPlaylistRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.RemoveTracksPlaylist(context.Background(), playlistId).RemoveTracksPlaylistRequest(removeTracksPlaylistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.RemoveTracksPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTracksPlaylist`: ReorderOrReplacePlaylistsTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.RemoveTracksPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTracksPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeTracksPlaylistRequest** | [**RemoveTracksPlaylistRequest**](RemoveTracksPlaylistRequest.md) |  | 

### Return type

[**ReorderOrReplacePlaylistsTracks200Response**](ReorderOrReplacePlaylistsTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTracksUser

> RemoveTracksUser(ctx).Ids(ids).RequestBody(requestBody).Execute()

Remove User's Saved Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TracksApi.RemoveTracksUser(context.Background()).Ids(ids).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.RemoveTracksUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTracksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **requestBody** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderOrReplacePlaylistsTracks

> ReorderOrReplacePlaylistsTracks200Response ReorderOrReplacePlaylistsTracks(ctx, playlistId).Uris(uris).RequestBody(requestBody).Execute()

Update Playlist Items 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    playlistId := "3cEYpjA9oz9GiPac4AsH4n" // string | 
    uris := "uris_example" // string |  (optional)
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TracksApi.ReorderOrReplacePlaylistsTracks(context.Background(), playlistId).Uris(uris).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.ReorderOrReplacePlaylistsTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReorderOrReplacePlaylistsTracks`: ReorderOrReplacePlaylistsTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `TracksApi.ReorderOrReplacePlaylistsTracks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderOrReplacePlaylistsTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uris** | **string** |  | 
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**ReorderOrReplacePlaylistsTracks200Response**](ReorderOrReplacePlaylistsTracks200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveTracksUser

> SaveTracksUser(ctx).Ids(ids).RequestBody(requestBody).Execute()

Save Tracks for Current User 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
    ids := "7ouMYWpwJ422jRcDASZB7P,4VqPOruhp5EdPBeR92t6lQ,2takcwOaAZWiXQijPHIx7B" // string | 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TracksApi.SaveTracksUser(context.Background()).Ids(ids).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TracksApi.SaveTracksUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveTracksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **requestBody** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

