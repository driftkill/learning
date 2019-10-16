// program to create a new json file, write into it, read it and finally print it as a string

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Details is holding names to be written into the json file
type Details struct {
	Firstname string
	Lastname  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(data Details) {
	file, err := json.MarshalIndent(data, "", "  ")
	check(err)
	err = ioutil.WriteFile("test.json", file, 0644)
	check(err)
}

func read() []byte {
	dat, err := ioutil.ReadFile("test.json")
	check(err)
	return dat
}

func main() {
	data := Details{
		Firstname: "Vishal",
		Lastname:  "Kumar",
	}
	write(data)
	dat := read()
	fmt.Println("File data -\n", string(dat))
}
