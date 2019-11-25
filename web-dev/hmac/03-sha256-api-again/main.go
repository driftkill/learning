package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

type store struct {
	Plain string
	Hash  string
}

var stores []store

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/check", check)
	http.ListenAndServe(":8000", nil)
	fmt.Println("Server started at port :8000")
}

func index(w http.ResponseWriter, r *http.Request) {
	c := r.FormValue("code")
	code := getCode(c)
	stores := store{
		Plain: c,
		Hash:  code,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", stores)
}

func check(w http.ResponseWriter, r *http.Request) {
	h := r.FormValue("hash")
	for _, v := range stores {
		if v.Hash == h {
			tpl.ExecuteTemplate(w, "check.gohtml", v.Plain)
		}
	}
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
