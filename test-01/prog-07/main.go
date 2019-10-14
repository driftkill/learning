package main

import (
	"fmt"
	"math"
)

func area(r float64) float64 {
	a := math.Pi * r * r
	return a
}

func circum(r float64) float64 {
	cir := 2 * math.Pi * r
	return cir
}

func main() {
	fmt.Println("Area of circle with radius 7 =", area(7))
	fmt.Println("Circumferrence of circle with radius 7 =", circum(7))
}
