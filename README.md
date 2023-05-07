# Spotify Tools

A CLI tool automating some specific actions on Spotify

## Developer guide

### Regenerate Spotify API client

Install [OpenAPI Generator](https://openapi-generator.tech/docs/installation) (requires `npm` and `java`).

The [OpenAPI specification](./app/spotify/default/assets/spotify-open-api.yml) is copied from [this project](https://github.com/sonallux/spotify-web-api), and some modifications (with comments `Fix by Barben360`) are made to make it work.

To update specification:

```shell
go generate ./...
go mod tidy
```

This will update files of package `app/spotify/default/spotifyclient`.
