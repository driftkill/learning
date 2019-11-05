package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dogs.jpg", seedogs)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utg=f-8")
	io.WriteString(w, `<img src="/dogs.jpg" width="1200" height="620">`)
}

func seedogs(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("dogs.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
