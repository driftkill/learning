// program to calculate average of 5 integers using array implementation

package main

import "fmt"

func average(x [5]int) float64 {
	l := len(x)
	s := 0
	for _, v := range x {
		s += v
	}
	var avg = float64(s) / float64(l)
	return avg
}

func main() {
	var ar [5]int
	ar[0] = 4
	ar[1] = 8
	ar[2] = 2
	ar[3] = 5
	ar[4] = 9
	res := average(ar)
	fmt.Println(res)
}
