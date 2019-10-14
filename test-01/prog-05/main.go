package main

import (
	"fmt"
	"sort"
)

// Int no comments
func iint(a []int) []int {
	sort.Ints(a)
	return a
}

//Float64 no comments
func Float64(a []float64) []float64 {
	sort.Float64s(a)
	return a
}

//String no comments
func String(a []string) []string {
	sort.Strings(a)
	return a
}

func main() {
	var x = []int{5, 9, 1, 34, 97, 12, 10, 55}
	var y = []float64{89.6, 43.22, 40.10, 1.02, 100.9, 54.54, 100.01}
	var z = []string{"a", "ty", "we", "bg", "jk", "ald", "ttt", "aaa", "ab", "vgh", "vv"}
	a := iint(x)

	for _, v := range a {
		fmt.Print(v, " ")
	}

	fmt.Print("\n")

	b := Float64(y)
	for _, v := range b {
		fmt.Print(v, " ")
	}

	fmt.Print("\n")

	c := String(z)
	for _, v := range c {
		fmt.Print(v, " ")
	}

	fmt.Print("\n")
}
