package main

import "testing"

func TestAll(t *testing.T) {
	type testint struct {
		idata   []int
		ianswer []int
	}

	testsint := []testint{
		testint{[]int{4, 2, 8, 5}, []int{2, 4, 5, 8}},
	}
	l := len([]int(testint.idata))
	for i, v := range testsint {
		Int(v.idata)
		if v.idata[i] != v.ianswer[i] {
			t.Error("Expecting", v.ianswer[i], "but got", v.idata[i])
		}
	}
}
