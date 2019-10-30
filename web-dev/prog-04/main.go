// Program demonstrate use of template.New function

package main

import (
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.New("Newt").Parse("It requires a string\n"))
	tpl.ExecuteTemplate(os.Stdout, "Newt", nil)
}
