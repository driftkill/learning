package main

import (
	"fmt"
	"index/suffixarray"
	"strings"
)

// Contain
func conCheck(a, b string) bool {
	res := strings.Contains(a, b)
	return res
}

// Count
func couCheck(a, b string) int {
	count := strings.Count(a, b)
	return count
}

// HasPrefix along with HasSuffix
func hasCheck(a, b, c string) (bool, bool) {
	pre := strings.HasPrefix(a, b)
	suf := strings.HasSuffix(a, c)
	return pre, suf
}

// Index
func indCheck(a, b string) {
	index := suffixarray.New([]byte("banana"))
	sets := index.Lookup([]byte("ana"), -1)
	for i, set := range sets {
		fmt.Println(i, set)
	}
}

// Index
func replace(a string) string {
	new := strings.Replace(a, "beer", "weed", -1)
	return new
}

// Split
//func split(a string) {
//	res := strings.Replace(a, ",")
//	for i, v := range res {
//		fmt.Println(i, v)
//	}
//}

//ToLower along with ToUpper
func lowup(a string) (string, string) {
	low := strings.ToLower(a)
	up := strings.ToUpper(a)
	return low, up
}

func main() {

	var con = "banana"
	var x = "ana"
	fmt.Println("is", x, "present in", con)
	conCheck(con, x)

	fmt.Println("How many a letters are in", con, "?", couCheck(con, "a"))

	var city = "New York"
	fmt.Println(city, "have New as prefix and York as suffix, right?")
	hasCheck(city, "New", "York")

	fmt.Println("What are the positions", x, "used in", con)
	indCheck(con, x)

	//fmt.Println(`Now let's try to split "I am grown"`)
	//split("I am grown")

	//fmt.Println("Let's convert the case of Abu Dhabi")
	//res1 := []string{lowup("Abu Dhabi")}
}
