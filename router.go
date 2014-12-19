package mercury

import (
	"net/http"
	"strings"
)

const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	RESTFUL = "RESTFUL"
)

/// Handle is just like "net/http" Handlers, only takes params.
type Handle func(http.ResponseWriter, *http.Request, map[string]string)

// Router name says it all.
type Router struct {
	tree *node
}

// New creates a new router. Take the root/fall through route
// like how the default mux works. Only difference is in this case,
// you have to specific one.
func New(rootHandler Handle) *Router {
	node := node{handler: rootHandler, component: "/"}
	return &Router{tree: &node}
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

// Restful just like rails "resources" it handles all the routes.
// Could be useful if you are wanting to use an MVC design (add controllers).
func (r *Router) Restful(path string, handler Handle) {
	r.Handle(RESTFUL, path, handler)
}

// Needed by "net/http" to handle http requests.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	params := make(map[string]string) // need to add merging form params in.
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if (req.Method == node.method || node.method == RESTFUL) && node.handler != nil {
		node.handler(w, req, params)
	} else {
		r.tree.handler(w, req, params)
	}
}
