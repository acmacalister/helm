package main

import (
	"fmt"
	"github.com/acmacalister/mercury"
	"net/http"
)

func main() {
	r := mercury.New(Root)
	r.Handle("GET", "/users", Users)
	r.Handle("GET", "/users/:name", UserShow)
	//r.Handle("GET", "/users/:name/blog", UserBlogShow)
	r.Handle("GET", "/blogs", Blogs)
	r.Handle("GET", "/blogs/:id", BlogShow)
	http.ListenAndServe(":8080", r)
}

func Root(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprint(w, "Root!\n")
}

func Users(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprint(w, "Users!\n")
}

func UserShow(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "Hi %s", params["name"])
}
func UserBlogShow(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "This is %s Blog", params["name"])
}

func Blogs(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprint(w, "Blogs!\n")
}

func BlogShow(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprintf(w, "Blog number: %s", params["id"])
}
