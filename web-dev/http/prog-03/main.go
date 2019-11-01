package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type hotdog int

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
