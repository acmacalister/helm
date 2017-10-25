package helm

import (
	"io"
	"log"
	"net/http"
	"time"
)

func NewLogger(w io.Writer, format string) HandlerFunc {
	logger := log.New(w, format, 0)
	return HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()
		logger.Printf(" Started %s %s", r.Method, r.URL.Path)
		rw := &responseWriter{w, 0}
		next(rw, r)
		logger.Printf(" Completed %d %s in %v", rw.status, http.StatusText(rw.status), time.Since(start))
	})
}
