package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Vishal",
		Last:  "Kumar",
		Age:   25,
	}

	p2 := person{
		First: "Jay",
		Last:  "Sharma",
		Age:   25,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(bs))
}
