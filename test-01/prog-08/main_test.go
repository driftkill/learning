package main

import "testing"

func TestFact(t *testing.T) {
	type test struct {
		data   int
		answer int
	}

	tests := []test{
		test{2, 2},
		test{3, 6},
		test{4, 24},
		test{5, 120},
		test{6, 720},
	}

	for _, v := range tests {
		if fact(v.data) != v.answer {
			t.Error("Expecting", v.answer, "but got", fact(v.data))
		}
	}
}
