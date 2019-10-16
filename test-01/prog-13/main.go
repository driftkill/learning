// program to add two matrices

package main

import "fmt"

func add(x, y [3][3]int) [3][3]int {
	var z [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			z[i][j] = x[i][j] + y[i][j]
		}
	}
	return z
}

func main() {
	a := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	b := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	c := add(a, b)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(c[i][j], "  ")
		}
		fmt.Print("\n")
	}
}
