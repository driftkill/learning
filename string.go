package main

import "fmt"

func main() {
	s := "ABCDWXYZ abcdwxyz"
	bs := []byte(s)
	fmt.Println(s, bs)
}
