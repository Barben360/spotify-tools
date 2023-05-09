# \PlaylistsApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTracksToPlaylist**](PlaylistsApi.md#AddTracksToPlaylist) | **Post** /playlists/{playlist_id}/tracks | Add Items to Playlist 
[**ChangePlaylistDetails**](PlaylistsApi.md#ChangePlaylistDetails) | **Put** /playlists/{playlist_id} | Change Playlist Details 
[**CheckIfUserFollowsPlaylist**](PlaylistsApi.md#CheckIfUserFollowsPlaylist) | **Get** /playlists/{playlist_id}/followers/contains | Check if Users Follow Playlist 
[**CreatePlaylist**](PlaylistsApi.md#CreatePlaylist) | **Post** /users/{user_id}/playlists | Create Playlist 
[**FollowPlaylist**](PlaylistsApi.md#FollowPlaylist) | **Put** /playlists/{playlist_id}/followers | Follow Playlist 
[**GetACategoriesPlaylists**](PlaylistsApi.md#GetACategoriesPlaylists) | **Get** /browse/categories/{category_id}/playlists | Get Category&#39;s Playlists 
[**GetAListOfCurrentUsersPlaylists**](PlaylistsApi.md#GetAListOfCurrentUsersPlaylists) | **Get** /me/playlists | Get Current User&#39;s Playlists 
[**GetFeaturedPlaylists**](PlaylistsApi.md#GetFeaturedPlaylists) | **Get** /browse/featured-playlists | Get Featured Playlists 
[**GetListUsersPlaylists**](PlaylistsApi.md#GetListUsersPlaylists) | **Get** /users/{user_id}/playlists | Get User&#39;s Playlists 
[**GetPlaylist**](PlaylistsApi.md#GetPlaylist) | **Get** /playlists/{playlist_id} | Get Playlist 
[**GetPlaylistCover**](PlaylistsApi.md#GetPlaylistCover) | **Get** /playlists/{playlist_id}/images | Get Playlist Cover Image 
[**GetPlaylistsTracks**](PlaylistsApi.md#GetPlaylistsTracks) | **Get** /playlists/{playlist_id}/tracks | Get Playlist Items 
[**RemoveTracksPlaylist**](PlaylistsApi.md#RemoveTracksPlaylist) | **Delete** /playlists/{playlist_id}/tracks | Remove Playlist Items 
[**ReorderOrReplacePlaylistsTracks**](PlaylistsApi.md#ReorderOrReplacePlaylistsTracks) | **Put** /playlists/{playlist_id}/tracks | Update Playlist Items 
[**UnfollowPlaylist**](PlaylistsApi.md#UnfollowPlaylist) | **Delete** /playlists/{playlist_id}/followers | Unfollow Playlist 
[**UploadCustomPlaylistCover**](PlaylistsApi.md#UploadCustomPlaylistCover) | **Put** /playlists/{playlist_id}/images | Add Custom Playlist Cover Image 



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
    resp, r, err := apiClient.PlaylistsApi.AddTracksToPlaylist(context.Background(), playlistId).Position(position).Uris(uris).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.AddTracksToPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTracksToPlaylist`: ReorderOrReplacePlaylistsTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.AddTracksToPlaylist`: %v\n", resp)
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


## ChangePlaylistDetails

> ChangePlaylistDetails(ctx, playlistId).RequestBody(requestBody).Execute()

Change Playlist Details 



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlaylistsApi.ChangePlaylistDetails(context.Background(), playlistId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.ChangePlaylistDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangePlaylistDetailsRequest struct via the builder pattern


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


## CheckIfUserFollowsPlaylist

> []bool CheckIfUserFollowsPlaylist(ctx, playlistId).Ids(ids).Execute()

Check if Users Follow Playlist 



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
    ids := "jmperezperez,thelinmichael,wizzler" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.CheckIfUserFollowsPlaylist(context.Background(), playlistId).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.CheckIfUserFollowsPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckIfUserFollowsPlaylist`: []bool
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.CheckIfUserFollowsPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckIfUserFollowsPlaylistRequest struct via the builder pattern


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


## CreatePlaylist

> PlaylistObject CreatePlaylist(ctx, userId).RequestBody(requestBody).Execute()

Create Playlist 



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
    userId := "smedjan" // string | 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.CreatePlaylist(context.Background(), userId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.CreatePlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlaylist`: PlaylistObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.CreatePlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**PlaylistObject**](PlaylistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowPlaylist

> FollowPlaylist(ctx, playlistId).RequestBody(requestBody).Execute()

Follow Playlist 



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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlaylistsApi.FollowPlaylist(context.Background(), playlistId).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.FollowPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFollowPlaylistRequest struct via the builder pattern


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


## GetACategoriesPlaylists

> PagingFeaturedPlaylistObject GetACategoriesPlaylists(ctx, categoryId).Country(country).Limit(limit).Offset(offset).Execute()

Get Category's Playlists 



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
    categoryId := "dinner" // string | 
    country := "SE" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetACategoriesPlaylists(context.Background(), categoryId).Country(country).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetACategoriesPlaylists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACategoriesPlaylists`: PagingFeaturedPlaylistObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.GetACategoriesPlaylists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetACategoriesPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingFeaturedPlaylistObject**](PagingFeaturedPlaylistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAListOfCurrentUsersPlaylists

> PagingPlaylistObject GetAListOfCurrentUsersPlaylists(ctx).Limit(limit).Offset(offset).Execute()

Get Current User's Playlists 



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
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetAListOfCurrentUsersPlaylists(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetAListOfCurrentUsersPlaylists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAListOfCurrentUsersPlaylists`: PagingPlaylistObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.GetAListOfCurrentUsersPlaylists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAListOfCurrentUsersPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingPlaylistObject**](PagingPlaylistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeaturedPlaylists

> PagingFeaturedPlaylistObject GetFeaturedPlaylists(ctx).Country(country).Locale(locale).Timestamp(timestamp).Limit(limit).Offset(offset).Execute()

Get Featured Playlists 



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
    country := "SE" // string |  (optional)
    locale := "sv_SE" // string |  (optional)
    timestamp := "2014-10-23T09:00:00" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetFeaturedPlaylists(context.Background()).Country(country).Locale(locale).Timestamp(timestamp).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetFeaturedPlaylists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeaturedPlaylists`: PagingFeaturedPlaylistObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.GetFeaturedPlaylists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturedPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 
 **locale** | **string** |  | 
 **timestamp** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingFeaturedPlaylistObject**](PagingFeaturedPlaylistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListUsersPlaylists

> PagingPlaylistObject GetListUsersPlaylists(ctx, userId).Limit(limit).Offset(offset).Execute()

Get User's Playlists 



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
    userId := "smedjan" // string | 
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetListUsersPlaylists(context.Background(), userId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetListUsersPlaylists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListUsersPlaylists`: PagingPlaylistObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.GetListUsersPlaylists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListUsersPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingPlaylistObject**](PagingPlaylistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylist

> PlaylistObject GetPlaylist(ctx, playlistId).Market(market).Fields(fields).AdditionalTypes(additionalTypes).Execute()

Get Playlist 



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
    additionalTypes := "additionalTypes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetPlaylist(context.Background(), playlistId).Market(market).Fields(fields).AdditionalTypes(additionalTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlaylist`: PlaylistObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.GetPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **fields** | **string** |  | 
 **additionalTypes** | **string** |  | 

### Return type

[**PlaylistObject**](PlaylistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistCover

> []ImageObject GetPlaylistCover(ctx, playlistId).Execute()

Get Playlist Cover Image 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaylistsApi.GetPlaylistCover(context.Background(), playlistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetPlaylistCover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlaylistCover`: []ImageObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.GetPlaylistCover`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistCoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ImageObject**](ImageObject.md)

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
    resp, r, err := apiClient.PlaylistsApi.GetPlaylistsTracks(context.Background(), playlistId).Market(market).Fields(fields).Limit(limit).Offset(offset).AdditionalTypes(additionalTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.GetPlaylistsTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlaylistsTracks`: PagingPlaylistTrackObject
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.GetPlaylistsTracks`: %v\n", resp)
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
    resp, r, err := apiClient.PlaylistsApi.RemoveTracksPlaylist(context.Background(), playlistId).RemoveTracksPlaylistRequest(removeTracksPlaylistRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.RemoveTracksPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTracksPlaylist`: ReorderOrReplacePlaylistsTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.RemoveTracksPlaylist`: %v\n", resp)
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
    resp, r, err := apiClient.PlaylistsApi.ReorderOrReplacePlaylistsTracks(context.Background(), playlistId).Uris(uris).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.ReorderOrReplacePlaylistsTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReorderOrReplacePlaylistsTracks`: ReorderOrReplacePlaylistsTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `PlaylistsApi.ReorderOrReplacePlaylistsTracks`: %v\n", resp)
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


## UnfollowPlaylist

> UnfollowPlaylist(ctx, playlistId).Execute()

Unfollow Playlist 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlaylistsApi.UnfollowPlaylist(context.Background(), playlistId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.UnfollowPlaylist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UploadCustomPlaylistCover

> UploadCustomPlaylistCover(ctx, playlistId).Body(body).Execute()

Add Custom Playlist Cover Image 



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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PlaylistsApi.UploadCustomPlaylistCover(context.Background(), playlistId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsApi.UploadCustomPlaylistCover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCustomPlaylistCoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: image/jpeg
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

