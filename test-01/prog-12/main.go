// sorting integer and string arrays in reverse order

package main

func isort(x [5]int) [5]int {
	for i, v := range x {
		if v[i] < v[i+1] {
			v[i] = v[i+1]
		}
	}
}
