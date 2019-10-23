package main

import "fmt"

func main() {
	ma := map[string]int{
		"Vishal": 10,
		"Rohit":  20,
		"Some":   50,
		"Other":  60,
		"What":   70,
	}
	mb := map[string]int{
		"Nishant": 30,
		"Jay":     40,
	}
	for k, v := range mb {
		ma[k] = v
	}

	for i, v := range ma {
		fmt.Print(i, v, "\n")
	}
}
