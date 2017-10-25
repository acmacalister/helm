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
  "os"

  "github.com/acmacalister/helm"
)

func main() {
  r := helm.New(fallThrough)                                         // Our fallthrough route.
  r.Use(helm.NewLogger(os.Stdout, "[helm]"), auth, helm.NewStatic()) // add global/router level middleware to run on every route.
  r.Handle("GET", "/", root, blah)
  r.Run(":8080")
}

func log(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  fmt.Println("Before")
  next(w, r)
  fmt.Println("After")
}

func auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  fmt.Println("Do sweet auth stuff")
  next(w, r)
}

func blah(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  fmt.Println("blah...")
  next(w, r)
}

func fallThrough(w http.ResponseWriter, r *http.Request) {
  http.Error(w, "You done messed up A-aron", http.StatusNotFound)
}

func root(w http.ResponseWriter, r *http.Request) {
  fmt.Println("root!")
  w.WriteHeader(200)
  w.Write([]byte("Root!"))
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
- [ ] Add support for something like the express.js `all` method.

## Contributing

If you are interested on helping out or have a feature suggestion, feel free to open an issue or do a PR.

## Additional middleware

helm's middleware is quite simple as it is standard `net/http` functions that provides pre-parsed params. If you would like would to include a middleware that is compatibility with helm, open an issue and we will get it added.

## License

MIT

## Contact


### Austin Cherry ###
* http://twitter.com/acmacalister
* http://austincherry.me

