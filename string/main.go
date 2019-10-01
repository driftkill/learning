package main

import "fmt"

func main() {
	s := "ABCDWX abcdwx"
	bs := []byte(s)
	fmt.Println(s, bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	sum := 0
	s1 := "vishal"
	for k := 0; k < len(s1); k++ {
		sum = sum + int(s1[k])
	}
	fmt.Println("")
	fmt.Println(sum)
}
