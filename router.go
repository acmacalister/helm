package helm

import (
	"net/http"
	"net/url"
	"strings"
)

const (
	GET    = "GET"
	HEAD   = "HEAD"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

/// Handle is just like "net/http" Handlers, only takes params.
type Handle func(http.ResponseWriter, *http.Request, url.Values) bool

// Router name says it all.
type Router struct {
	tree        *node
	rootHandler Handle
	middleware  []Handle
}

// New creates a new router. Take the root/fall through route
// like how the default mux works. Only difference is in this case,
// you have to specific one.
func New(rootHandler Handle) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]*route)}
	return &Router{tree: &node, rootHandler: rootHandler}
}

// Handle takes an http handler, method and pattern for a route.
func (r *Router) Handle(method, path string, handler Handle, middleware ...Handle) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.tree.addNode(method, path, handler, middleware...)
}

// GET same as Handle only the method is already implied.
func (r *Router) GET(path string, handler Handle, middleware ...Handle) {
	r.Handle(GET, path, handler, middleware...)
}

// HEAD same as Handle only the method is already implied.
func (r *Router) HEAD(path string, handler Handle, middleware ...Handle) {
	r.Handle(HEAD, path, handler, middleware...)
}

// POST same as Handle only the method is already implied.
func (r *Router) POST(path string, handler Handle, middleware ...Handle) {
	r.Handle(POST, path, handler, middleware...)
}

// PUT same as Handle only the method is already implied.
func (r *Router) PUT(path string, handler Handle, middleware ...Handle) {
	r.Handle(PUT, path, handler, middleware...)
}

// PATCH same as Handle only the method is already implied.
func (r *Router) PATCH(path string, handler Handle, middleware ...Handle) { // might make this and put one.
	r.Handle(PATCH, path, handler, middleware...)
}

// DELETE same as Handle only the method is already implied.
func (r *Router) DELETE(path string, handler Handle, middleware ...Handle) {
	r.Handle(DELETE, path, handler, middleware...)
}

// Add Middleware adds middleware to all of the routes.
func (r *Router) AddMiddleware(middleware ...Handle) {
	r.middleware = append(r.middleware, middleware...)
}

// runMiddleware loops over the slice of middleware and call to each of the middleware handlers.
func runMiddleware(w http.ResponseWriter, req *http.Request, params url.Values, middleware ...Handle) {
	for _, m := range middleware {
		if !m(w, req, params) {
			break // the middleware returned false, so end processing the chain.
		}
	}
}

// Needed by "net/http" to handle http requests and be a mux to http.ListenAndServe.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	runMiddleware(w, req, params, r.middleware...)
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		runMiddleware(w, req, params, handler.middleware...)
		handler.handler(w, req, params)
	} else {
		r.rootHandler(w, req, params)
	}
}
