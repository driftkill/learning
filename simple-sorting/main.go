package main

import "fmt"

func main() {
	var x [10]int
	x[0] = 12
	x[1] = 43
	x[2] = 9
	x[3] = 123
	x[4] = 32
	x[5] = 2
	x[6] = 19
	x[7] = 99
	x[8] = 76
	x[9] = 22
	fmt.Println(x)
	var temp int
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if x[i] > x[j] {
				temp = x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
	fmt.Println(x)
}
