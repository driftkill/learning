package main

import "fmt"

var x int = 99
var y int = x << 1

func main() {
	fmt.Printf("%d\t", x)
	fmt.Printf("%b\t\t", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%d\t%b\t\t%x", y, y, y)
}
