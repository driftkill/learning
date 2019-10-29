package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Vishal Kumar"

	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html>
	<head> <title> Hello World </title> </head>
	<body> <h1> ` + name + `</h1> </body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating index file: ", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
