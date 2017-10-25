package helm

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/schema"
)

type key string

const kparams key = "params"

type Handler interface {
	ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc)

// Router name says it all.
type Router struct {
	tree        *node
	rootHandler http.HandlerFunc
	middleware  []HandlerFunc
	URIVersion  string
}

type Param struct {
	Name     string
	Required bool
}

type mware struct {
	handler Handler
	next    *mware
}

func (h HandlerFunc) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h(rw, r, next)
}

func (m mware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(rw, r, m.next.ServeHTTP)
}

// New creates a new router. Take the root/fall through route
// like how the default mux works. Only difference is in this case,
// you have to specific one.
func New(rootHandler http.HandlerFunc) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]*route)}
	return &Router{tree: &node, rootHandler: rootHandler, URIVersion: ""}
}

// Handle takes an http handler, method and pattern for a route.
func (r *Router) Handle(method, path string, handler http.HandlerFunc, middleware ...HandlerFunc) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.tree.addNode(method, r.URIVersion+path, handler, middleware...)
}

// GET same as Handle only the method is already implied.
func (r *Router) GET(path string, handler http.HandlerFunc, middleware ...HandlerFunc) {
	r.Handle(http.MethodGet, path, handler, middleware...)
}

// HEAD same as Handle only the method is already implied.
func (r *Router) HEAD(path string, handler http.HandlerFunc, middleware ...HandlerFunc) {
	r.Handle(http.MethodHead, path, handler, middleware...)
}

// POST same as Handle only the method is already implied.
func (r *Router) POST(path string, handler http.HandlerFunc, middleware ...HandlerFunc) {
	r.Handle(http.MethodPost, path, handler, middleware...)
}

// PUT same as Handle only the method is already implied.
func (r *Router) PUT(path string, handler http.HandlerFunc, middleware ...HandlerFunc) {
	r.Handle(http.MethodPut, path, handler, middleware...)
}

// PATCH same as Handle only the method is already implied.
func (r *Router) PATCH(path string, handler http.HandlerFunc, middleware ...HandlerFunc) { // might make this and put one.
	r.Handle(http.MethodPatch, path, handler, middleware...)
}

// DELETE same as Handle only the method is already implied.
func (r *Router) DELETE(path string, handler http.HandlerFunc, middleware ...HandlerFunc) {
	r.Handle(http.MethodDelete, path, handler, middleware...)
}

// Use adds middleware to all of the routes.
func (r *Router) Use(middleware ...HandlerFunc) {
	r.middleware = append(r.middleware, middleware...)
}

// Run is a simple wrapper around http.ListenAndServe.
func (r *Router) Run(address string) {
	http.ListenAndServe(address, r)
}

// runMiddleware loops over the slice of middleware and call to each of the middleware handlers.
func runMiddleware(w http.ResponseWriter, req *http.Request, middleware mware) {
	middleware.ServeHTTP(w, req)
}

// buildMList converts/builds the list of middleware handlers to a list we can use.
func buildMList(middleware []HandlerFunc, handler http.HandlerFunc) mware {
	var next mware

	if len(middleware) == 0 {
		return finalHandler(handler)
	} else if len(middleware) > 1 {
		next = buildMList(middleware[1:], handler)
	} else {
		next = finalHandler(handler)
	}

	return mware{middleware[0], &next}
}

func finalHandler(handler http.HandlerFunc) mware {
	return mware{
		HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			handler(w, r)
		}),
		&mware{},
	}
}

// Needed by "net/http" to handle http requests and be a mux to http.ListenAndServe.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 * 1024 * 1024) // 10MB. Should probably make this configurable...
	params := req.Form

	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	req = contextSet(req, kparams, params) // set all the params we have collected.

	if h := node.methods[req.Method]; h != nil {
		runMiddleware(w, req, buildMList(append(r.middleware, h.middleware...), h.handler))
	} else {
		runMiddleware(w, req, buildMList(r.middleware, r.rootHandler))
	}
}

func contextGet(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}

func contextSet(r *http.Request, key, val interface{}) *http.Request {
	if val == nil {
		return r
	}
	return r.WithContext(context.WithValue(r.Context(), key, val))
}

func decodeFormData(r *http.Request, v interface{}) error {
	contentType := r.Header.Get("Content-type")
	if strings.Contains(contentType, "application/json") {
		return json.NewDecoder(r.Body).Decode(v)
	}
	return schema.NewDecoder().Decode(v, r.PostForm) // we don't need to parse the form as it has already been done.
}
