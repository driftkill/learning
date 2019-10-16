package main

import "testing"

func TestAll(t *testing.T) {
	type itest struct {
		idata   [5]int
		ianswer [5]int
	}

	itests := []itest{
		itest{[5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
		itest{[5]int{65, 32, 677, 12, 56}, [5]int{677, 65, 56, 32, 12}},
		itest{[5]int{56, 78, 12, 01, 23}, [5]int{78, 56, 23, 12, 1}},
		itest{[5]int{45, 2, 56, 1, 65}, [5]int{65, 56, 45, 2, 1}},
		itest{[5]int{98, 100, 21, 11, 400}, [5]int{400, 100, 98, 21, 11}},
	}

	for _, v := range itests {
		x := isort(v.idata)
		for f := 0; f < 5; f++ {
			if x[f] != v.ianswer[f] {
				t.Error("Expecting", v.ianswer, "but got", x)
				break
			}
		}

	}
	type stest struct {
		sdata   [5]string
		sanswer [5]string
	}

	stests := []stest{
		stest{[5]string{"ab", "ac", "bd", "da", "bc"}, [5]string{"da", "bd", "bc", "ac", "ab"}},
		stest{[5]string{"xd", "is", "so", "idk", "oh"}, [5]string{"xd", "so", "oh", "is", "idk"}},
	}

	for _, v := range stests {
		x := ssort(v.sdata)
		for f := 0; f < 5; f++ {
			if x[f] != v.sanswer[f] {
				t.Error("Expecting", v.sanswer, "but got", x)
				break
			}
		}

	}
}
