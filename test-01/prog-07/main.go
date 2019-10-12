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
	fmt.Println("Area of circle with radius 10 =", area(10))
	fmt.Println("Circumferrence of circle with radius 10 =", circum(10))
}
