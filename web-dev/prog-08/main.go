package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now().Format(time.Stamp)
	fmt.Println(t)
}
