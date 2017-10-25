package helm

import (
	"net/http"
	"os"
	"strings"
)

// NewStatic is a builtin middleware for serving static assets. If no directories are added,
// public is used. If directories contain the same file paths, the first one is used.
func NewStatic(directories ...string) HandlerFunc {
	if len(directories) <= 0 {
		directories = append(directories, "public")
	}
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.Method != "GET" && r.Method != "HEAD" {
			next(w, r)
			return // bail out.
		}

		for i, dir := range directories {
			file := dir + r.URL.Path
			if strings.HasSuffix(r.URL.Path, "/") {
				file = file + "index.html"
			}

			if _, err := os.Stat(file); err != nil {
				if i+1 == len(directories) {
					next(w, r)
					return
				}
				continue
			}
			http.ServeFile(w, r, file)
			return
		}
		next(w, r)
	})
}
