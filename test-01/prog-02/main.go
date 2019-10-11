package main

import "fmt"

func main() {
	fmt.Println("2 + 5 =", sum(2, 5))
	fmt.Println("2 + 0 + 7 =", sum(2, 0, 7))
	fmt.Println("2 + 9 + 11 =", sum(2, 9, 11))
}

func sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		if v < 1 {
			fmt.Println("Voilation - Non natural number detected !")
			return 0
		}
		s += v
	}
	return s
}
