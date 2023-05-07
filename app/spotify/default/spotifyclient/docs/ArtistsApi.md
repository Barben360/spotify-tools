# \ArtistsApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCurrentUserFollows**](ArtistsApi.md#CheckCurrentUserFollows) | **Get** /me/following/contains | Check If User Follows Artists or Users 
[**FollowArtistsUsers**](ArtistsApi.md#FollowArtistsUsers) | **Put** /me/following | Follow Artists or Users 
[**GetAnArtist**](ArtistsApi.md#GetAnArtist) | **Get** /artists/{id} | Get Artist 
[**GetAnArtistsAlbums**](ArtistsApi.md#GetAnArtistsAlbums) | **Get** /artists/{id}/albums | Get Artist&#39;s Albums 
[**GetAnArtistsRelatedArtists**](ArtistsApi.md#GetAnArtistsRelatedArtists) | **Get** /artists/{id}/related-artists | Get Artist&#39;s Related Artists 
[**GetAnArtistsTopTracks**](ArtistsApi.md#GetAnArtistsTopTracks) | **Get** /artists/{id}/top-tracks | Get Artist&#39;s Top Tracks 
[**GetFollowed**](ArtistsApi.md#GetFollowed) | **Get** /me/following | Get Followed Artists 
[**GetMultipleArtists**](ArtistsApi.md#GetMultipleArtists) | **Get** /artists | Get Several Artists 
[**GetUsersTopArtists**](ArtistsApi.md#GetUsersTopArtists) | **Get** /me/top/artists | Get User&#39;s Top Artists 
[**UnfollowArtistsUsers**](ArtistsApi.md#UnfollowArtistsUsers) | **Delete** /me/following | Unfollow Artists or Users 



## CheckCurrentUserFollows

> []bool CheckCurrentUserFollows(ctx).Type_(type_).Ids(ids).Execute()

Check If User Follows Artists or Users 



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
    type_ := "artist" // string | 
    ids := "2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistsApi.CheckCurrentUserFollows(context.Background()).Type_(type_).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.CheckCurrentUserFollows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckCurrentUserFollows`: []bool
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.CheckCurrentUserFollows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckCurrentUserFollowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
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


## FollowArtistsUsers

> FollowArtistsUsers(ctx).Type_(type_).Ids(ids).RequestBody(requestBody).Execute()

Follow Artists or Users 



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
    type_ := "artist" // string | 
    ids := "2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6" // string | 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtistsApi.FollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.FollowArtistsUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFollowArtistsUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
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


## GetAnArtist

> ArtistObject GetAnArtist(ctx, id).Execute()

Get Artist 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistsApi.GetAnArtist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.GetAnArtist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnArtist`: ArtistObject
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.GetAnArtist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnArtistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArtistObject**](ArtistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnArtistsAlbums

> PagingSimplifiedAlbumObject GetAnArtistsAlbums(ctx, id).IncludeGroups(includeGroups).Market(market).Limit(limit).Offset(offset).Execute()

Get Artist's Albums 



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
    includeGroups := "single,appears_on" // string |  (optional)
    market := "ES" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistsApi.GetAnArtistsAlbums(context.Background(), id).IncludeGroups(includeGroups).Market(market).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.GetAnArtistsAlbums``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnArtistsAlbums`: PagingSimplifiedAlbumObject
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.GetAnArtistsAlbums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnArtistsAlbumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeGroups** | **string** |  | 
 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSimplifiedAlbumObject**](PagingSimplifiedAlbumObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnArtistsRelatedArtists

> GetMultipleArtists200Response GetAnArtistsRelatedArtists(ctx, id).Execute()

Get Artist's Related Artists 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistsApi.GetAnArtistsRelatedArtists(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.GetAnArtistsRelatedArtists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnArtistsRelatedArtists`: GetMultipleArtists200Response
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.GetAnArtistsRelatedArtists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnArtistsRelatedArtistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMultipleArtists200Response**](GetMultipleArtists200Response.md)

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
    resp, r, err := apiClient.ArtistsApi.GetAnArtistsTopTracks(context.Background(), id).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.GetAnArtistsTopTracks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnArtistsTopTracks`: GetAnArtistsTopTracks200Response
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.GetAnArtistsTopTracks`: %v\n", resp)
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


## GetFollowed

> GetFollowed200Response GetFollowed(ctx).Type_(type_).After(after).Limit(limit).Execute()

Get Followed Artists 



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
    type_ := "artist" // string | 
    after := "0I2XqVXqHScXjHhk6AYYRe" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistsApi.GetFollowed(context.Background()).Type_(type_).After(after).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.GetFollowed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFollowed`: GetFollowed200Response
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.GetFollowed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **after** | **string** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**GetFollowed200Response**](GetFollowed200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleArtists

> GetMultipleArtists200Response GetMultipleArtists(ctx).Ids(ids).Execute()

Get Several Artists 



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
    ids := "2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtistsApi.GetMultipleArtists(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.GetMultipleArtists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleArtists`: GetMultipleArtists200Response
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.GetMultipleArtists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleArtistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

### Return type

[**GetMultipleArtists200Response**](GetMultipleArtists200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersTopArtists

> PagingArtistObject GetUsersTopArtists(ctx).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()

Get User's Top Artists 



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
    resp, r, err := apiClient.ArtistsApi.GetUsersTopArtists(context.Background()).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.GetUsersTopArtists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersTopArtists`: PagingArtistObject
    fmt.Fprintf(os.Stdout, "Response from `ArtistsApi.GetUsersTopArtists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersTopArtistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeRange** | **string** |  | [default to &quot;medium_term&quot;]
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingArtistObject**](PagingArtistObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnfollowArtistsUsers

> UnfollowArtistsUsers(ctx).Type_(type_).Ids(ids).RequestBody(requestBody).Execute()

Unfollow Artists or Users 



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
    type_ := "artist" // string | 
    ids := "2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6" // string | 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtistsApi.UnfollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtistsApi.UnfollowArtistsUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowArtistsUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
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

