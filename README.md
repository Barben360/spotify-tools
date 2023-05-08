# Spotify Tools

A CLI tool automating some specific actions on Spotify

## Usage

Use integrated CLI contextual help for usage:

```
> ./spotify-tools --help

A Spotify automation tool

Usage:
  spotify-tools [command]

Available Commands:
  auth-test        Test Spotify user authentication and token refresh
  completion       Generate the autocompletion script for the specified shell
  filter-playlists Update Spotify playlists based on filters on other playlists/shows
  get-playlist     Get basic info about a Spotify playlist and its items
  get-show         Get basic info of a Spotify show and its episodes
  help             Help about any command
  reset            Reset Spotify user authentication

Flags:
      --dev                                Dev mode
  -h, --help                               help for spotify-tools
      --public-api-endpoint string         Public API endpoint (default "http://localhost:8080")
      --server-listen-port uint16          Server listen port (default 8080)
  -i, --spotify-app-client-id string       Spotify App Client ID
  -s, --spotify-app-client-secret string   Spotify App Client Secret
  -v, --version                            version for spotify-tools

Use "spotify-tools [command] --help" for more information about a command.
```

When running a command for the first time, you will get the following messages:

```
2023-05-08T08:40:16.284Z        INFO    default/authtest.go:22  Testing user auth and token refresh
2023-05-08T08:40:16.284Z        DEBUG   default/auth.go:41      no refresh token, doing authorization flow
2023-05-08T08:40:16.284Z        INFO    default/auth.go:85      Please open this URL in a web browser to authorize the app to access your Spotify account       {"url": "https://accounts.spotify.com/authorize?response_type=code&client_id=d9cbe403222b4aa2ab4a6e952bb2abe0&scope=playlist-read-private%20playlist-read-collaborative%20playlist-modify-private%20playlist-modify-public&redirect_uri=http:%2F%2Flocalhost:8080%2Fauthorize&state=fce11587-ebf3-40cc-bafc-bc37f80b2c57"}
2023-05-08T08:40:16.284Z        INFO    default/auth.go:126     starting authentication server
⇨ http server started on [::]:8080
```

Open the link in a web browser, on a machine that has access to your service through the `Redirect URI` (see developer guide) and you should get a web page with a text:

> Authorization successful! You may close this window now.

Then the command you entered will continue.

Note that your access/refresh tokens are written in clear in `/tmp/.spotify-tools-cache.json`. If you have any trouble, you can run the `reset` subcommand which will reset the authentication system.

## Run in Docker

Using this tool in Docker is strongly advised (access and refresh tokens are written in clear on the disk). You can use [this public Docker image](https://hub.docker.com/r/barben360/spotify-tools):

```
docker run -p 8080:8080 barben360/spotify-tools:v1.0.0 <your-command>
```

Note that you can pass your Spotify client ID/secret through the command line or through the environment (recommended as it should not appear in you Docker command).

## Developer guide

### Requirements

- You must create an app from [Spotify developer website](https://developer.spotify.com/dashboard) and get the Client ID and Client secret that you will need when running the tool
- In your app, you must add the following URI in `Redirect URIs`: `http://localhost:8080/authorize` (or you can replace `http://localhost:8080` by the public endpoint of where you run spotify-tools)

### Run debug configurations

Copy `.env.example` to `.env` and fill values. Then, use VS Code launch configurations or get inspired by them for another IDE.

### Build Docker image

For production:

```shell
export VERSION="<version>"

docker build --build-arg=SPOTIFY_TOOLS_VERSION=$VERSION -t barben360/spotify-tools:$VERSION -f ./docker/Dockerfile .
```

For development, you don't need the build arg and you can of course choose any tag you want.

### Regenerate Spotify API client

Install [OpenAPI Generator](https://openapi-generator.tech/docs/installation) (requires `npm` and `java`).

The [OpenAPI specification](./app/spotify/default/assets/spotify-open-api.yml) is copied from [this project](https://github.com/sonallux/spotify-web-api), and some modifications (with comments `Fix by Barben360`) are made to make it work.

To update specification:

```shell
go generate ./...
go mod tidy
```

This will update files of package `app/spotify/default/spotifyclient`.
