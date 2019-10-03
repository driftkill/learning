package main

import "fmt"

func main() {
	m := map[string]int{
		"Vishal": 25,
		"Kumar":  26,
	}
	if v, ok := m["Vishal"]; ok {
		fmt.Println("Search result found", v)
	}
	if v, ok := m["vishal"]; ok != true {
		fmt.Println(v, "Entry not found")
	}
}
