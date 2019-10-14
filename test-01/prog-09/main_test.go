package main

import "testing"

func TestCalc(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2, 3}, 3},
		test{[]int{6, 4, 8}, 8},
		test{[]int{11, 22, 33}, 33},
		test{[]int{100, 20, 39}, 100},
		test{[]int{3445, 2344, 3455}, 3455},
	}

	for _, v := range tests {
		x := Calc(v.data[0], v.data[1], v.data[2])
		if x != v.answer {
			t.Error("Expecting", v.answer, "but got", x)
		}
	}
}
