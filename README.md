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
  auth             Manage Spotify authentication
  completion       Generate the autocompletion script for the specified shell
  filter-playlists Update Spotify playlists based on filters on other playlists/shows
  get-playlist     Get basic info about a Spotify playlist and its items
  get-show         Get basic info of a Spotify show and its episodes
  help             Help about any command

Flags:
      --dev                                Dev mode
  -h, --help                               help for spotify-tools
      --public-api-endpoint string         Public API endpoint (default "http://127.0.0.1:8080")
      --server-listen-port uint16          Server listen port (default 8080)
  -i, --spotify-app-client-id string       Spotify App Client ID
  -s, --spotify-app-client-secret string   Spotify App Client Secret
  -v, --version                            version for spotify-tools

Use "spotify-tools [command] --help" for more information about a command.
```

### Authentication

Authentication is managed explicitly via the `auth` subcommand. Commands that call the Spotify API (e.g. `get-playlist`, `filter-playlists`) will **not** trigger an OAuth flow automatically — you must log in first.

```
> ./spotify-tools auth --help

Manage Spotify authentication

Usage:
  spotify-tools auth [command]

Available Commands:
  login   Log in to Spotify via OAuth authorization flow
  logout  Log out of Spotify by removing cached tokens
  status  Check Spotify authentication status
```

#### Login

```
./spotify-tools auth login -i <client-id> -s <client-secret>
```

This starts a local web server and opens the Spotify authorization URL. Visit the URL in a browser:

```
2024-01-01T00:00:00.000Z  INFO  Please open this URL in a web browser to authorize the app...
    {"url": "https://accounts.spotify.com/authorize?..."}
⇨ http server started on [::]:8080
```

After authorizing, the browser will show:

> Authorization successful! You may close this window now.

Tokens are then cached in `/tmp/.spotify-tools-cache.json` for subsequent runs (configurable via the `SPOTIFY_TOOLS_CACHE_FILE` environment variable).

#### Status

```
./spotify-tools auth status -i <client-id> -s <client-secret>
```

Prints the authentication state and exits with:
- **0** — authenticated and tokens are valid
- **2** — not authenticated (no tokens, or token refresh failed)

#### Logout

```
./spotify-tools auth logout
```

Removes the cached token file. Does not require Spotify app credentials. Never fails.

## Run in Docker

Using this tool in Docker is strongly advised (access and refresh tokens are written in clear on the disk). You can use [this public Docker image](https://ghcr.io/barben360/spotify-tools):

```
docker run -p 8080:8080 ghcr.io/barben360/spotify-tools:latest <your-command>
```

Note that you can pass your Spotify client ID/secret through the command line or through the environment (recommended as it should not appear in your Docker command).

## Developer guide

### Requirements

- You must create an app from [Spotify developer website](https://developer.spotify.com/dashboard) and get the Client ID and Client secret that you will need when running the tool
- In your app, you must add the following URI in `Redirect URIs`: `http://127.0.0.1:8080/authorize` (or you can replace `http://127.0.0.1:8080` by the public endpoint of where you run spotify-tools)

### Run debug configurations

Copy `.env.example` to `.env` and fill values. Then, use VS Code launch configurations or get inspired by them for another IDE.

### Build Docker image

For production:

```shell
export VERSION="<version>"

docker build --build-arg=SPOTIFY_TOOLS_VERSION=$VERSION -t ghcr.io/barben360/spotify-tools:$VERSION -f ./docker/Dockerfile .
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
