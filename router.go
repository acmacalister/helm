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
type Handle func(http.ResponseWriter, *http.Request, url.Values)

// Router name says it all.
type Router struct {
	tree        *node
	rootHandler Handle
}

// New creates a new router. Take the root/fall through route
// like how the default mux works. Only difference is in this case,
// you have to specific one.
func New(rootHandler Handle) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]Handle)}
	return &Router{tree: &node, rootHandler: rootHandler}
}

// Handle takes an http handler, method and pattern for a route.
func (r *Router) Handle(method, path string, handler Handle) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.tree.addNode(method, path, handler)
}

// GET same as Handle only the method is already implied.
func (r *Router) GET(path string, handler Handle) {
	r.Handle(GET, path, handler)
}

// GET same as Handle only the method is already implied.
func (r *Router) HEAD(path string, handler Handle) {
	r.Handle(HEAD, path, handler)
}

// GET same as Handle only the method is already implied.
func (r *Router) POST(path string, handler Handle) {
	r.Handle(POST, path, handler)
}

// GET same as Handle only the method is already implied.
func (r *Router) PUT(path string, handler Handle) {
	r.Handle(PUT, path, handler)
}

// GET same as Handle only the method is already implied.
func (r *Router) PATCH(path string, handler Handle) { // might make this and put one.
	r.Handle(PATCH, path, handler)
}

// GET same as Handle only the method is already implied.
func (r *Router) DELETE(path string, handler Handle) {
	r.Handle(DELETE, path, handler)
}

// Needed by "net/http" to handle http requests.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		handler(w, req, params)
	} else {
		r.rootHandler(w, req, params)
	}
}
