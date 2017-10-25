// Package helm is a simple, fast and minimalist router for writing web applications in Go. It builds on top of `net/http` and aims to be an elegant addition, by removing some of the cumbersome work involved with using the default `net/http` mux.
//
// For more information, see https://github.com/acmacalister/helm
//
// package main
//
// import (
// 	"fmt"
// 	"net/http"
//
// 	"github.com/acmacalister/helm"
// )
//
// type user struct {
// 	name string
// }
//
// type server struct {
// 	db string
// }
//
// func main() {
// 	//s := server{db: "austin you are awesome!"}
// 	// helm.NewStatic()
//
// 	r := helm.New(fallThrough) // Our fallthrough route.
// 	r.Use(log, auth)           // add global/router level middleware to run on every route.
// 	r.Handle("GET", "/", root, blah)
// 	r.Run(":8080")
// }
//
// func log(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	fmt.Println("Before")
// 	next(w, r)
// 	fmt.Println("After")
// }
//
// func auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	fmt.Println("Do sweet auth stuff")
// 	next(w, r)
// }
//
// func blah(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	fmt.Println("blah...")
// 	next(w, r)
// }
//
// func fallThrough(w http.ResponseWriter, r *http.Request) {
// 	http.Error(w, "You done messed up A-aron", http.StatusNotFound)
// }
//
// func root(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("root!")
// 	w.WriteHeader(200)
// 	w.Write([]byte("Root!"))
// }
package helm
