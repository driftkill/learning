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
	Hash  []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	hash := (r.FormValue("hash"))

	if code != nil {
		c := []string(getCode(code))

		stores := store{
			Plain: code,
			Hash:  c,
		}

		tpl.ExecuteTemplate(w, "index.gohtml", stores)
	}

	if hash != nil {
		for _, v := range stores {
			for i := range v.Hash {
				if v.Hash[i] != hash[i] {
					return
				}
			}
			tpl.ExecuteTemplate(w, "index.html", v.Plain)
			break
		}
	}
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
