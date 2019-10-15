package main

import "testing"

func TestAverage(t *testing.T) {
	type test struct {
		data   [5]int
		answer float64
	}

	tests := []test{
		test{[5]int{1, 2, 3, 4, 5}, 3},
		test{[5]int{5, 3, 6, 8, 3}, 5},
		test{[5]int{10, 23, 65, 43, 22}, 32.6},
		test{[5]int{100, 200, 300, 400, 500}, 300},
		test{[5]int{10, 20, 30, 40, 50}, 30},
	}

	for _, v := range tests {
		x := average(v.data)
		if x != v.answer {
			t.Error("Expecting", v.answer, "but got", x)
		}
	}
}
