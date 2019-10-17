// program to convert map key/values to json

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	js := maptojson()
	fmt.Println(string(js))
}

func maptojson() []byte {
	m := map[string]string{
		"Vishal":  "Kumar",
		"Rohit":   "Mishra",
		"Nishant": "Kumar",
	}
	fmt.Println(m)
	file, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return file
}
