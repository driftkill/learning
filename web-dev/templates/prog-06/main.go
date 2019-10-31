// Demonstration of map data passing to a template

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
	m := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	if err := tpl.Execute(os.Stdout, m); err != nil {
		log.Fatalln(err)
	}
}
