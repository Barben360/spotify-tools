# \ChaptersApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAChapter**](ChaptersApi.md#GetAChapter) | **Get** /chapters/{id} | Get a Chapter 
[**GetAudiobookChapters**](ChaptersApi.md#GetAudiobookChapters) | **Get** /audiobooks/{id}/chapters | Get Audiobook Chapters 
[**GetSeveralChapters**](ChaptersApi.md#GetSeveralChapters) | **Get** /chapters | Get Several Chapters 



## GetAChapter

> ChapterObject GetAChapter(ctx, id).Market(market).Execute()

Get a Chapter 



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
    id := "0D5wENdkdwbqlrHoaJ9g29" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChaptersApi.GetAChapter(context.Background(), id).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChaptersApi.GetAChapter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAChapter`: ChapterObject
    fmt.Fprintf(os.Stdout, "Response from `ChaptersApi.GetAChapter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAChapterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 

### Return type

[**ChapterObject**](ChapterObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiobookChapters

> PagingSimplifiedChapterObject GetAudiobookChapters(ctx, id).Market(market).Limit(limit).Offset(offset).Execute()

Get Audiobook Chapters 



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
    id := "7iHfbu1YPACw6oZPAFJtqe" // string | 
    market := "ES" // string |  (optional)
    limit := int32(10) // int32 |  (optional) (default to 20)
    offset := int32(5) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChaptersApi.GetAudiobookChapters(context.Background(), id).Market(market).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChaptersApi.GetAudiobookChapters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudiobookChapters`: PagingSimplifiedChapterObject
    fmt.Fprintf(os.Stdout, "Response from `ChaptersApi.GetAudiobookChapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiobookChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSimplifiedChapterObject**](PagingSimplifiedChapterObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeveralChapters

> GetSeveralChapters200Response GetSeveralChapters(ctx).Ids(ids).Market(market).Execute()

Get Several Chapters 



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
    ids := "0IsXVP0JmcB2adSE338GkK,3ZXb8FKZGU0EHALYX6uCzU,0D5wENdkdwbqlrHoaJ9g29" // string | 
    market := "ES" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChaptersApi.GetSeveralChapters(context.Background()).Ids(ids).Market(market).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChaptersApi.GetSeveralChapters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSeveralChapters`: GetSeveralChapters200Response
    fmt.Fprintf(os.Stdout, "Response from `ChaptersApi.GetSeveralChapters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeveralChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **market** | **string** |  | 

### Return type

[**GetSeveralChapters200Response**](GetSeveralChapters200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

