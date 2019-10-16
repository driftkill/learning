// sorting integer and string arrays in reverse order

package main

import (
	"fmt"
)

func isort(x [5]int) [5]int {

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if x[i] < x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	return x
}

func ssort(x [5]string) [5]string {

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if x[i] < x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	return x
}

func main() {
	iar := [5]int{6, 3, 9, 12, 4}
	sar := [5]string{"ab", "cd", "aa", "bc", "bd"}
	iar = isort(iar)
	sar = ssort(sar)
	fmt.Println(iar)
	fmt.Println(sar)
}
