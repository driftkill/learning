package main

import "testing"

func testCheck(t *testing.T) {
	type test struct {
		data   int
		answer bool
	}

	tests := []test{
		test{153, true},
		test{370, true},
		test{371, true},
		test{1634, true},
		test{100, false},
	}

	for _, v := range tests {
		x := check(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
