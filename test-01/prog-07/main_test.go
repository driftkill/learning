package main

import "testing"

func TestAll(t *testing.T) {
	type test struct {
		rad float64
		ar  float64
		cir float64
	}

	tests := []test{
		test{10, 314.1592653589793, 62.83185307179586},
		test{5, 78.53981633974483, 31.41592653589793},
	}
	for _, v := range tests {
		if area(v.rad) != v.ar || circum(v.rad) != v.cir {
			t.Error("Expenting", v.ar, "and", v.cir, "but got", area(v.rad), "and", circum(v.rad))
		}
	}
}
