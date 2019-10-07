package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 5, 76, 34, 90, 3, 21}
	xs := []string{"bg", "rt", "as", "kj", "se", "az", "aa", "ew"}
	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)
}
