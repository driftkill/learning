package main

import "fmt"

//Calc calculates the largest integer among 3 integers and returns it
func Calc(x, y, z int) int {
	if x > y && x > z {
		return x
	} else if y > z {
		return y
	} else {
		return z
	}
}

func main() {
	var a, b, c int
	fmt.Print("Enter 1st integer: ")
	fmt.Scan(&a)
	fmt.Print("Enter 2nd integer: ")
	fmt.Scan(&b)
	fmt.Print("Enter 3rd integer: ")
	fmt.Scan(&c)
	fmt.Println(Calc(a, b, c), "is the largest.")
}
