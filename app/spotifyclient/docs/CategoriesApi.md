# \CategoriesApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetACategoriesPlaylists**](CategoriesApi.md#GetACategoriesPlaylists) | **Get** /browse/categories/{category_id}/playlists | Get Category&#39;s Playlists 
[**GetACategory**](CategoriesApi.md#GetACategory) | **Get** /browse/categories/{category_id} | Get Single Browse Category 
[**GetCategories**](CategoriesApi.md#GetCategories) | **Get** /browse/categories | Get Several Browse Categories 



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
    openapiclient "github.com/Barben360/spotify-playlist-filters/app/spotifyclient"
)

func main() {
    categoryId := "dinner" // string | 
    country := "SE" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesApi.GetACategoriesPlaylists(context.Background(), categoryId).Country(country).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetACategoriesPlaylists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACategoriesPlaylists`: PagingFeaturedPlaylistObject
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetACategoriesPlaylists`: %v\n", resp)
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


## GetACategory

> CategoryObject GetACategory(ctx, categoryId).Country(country).Locale(locale).Execute()

Get Single Browse Category 



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
    categoryId := "dinner" // string | 
    country := "SE" // string |  (optional)
    locale := "sv_SE" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesApi.GetACategory(context.Background(), categoryId).Country(country).Locale(locale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetACategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACategory`: CategoryObject
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetACategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetACategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** |  | 
 **locale** | **string** |  | 

### Return type

[**CategoryObject**](CategoryObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategories

> GetCategories200Response GetCategories(ctx).Country(country).Locale(locale).Limit(limit).Offset(offset).Execute()

Get Several Browse Categories 



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
    country := "SE" // string |  (optional)
    locale := "sv_SE" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CategoriesApi.GetCategories(context.Background()).Country(country).Locale(locale).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategories`: GetCategories200Response
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 
 **locale** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**GetCategories200Response**](GetCategories200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

