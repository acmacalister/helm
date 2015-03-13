package helm

import (
	"bufio"
	"net"
	"net/http"
)

// custom response writer that "implements http.ResponseWriter inteface"
// so we can store the status.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader just implements http ResponseWriter, but stores the status.
func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, _ := w.ResponseWriter.(http.Hijacker)
	return h.Hijack()
}
