package main

import (
	"sort"
	"testing"
)

func TestAll(t *testing.T) {
	type testint struct {
		idata   []int
		ianswer []int
	}

	testsint := []testint{
		testint{[]int{4, 2, 8, 5}, []int{2, 4, 5, 8}},
		testint{[]int{7, 44, 8, 1}, []int{1, 7, 8, 44}},
		testint{[]int{98, 200, 23, 456}, []int{23, 98, 200, 456}},
		testint{[]int{234, 2, 45, 11}, []int{2, 11, 45, 234}},
		testint{[]int{100, 29, 78, 34}, []int{29, 34, 78, 100}},
	}

	var l int
	for _, v := range testsint {
		l = len(v.idata)
		sort.Ints(v.idata)
		for i := 0; i < l; i++ {
			if v.idata[i] != v.ianswer[i] {
				t.Error("Expecting", v.ianswer, "but got", v.idata)
				break
			}
		}
	}

	type teststr struct {
		sdata   []string
		sanswer []string
	}

	testsstr := []teststr{
		teststr{[]string{"c", "a", "d", "b"}, []string{"a", "b", "c", "d"}},
		teststr{[]string{"aan", "hfh", "abb", "thr"}, []string{"aan", "abb", "hfh", "thr"}},
		teststr{[]string{"hi", "hey", "whr", "its"}, []string{"hey", "hi", "its", "whr"}},
		teststr{[]string{"thr", "whr", "her", "his"}, []string{"her", "his", "thr", "whr"}},
		teststr{[]string{"iam", "you", "so", "what"}, []string{"iam", "so", "what", "you"}},
	}

	for _, v := range testsstr {
		l = len(v.sdata)
		sort.Strings(v.sdata)
		for i := 0; i < l; i++ {
			if v.sdata[i] != v.sanswer[i] {
				t.Error("Expecting", v.sanswer, "but got", v.sdata)
				break
			}
		}
	}

	type testfloat struct {
		fdata   []float64
		fanswer []float64
	}

	testsfloat := []testfloat{
		testfloat{[]float64{1.7, 1.3, 1.9, 1.1}, []float64{1.1, 1.3, 1.7, 1.9}},
		testfloat{[]float64{56.6, 23.4, 11.11, 11.1}, []float64{11.1, 11.11, 23.4, 56.6}},
		testfloat{[]float64{1.2, 3.2, 2.2, 1.11}, []float64{1.11, 1.2, 2.2, 3.2}},
		testfloat{[]float64{22.3, 55.4, 11.3, 77.2}, []float64{11.3, 22.3, 55.4, 77.2}},
		testfloat{[]float64{1.11, 2.11, 1.12, 2.02}, []float64{1.11, 1.12, 2.02, 2.11}},
	}

	for _, v := range testsfloat {
		l = len(v.fdata)
		sort.Float64s(v.fdata)
		for i := 0; i < l; i++ {
			if v.fdata[i] != v.fanswer[i] {
				t.Error("Expecting", v.fanswer, "but got", v.fdata)
				break
			}
		}
	}
}
