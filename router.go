package mercury

import (
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request, map[string]string)

type Router struct {
	tree *node
}

func New(rootHandler Handle) *Router {
	node := node{path: "/", handler: rootHandler, component: "/"}
	return &Router{tree: &node}
}

func (r *Router) Handle(method, path string, handler Handle) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.tree.addNode(path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	params := make(map[string]string)
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	node.handler(w, req, params)
}
