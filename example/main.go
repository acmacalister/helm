package main

import (
	"fmt"
	"github.com/acmacalister/helm"
	"net/http"
	"net/url"
)

func main() {
	r := helm.New(Root)
	r.Handle("GET", "/", Root)
	r.Handle("GET", "/users", Users)
	r.Handle("GET", "/users/:name", UserShow)
	r.Handle("GET", "/users/:name/blog/new", UserBlogShow)
	r.Handle("GET", "/blogs", Blogs)
	r.Handle("GET", "/blogs/:id", BlogShow)
	http.ListenAndServe(":8080", r)
}

func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Root!\n")
}

func Users(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Users!\n")
}

func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "Hi %s", params["name"])
}

func UserBlogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "This is %s Blog", params["name"])
}

func Blogs(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprint(w, "Blogs!\n")
}

func BlogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
	fmt.Fprintf(w, "Blog number: %s", params["id"])
}
