package main

import "fmt"

func swap(x, y int) (int, int) {
	fmt.Println("Before swapping\nx =", x, "and y =", y)
	x, y = y, x
	return x, y
}

func main() {
	x, y := swap(10, 50)
	fmt.Println("After swapping\nx =", x, "and y =", y)
}
