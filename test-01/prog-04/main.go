package main

import "fmt"

func floyd(x int) [][]int {

	var c = 1
	var y []int
	var xy [][]int

	for i := 1; i <= x; i++ {
		y = nil
		for j := 1; j <= i; j++ {
			y = append(y, c)
			c++
		}
		xy = append(xy, y)
	}
	return xy
}

func main() {
	res := floyd(7)
	for _, v := range res {
		fmt.Println(v)
	}
}
