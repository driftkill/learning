package main

import "testing"

func TestSum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2, 3}, 6},
		test{[]int{1, 2}, 3},
		test{[]int{4, 5, 6}, 15},
		test{[]int{1, 2, 3, 4}, 10},
		test{[]int{1, 2, 6, 9}, 18},
	}

	for _, v := range tests {
		x := sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
