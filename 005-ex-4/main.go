package main

import "fmt"

func main() {
	x1 := []string{"James", "Bond", "Nothing"}
	x2 := []string{"Sandra", "Lois", "Something"}
	x := [][]string{x1, x2}
	fmt.Println(x)
}
