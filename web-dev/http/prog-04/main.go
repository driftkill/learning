package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	w.Header().Set("Vishal key", "This is just for messing around")
	fmt.Fprintf(w, "<b>I am BOLD</b><br><strong>I am STRONG</strong><br><i>It is ITALIC</i><br><em>It is EMPHASISED</em><br><big><code>This is CODE</code><br><samp>This is SAMPLE</samp></big>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
