package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int
type simple int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "It's a dog")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "It's a cat")
}

func (s simple) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	io.WriteString(w, `Please goto <a href="localhost:8080/dog">DOG</a>`)
}

func main() {
	var d hotdog
	var c hotcat
	var s simple

	mux := http.NewServeMux()
	mux.Handle("/", s)
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
