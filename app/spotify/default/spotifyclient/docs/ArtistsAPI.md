# \ArtistsAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCurrentUserFollows**](ArtistsAPI.md#CheckCurrentUserFollows) | **Get** /me/following/contains | Check If User Follows Artists or Users 
[**FollowArtistsUsers**](ArtistsAPI.md#FollowArtistsUsers) | **Put** /me/following | Follow Artists or Users 
[**GetAnArtist**](ArtistsAPI.md#GetAnArtist) | **Get** /artists/{id} | Get Artist 
[**GetAnArtistsAlbums**](ArtistsAPI.md#GetAnArtistsAlbums) | **Get** /artists/{id}/albums | Get Artist&#39;s Albums 
[**GetAnArtistsRelatedArtists**](ArtistsAPI.md#GetAnArtistsRelatedArtists) | **Get** /artists/{id}/related-artists | Get Artist&#39;s Related Artists 
[**GetAnArtistsTopTracks**](ArtistsAPI.md#GetAnArtistsTopTracks) | **Get** /artists/{id}/top-tracks | Get Artist&#39;s Top Tracks 
[**GetFollowed**](ArtistsAPI.md#GetFollowed) | **Get** /me/following | Get Followed Artists 
[**GetMultipleArtists**](ArtistsAPI.md#GetMultipleArtists) | **Get** /artists | Get Several Artists 
[**GetUsersTopArtists**](ArtistsAPI.md#GetUsersTopArtists) | **Get** /me/top/artists | Get User&#39;s Top Artists 
[**UnfollowArtistsUsers**](ArtistsAPI.md#UnfollowArtistsUsers) | **Delete** /me/following | Unfollow Artists or Users 



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
	resp, r, err := apiClient.ArtistsAPI.CheckCurrentUserFollows(context.Background()).Type_(type_).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.CheckCurrentUserFollows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCurrentUserFollows`: []bool
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.CheckCurrentUserFollows`: %v\n", resp)
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

> FollowArtistsUsers(ctx).Type_(type_).Ids(ids).FollowArtistsUsersRequest(followArtistsUsersRequest).Execute()

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
	followArtistsUsersRequest := *openapiclient.NewFollowArtistsUsersRequest([]string{"Ids_example"}) // FollowArtistsUsersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtistsAPI.FollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).FollowArtistsUsersRequest(followArtistsUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.FollowArtistsUsers``: %v\n", err)
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
 **followArtistsUsersRequest** | [**FollowArtistsUsersRequest**](FollowArtistsUsersRequest.md) |  | 

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
	resp, r, err := apiClient.ArtistsAPI.GetAnArtist(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetAnArtist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnArtist`: ArtistObject
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetAnArtist`: %v\n", resp)
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

> PagingArtistDiscographyAlbumObject GetAnArtistsAlbums(ctx, id).IncludeGroups(includeGroups).Market(market).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.ArtistsAPI.GetAnArtistsAlbums(context.Background(), id).IncludeGroups(includeGroups).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetAnArtistsAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnArtistsAlbums`: PagingArtistDiscographyAlbumObject
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetAnArtistsAlbums`: %v\n", resp)
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

[**PagingArtistDiscographyAlbumObject**](PagingArtistDiscographyAlbumObject.md)

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
	resp, r, err := apiClient.ArtistsAPI.GetAnArtistsRelatedArtists(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetAnArtistsRelatedArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnArtistsRelatedArtists`: GetMultipleArtists200Response
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetAnArtistsRelatedArtists`: %v\n", resp)
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
	resp, r, err := apiClient.ArtistsAPI.GetAnArtistsTopTracks(context.Background(), id).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetAnArtistsTopTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnArtistsTopTracks`: GetAnArtistsTopTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetAnArtistsTopTracks`: %v\n", resp)
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
	resp, r, err := apiClient.ArtistsAPI.GetFollowed(context.Background()).Type_(type_).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetFollowed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowed`: GetFollowed200Response
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetFollowed`: %v\n", resp)
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
	resp, r, err := apiClient.ArtistsAPI.GetMultipleArtists(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetMultipleArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultipleArtists`: GetMultipleArtists200Response
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetMultipleArtists`: %v\n", resp)
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
	resp, r, err := apiClient.ArtistsAPI.GetUsersTopArtists(context.Background()).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetUsersTopArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersTopArtists`: PagingArtistObject
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetUsersTopArtists`: %v\n", resp)
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

> UnfollowArtistsUsers(ctx).Type_(type_).Ids(ids).UnfollowArtistsUsersRequest(unfollowArtistsUsersRequest).Execute()

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
	unfollowArtistsUsersRequest := *openapiclient.NewUnfollowArtistsUsersRequest() // UnfollowArtistsUsersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtistsAPI.UnfollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).UnfollowArtistsUsersRequest(unfollowArtistsUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.UnfollowArtistsUsers``: %v\n", err)
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
 **unfollowArtistsUsersRequest** | [**UnfollowArtistsUsersRequest**](UnfollowArtistsUsersRequest.md) |  | 

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

