package main

import "fmt"

func main() {
	var x int = 1
	var i, j int
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}

	x = 220
	for i = x {
		for j=2; j<=i; j++ {
			if i%j == 0 {
				fmt.Printf("%d\t", j)
				break
			}
		}
		x=x%10
	}