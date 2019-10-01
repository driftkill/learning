package main

import "fmt"

const (
	a = 2019 - iota
	b = 2019 - iota
	c = 2019 - iota
	d = 2019 - iota
)

func main() {
	fmt.Println(a, b, c, d)
}
