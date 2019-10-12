package main

import "fmt"

func fact(x int) int {
	var fac = 1
	for i := x; i > 1; i-- {
		fac *= i
	}
	return fac
}

func main() {
	fmt.Println("Factorial of 10 = ", fact(10))
}
