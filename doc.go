// helm is a simple, fast and minimalist router for writing web applications in Go. It builds on top of `net/http` and aims to be an elegant addition, by removing some of the cumbersome work involved with using the default `net/http` mux.
//
// For more information, see https://github.com/acmacalister/helm
//
//    package main
//
//    import (
//      "fmt"
//      "net/http"
//      "net/url"
//
//      "github.com/acmacalister/helm"
//    )
//
//    func main() {
//      r := helm.New(FallThrough)                    // Our fallthrough route.
//      r.AddMiddleware(FooMiddleware, BarMiddleware) // add global/router level middleware to run on every route.
//      r.Handle("GET", "/", Root)
//      r.Handle("GET", "/users", Users, AuthMiddleware)          // local/route specific middleware that only runs on this route.
//      r.Handle("GET", "/users/:name", UserShow, AuthMiddleware) // same as above, but with a named param.
//      r.Handle("GET", "/users/:name/blog/new", UserBlogShow, AuthMiddleware)
//      r.GET("/blogs", Blogs) // convenience method for HTTP verb. Beside GET, there is the whole RESTful gang (POST, PUT, PATCH, DELETE, etc)
//      r.GET("/blogs/:id", BlogShow)
//      http.ListenAndServe(":8080", r) // use our router as the mux!
//    }
//
//    func FooMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Println("Foo!")
//    }
//
//    func BarMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Println("Bar!")
//    }
//
//    func AuthMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Println("Doing Auth here")
//    }
//
//    func FallThrough(w http.ResponseWriter, r *http.Request, params url.Values) {
//      http.Error(w, "You done messed up A-aron", http.StatusNotFound)
//    }
//
//    func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprint(w, "Root!\n")
//    }
//
//    func Users(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprint(w, "Users!\n")
//    }
//
//    func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprintf(w, "Hi %s", params["name"]) // Notice we are able to get the username from the url resource. Quite handy!
//    }
//
//    func UserBlogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprintf(w, "This is %s Blog", params["name"])
//    }
//
//    func Blogs(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprint(w, "Blogs!\n")
//    }
//
//    func BlogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprintf(w, "Blog number: %s", params["id"])
//    }
package helm
