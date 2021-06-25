// This example demonstrates how to embed static files into the Go binary.
// This works from Go 1.16 and up.

package main

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
)

//go:embed frontend/build/*
var assets embed.FS

// GetFilesystem creates a scoped filesystem so we resolve file names relative to the build directory.
func GetFilesystem() http.FileSystem {
	return http.FS(&scopedFilesystem{
		backend: assets,
		scope:   "frontend/build/",
	})
}

type scopedFilesystem struct {
	backend embed.FS
	scope   string
}

func (s scopedFilesystem) Open(name string) (fs.File, error) {
	return s.backend.Open(path.Join(s.scope, name))
}

func main() {
	fileSystem := http.FileServer(GetFilesystem())
	http.Handle("/", fileSystem)
	panic(http.ListenAndServe(":8080", nil))
}
