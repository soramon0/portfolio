package template

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/soramon0/portfolio/src/lib"
)

//go:embed all:dist
var content embed.FS

// All returns the content of the all directory.
func Dist(l *lib.AppLogger) http.FileSystem {
	dist, err := fs.Sub(content, "dist")
	if err != nil {
		l.ErrorFatal(err)
	}
	return http.FS(dist)
}
