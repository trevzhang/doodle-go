package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 2
	for i := 0; i < 10; i++ {
		x, y = y, x+y
		fmt.Printf("x: %d\n", x)
		fmt.Printf("y: %d\n", y)
		fmt.Println("------")
	}
}
