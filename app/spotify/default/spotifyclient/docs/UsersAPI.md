# \UsersAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCurrentUserFollows**](UsersAPI.md#CheckCurrentUserFollows) | **Get** /me/following/contains | Check If User Follows Artists or Users 
[**CheckIfUserFollowsPlaylist**](UsersAPI.md#CheckIfUserFollowsPlaylist) | **Get** /playlists/{playlist_id}/followers/contains | Check if Current User Follows Playlist 
[**FollowArtistsUsers**](UsersAPI.md#FollowArtistsUsers) | **Put** /me/following | Follow Artists or Users 
[**FollowPlaylist**](UsersAPI.md#FollowPlaylist) | **Put** /playlists/{playlist_id}/followers | Follow Playlist 
[**GetCurrentUsersProfile**](UsersAPI.md#GetCurrentUsersProfile) | **Get** /me | Get Current User&#39;s Profile 
[**GetFollowed**](UsersAPI.md#GetFollowed) | **Get** /me/following | Get Followed Artists 
[**GetListUsersPlaylists**](UsersAPI.md#GetListUsersPlaylists) | **Get** /users/{user_id}/playlists | Get User&#39;s Playlists 
[**GetUsersProfile**](UsersAPI.md#GetUsersProfile) | **Get** /users/{user_id} | Get User&#39;s Profile 
[**GetUsersTopArtists**](UsersAPI.md#GetUsersTopArtists) | **Get** /me/top/artists | Get User&#39;s Top Artists 
[**GetUsersTopTracks**](UsersAPI.md#GetUsersTopTracks) | **Get** /me/top/tracks | Get User&#39;s Top Tracks 
[**UnfollowArtistsUsers**](UsersAPI.md#UnfollowArtistsUsers) | **Delete** /me/following | Unfollow Artists or Users 
[**UnfollowPlaylist**](UsersAPI.md#UnfollowPlaylist) | **Delete** /playlists/{playlist_id}/followers | Unfollow Playlist 



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
	resp, r, err := apiClient.UsersAPI.CheckCurrentUserFollows(context.Background()).Type_(type_).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CheckCurrentUserFollows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCurrentUserFollows`: []bool
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CheckCurrentUserFollows`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.CheckIfUserFollowsPlaylist(context.Background(), playlistId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CheckIfUserFollowsPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckIfUserFollowsPlaylist`: []bool
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CheckIfUserFollowsPlaylist`: %v\n", resp)
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
	r, err := apiClient.UsersAPI.FollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).FollowArtistsUsersRequest(followArtistsUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.FollowArtistsUsers``: %v\n", err)
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
	r, err := apiClient.UsersAPI.FollowPlaylist(context.Background(), playlistId).FollowPlaylistRequest(followPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.FollowPlaylist``: %v\n", err)
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


## GetCurrentUsersProfile

> PrivateUserObject GetCurrentUsersProfile(ctx).Execute()

Get Current User's Profile 



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
	resp, r, err := apiClient.UsersAPI.GetCurrentUsersProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUsersProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUsersProfile`: PrivateUserObject
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUsersProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUsersProfileRequest struct via the builder pattern


### Return type

[**PrivateUserObject**](PrivateUserObject.md)

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
	resp, r, err := apiClient.UsersAPI.GetFollowed(context.Background()).Type_(type_).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetFollowed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowed`: GetFollowed200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetFollowed`: %v\n", resp)
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
	resp, r, err := apiClient.UsersAPI.GetListUsersPlaylists(context.Background(), userId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetListUsersPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListUsersPlaylists`: PagingPlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetListUsersPlaylists`: %v\n", resp)
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


## GetUsersProfile

> PublicUserObject GetUsersProfile(ctx, userId).Execute()

Get User's Profile 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsersProfile(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersProfile`: PublicUserObject
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicUserObject**](PublicUserObject.md)

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
	resp, r, err := apiClient.UsersAPI.GetUsersTopArtists(context.Background()).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersTopArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersTopArtists`: PagingArtistObject
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersTopArtists`: %v\n", resp)
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


## GetUsersTopTracks

> PagingTrackObject GetUsersTopTracks(ctx).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()

Get User's Top Tracks 



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
	resp, r, err := apiClient.UsersAPI.GetUsersTopTracks(context.Background()).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsersTopTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersTopTracks`: PagingTrackObject
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsersTopTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersTopTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeRange** | **string** |  | [default to &quot;medium_term&quot;]
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingTrackObject**](PagingTrackObject.md)

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
	r, err := apiClient.UsersAPI.UnfollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).UnfollowArtistsUsersRequest(unfollowArtistsUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnfollowArtistsUsers``: %v\n", err)
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
	r, err := apiClient.UsersAPI.UnfollowPlaylist(context.Background(), playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UnfollowPlaylist``: %v\n", err)
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

