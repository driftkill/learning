package main

import (
	"fmt"
	"sort"
)

// Int no comments
func Int(a []int) {
	sort.Ints(a)
	for _, v := range a {
		fmt.Print(v, " ")
	}
}

//Float64 no comments
func Float64(a []float64) {
	sort.Float64s(a)
	for _, v := range a {
		fmt.Print(v, " ")
	}
}

//String no comments
func String(a []string) {
	sort.Strings(a)
	for _, v := range a {
		fmt.Print(v, " ")
	}
}

func main() {
	var x = []int{5, 9, 1, 34, 97, 12, 10, 55}
	var y = []float64{89.6, 43.22, 40.10, 1.02, 100.9, 54.54, 100.01}
	var z = []string{"a", "ty", "we", "bg", "jk", "ald", "ttt", "aaa", "ab", "vgh", "vv"}
	Int(x)
	fmt.Print("\n")
	Float64(y)
	fmt.Print("\n")
	String(z)
	fmt.Print("\n")
}
