package main

import (
	"fmt"
)

// type test struct {
// 	First string
// }

func main() {
	// tests := []test{
	// 	test{"ab"},
	// 	test{"cd"},
	// 	test{"ef"},
	// 	test{"gh"},
	// 	test{"ij"},
	// }
	m := map[string]string{
		"a": "b",
		"c": "d",
		"e": "f",
		"g": "h",
	}
	for i, v := range m && range tests {
		fmt.Println(i, v)
	}
}
