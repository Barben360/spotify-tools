# \GenresAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecommendationGenres**](GenresAPI.md#GetRecommendationGenres) | **Get** /recommendations/available-genre-seeds | Get Available Genre Seeds 



## GetRecommendationGenres

> GetRecommendationGenres200Response GetRecommendationGenres(ctx).Execute()

Get Available Genre Seeds 



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
	resp, r, err := apiClient.GenresAPI.GetRecommendationGenres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenresAPI.GetRecommendationGenres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecommendationGenres`: GetRecommendationGenres200Response
	fmt.Fprintf(os.Stdout, "Response from `GenresAPI.GetRecommendationGenres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendationGenresRequest struct via the builder pattern


### Return type

[**GetRecommendationGenres200Response**](GetRecommendationGenres200Response.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

