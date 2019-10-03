package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fwd bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		fwd: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(s1)
}
