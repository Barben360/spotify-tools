# \MarketsApi

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableMarkets**](MarketsApi.md#GetAvailableMarkets) | **Get** /markets | Get Available Markets 



## GetAvailableMarkets

> GetAvailableMarkets200Response GetAvailableMarkets(ctx).Execute()

Get Available Markets 



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
    resp, r, err := apiClient.MarketsApi.GetAvailableMarkets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketsApi.GetAvailableMarkets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableMarkets`: GetAvailableMarkets200Response
    fmt.Fprintf(os.Stdout, "Response from `MarketsApi.GetAvailableMarkets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableMarketsRequest struct via the builder pattern


### Return type

[**GetAvailableMarkets200Response**](GetAvailableMarkets200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

