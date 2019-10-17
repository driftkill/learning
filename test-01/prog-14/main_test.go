package main

import (
	"encoding/json"
	"testing"
)

func TestAll(t *testing.T) {

	type Details struct {
		Firstname string
		Lastname  string
	}

	type test struct {
		data   *Details
		answer string
	}

	tests := []test{
		test{data: &Details{Firstname: "Vishal", Lastname: "Kumar"},
			answer: `{"Firstname":"Vishal","Lastname":"Kumar"}`},
	}
	for _, v := range tests {
		x, _ := json.Marshal(v.data)
		write(x)
		y := read()
		sy := string(y[:])
		l := len(v.answer)
		for i := 0; i < l; i++ {
			if sy[i] != v.answer[i] {
				t.Error("Expecting", v.answer, "but got", sy)
				break
			}
		}
	}
}
