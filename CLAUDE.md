# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`spotify-tools` is a Go CLI tool that automates actions on Spotify via its Web API. It uses OAuth2 authorization code flow to authenticate users and exposes commands for inspecting and managing playlists/shows.

## Repository Layout

```
.
├── main.go                          # Entrypoint: wires up CLI with signal handling
├── go.mod / go.sum                  # Go module definition (module: github.com/Barben360/spotify-tools)
├── .env.example                     # Template for local env vars
├── cli/
│   ├── cli.go                       # CLI struct, cobra command registration, flag definitions
│   ├── handlers.go                  # Per-command handler functions (delegates to app layer)
│   └── utils.go                     # Config file/JSON loading helper
├── app/
│   ├── app.go                       # App struct, constructor, public API methods
│   ├── models.go                    # AppConfig struct and internal features aggregation
│   ├── spotify/
│   │   ├── spotify.go              # Spotifier interface (composed of sub-interfaces)
│   │   ├── models.go               # Domain models: Playlist, Show, Item, PlaylistFilterConfig, etc.
│   │   └── default/
│   │       ├── spotify.go          # Spotify struct, constructor, token caching from disk
│   │       ├── auth.go             # OAuth2 flow, token refresh, local callback server
│   │       ├── playlist.go         # Playlist CRUD + filter/reorder logic
│   │       ├── show.go             # Show/episode fetching with pagination
│   │       └── spotifyclient/      # Auto-generated OpenAPI client (do not edit manually)
│   │           └── assets/
│   │               ├── spotify-open-api.yml        # OpenAPI spec (with local patches)
│   │               └── patches/0001-Make-oneof-less-strict.patch
│   ├── services/
│   │   └── logger/logger.go        # Zap-based logger with context propagation
│   └── utils/
│       └── httpclient.go           # Custom http.RoundTripper with logging & request transforms
├── docker/
│   └── Dockerfile                  # Multi-stage build (golang:1.26-alpine → scratch)
└── examples/
    ├── docker-compose.yml
    └── playlist-filter/
        ├── name-filter.json         # Example: filter by item name regexp
        └── grosses-tetes.json       # Example: filter by regexp + min duration
```

## Architecture

The project follows a clean layered architecture:

```
main.go
  └── cli/           (cobra CLI layer — parses flags, loads config, builds App)
        └── app/     (application layer — orchestrates features, validates config)
              └── app/spotify/default/  (infrastructure layer — Spotify API calls)
```

### Key Design Patterns

- **Interface-driven**: `spotify.Spotifier` is an interface composed of `SpotifiyAuther`, `SpotifyPlaylister`, and `SpotifyShower`. The CLI and app layers depend only on the interface, not the concrete implementation.
- **Context propagation**: The logger is stored in and retrieved from `context.Context` via `logger.ToContext` / `logger.FromContext`.
- **Token caching on disk**: Access and refresh tokens are persisted to `/tmp/.spotify-tools-cache.json` to avoid re-authorization on every run.
- **Auto-retry on 401**: `authExec` in `app/spotify/default/spotify.go` wraps all API calls with automatic token refresh on unauthorized responses.
- **Pagination**: Both playlist and show fetching loop over paginated API responses transparently.

## Available Commands

| Command | Description |
|---|---|
| `auth-test` | Runs OAuth flow (if needed) then verifies token refresh works |
| `reset` | Deletes the token cache file to force re-authorization |
| `get-playlist <playlistID>` | Prints JSON of a playlist and all its items |
| `get-show <showID>` | Prints JSON of a show and all its episodes |
| `filter-playlists` | Reads a JSON config and populates/reorders target playlists |

### Global Flags

| Flag | Env var | Default | Description |
|---|---|---|---|
| `-i / --spotify-app-client-id` | `SPOTIFY_TOOLS_CLIENT_ID` | — | Spotify OAuth app Client ID |
| `-s / --spotify-app-client-secret` | `SPOTIFY_TOOLS_CLIENT_SECRET` | — | Spotify OAuth app Client Secret |
| `--dev` | — | `false` | Dev mode (verbose/development zap logger) |
| `--public-api-endpoint` | — | `http://localhost:8080` | Public URL for the OAuth redirect URI |
| `--server-listen-port` | — | `8080` | Port for the local OAuth callback server |

### `filter-playlists` Extra Flags

| Flag | Default | Description |
|---|---|---|
| `-f / --config-file` | — | Path to JSON config file (mutually exclusive with `--config`) |
| `-c / --config` | — | Inline JSON config string |
| `--daemon` | `false` | Run repeatedly on a timer |
| `-p / --period` | `1h` | Interval between runs in daemon mode |

## filter-playlists Config Schema

The config is a JSON array of `PlaylistFilterConfig` objects:

```json
[
  {
    "target_playlist_id": "<Spotify playlist ID to update>",
    "description": "Human-readable label (used in logs)",
    "order_by": "added_at" | "release_date",
    "add_latest_update_date_to_description": true | false,
    "latest_update_date_location": "Europe/Paris",
    "sources": [
      {
        "playlist_id": "<source playlist ID>",   // mutually exclusive with show_id
        "show_id": "<source show ID>",            // mutually exclusive with playlist_id
        "filters": {
          "item_name_regexp": "<Go regexp>",      // optional
          "min_duration": 3600,                  // seconds, optional
          "max_duration": 7200                   // seconds, optional
        }
      }
    ]
  }
]
```

Filter behaviour:
1. Items already present in the target playlist are skipped (idempotent).
2. Items that do not match `item_name_regexp` are skipped.
3. Items outside `[min_duration, max_duration]` are skipped.
4. Surviving items are appended to the target playlist, then the entire playlist is reordered according to `order_by`.

## Development Setup

### Prerequisites

- Go 1.26+
- A Spotify developer app with the redirect URI `http://localhost:8080/authorize` (or your custom endpoint) configured.

### Local Environment

```sh
cp .env.example .env
# Fill in SPOTIFY_TOOLS_CLIENT_ID and SPOTIFY_TOOLS_CLIENT_SECRET
```

The `.env` file is gitignored and used by VS Code launch configs.

### Build

```sh
go build -o spotify-tools .
```

### Run (examples)

```sh
# Auth test
./spotify-tools auth-test -i <client-id> -s <client-secret>

# Or via env vars
export SPOTIFY_TOOLS_CLIENT_ID=<id>
export SPOTIFY_TOOLS_CLIENT_SECRET=<secret>
./spotify-tools get-playlist 37i9dQZF1DXdrln2UyZD7F

# Filter playlists from a config file
./spotify-tools filter-playlists -f examples/playlist-filter/name-filter.json

# Daemon mode, runs every 5 minutes
./spotify-tools filter-playlists -f examples/playlist-filter/grosses-tetes.json --daemon -p 5m
```

### Testing

There are no automated tests in this repository. Verify changes manually using the VS Code launch configurations in `.vscode/launch.json` or by running the binary directly.

## CI/CD

Two GitHub Actions workflows are defined in `.github/workflows/`:

- **`ci.yml`**: Runs `go build ./...` on every push and pull request to any branch.
- **`docker.yml`**: Builds and pushes a multi-platform (`linux/amd64`, `linux/arm64`) Docker image to `ghcr.io/barben360/spotify-tools`. Triggered manually (tags with short commit SHA) or automatically on a GitHub Release (tags with the release tag and `latest`).

## Docker

### Build

```sh
export VERSION="v1.0.0"
docker build --build-arg=SPOTIFY_TOOLS_VERSION=$VERSION \
  -t ghcr.io/barben360/spotify-tools:$VERSION \
  -f ./docker/Dockerfile .
```

The Dockerfile uses a two-stage build (builder on `golang:1.26-alpine`, final image on `scratch`) and applies `upx` compression to the binary.

### Run

```sh
docker run -p 8080:8080 \
  -e SPOTIFY_TOOLS_CLIENT_ID=<id> \
  -e SPOTIFY_TOOLS_CLIENT_SECRET=<secret> \
  ghcr.io/barben360/spotify-tools:v1.0.0 filter-playlists -f /playlist-filter/config.json
```

See `examples/docker-compose.yml` for a daemon-mode compose setup.

## Regenerating the Spotify OpenAPI Client

The `app/spotify/default/spotifyclient/` directory is auto-generated and **must not be edited by hand**.

To regenerate:

1. Install [OpenAPI Generator](https://openapi-generator.tech/docs/installation) (requires `npm` and `java`).
2. Update or replace `app/spotify/default/assets/spotify-open-api.yml` if needed (sourced from [sonallux/spotify-web-api](https://github.com/sonallux/spotify-web-api)). Lines marked `Fix by Barben360` contain manual corrections that must be re-applied via the patch.
3. Run:

```sh
go generate ./...
go mod tidy
```

This re-runs the generator and automatically applies `assets/patches/0001-Make-oneof-less-strict.patch`.

## Key Files for Common Tasks

| Task | File(s) |
|---|---|
| Add a new CLI command | `cli/cli.go` (register), `cli/handlers.go` (handler), `app/app.go` (public method) |
| Add a new Spotify API operation | `app/spotify/spotify.go` (interface), `app/spotify/default/` (implementation) |
| Modify playlist filter logic | `app/spotify/default/playlist.go` → `UpdatePlaylistFilter` |
| Change domain models | `app/spotify/models.go` |
| Adjust logging | `app/services/logger/logger.go`, pass logger through context |
| Change HTTP transport behaviour | `app/utils/httpclient.go` |

## Conventions

- **Error wrapping**: Use `fmt.Errorf("context: %w", err)` for error wrapping throughout.
- **Validation**: Struct validation uses `github.com/go-playground/validator/v10` tags (`validate:"required"`, `validate:"omitempty"`, etc.).
- **Logging**: Always retrieve the logger from context (`logger.FromContext(ctx)`). Never create a new logger inside a function — propagate it via context.
- **No global state**: All state lives on the `Spotify` struct, protected by `authLock` for token operations.
- **Batch size**: Spotify API calls that mutate playlists use a batch size of 100 items (API limit).
- **Token storage**: Tokens are written to `/tmp/.spotify-tools-cache.json` (plaintext). Using Docker is recommended to limit exposure.
- **Version injection**: The binary version is set at build time via `-ldflags "-X github.com/Barben360/spotify-tools/cli.version=<version>"`.
