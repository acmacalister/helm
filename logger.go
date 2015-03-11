package helm

import (
	"log"
	"net/http"
	"net/url"
)

func Logger(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	log.Printf("[helm] Started %s %s\n", r.Method, r.URL.Path)
	return true
}
