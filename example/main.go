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
