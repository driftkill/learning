// Demonstration of .ParseGLob and .Must
// and we'll be passing some data into templates too...

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// template.Must takes pointer to template and an error as parameter
	// template.ParseGlob takes templates as parameter and returns pointer to template
	// and an error which is the exact same paramter .Must takes
	// .Must does the error handling within itself.
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(nf, "tpl.vishal", nil); err != nil {
		log.Fatalln(err)
	}
	if err = tpl.ExecuteTemplate(nf, "two.gogo", 768); err != nil {
		log.Fatalln(err)
	}
	if err = tpl.ExecuteTemplate(nf, "three.gogo", "just do it"); err != nil {
		log.Fatalln(err)
	}
	if err = tpl.Execute(nf, "OVERRIDDEN"); err != nil {
		log.Fatalln(err)
	}
}
