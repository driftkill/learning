package main

import "fmt"

var x = 99
var y = x << 1

func main() {
	fmt.Printf("%d\t", x)
	fmt.Printf("%b\t\t", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%d\t%b\t\t%x", y, y, y)
}
