package app

//go:generate openapi-generator-cli generate -g go --additional-properties=disallowAdditionalPropertiesIfNotPresent=false,packageName=spotifyclient,isGoSubmodule=true,structPrefix=true -o spotifyclient -i openapi/fixed-spotify-open-api.yml --git-repo-id=spotify-playlist-filters/app --git-user-id Barben360
//go:generate rm -rf spotifyclient/go.mod spotifyclient/go.sum spotifyclient/git_push.sh spotifyclient/.openapi-generator/VERSION
type App struct {
}
