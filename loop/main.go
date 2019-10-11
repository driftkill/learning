package main

import "fmt"

func main() {
	var x = 1
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

	i = 27
	for i > 0 {
		if i == 1 {
			break
		}
		for j = 2; j <= i; j++ {
			if i%j == 0 {
				println(j)
				i = i / j
				break
			}
		}
	}
}
