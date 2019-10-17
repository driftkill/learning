package main

import "fmt"

func main() {
	ma := map[string]int{
		"Vishal": 10,
		"Rohit":  20,
	}
	mb := map[string]int{
		"Nishant": 30,
		"Jay":     40,
	}
	for k, v := range mb {
		ma[k] = v
	}
	fmt.Println(ma)
}
