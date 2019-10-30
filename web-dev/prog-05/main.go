// Demontrating template variable

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("temp.vk"))
}

func main() {
	if err := tpl.Execute(os.Stdout, "TEST"); err != nil {
		log.Fatalln(err)
	}
}
