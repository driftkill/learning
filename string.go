package main

import "fmt"

func main() {
	s := "ABCDWXYZ abcdwxyz"
	bs := []byte(s)
	fmt.Println(s, bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
}
