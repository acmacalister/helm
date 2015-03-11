helm
=======

helm is a simple, fast and minimalist router for writing web applications in Go. It builds on top of `net/http` and aims to be an elegant addition by removing some of the cumbersome work involved with using the default `net/http` mux.

## Features

- Simple API.
- Middleware support built right in.
- Great for building API servers.
- Minimalist codebase at just a few hundred LOC. Great way to learn how to write your own router.
- Inspired by Express.js.

## Installation

`go get github.com/acmacalister/helm`

## Example

```go
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

// Notice the Middleware has a return type. True means go to the next middleware. False
// means to stop right here. If you return false to end the request-response cycle you MUST
// write something back to the client, otherwise it will be left hanging.
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

func FallThrough(w http.ResponseWriter, r *http.Request, params url.Values) {
  http.Error(w, "You done messed up A-aron", http.StatusNotFound)
}

func Root(w http.ResponseWriter, r *http.Request, params url.Values) {
  fmt.Fprint(w, "Root!\n")
}

func Users(w http.ResponseWriter, r *http.Request, params url.Values) {
  fmt.Fprint(w, "Users!\n")
}

func UserShow(w http.ResponseWriter, r *http.Request, params url.Values) {
  fmt.Fprintf(w, "Hi %s", params["name"]) // Notice we are able to get the username from the url resource. Quite handy!
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

```

## Docs

[godoc](http://godoc.org/github.com/acmacalister/helm)

## Example Project

Check out the example directory for a simple example.

## Why?

There are already a number of great routers and middleware out there for Go, but since most of them are either middlware or a router, getting them to work together felt clumsy to me. Helm's goal is to provide a minimalist set of tools to make building web services a breeze.

## TODOs

- [ ] Add Unit Tests
- [ ] Add some handy default middleware. (static, logging, etc)
- [ ] Add support for something like the express.js `all` method.

## Contributing

If you are interested on helping out or have a feature suggestion, feel free to open an issue or do a PR.

## Additional middleware

helm's middleware is quite simple as it is standard `net/http` functions that provides pre-parsed params. If you would like would to include a middleware that is compatibility with helm, open an issue and we will get it added.

## License

MIT

## Contact


### Austin Cherry ###
* https://github.com/acmacalister
* http://twitter.com/acmacalister
* http://austincherry.me

