package main

import (
	"fmt"
	"math"
)

func check(x int) bool {
	count := 0.00
	c := x
	for c > 0 {
		count++
		c /= 10
	}
	a := x
	var s, r int
	for a > 0 {
		r = a % 10
		s = s + int(math.Pow(float64(r), count))
		a = a / 10
	}
	if s == x {
		return true
	}
	return false
}

func main() {
	var a = 1634
	b := check(a)
	if b == true {
		fmt.Println(a, "is an armstrong number.")
	} else {
		fmt.Println("Wrong number")
	}
}
