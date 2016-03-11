package helm

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	get     = "GET"
	head    = "HEAD"
	post    = "POST"
	put     = "PUT"
	patch   = "PATCH"
	deleteh = "DELETE"
)

// Handle is just like "net/http" Handlers, only takes params.
type Handle func(http.ResponseWriter, *http.Request, url.Values)

// Middleware is just like the Handle type, but has a boolean return. True
// means to keep processing the rest of the middleware chain, false means end.
// If you return false to end the request-response cycle you MUST
// write something back to the client, otherwise it will be left hanging.
type Middleware func(http.ResponseWriter, *http.Request, url.Values) bool

// Router name says it all.
type Router struct {
	tree           *node
	rootHandler    Handle
	middleware     []Middleware
	l              *log.Logger
	LoggingEnabled bool
	URIVersion     string
}

type Param struct {
	Name     string
	Required bool
}

// New creates a new router. Take the root/fall through route
// like how the default mux works. Only difference is in this case,
// you have to specific one.
func New(rootHandler Handle) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]*route)}
	return &Router{tree: &node, rootHandler: rootHandler, URIVersion: ""}
}

// EnableLogging sets logging to supplied writer.
func (r *Router) EnableLogging(w io.Writer) {
	r.l = log.New(w, "[helm] ", 0)
	r.LoggingEnabled = true
}

// Handle takes an http handler, method and pattern for a route.
func (r *Router) Handle(method, path string, handler Handle, middleware ...Middleware) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.tree.addNode(method, r.URIVersion+path, handler, middleware...)
}

// GET same as Handle only the method is already implied.
func (r *Router) GET(path string, handler Handle, middleware ...Middleware) {
	r.Handle(get, path, handler, middleware...)
}

// HEAD same as Handle only the method is already implied.
func (r *Router) HEAD(path string, handler Handle, middleware ...Middleware) {
	r.Handle(head, path, handler, middleware...)
}

// POST same as Handle only the method is already implied.
func (r *Router) POST(path string, handler Handle, middleware ...Middleware) {
	r.Handle(post, path, handler, middleware...)
}

// PUT same as Handle only the method is already implied.
func (r *Router) PUT(path string, handler Handle, middleware ...Middleware) {
	r.Handle(put, path, handler, middleware...)
}

// PATCH same as Handle only the method is already implied.
func (r *Router) PATCH(path string, handler Handle, middleware ...Middleware) { // might make this and put one.
	r.Handle(patch, path, handler, middleware...)
}

// DELETE same as Handle only the method is already implied.
func (r *Router) DELETE(path string, handler Handle, middleware ...Middleware) {
	r.Handle(deleteh, path, handler, middleware...)
}

// Use adds middleware to all of the routes.
func (r *Router) Use(middleware ...Middleware) {
	r.middleware = append(r.middleware, middleware...)
}

// Run is a simple wrapper around http.ListenAndServe.
func (r *Router) Run(address string) error {
	if r.LoggingEnabled {
		r.l.Println("Running on", address)
	}
	return http.ListenAndServe(address, r)
}

// runMiddleware loops over the slice of middleware and call to each of the middleware handlers.
func runMiddleware(w http.ResponseWriter, req *http.Request, params url.Values, middleware ...Middleware) bool {
	for _, m := range middleware {
		if !m(w, req, params) {
			return false // the middleware returned false, so end processing the chain.
		}
	}
	return true
}

// Needed by "net/http" to handle http requests and be a mux to http.ListenAndServe.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer clear(req)
	cw := w
	if r.LoggingEnabled {
		r.l.Printf("Started %s %s", req.Method, req.URL.Path)
		cw = &responseWriter{w, 200}
		start := time.Now()
		defer func(time.Time) {
			status := cw.(*responseWriter).status
			r.l.Printf("Completed %d %s in %v", status, http.StatusText(status), time.Since(start))
		}(start)
	}

	req.ParseMultipartForm(10 * 1024 * 1024) // 10MB. Should probably make this configurable...
	params := req.Form
	if !runMiddleware(cw, req, params, r.middleware...) {
		return // end the chain.
	}
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		if !runMiddleware(cw, req, params, handler.middleware...) {
			return
		}
		handler.handler(cw, req, params)
	} else {
		r.rootHandler(cw, req, params)
	}
}

// ValidateParams is used for validating and sanizating params. Since HTTP params can have
// same name for multiple params, if this happens it will just use the first one.
func ValidateParams(params url.Values, desiredParams []Param) (map[string]string, error) {
	paramValues := make(map[string]string)
	for _, param := range desiredParams {
		p, ok := params[param.Name]
		if param.Required && (!ok || p[0] == "") {
			return nil, errors.New(fmt.Sprintf("Required parameter (%s) not valid", param.Name))
		} else if !ok || p[0] == "" {
			continue		
		} else {
			paramValues[param.Name] = p[0]	
		}		
	}
	return paramValues, nil
}
