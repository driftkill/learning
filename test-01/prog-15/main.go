// program to convert map key/values to json

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := map[string]string{
		"Vishal":  "Kumar",
		"Rohit":   "Mishra",
		"Nishant": "Kumar",
	}
	js := maptojson(m)
	fmt.Println(string(js))
}

func maptojson(mp map[string]string) []byte {

	file, err := json.Marshal(mp)
	if err != nil {
		panic(err)
	}
	return file
}
