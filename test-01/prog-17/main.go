package main

import "fmt"

func insert() {
	m := map[string]string{}
	x := true
	var key, val, res string
	for x == true {
		fmt.Print("\nInsert key (Type-String)")
		fmt.Scan(&key)
		fmt.Print("\nInsert value (Type-String)")
		fmt.Scan(&val)
		m[key] = val
		fmt.Print("\nDo you want to insert more?(Y/y)")
		fmt.Scan(&res)
		if res == "y" || res == "Y" {
			insert()
		}
		break
	}
	fmt.Println("Here's your result\n", m)
}

func main() {
	insert()
}
