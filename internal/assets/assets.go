package assets

import (
	"embed"
	resource "github.com/quasilyte/ebitengine-resource"
	"io"
)

//go:embed all:_data
var gameAssets embed.FS

// OpenAsset opens an asset from the game's data directory.
func OpenAsset(path string) io.ReadCloser {
	f, err := gameAssets.Open("_data/" + path)
	if err != nil {
		panic(err)
	}
	return f
}

func RegisterResources(loader *resource.Loader) {
	registerImageResources(loader)
}
