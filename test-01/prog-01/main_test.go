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
		
		if {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		swap(4, 5)
	}
}
