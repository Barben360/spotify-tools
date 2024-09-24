# \PlaylistsAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTracksToPlaylist**](PlaylistsAPI.md#AddTracksToPlaylist) | **Post** /playlists/{playlist_id}/tracks | Add Items to Playlist 
[**ChangePlaylistDetails**](PlaylistsAPI.md#ChangePlaylistDetails) | **Put** /playlists/{playlist_id} | Change Playlist Details 
[**CheckIfUserFollowsPlaylist**](PlaylistsAPI.md#CheckIfUserFollowsPlaylist) | **Get** /playlists/{playlist_id}/followers/contains | Check if Current User Follows Playlist 
[**CreatePlaylist**](PlaylistsAPI.md#CreatePlaylist) | **Post** /users/{user_id}/playlists | Create Playlist 
[**FollowPlaylist**](PlaylistsAPI.md#FollowPlaylist) | **Put** /playlists/{playlist_id}/followers | Follow Playlist 
[**GetACategoriesPlaylists**](PlaylistsAPI.md#GetACategoriesPlaylists) | **Get** /browse/categories/{category_id}/playlists | Get Category&#39;s Playlists 
[**GetAListOfCurrentUsersPlaylists**](PlaylistsAPI.md#GetAListOfCurrentUsersPlaylists) | **Get** /me/playlists | Get Current User&#39;s Playlists 
[**GetFeaturedPlaylists**](PlaylistsAPI.md#GetFeaturedPlaylists) | **Get** /browse/featured-playlists | Get Featured Playlists 
[**GetListUsersPlaylists**](PlaylistsAPI.md#GetListUsersPlaylists) | **Get** /users/{user_id}/playlists | Get User&#39;s Playlists 
[**GetPlaylist**](PlaylistsAPI.md#GetPlaylist) | **Get** /playlists/{playlist_id} | Get Playlist 
[**GetPlaylistCover**](PlaylistsAPI.md#GetPlaylistCover) | **Get** /playlists/{playlist_id}/images | Get Playlist Cover Image 
[**GetPlaylistsTracks**](PlaylistsAPI.md#GetPlaylistsTracks) | **Get** /playlists/{playlist_id}/tracks | Get Playlist Items 
[**RemoveTracksPlaylist**](PlaylistsAPI.md#RemoveTracksPlaylist) | **Delete** /playlists/{playlist_id}/tracks | Remove Playlist Items 
[**ReorderOrReplacePlaylistsTracks**](PlaylistsAPI.md#ReorderOrReplacePlaylistsTracks) | **Put** /playlists/{playlist_id}/tracks | Update Playlist Items 
[**UnfollowPlaylist**](PlaylistsAPI.md#UnfollowPlaylist) | **Delete** /playlists/{playlist_id}/followers | Unfollow Playlist 
[**UploadCustomPlaylistCover**](PlaylistsAPI.md#UploadCustomPlaylistCover) | **Put** /playlists/{playlist_id}/images | Add Custom Playlist Cover Image 



## AddTracksToPlaylist

> ReorderOrReplacePlaylistsTracks200Response AddTracksToPlaylist(ctx, playlistId).Position(position).Uris(uris).AddTracksToPlaylistRequest(addTracksToPlaylistRequest).Execute()

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
	addTracksToPlaylistRequest := *openapiclient.NewAddTracksToPlaylistRequest() // AddTracksToPlaylistRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.AddTracksToPlaylist(context.Background(), playlistId).Position(position).Uris(uris).AddTracksToPlaylistRequest(addTracksToPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.AddTracksToPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTracksToPlaylist`: ReorderOrReplacePlaylistsTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.AddTracksToPlaylist`: %v\n", resp)
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
 **addTracksToPlaylistRequest** | [**AddTracksToPlaylistRequest**](AddTracksToPlaylistRequest.md) |  | 

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

> ChangePlaylistDetails(ctx, playlistId).ChangePlaylistDetailsRequest(changePlaylistDetailsRequest).Execute()

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
	changePlaylistDetailsRequest := *openapiclient.NewChangePlaylistDetailsRequest() // ChangePlaylistDetailsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.ChangePlaylistDetails(context.Background(), playlistId).ChangePlaylistDetailsRequest(changePlaylistDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.ChangePlaylistDetails``: %v\n", err)
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

 **changePlaylistDetailsRequest** | [**ChangePlaylistDetailsRequest**](ChangePlaylistDetailsRequest.md) |  | 

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

Check if Current User Follows Playlist 



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
	ids := "jmperezperez" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.CheckIfUserFollowsPlaylist(context.Background(), playlistId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.CheckIfUserFollowsPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckIfUserFollowsPlaylist`: []bool
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.CheckIfUserFollowsPlaylist`: %v\n", resp)
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

> PlaylistObject CreatePlaylist(ctx, userId).CreatePlaylistRequest(createPlaylistRequest).Execute()

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
	createPlaylistRequest := *openapiclient.NewCreatePlaylistRequest("Name_example") // CreatePlaylistRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.CreatePlaylist(context.Background(), userId).CreatePlaylistRequest(createPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.CreatePlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlaylist`: PlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.CreatePlaylist`: %v\n", resp)
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

 **createPlaylistRequest** | [**CreatePlaylistRequest**](CreatePlaylistRequest.md) |  | 

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

> FollowPlaylist(ctx, playlistId).FollowPlaylistRequest(followPlaylistRequest).Execute()

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
	followPlaylistRequest := *openapiclient.NewFollowPlaylistRequest() // FollowPlaylistRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.FollowPlaylist(context.Background(), playlistId).FollowPlaylistRequest(followPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.FollowPlaylist``: %v\n", err)
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

 **followPlaylistRequest** | [**FollowPlaylistRequest**](FollowPlaylistRequest.md) |  | 

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

> PagingFeaturedPlaylistObject GetACategoriesPlaylists(ctx, categoryId).Limit(limit).Offset(offset).Execute()

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
	limit := int32(10) // int32 |  (optional) (default to 20)
	offset := int32(5) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.GetACategoriesPlaylists(context.Background(), categoryId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetACategoriesPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetACategoriesPlaylists`: PagingFeaturedPlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetACategoriesPlaylists`: %v\n", resp)
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
	resp, r, err := apiClient.PlaylistsAPI.GetAListOfCurrentUsersPlaylists(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetAListOfCurrentUsersPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAListOfCurrentUsersPlaylists`: PagingPlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetAListOfCurrentUsersPlaylists`: %v\n", resp)
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

> PagingFeaturedPlaylistObject GetFeaturedPlaylists(ctx).Locale(locale).Limit(limit).Offset(offset).Execute()

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
	locale := "sv_SE" // string |  (optional)
	limit := int32(10) // int32 |  (optional) (default to 20)
	offset := int32(5) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.GetFeaturedPlaylists(context.Background()).Locale(locale).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetFeaturedPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeaturedPlaylists`: PagingFeaturedPlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetFeaturedPlaylists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturedPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locale** | **string** |  | 
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
	resp, r, err := apiClient.PlaylistsAPI.GetListUsersPlaylists(context.Background(), userId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetListUsersPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListUsersPlaylists`: PagingPlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetListUsersPlaylists`: %v\n", resp)
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
	resp, r, err := apiClient.PlaylistsAPI.GetPlaylist(context.Background(), playlistId).Market(market).Fields(fields).AdditionalTypes(additionalTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylist`: PlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetPlaylist`: %v\n", resp)
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
	resp, r, err := apiClient.PlaylistsAPI.GetPlaylistCover(context.Background(), playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetPlaylistCover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistCover`: []ImageObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetPlaylistCover`: %v\n", resp)
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
	resp, r, err := apiClient.PlaylistsAPI.GetPlaylistsTracks(context.Background(), playlistId).Market(market).Fields(fields).Limit(limit).Offset(offset).AdditionalTypes(additionalTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetPlaylistsTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistsTracks`: PagingPlaylistTrackObject
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetPlaylistsTracks`: %v\n", resp)
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
	resp, r, err := apiClient.PlaylistsAPI.RemoveTracksPlaylist(context.Background(), playlistId).RemoveTracksPlaylistRequest(removeTracksPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.RemoveTracksPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveTracksPlaylist`: ReorderOrReplacePlaylistsTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.RemoveTracksPlaylist`: %v\n", resp)
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

> ReorderOrReplacePlaylistsTracks200Response ReorderOrReplacePlaylistsTracks(ctx, playlistId).Uris(uris).ReorderOrReplacePlaylistsTracksRequest(reorderOrReplacePlaylistsTracksRequest).Execute()

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
	reorderOrReplacePlaylistsTracksRequest := *openapiclient.NewReorderOrReplacePlaylistsTracksRequest() // ReorderOrReplacePlaylistsTracksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.ReorderOrReplacePlaylistsTracks(context.Background(), playlistId).Uris(uris).ReorderOrReplacePlaylistsTracksRequest(reorderOrReplacePlaylistsTracksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.ReorderOrReplacePlaylistsTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderOrReplacePlaylistsTracks`: ReorderOrReplacePlaylistsTracks200Response
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.ReorderOrReplacePlaylistsTracks`: %v\n", resp)
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
 **reorderOrReplacePlaylistsTracksRequest** | [**ReorderOrReplacePlaylistsTracksRequest**](ReorderOrReplacePlaylistsTracksRequest.md) |  | 

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
	r, err := apiClient.PlaylistsAPI.UnfollowPlaylist(context.Background(), playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.UnfollowPlaylist``: %v\n", err)
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
	r, err := apiClient.PlaylistsAPI.UploadCustomPlaylistCover(context.Background(), playlistId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.UploadCustomPlaylistCover``: %v\n", err)
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

