package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "My-Cookie",
		Value: "Any-Value",
	})
	http.Redirect(w, r, "/read", http.StatusSeeOther)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("My-Cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "Your Cookie: ", c)
}
