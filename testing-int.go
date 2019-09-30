package main

import "fmt"

var x int8
var y int16
var z int32
var y float32

func main() {
	x = 100
	y = 10000
	z = 100000000
	y = 42.45
	x = int(y)
	fmt.Printf("%T\n%T\n%T", x, y, z)
	fmt.Println(x)
}
