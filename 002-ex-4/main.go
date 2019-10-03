package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:4])
	fmt.Println(x[6:])
	fmt.Println(x[1:5])
	fmt.Println(x[5:9])
}
