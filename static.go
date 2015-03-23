package helm

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Static(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	if r.Method != "GET" && r.Method != "HEAD" {
		return true // bail out.
	}

	file := "public" + r.URL.Path
	if strings.HasSuffix(r.URL.Path, "/") {
		file = file + "index.html"
	}

	if _, err := os.Stat(file); err != nil {
		return true // bail out.
	}
	http.ServeFile(w, r, file)
	return false // Since we serve a file to the client, we can end the middleware and don't need to call a handler.
}
