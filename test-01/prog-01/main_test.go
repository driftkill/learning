package main

import "testing"

func TestSwap(t *testing.T) {
	type test struct {
		data   []int
		answer []int
	}

	tests := []test{
		test{[]int{1, 2}, []int{2, 1}},
		test{[]int{6, 3}, []int{3, 6}},
		test{[]int{11, 22}, []int{22, 11}},
		test{[]int{100, 200}, []int{200, 100}},
		test{[]int{33, 66}, []int{66, 33}},
	}

	for _, v := range tests {
		x, y := swap(v.data[0], v.data[1])
		if x != v.answer[0] || y != v.answer[1] {
			t.Error("Expected", v.answer[0], v.answer[1], "but got", x, y)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		swap(4, 5)
	}
}
