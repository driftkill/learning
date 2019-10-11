package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	p1 := person{
		first: "Vishal",
		last:  "Kumar",
		fav:   []string{"Chocolate", "Vanila"},
	}
	p2 := person{
		first: "Jay",
		last:  "Sharma",
		fav:   []string{"Curent", "Butter"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
