package main

import "fmt"

func main() {
	var x = 8
	fmt.Printf("%d\t%b\n", x, x)
	var y = x << 1
	fmt.Printf("%d\t%b", y, y)
}
