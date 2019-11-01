package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "<big><em>Hey it's page Dog</em></big>")
	case "/cat":
		io.WriteString(w, "<big><em>Hey it's page Cat</em></big>")
	default:
		io.WriteString(w, `<big>Go-to <em><a href="localhost:8080/dog">Dog</a></em> or <em><a href="localhost:8080/cat">Cat</a></em></big>`)
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
