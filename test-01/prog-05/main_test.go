package main

import "testing"

func TestIint(t *testing.T) {

	type test struct {
		idata   []int
		ianswer []int
	}

	tests := []test{
		test{[]int{4, 7, 1, 9}, []int{1, 4, 7, 9}},
		test{[]int{7, 3, 9, 5}, []int{3, 5, 7, 9}},
		test{[]int{30, 60, 10, 5}, []int{5, 10, 30, 60}},
		test{[]int{65, 899, 23, 34}, []int{23, 34, 65, 899}},
		test{[]int{34, 23, 56, 4}, []int{4, 23, 34, 56}},
	}
	for i, v := range tests {
		iint(v.idata)
		if v.idata[i] != v.ianswer[i] {
			t.Error("Expenting", v.ianswer, "but got", v.idata)
		}
	}
}

// func TestFloat64(t *testing.T) {

// 	type test struct {
// 		fdata   []float64
// 		fanswer []float64
// 	}

// 	tests := []test{
// 		test{[]float64{4.4, 2.2, 8.8, 1.1}, []float64{1, 4, 7, 9}},
// 		test{[]float64{1.11, 7.77, 4.44, 2.22}, []float64{1.11, 2.22, 4.44, 7.77}},
// 		test{[]float64{4.32, 11.11, 2.12}, []float64{2.12, 4.32, 11.11}},
// 		test{[]float64{3.43, 12.2, 56.43}, []float64{3.43, 12.2, 56.43}},
// 		test{[]float64{76.12, 34.22, 12.32, 66.3}, []float64{12.32, 34.22, 66.3, 76.12}},
// 	}
// 	for i, v := range tests {
// 		Float64(v.fdata)
// 		if v.fdata[i] != v.fanswer[i] {
// 			t.Error("Expenting", v.fanswer, "but got", v.fdata)
// 		}
// 	}
// }

// func TestString(t *testing.T) {

// 	type test struct {
// 		sdata   []string
// 		sanswer []string
// 	}

// 	tests := []test{
// 		test{[]string{"g", "d", "a", "t"}, []string{"a", "d", "g", "t"}},
// 		test{[]string{"dd", "ss", "cc"}, []string{"cc", "dd", "ss"}},
// 		test{[]string{"hhh", "sss", "aaa"}, []string{"aaa", "hhh", "sss"}},
// 		test{[]string{"d", "c", "bb", "aa"}, []string{"aa", "bb", "c", "d"}},
// 		test{[]string{"v", "hj", "fd", "aa"}, []string{"aa", "fd", "hj", "v"}},
// 	}
// 	for i, v := range tests {
// 		String(v.sdata)
// 		if v.sdata[i] != v.sanswer[i] {
// 			t.Error("Expenting", v.sanswer, "but got", v.sdata)
// 		}
// 	}
// }
