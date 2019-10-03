package main

import "fmt"

func (s person) print() {
	fmt.Println(s.first, s.last)
}

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "Vishal",
		last:  "Kumar",
	}
	p.print()
}
