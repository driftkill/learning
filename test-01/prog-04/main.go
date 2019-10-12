package main

import "fmt"

func floyd(x int) {
	var c = 1
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(c, " ")
			c++
		}
		fmt.Print("\n")
	}
}

func main() {
	floyd(7)
}
