package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"text/template"
)

type stores struct {
	Data string
	Hash string
}

var store []stores

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/check", check)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	code := r.FormValue("code")

	hash := getCode(code)

	store := stores{
		Data: code,
		Hash: hash,
	}

	fmt.Println(store)
	tpl.ExecuteTemplate(w, "index.gohtml", store)
}

func check(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("hash")

	for _, v := range stores {
		for i := range v.Hash {
			if v.Hash[i] != hash[i] {
				return
			}
		}
		tpl.ExecuteTemplate(w, "check.gohtml", v.Data)
	}
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
