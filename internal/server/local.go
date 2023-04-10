//go:build !pi
// +build !pi

package server

import (
	"io/fs"
	"os"
)

func getFrontendAssets() fs.FS {
	return os.DirFS("internal/server/resources")
}
