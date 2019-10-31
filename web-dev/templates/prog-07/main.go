// Demontration of FuncMap

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template
var start, end time.Time

type one struct {
	First  string
	Second string
}

type two struct {
	Third string
	Forth string
}

var fm = template.FuncMap{
	"tu": strings.ToUpper,
	"tl": strings.ToLower,
}

func init() {
	start = time.Now()
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("temp.vk"))
}

func main() {
	a := one{
		First:  "here",
		Second: "",
	}
	b := one{
		First:  "there",
		Second: "where",
	}
	c := one{
		First:  "somewhere",
		Second: "nowhere",
	}
	d := two{
		Third: "NOW",
		Forth: "THEN",
	}
	e := two{
		Third: "TOMORROW",
		Forth: "TODAY",
	}
	f := two{
		Third: "YESTERDAY",
		Forth: "WEEKEND",
	}

	ones := []one{a, b, c}
	twos := []two{d, e, f}

	data := struct {
		Lotone []one
		Lottwo []two
	}{
		ones,
		twos,
	}

	if err := tpl.ExecuteTemplate(os.Stdout, "temp.vk", data); err != nil {
		log.Fatalln(err)
	}
	end = time.Now()
	fmt.Println("\nExecution time: ", end.Sub(start))
}
