package main

import "fmt"

func foo(y int) {
	fmt.Println("foo before", y)
	y = 20
	fmt.Println("foo after", y)
}

func bar(y *int) {
	fmt.Println("bar before", *y)
	*y = 30
	fmt.Println("bar after", *y)
}

func main() {
	x := 10
	fmt.Println("x before", x)
	foo(x)
	fmt.Println("x after foo", x)
	bar(&x)
	fmt.Println("x after bar", x)
}
