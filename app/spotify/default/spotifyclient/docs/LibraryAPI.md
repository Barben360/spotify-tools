# \LibraryAPI

All URIs are relative to *https://api.spotify.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePlaylistDetails**](LibraryAPI.md#ChangePlaylistDetails) | **Put** /playlists/{playlist_id} | Change Playlist Details 
[**CheckCurrentUserFollows**](LibraryAPI.md#CheckCurrentUserFollows) | **Get** /me/following/contains | Check If User Follows Artists or Users 
[**CheckUsersSavedAlbums**](LibraryAPI.md#CheckUsersSavedAlbums) | **Get** /me/albums/contains | Check User&#39;s Saved Albums 
[**CheckUsersSavedAudiobooks**](LibraryAPI.md#CheckUsersSavedAudiobooks) | **Get** /me/audiobooks/contains | Check User&#39;s Saved Audiobooks 
[**CheckUsersSavedEpisodes**](LibraryAPI.md#CheckUsersSavedEpisodes) | **Get** /me/episodes/contains | Check User&#39;s Saved Episodes 
[**CheckUsersSavedShows**](LibraryAPI.md#CheckUsersSavedShows) | **Get** /me/shows/contains | Check User&#39;s Saved Shows 
[**CheckUsersSavedTracks**](LibraryAPI.md#CheckUsersSavedTracks) | **Get** /me/tracks/contains | Check User&#39;s Saved Tracks 
[**CreatePlaylist**](LibraryAPI.md#CreatePlaylist) | **Post** /users/{user_id}/playlists | Create Playlist 
[**FollowArtistsUsers**](LibraryAPI.md#FollowArtistsUsers) | **Put** /me/following | Follow Artists or Users 
[**GetAListOfCurrentUsersPlaylists**](LibraryAPI.md#GetAListOfCurrentUsersPlaylists) | **Get** /me/playlists | Get Current User&#39;s Playlists 
[**GetFollowed**](LibraryAPI.md#GetFollowed) | **Get** /me/following | Get Followed Artists 
[**GetUsersSavedAlbums**](LibraryAPI.md#GetUsersSavedAlbums) | **Get** /me/albums | Get User&#39;s Saved Albums 
[**GetUsersSavedAudiobooks**](LibraryAPI.md#GetUsersSavedAudiobooks) | **Get** /me/audiobooks | Get User&#39;s Saved Audiobooks 
[**GetUsersSavedEpisodes**](LibraryAPI.md#GetUsersSavedEpisodes) | **Get** /me/episodes | Get User&#39;s Saved Episodes 
[**GetUsersSavedShows**](LibraryAPI.md#GetUsersSavedShows) | **Get** /me/shows | Get User&#39;s Saved Shows 
[**GetUsersSavedTracks**](LibraryAPI.md#GetUsersSavedTracks) | **Get** /me/tracks | Get User&#39;s Saved Tracks 
[**GetUsersTopArtists**](LibraryAPI.md#GetUsersTopArtists) | **Get** /me/top/artists | Get User&#39;s Top Artists 
[**GetUsersTopTracks**](LibraryAPI.md#GetUsersTopTracks) | **Get** /me/top/tracks | Get User&#39;s Top Tracks 
[**RemoveAlbumsUser**](LibraryAPI.md#RemoveAlbumsUser) | **Delete** /me/albums | Remove Users&#39; Saved Albums 
[**RemoveAudiobooksUser**](LibraryAPI.md#RemoveAudiobooksUser) | **Delete** /me/audiobooks | Remove User&#39;s Saved Audiobooks 
[**RemoveEpisodesUser**](LibraryAPI.md#RemoveEpisodesUser) | **Delete** /me/episodes | Remove User&#39;s Saved Episodes 
[**RemoveShowsUser**](LibraryAPI.md#RemoveShowsUser) | **Delete** /me/shows | Remove User&#39;s Saved Shows 
[**RemoveTracksUser**](LibraryAPI.md#RemoveTracksUser) | **Delete** /me/tracks | Remove User&#39;s Saved Tracks 
[**SaveAlbumsUser**](LibraryAPI.md#SaveAlbumsUser) | **Put** /me/albums | Save Albums for Current User 
[**SaveAudiobooksUser**](LibraryAPI.md#SaveAudiobooksUser) | **Put** /me/audiobooks | Save Audiobooks for Current User 
[**SaveEpisodesUser**](LibraryAPI.md#SaveEpisodesUser) | **Put** /me/episodes | Save Episodes for Current User 
[**SaveShowsUser**](LibraryAPI.md#SaveShowsUser) | **Put** /me/shows | Save Shows for Current User 
[**SaveTracksUser**](LibraryAPI.md#SaveTracksUser) | **Put** /me/tracks | Save Tracks for Current User 
[**UnfollowArtistsUsers**](LibraryAPI.md#UnfollowArtistsUsers) | **Delete** /me/following | Unfollow Artists or Users 



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
	r, err := apiClient.LibraryAPI.ChangePlaylistDetails(context.Background(), playlistId).ChangePlaylistDetailsRequest(changePlaylistDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.ChangePlaylistDetails``: %v\n", err)
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
	resp, r, err := apiClient.LibraryAPI.CheckCurrentUserFollows(context.Background()).Type_(type_).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.CheckCurrentUserFollows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCurrentUserFollows`: []bool
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.CheckCurrentUserFollows`: %v\n", resp)
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


## CheckUsersSavedAlbums

> []bool CheckUsersSavedAlbums(ctx).Ids(ids).Execute()

Check User's Saved Albums 



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
	ids := "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.CheckUsersSavedAlbums(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.CheckUsersSavedAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedAlbums`: []bool
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.CheckUsersSavedAlbums`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedAlbumsRequest struct via the builder pattern


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


## CheckUsersSavedAudiobooks

> []bool CheckUsersSavedAudiobooks(ctx).Ids(ids).Execute()

Check User's Saved Audiobooks 



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
	ids := "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.CheckUsersSavedAudiobooks(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.CheckUsersSavedAudiobooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedAudiobooks`: []bool
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.CheckUsersSavedAudiobooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedAudiobooksRequest struct via the builder pattern


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
	resp, r, err := apiClient.LibraryAPI.CheckUsersSavedEpisodes(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.CheckUsersSavedEpisodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedEpisodes`: []bool
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.CheckUsersSavedEpisodes`: %v\n", resp)
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
	resp, r, err := apiClient.LibraryAPI.CheckUsersSavedShows(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.CheckUsersSavedShows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedShows`: []bool
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.CheckUsersSavedShows`: %v\n", resp)
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


## CheckUsersSavedTracks

> []bool CheckUsersSavedTracks(ctx).Ids(ids).Execute()

Check User's Saved Tracks 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.CheckUsersSavedTracks(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.CheckUsersSavedTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUsersSavedTracks`: []bool
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.CheckUsersSavedTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUsersSavedTracksRequest struct via the builder pattern


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
	resp, r, err := apiClient.LibraryAPI.CreatePlaylist(context.Background(), userId).CreatePlaylistRequest(createPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.CreatePlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlaylist`: PlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.CreatePlaylist`: %v\n", resp)
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
	r, err := apiClient.LibraryAPI.FollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).FollowArtistsUsersRequest(followArtistsUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.FollowArtistsUsers``: %v\n", err)
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
	resp, r, err := apiClient.LibraryAPI.GetAListOfCurrentUsersPlaylists(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetAListOfCurrentUsersPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAListOfCurrentUsersPlaylists`: PagingPlaylistObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetAListOfCurrentUsersPlaylists`: %v\n", resp)
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
	resp, r, err := apiClient.LibraryAPI.GetFollowed(context.Background()).Type_(type_).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetFollowed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowed`: GetFollowed200Response
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetFollowed`: %v\n", resp)
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


## GetUsersSavedAlbums

> PagingSavedAlbumObject GetUsersSavedAlbums(ctx).Limit(limit).Offset(offset).Market(market).Execute()

Get User's Saved Albums 



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
	market := "ES" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetUsersSavedAlbums(context.Background()).Limit(limit).Offset(offset).Market(market).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetUsersSavedAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedAlbums`: PagingSavedAlbumObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetUsersSavedAlbums`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedAlbumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **market** | **string** |  | 

### Return type

[**PagingSavedAlbumObject**](PagingSavedAlbumObject.md)

### Authorization

[oauth_2_0](../README.md#oauth_2_0)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersSavedAudiobooks

> PagingSavedAudiobookObject GetUsersSavedAudiobooks(ctx).Limit(limit).Offset(offset).Execute()

Get User's Saved Audiobooks 



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
	resp, r, err := apiClient.LibraryAPI.GetUsersSavedAudiobooks(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetUsersSavedAudiobooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedAudiobooks`: PagingSavedAudiobookObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetUsersSavedAudiobooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedAudiobooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSavedAudiobookObject**](PagingSavedAudiobookObject.md)

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
	resp, r, err := apiClient.LibraryAPI.GetUsersSavedEpisodes(context.Background()).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetUsersSavedEpisodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedEpisodes`: PagingSavedEpisodeObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetUsersSavedEpisodes`: %v\n", resp)
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
	resp, r, err := apiClient.LibraryAPI.GetUsersSavedShows(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetUsersSavedShows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedShows`: PagingSavedShowObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetUsersSavedShows`: %v\n", resp)
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


## GetUsersSavedTracks

> PagingSavedTrackObject GetUsersSavedTracks(ctx).Market(market).Limit(limit).Offset(offset).Execute()

Get User's Saved Tracks 



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
	resp, r, err := apiClient.LibraryAPI.GetUsersSavedTracks(context.Background()).Market(market).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetUsersSavedTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersSavedTracks`: PagingSavedTrackObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetUsersSavedTracks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersSavedTracksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **market** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**PagingSavedTrackObject**](PagingSavedTrackObject.md)

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
	resp, r, err := apiClient.LibraryAPI.GetUsersTopArtists(context.Background()).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetUsersTopArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersTopArtists`: PagingArtistObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetUsersTopArtists`: %v\n", resp)
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
	resp, r, err := apiClient.LibraryAPI.GetUsersTopTracks(context.Background()).TimeRange(timeRange).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetUsersTopTracks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersTopTracks`: PagingTrackObject
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetUsersTopTracks`: %v\n", resp)
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


## RemoveAlbumsUser

> RemoveAlbumsUser(ctx).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()

Remove Users' Saved Albums 



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
	ids := "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc" // string | 
	saveAlbumsUserRequest := *openapiclient.NewSaveAlbumsUserRequest() // SaveAlbumsUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.RemoveAlbumsUser(context.Background()).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RemoveAlbumsUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAlbumsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveAlbumsUserRequest** | [**SaveAlbumsUserRequest**](SaveAlbumsUserRequest.md) |  | 

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


## RemoveAudiobooksUser

> RemoveAudiobooksUser(ctx).Ids(ids).Execute()

Remove User's Saved Audiobooks 



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
	ids := "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.RemoveAudiobooksUser(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RemoveAudiobooksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAudiobooksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

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


## RemoveEpisodesUser

> RemoveEpisodesUser(ctx).Ids(ids).RemoveEpisodesUserRequest(removeEpisodesUserRequest).Execute()

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
	removeEpisodesUserRequest := *openapiclient.NewRemoveEpisodesUserRequest() // RemoveEpisodesUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.RemoveEpisodesUser(context.Background()).Ids(ids).RemoveEpisodesUserRequest(removeEpisodesUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RemoveEpisodesUser``: %v\n", err)
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
 **removeEpisodesUserRequest** | [**RemoveEpisodesUserRequest**](RemoveEpisodesUserRequest.md) |  | 

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
	r, err := apiClient.LibraryAPI.RemoveShowsUser(context.Background()).Ids(ids).Market(market).SaveShowsUserRequest(saveShowsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RemoveShowsUser``: %v\n", err)
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


## RemoveTracksUser

> RemoveTracksUser(ctx).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()

Remove User's Saved Tracks 



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
	saveAlbumsUserRequest := *openapiclient.NewSaveAlbumsUserRequest() // SaveAlbumsUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.RemoveTracksUser(context.Background()).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RemoveTracksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTracksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveAlbumsUserRequest** | [**SaveAlbumsUserRequest**](SaveAlbumsUserRequest.md) |  | 

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


## SaveAlbumsUser

> SaveAlbumsUser(ctx).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()

Save Albums for Current User 



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
	ids := "382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc" // string | 
	saveAlbumsUserRequest := *openapiclient.NewSaveAlbumsUserRequest() // SaveAlbumsUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.SaveAlbumsUser(context.Background()).Ids(ids).SaveAlbumsUserRequest(saveAlbumsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.SaveAlbumsUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAlbumsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveAlbumsUserRequest** | [**SaveAlbumsUserRequest**](SaveAlbumsUserRequest.md) |  | 

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


## SaveAudiobooksUser

> SaveAudiobooksUser(ctx).Ids(ids).Execute()

Save Audiobooks for Current User 



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
	ids := "18yVqkdbdRvS24c0Ilj2ci,1HGw3J3NxZO1TP1BTtVhpZ,7iHfbu1YPACw6oZPAFJtqe" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.SaveAudiobooksUser(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.SaveAudiobooksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAudiobooksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 

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


## SaveEpisodesUser

> SaveEpisodesUser(ctx).Ids(ids).SaveEpisodesUserRequest(saveEpisodesUserRequest).Execute()

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
	saveEpisodesUserRequest := *openapiclient.NewSaveEpisodesUserRequest([]string{"Ids_example"}) // SaveEpisodesUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.SaveEpisodesUser(context.Background()).Ids(ids).SaveEpisodesUserRequest(saveEpisodesUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.SaveEpisodesUser``: %v\n", err)
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
 **saveEpisodesUserRequest** | [**SaveEpisodesUserRequest**](SaveEpisodesUserRequest.md) |  | 

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
	r, err := apiClient.LibraryAPI.SaveShowsUser(context.Background()).Ids(ids).SaveShowsUserRequest(saveShowsUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.SaveShowsUser``: %v\n", err)
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


## SaveTracksUser

> SaveTracksUser(ctx).Ids(ids).SaveTracksUserRequest(saveTracksUserRequest).Execute()

Save Tracks for Current User 



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
	saveTracksUserRequest := *openapiclient.NewSaveTracksUserRequest([]string{"Ids_example"}) // SaveTracksUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.SaveTracksUser(context.Background()).Ids(ids).SaveTracksUserRequest(saveTracksUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.SaveTracksUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveTracksUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** |  | 
 **saveTracksUserRequest** | [**SaveTracksUserRequest**](SaveTracksUserRequest.md) |  | 

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
	r, err := apiClient.LibraryAPI.UnfollowArtistsUsers(context.Background()).Type_(type_).Ids(ids).UnfollowArtistsUsersRequest(unfollowArtistsUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.UnfollowArtistsUsers``: %v\n", err)
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

