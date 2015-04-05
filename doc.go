// Package helm is a simple, fast and minimalist router for writing web applications in Go. It builds on top of `net/http` and aims to be an elegant addition, by removing some of the cumbersome work involved with using the default `net/http` mux.
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
//      r := helm.New(fallThrough)                         // Our fallthrough route.
//      r.Use(fooMiddleware, barMiddleware, helm.Static()) // add global/router level middleware to run on every route.
//      r.Handle("GET", "/", root)
//      r.Handle("GET", "/users", users, authMiddleware) // local/route specific middleware that only runs on this route.
//      r.GET("/users/edit", root)
//      r.Handle("GET", "/users/:name", userShow, authMiddleware) // same as above, but with a named param.
//      r.Handle("GET", "/users/:name/blog/new", userBlogShow, authMiddleware)
//      r.GET("/blogs", blogs) // convenience method for HTTP verb. Beside GET, there is the whole RESTful gang (POST, PUT, PATCH, DELETE, etc)
//      r.GET("/blogs/:id", blogShow)
//      r.Run(":8080")
//    }
//
//    // Notice the Middleware has a return type. True means go to the next middleware. False
//    // means to stop right here. If you return false to end the request-response cycle you MUST
//    // write something back to the client, otherwise it will be left hanging.
//    func fooMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
//      fmt.Println("Foo!")
//      return true
//    }
//
//    func barMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
//      fmt.Println("Bar!")
//      return true
//    }
//
//    func authMiddleware(w http.ResponseWriter, r *http.Request, params url.Values) bool {
//      fmt.Println("Doing Auth here")
//      return true
//    }
//
//    func fallThrough(w http.ResponseWriter, r *http.Request, params url.Values) {
//      http.Error(w, "You done messed up A-aron", http.StatusNotFound)
//    }
//
//    func root(w http.ResponseWriter, r *http.Request, params url.Values) {
//      w.WriteHeader(200)
//      w.Write([]byte("Root!"))
//    }
//
//    func users(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprint(w, "Users!\n")
//    }
//
//    func userShow(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprintf(w, "Hi %s", params["name"]) // Notice we are able to get the username from the url resource. Quite handy!
//    }
//
//    func userBlogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprintf(w, "This is %s Blog", params["name"])
//    }
//
//    func blogs(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprint(w, "Blogs!\n")
//    }
//
//    func blogShow(w http.ResponseWriter, r *http.Request, params url.Values) {
//      fmt.Fprintf(w, "Blog number: %s", params["id"])
//    }
package helm
