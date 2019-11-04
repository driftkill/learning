package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "It's a dog page")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "It's a cat page")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
