package main

import (
	"fmt"
	"strconv"
)

func check(x int) bool {
	l := len(strconv.Itoa(x))
	a := x
	var s, sl int
	for a > 0 {
		if a < 1 {
			break
		}
		sl = a % 10
		s += sl ^ l
		a %= 10
	}
	if s == x {
		return true
	}
	return false
}

func main() {
	var a = 153
	b := check(a)
	if b == true {
		fmt.Println(a, "is an armstrong number.")
	}
}
