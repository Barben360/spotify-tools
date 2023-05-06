# spotify-playlist-filters

A CLI tool to update Spotify playlists from other playlists with configurable filters

## Developer guide

### Regenerate Spotify API client

Install [OpenAPI Generator](https://openapi-generator.tech/docs/installation) (requires `npm` and `java`).

The OpenAPI specification is included as a git submodule (see project [here](https://github.com/sonallux/spotify-web-api)). Therefore, think to init submodules in your repository or make a recursive clone at first.

To update specification:

```shell
cd ./app/openapi
git checkout <insert-tag-here>
cd ../..
go generate ./...
go mod tidy
```

This will update files in `app/spotifyclient`


