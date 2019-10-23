package main

import "testing"

func TestMaptojson(t *testing.T) {
	m := map[string]string{
		"Vivek":  "Human",
		"Sparky": "Animal",
	}
	s1 := `{"Sparky":"Animal","Vivek":"Human"}`
	ms := string(maptojson(m))
	l := len(s1)
	for i := 0; i < l; i++ {
		if s1[i] != ms[i] {
			t.Error("Expeced", s1, "but got", ms)
			return
		}
	}
}
