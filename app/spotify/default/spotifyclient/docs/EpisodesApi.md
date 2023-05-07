# \EpisodesApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUsersSavedEpisodes**](EpisodesApi.md#CheckUsersSavedEpisodes) | **Get** /me/episodes/contains | Check User&#39;s Saved Episodes 
[**GetAShowsEpisodes**](EpisodesApi.md#GetAShowsEpisodes) | **Get** /shows/{id}/episodes | Get Show Episodes 
[**GetAnEpisode**](EpisodesApi.md#GetAnEpisode) | **Get** /episodes/{id} | Get Episode 
[**GetMultipleEpisodes**](EpisodesApi.md#GetMultipleEpisodes) | **Get** /episodes | Get Several Episodes 
[**GetUsersSavedEpisodes**](EpisodesApi.md#GetUsersSavedEpisodes) | **Get** /me/episodes | Get User&#39;s Saved Episodes 
[**RemoveEpisodesUser**](EpisodesApi.md#RemoveEpisodesUser) | **Delete** /me/episodes | Remove User&#39;s Saved Episodes 
[**SaveEpisodesUser**](EpisodesApi.md#SaveEpisodesUser) | **Put** /me/episodes | Save Episodes for Current User 



## CheckUsersSavedEpisodes

> []bool CheckUsersSavedEpisodes(ctx).Ids(ids).Execute()

Check User's Saved Episodes 



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
    ids := "77o6BIVlYM3msb4MMIL1jH,0Q86acNRm6V9GYx55SXKwf" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodesApi.CheckUsersSavedEpisodes(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodesApi.CheckUsersSavedEpisodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckUsersSavedEpisodes`: []bool
    fmt.Fprintf(os.Stdout, "Response from `EpisodesApi.CheckUsersSavedEpisodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedEpisodesRequest struct via the builder pattern


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


## GetAShowsEpisodes

> PagingSimplifiedEpisodeObject GetAShowsEpisodes(ctx, id).Market(market).Limit(limit).Offset(offset).Execute()

Get Show Episodes 



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
    id := "38bS44xjbVVZ3No3ByF1dJ" // string | 
    market := "ES" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodesApi.GetAShowsEpisodes(context.Background(), id).Market(market).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodesApi.GetAShowsEpisodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAShowsEpisodes`: PagingSimplifiedEpisodeObject
    fmt.Fprintf(os.Stdout, "Response from `EpisodesApi.GetAShowsEpisodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAShowsEpisodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSimplifiedEpisodeObject**](PagingSimplifiedEpisodeObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnEpisode

> EpisodeObject GetAnEpisode(ctx, id).Market(market).Execute()

Get Episode 



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
    id := "512ojhOuo1ktJprKbVcKyQ" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodesApi.GetAnEpisode(context.Background(), id).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodesApi.GetAnEpisode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnEpisode`: EpisodeObject
    fmt.Fprintf(os.Stdout, "Response from `EpisodesApi.GetAnEpisode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnEpisodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**EpisodeObject**](EpisodeObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleEpisodes

> GetMultipleEpisodes200Response GetMultipleEpisodes(ctx).Ids(ids).Market(market).Execute()

Get Several Episodes 



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
    ids := "77o6BIVlYM3msb4MMIL1jH,0Q86acNRm6V9GYx55SXKwf" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EpisodesApi.GetMultipleEpisodes(context.Background()).Ids(ids).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodesApi.GetMultipleEpisodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleEpisodes`: GetMultipleEpisodes200Response
    fmt.Fprintf(os.Stdout, "Response from `EpisodesApi.GetMultipleEpisodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleEpisodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 

### Return type

[**GetMultipleEpisodes200Response**](GetMultipleEpisodes200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSavedEpisodes

> PagingSavedEpisodeObject GetUsersSavedEpisodes(ctx).Market(market).Limit(limit).Offset(offset).Execute()

Get User's Saved Episodes 



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
    resp, r, err := apiClient.EpisodesApi.GetUsersSavedEpisodes(context.Background()).Market(market).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodesApi.GetUsersSavedEpisodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersSavedEpisodes`: PagingSavedEpisodeObject
    fmt.Fprintf(os.Stdout, "Response from `EpisodesApi.GetUsersSavedEpisodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedEpisodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSavedEpisodeObject**](PagingSavedEpisodeObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveEpisodesUser

> RemoveEpisodesUser(ctx).Ids(ids).RequestBody(requestBody).Execute()

Remove User's Saved Episodes 



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
    r, err := apiClient.EpisodesApi.RemoveEpisodesUser(context.Background()).Ids(ids).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodesApi.RemoveEpisodesUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveEpisodesUserRequest struct via the builder pattern


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


## SaveEpisodesUser

> SaveEpisodesUser(ctx).Ids(ids).RequestBody(requestBody).Execute()

Save Episodes for Current User 



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
    ids := "77o6BIVlYM3msb4MMIL1jH,0Q86acNRm6V9GYx55SXKwf" // string | 
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EpisodesApi.SaveEpisodesUser(context.Background()).Ids(ids).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EpisodesApi.SaveEpisodesUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveEpisodesUserRequest struct via the builder pattern


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

