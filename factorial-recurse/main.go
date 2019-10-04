package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	x := 5
	n := factorial(x)
	fmt.Printf("Factorial of %d is %d", x, n)
}
