package main

import (
	"html/template"
	"log"
	"net/http"
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
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
