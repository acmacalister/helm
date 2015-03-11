package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/acmacalister/helm"
)

func main() {
	r := helm.New(FallThrough)                    // Our fallthrough route.
	r.AddMiddleware(FooMiddleware, BarMiddleware) // add global/router level middleware to run on every route.
	r.Handle("GET", "/", Root)
	r.Handle("GET", "/users", Users, AuthMiddleware)          // local/route specific middleware that only runs on this route.
	r.Handle("GET", "/users/:name", UserShow, AuthMiddleware) // same as above, but with a named param.
	r.Handle("GET", "/users/:name/blog/new", UserBlogShow, AuthMiddleware)
	r.GET("/blogs", Blogs) // convenience method for HTTP verb. Beside GET, there is the whole RESTful gang (POST, PUT, PATCH, DELETE, etc)
	r.GET("/blogs/:id", BlogShow)
	http.ListenAndServe(":8080", r) // use our router as the mux!
}

func FooMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Println("Foo!")
	return true
}

func BarMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Println("Bar!")
	return true
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Println("Doing Auth here")
	return true
}

func FallThrough(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	http.Error(w, "You done messed up A-aron", http.StatusNotFound)
	return true
}

func Root(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Fprint(w, "Root!\n")
	return true
}

func Users(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Fprint(w, "Users!\n")
	return true
}

func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Fprintf(w, "Hi %s", params["name"]) // Notice we are able to get the username from the url resource. Quite handy!
	return true
}

func UserBlogShow(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Fprintf(w, "This is %s Blog", params["name"])
	return true
}

func Blogs(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Fprint(w, "Blogs!\n")
	return true
}

func BlogShow(w http.ResponseWriter, r *http.Request, params url.Values) bool {
	fmt.Fprintf(w, "Blog number: %s", params["id"])
	return true
}
