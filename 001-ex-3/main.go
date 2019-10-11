package main

import "fmt"

func main() {
	var x = 1
	for i := 65; i <= 90; i++ {
		fmt.Println(x)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		x++
	}
}
