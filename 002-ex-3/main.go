package main

import "fmt"

func main() {
	var yob = 1995
	var sum int
	for yob <= 2019 {
		sum = sum + 1
		yob++
	}
	fmt.Println("My age is", sum)
}
