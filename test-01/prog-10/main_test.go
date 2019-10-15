package main

import "testing"

func TestGen(t *testing.T) {
	type test struct {
		data   int
		answer []int
	}

	tests := []test{
		test{2, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
		test{3, []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30}},
		test{4, []int{4, 8, 12, 16, 20, 24, 28, 32, 36, 40}},
		test{5, []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}},
		test{6, []int{6, 12, 18, 24, 30, 36, 42, 48, 54, 60}},
	}

	var l int

	for _, v := range tests {
		x := gen(v.data)
		l = len(x)
		for i := 0; i < l; i++ {
			if x[i] != v.answer[i] {
				t.Error("Expecting", v.answer, "but got", x)
				break
			}
		}
	}
}
