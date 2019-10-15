package main

import "testing"

func TestFloyd(t *testing.T) {
	type test struct {
		data   int
		answer [][]int
	}

	tests := []test{
		test{2, [][]int{{1}, {2, 3}}},
		test{3, [][]int{{1}, {2, 3}, {4, 5, 6}}},
		test{4, [][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}}},
		test{5, [][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}, {11, 12, 13, 14, 15}}},
		test{5, [][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20, 21}}},
	}

	for i, v := range tests {
		x := floyd(v.data)
		for j := 0; j <= i; j++ {
			for k := 0; k <= j; k++ {
				if x[j][k] != v.answer[j][k] {
					t.Error("Expecting", v.answer, "but got", x)
					break
				}
			}
		}
	}
}
