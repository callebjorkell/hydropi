//go:build pi
// +build pi

package server

import (
	"embed"
	"io/fs"
)

//go:embed resources
var embedFrontend embed.FS

func getFrontendAssets() fs.FS {
	f, err := fs.Sub(embedFrontend, "resources")
	if err != nil {
		panic(err)
	}

	return f
}
