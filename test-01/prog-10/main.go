package main

import "fmt"

func gen(x int) []int {
	var s []int
	for i := 1; i <= 10; i++ {
		s = append(s, x*i)
	}
	return s
}

func main() {
	fmt.Println("Multiplication table of 5 -")
	x := gen(5)
	for _, v := range x {
		fmt.Println(v)
	}
}
