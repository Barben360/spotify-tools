# \PlayerApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToQueue**](PlayerApi.md#AddToQueue) | **Post** /me/player/queue | Add Item to Playback Queue 
[**GetAUsersAvailableDevices**](PlayerApi.md#GetAUsersAvailableDevices) | **Get** /me/player/devices | Get Available Devices 
[**GetInformationAboutTheUsersCurrentPlayback**](PlayerApi.md#GetInformationAboutTheUsersCurrentPlayback) | **Get** /me/player | Get Playback State 
[**GetQueue**](PlayerApi.md#GetQueue) | **Get** /me/player/queue | Get the User&#39;s Queue 
[**GetRecentlyPlayed**](PlayerApi.md#GetRecentlyPlayed) | **Get** /me/player/recently-played | Get Recently Played Tracks 
[**GetTheUsersCurrentlyPlayingTrack**](PlayerApi.md#GetTheUsersCurrentlyPlayingTrack) | **Get** /me/player/currently-playing | Get Currently Playing Track 
[**PauseAUsersPlayback**](PlayerApi.md#PauseAUsersPlayback) | **Put** /me/player/pause | Pause Playback 
[**SeekToPositionInCurrentlyPlayingTrack**](PlayerApi.md#SeekToPositionInCurrentlyPlayingTrack) | **Put** /me/player/seek | Seek To Position 
[**SetRepeatModeOnUsersPlayback**](PlayerApi.md#SetRepeatModeOnUsersPlayback) | **Put** /me/player/repeat | Set Repeat Mode 
[**SetVolumeForUsersPlayback**](PlayerApi.md#SetVolumeForUsersPlayback) | **Put** /me/player/volume | Set Playback Volume 
[**SkipUsersPlaybackToNextTrack**](PlayerApi.md#SkipUsersPlaybackToNextTrack) | **Post** /me/player/next | Skip To Next 
[**SkipUsersPlaybackToPreviousTrack**](PlayerApi.md#SkipUsersPlaybackToPreviousTrack) | **Post** /me/player/previous | Skip To Previous 
[**StartAUsersPlayback**](PlayerApi.md#StartAUsersPlayback) | **Put** /me/player/play | Start/Resume Playback 
[**ToggleShuffleForUsersPlayback**](PlayerApi.md#ToggleShuffleForUsersPlayback) | **Put** /me/player/shuffle | Toggle Playback Shuffle 
[**TransferAUsersPlayback**](PlayerApi.md#TransferAUsersPlayback) | **Put** /me/player | Transfer Playback 



## AddToQueue

> AddToQueue(ctx).Uri(uri).DeviceId(deviceId).Execute()

Add Item to Playback Queue 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    uri := "spotify:track:4iV5W9uYEdYUVa79Axb7Rh" // string | 
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.AddToQueue(context.Background()).Uri(uri).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.AddToQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddToQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uri** | **string** |  | 
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAUsersAvailableDevices

> DevicesObject GetAUsersAvailableDevices(ctx).Execute()

Get Available Devices 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.GetAUsersAvailableDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.GetAUsersAvailableDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAUsersAvailableDevices`: DevicesObject
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.GetAUsersAvailableDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAUsersAvailableDevicesRequest struct via the builder pattern


### Return type

[**DevicesObject**](DevicesObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInformationAboutTheUsersCurrentPlayback

> CurrentlyPlayingContextObject GetInformationAboutTheUsersCurrentPlayback(ctx).Market(market).AdditionalTypes(additionalTypes).Execute()

Get Playback State 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    market := "ES" // string |  (optional)
    additionalTypes := "additionalTypes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.GetInformationAboutTheUsersCurrentPlayback(context.Background()).Market(market).AdditionalTypes(additionalTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.GetInformationAboutTheUsersCurrentPlayback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInformationAboutTheUsersCurrentPlayback`: CurrentlyPlayingContextObject
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.GetInformationAboutTheUsersCurrentPlayback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInformationAboutTheUsersCurrentPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **additionalTypes** | **string** |  | 

### Return type

[**CurrentlyPlayingContextObject**](CurrentlyPlayingContextObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueue

> QueueObject GetQueue(ctx).Execute()

Get the User's Queue 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.GetQueue(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.GetQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQueue`: QueueObject
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.GetQueue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueRequest struct via the builder pattern


### Return type

[**QueueObject**](QueueObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentlyPlayed

> CursorPagingPlayHistoryObject GetRecentlyPlayed(ctx).Limit(limit).After(after).Before(before).Execute()

Get Recently Played Tracks 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    limit := int32(10) // int32 |  (optional) (default to 20)
    after := int32(1484811043508) // int32 |  (optional)
    before := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.GetRecentlyPlayed(context.Background()).Limit(limit).After(after).Before(before).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.GetRecentlyPlayed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecentlyPlayed`: CursorPagingPlayHistoryObject
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.GetRecentlyPlayed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentlyPlayedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **after** | **int32** |  | 
 **before** | **int32** |  | 

### Return type

[**CursorPagingPlayHistoryObject**](CursorPagingPlayHistoryObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTheUsersCurrentlyPlayingTrack

> CurrentlyPlayingObject GetTheUsersCurrentlyPlayingTrack(ctx).Market(market).AdditionalTypes(additionalTypes).Execute()

Get Currently Playing Track 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    market := "ES" // string |  (optional)
    additionalTypes := "additionalTypes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayerApi.GetTheUsersCurrentlyPlayingTrack(context.Background()).Market(market).AdditionalTypes(additionalTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.GetTheUsersCurrentlyPlayingTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTheUsersCurrentlyPlayingTrack`: CurrentlyPlayingObject
    fmt.Fprintf(os.Stdout, "Response from `PlayerApi.GetTheUsersCurrentlyPlayingTrack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTheUsersCurrentlyPlayingTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **additionalTypes** | **string** |  | 

### Return type

[**CurrentlyPlayingObject**](CurrentlyPlayingObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseAUsersPlayback

> PauseAUsersPlayback(ctx).DeviceId(deviceId).Execute()

Pause Playback 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.PauseAUsersPlayback(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.PauseAUsersPlayback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPauseAUsersPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeekToPositionInCurrentlyPlayingTrack

> SeekToPositionInCurrentlyPlayingTrack(ctx).PositionMs(positionMs).DeviceId(deviceId).Execute()

Seek To Position 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    positionMs := int32(25000) // int32 | 
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.SeekToPositionInCurrentlyPlayingTrack(context.Background()).PositionMs(positionMs).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.SeekToPositionInCurrentlyPlayingTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeekToPositionInCurrentlyPlayingTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **positionMs** | **int32** |  | 
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRepeatModeOnUsersPlayback

> SetRepeatModeOnUsersPlayback(ctx).State(state).DeviceId(deviceId).Execute()

Set Repeat Mode 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    state := "context" // string | 
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.SetRepeatModeOnUsersPlayback(context.Background()).State(state).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.SetRepeatModeOnUsersPlayback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRepeatModeOnUsersPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVolumeForUsersPlayback

> SetVolumeForUsersPlayback(ctx).VolumePercent(volumePercent).DeviceId(deviceId).Execute()

Set Playback Volume 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    volumePercent := int32(50) // int32 | 
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.SetVolumeForUsersPlayback(context.Background()).VolumePercent(volumePercent).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.SetVolumeForUsersPlayback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetVolumeForUsersPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumePercent** | **int32** |  | 
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipUsersPlaybackToNextTrack

> SkipUsersPlaybackToNextTrack(ctx).DeviceId(deviceId).Execute()

Skip To Next 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.SkipUsersPlaybackToNextTrack(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.SkipUsersPlaybackToNextTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSkipUsersPlaybackToNextTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SkipUsersPlaybackToPreviousTrack

> SkipUsersPlaybackToPreviousTrack(ctx).DeviceId(deviceId).Execute()

Skip To Previous 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.SkipUsersPlaybackToPreviousTrack(context.Background()).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.SkipUsersPlaybackToPreviousTrack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSkipUsersPlaybackToPreviousTrackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartAUsersPlayback

> StartAUsersPlayback(ctx).DeviceId(deviceId).RequestBody(requestBody).Execute()

Start/Resume Playback 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.StartAUsersPlayback(context.Background()).DeviceId(deviceId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.StartAUsersPlayback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartAUsersPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** |  | 
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


## ToggleShuffleForUsersPlayback

> ToggleShuffleForUsersPlayback(ctx).State(state).DeviceId(deviceId).Execute()

Toggle Playback Shuffle 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    state := true // bool | 
    deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.ToggleShuffleForUsersPlayback(context.Background()).State(state).DeviceId(deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.ToggleShuffleForUsersPlayback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiToggleShuffleForUsersPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **bool** |  | 
 **deviceId** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferAUsersPlayback

> TransferAUsersPlayback(ctx).RequestBody(requestBody).Execute()

Transfer Playback 



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlayerApi.TransferAUsersPlayback(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayerApi.TransferAUsersPlayback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferAUsersPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

