# \ShowsApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUsersSavedShows**](ShowsApi.md#CheckUsersSavedShows) | **Get** /me/shows/contains | Check User&#39;s Saved Shows 
[**GetAShow**](ShowsApi.md#GetAShow) | **Get** /shows/{id} | Get Show 
[**GetAShowsEpisodes**](ShowsApi.md#GetAShowsEpisodes) | **Get** /shows/{id}/episodes | Get Show Episodes 
[**GetMultipleShows**](ShowsApi.md#GetMultipleShows) | **Get** /shows | Get Several Shows 
[**GetUsersSavedShows**](ShowsApi.md#GetUsersSavedShows) | **Get** /me/shows | Get User&#39;s Saved Shows 
[**RemoveShowsUser**](ShowsApi.md#RemoveShowsUser) | **Delete** /me/shows | Remove User&#39;s Saved Shows 
[**SaveShowsUser**](ShowsApi.md#SaveShowsUser) | **Put** /me/shows | Save Shows for Current User 



## CheckUsersSavedShows

> []bool CheckUsersSavedShows(ctx).Ids(ids).Execute()

Check User's Saved Shows 



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
    ids := "5CfCWKI5pZ28U0uOzXkDHe,5as3aKmN2k11yfDDDSrvaZ" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowsApi.CheckUsersSavedShows(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowsApi.CheckUsersSavedShows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckUsersSavedShows`: []bool
    fmt.Fprintf(os.Stdout, "Response from `ShowsApi.CheckUsersSavedShows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedShowsRequest struct via the builder pattern


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


## GetAShow

> ShowObject GetAShow(ctx, id).Market(market).Execute()

Get Show 



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowsApi.GetAShow(context.Background(), id).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowsApi.GetAShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAShow`: ShowObject
    fmt.Fprintf(os.Stdout, "Response from `ShowsApi.GetAShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**ShowObject**](ShowObject.md)

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
    resp, r, err := apiClient.ShowsApi.GetAShowsEpisodes(context.Background(), id).Market(market).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowsApi.GetAShowsEpisodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAShowsEpisodes`: PagingSimplifiedEpisodeObject
    fmt.Fprintf(os.Stdout, "Response from `ShowsApi.GetAShowsEpisodes`: %v\n", resp)
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


## GetMultipleShows

> GetMultipleShows200Response GetMultipleShows(ctx).Ids(ids).Market(market).Execute()

Get Several Shows 



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
    ids := "5CfCWKI5pZ28U0uOzXkDHe,5as3aKmN2k11yfDDDSrvaZ" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowsApi.GetMultipleShows(context.Background()).Ids(ids).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowsApi.GetMultipleShows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleShows`: GetMultipleShows200Response
    fmt.Fprintf(os.Stdout, "Response from `ShowsApi.GetMultipleShows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleShowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 

### Return type

[**GetMultipleShows200Response**](GetMultipleShows200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSavedShows

> PagingSavedShowObject GetUsersSavedShows(ctx).Limit(limit).Offset(offset).Execute()

Get User's Saved Shows 



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
    resp, r, err := apiClient.ShowsApi.GetUsersSavedShows(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowsApi.GetUsersSavedShows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersSavedShows`: PagingSavedShowObject
    fmt.Fprintf(os.Stdout, "Response from `ShowsApi.GetUsersSavedShows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedShowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSavedShowObject**](PagingSavedShowObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveShowsUser

> RemoveShowsUser(ctx).Ids(ids).Market(market).SaveShowsUserRequest(saveShowsUserRequest).Execute()

Remove User's Saved Shows 



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
    ids := "5CfCWKI5pZ28U0uOzXkDHe,5as3aKmN2k11yfDDDSrvaZ" // string | 
    market := "ES" // string |  (optional)
    saveShowsUserRequest := *openapiclient.NewSaveShowsUserRequest() // SaveShowsUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowsApi.RemoveShowsUser(context.Background()).Ids(ids).Market(market).SaveShowsUserRequest(saveShowsUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowsApi.RemoveShowsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveShowsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 
 **saveShowsUserRequest** | [**SaveShowsUserRequest**](SaveShowsUserRequest.md) |  | 

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


## SaveShowsUser

> SaveShowsUser(ctx).Ids(ids).SaveShowsUserRequest(saveShowsUserRequest).Execute()

Save Shows for Current User 



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
    ids := "5CfCWKI5pZ28U0uOzXkDHe,5as3aKmN2k11yfDDDSrvaZ" // string | 
    saveShowsUserRequest := *openapiclient.NewSaveShowsUserRequest() // SaveShowsUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowsApi.SaveShowsUser(context.Background()).Ids(ids).SaveShowsUserRequest(saveShowsUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowsApi.SaveShowsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveShowsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveShowsUserRequest** | [**SaveShowsUserRequest**](SaveShowsUserRequest.md) |  | 

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

