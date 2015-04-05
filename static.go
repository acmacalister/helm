package helm

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

type static struct {
	directories []string
}

// serveStaticFiles actually serves the assets by looping over the provided directories.
func (s *static) serveStaticFiles(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	if r.Method != "GET" && r.Method != "HEAD" {
		return true // bail out.
	}

	for i, dir := range s.directories {
		file := dir + r.URL.Path
		if strings.HasSuffix(r.URL.Path, "/") {
			file = file + "index.html"
		}

		if _, err := os.Stat(file); err != nil {
			if i+1 == len(s.directories) {
				return true // bail out.
			}
			continue
		}
		http.ServeFile(w, r, file)
		return false // Since we serve a file to the client, we can end the middleware and don't need to call a handler.
	}
	return true // bail out.
}

// Static is a builtin middleware for serving static assets. If no directories are added,
// public is used. If directories contain the same file paths, the first one is used.
func Static(directories ...string) Middleware {
	s := static{}
	s.directories = append(s.directories, directories...)
	if len(s.directories) <= 0 {
		s.directories = append(s.directories, "public")
	}
	return s.serveStaticFiles
}
