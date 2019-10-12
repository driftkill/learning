package main

import "fmt"

func gen(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(x, "x", i, "=", x*i)
	}
}

func main() {
	fmt.Println("Multiplication table of 5 -")
	gen(5)
}
