helm
=======

yet another golang router (mux). Simple and fast.

## Features

- Simple API.
- Other Nifty features I still need to write about.
- Simple concise code-base at just a few hundred LOC. Great way to learn how to write your own router.

## Installation

`go get github.com/acmacalister/helm`

## Example

```go
package main

import (
  "fmt"
  "github.com/acmacalister/helm"
  "net/http"
)

func main() {
  r := helm.New(Root)
  r.Handle("GET", "/users", Users)
  r.Handle("GET", "/users/:name", UserShow)
  r.Handle("GET", "/users/:name/blog/new", UserBlogShow)
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
```

## Docs

[godoc](http://godoc.org/github.com/acmacalister/helm)

## Example Project

Check out the example directory for a simple example.


## TODOs

- [ ] Complete Docs
- [ ] Add Unit Tests
- [ ] Stable API
- [ ] net/http compatibility

## License

MIT

## Contact


### Austin Cherry ###
* https://github.com/acmacalister
* http://twitter.com/acmacalister
* http://austincherry.me

