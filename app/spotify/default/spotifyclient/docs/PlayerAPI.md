# \PlayerAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToQueue**](PlayerAPI.md#AddToQueue) | **Post** /me/player/queue | Add Item to Playback Queue 
[**GetAUsersAvailableDevices**](PlayerAPI.md#GetAUsersAvailableDevices) | **Get** /me/player/devices | Get Available Devices 
[**GetInformationAboutTheUsersCurrentPlayback**](PlayerAPI.md#GetInformationAboutTheUsersCurrentPlayback) | **Get** /me/player | Get Playback State 
[**GetQueue**](PlayerAPI.md#GetQueue) | **Get** /me/player/queue | Get the User&#39;s Queue 
[**GetRecentlyPlayed**](PlayerAPI.md#GetRecentlyPlayed) | **Get** /me/player/recently-played | Get Recently Played Tracks 
[**GetTheUsersCurrentlyPlayingTrack**](PlayerAPI.md#GetTheUsersCurrentlyPlayingTrack) | **Get** /me/player/currently-playing | Get Currently Playing Track 
[**PauseAUsersPlayback**](PlayerAPI.md#PauseAUsersPlayback) | **Put** /me/player/pause | Pause Playback 
[**SeekToPositionInCurrentlyPlayingTrack**](PlayerAPI.md#SeekToPositionInCurrentlyPlayingTrack) | **Put** /me/player/seek | Seek To Position 
[**SetRepeatModeOnUsersPlayback**](PlayerAPI.md#SetRepeatModeOnUsersPlayback) | **Put** /me/player/repeat | Set Repeat Mode 
[**SetVolumeForUsersPlayback**](PlayerAPI.md#SetVolumeForUsersPlayback) | **Put** /me/player/volume | Set Playback Volume 
[**SkipUsersPlaybackToNextTrack**](PlayerAPI.md#SkipUsersPlaybackToNextTrack) | **Post** /me/player/next | Skip To Next 
[**SkipUsersPlaybackToPreviousTrack**](PlayerAPI.md#SkipUsersPlaybackToPreviousTrack) | **Post** /me/player/previous | Skip To Previous 
[**StartAUsersPlayback**](PlayerAPI.md#StartAUsersPlayback) | **Put** /me/player/play | Start/Resume Playback 
[**ToggleShuffleForUsersPlayback**](PlayerAPI.md#ToggleShuffleForUsersPlayback) | **Put** /me/player/shuffle | Toggle Playback Shuffle 
[**TransferAUsersPlayback**](PlayerAPI.md#TransferAUsersPlayback) | **Put** /me/player | Transfer Playback 



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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	uri := "spotify:track:4iV5W9uYEdYUVa79Axb7Rh" // string | 
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.AddToQueue(context.Background()).Uri(uri).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.AddToQueue``: %v\n", err)
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

> GetAUsersAvailableDevices200Response GetAUsersAvailableDevices(ctx).Execute()

Get Available Devices 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayerAPI.GetAUsersAvailableDevices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.GetAUsersAvailableDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAUsersAvailableDevices`: GetAUsersAvailableDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `PlayerAPI.GetAUsersAvailableDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAUsersAvailableDevicesRequest struct via the builder pattern


### Return type

[**GetAUsersAvailableDevices200Response**](GetAUsersAvailableDevices200Response.md)

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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	market := "ES" // string |  (optional)
	additionalTypes := "additionalTypes_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayerAPI.GetInformationAboutTheUsersCurrentPlayback(context.Background()).Market(market).AdditionalTypes(additionalTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.GetInformationAboutTheUsersCurrentPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInformationAboutTheUsersCurrentPlayback`: CurrentlyPlayingContextObject
	fmt.Fprintf(os.Stdout, "Response from `PlayerAPI.GetInformationAboutTheUsersCurrentPlayback`: %v\n", resp)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayerAPI.GetQueue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.GetQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueue`: QueueObject
	fmt.Fprintf(os.Stdout, "Response from `PlayerAPI.GetQueue`: %v\n", resp)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	limit := int32(10) // int32 |  (optional) (default to 20)
	after := int32(1484811043508) // int32 |  (optional)
	before := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayerAPI.GetRecentlyPlayed(context.Background()).Limit(limit).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.GetRecentlyPlayed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecentlyPlayed`: CursorPagingPlayHistoryObject
	fmt.Fprintf(os.Stdout, "Response from `PlayerAPI.GetRecentlyPlayed`: %v\n", resp)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	market := "ES" // string |  (optional)
	additionalTypes := "additionalTypes_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayerAPI.GetTheUsersCurrentlyPlayingTrack(context.Background()).Market(market).AdditionalTypes(additionalTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.GetTheUsersCurrentlyPlayingTrack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTheUsersCurrentlyPlayingTrack`: CurrentlyPlayingObject
	fmt.Fprintf(os.Stdout, "Response from `PlayerAPI.GetTheUsersCurrentlyPlayingTrack`: %v\n", resp)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.PauseAUsersPlayback(context.Background()).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.PauseAUsersPlayback``: %v\n", err)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	positionMs := int32(25000) // int32 | 
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.SeekToPositionInCurrentlyPlayingTrack(context.Background()).PositionMs(positionMs).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.SeekToPositionInCurrentlyPlayingTrack``: %v\n", err)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	state := "context" // string | 
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.SetRepeatModeOnUsersPlayback(context.Background()).State(state).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.SetRepeatModeOnUsersPlayback``: %v\n", err)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	volumePercent := int32(50) // int32 | 
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.SetVolumeForUsersPlayback(context.Background()).VolumePercent(volumePercent).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.SetVolumeForUsersPlayback``: %v\n", err)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.SkipUsersPlaybackToNextTrack(context.Background()).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.SkipUsersPlaybackToNextTrack``: %v\n", err)
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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.SkipUsersPlaybackToPreviousTrack(context.Background()).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.SkipUsersPlaybackToPreviousTrack``: %v\n", err)
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

> StartAUsersPlayback(ctx).DeviceId(deviceId).StartAUsersPlaybackRequest(startAUsersPlaybackRequest).Execute()

Start/Resume Playback 



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
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)
	startAUsersPlaybackRequest := *openapiclient.NewStartAUsersPlaybackRequest() // StartAUsersPlaybackRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.StartAUsersPlayback(context.Background()).DeviceId(deviceId).StartAUsersPlaybackRequest(startAUsersPlaybackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.StartAUsersPlayback``: %v\n", err)
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
 **startAUsersPlaybackRequest** | [**StartAUsersPlaybackRequest**](StartAUsersPlaybackRequest.md) |  | 

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
	openapiclient "github.com/Barben360/spotify-tools/app/default/spotifyclient"
)

func main() {
	state := true // bool | 
	deviceId := "0d1841b0976bae2a3a310dd74c0f3df354899bc8" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.ToggleShuffleForUsersPlayback(context.Background()).State(state).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.ToggleShuffleForUsersPlayback``: %v\n", err)
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

> TransferAUsersPlayback(ctx).TransferAUsersPlaybackRequest(transferAUsersPlaybackRequest).Execute()

Transfer Playback 



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
	transferAUsersPlaybackRequest := *openapiclient.NewTransferAUsersPlaybackRequest([]string{"DeviceIds_example"}) // TransferAUsersPlaybackRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlayerAPI.TransferAUsersPlayback(context.Background()).TransferAUsersPlaybackRequest(transferAUsersPlaybackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayerAPI.TransferAUsersPlayback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferAUsersPlaybackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferAUsersPlaybackRequest** | [**TransferAUsersPlaybackRequest**](TransferAUsersPlaybackRequest.md) |  | 

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

