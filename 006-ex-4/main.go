package main

import "fmt"

func main() {
	m := map[string][]string{
		"Vishal": []string{"PUBG", "COD-M"},
		"Jay":    []string{"KTM", "Alexa"},
		"Tushar": []string{"Shopping", "Dating"},
	}
	m["Gaurav"] = []string{"Beer", "Travel"}
	for k, v := range m {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Println("\t", i+1, v1)
		}
	}
	delete(m, "Gaurav")
	fmt.Println("After deletion")
	for k, v := range m {
		fmt.Println(k)
		for i, v1 := range v {
			fmt.Println("\t", i+1, v1)
		}
	}
}
