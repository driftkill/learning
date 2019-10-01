package main

import "fmt"

func main() {
	var i int
	for i = 33; i < 134; i++ {
		fmt.Printf("%#U\n", i)
	}
}
