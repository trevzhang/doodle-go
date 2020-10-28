package main

import "fmt"

func main() {
	sum, sub := dualReturn(5, 1)
	fmt.Println("sum", sum)
	fmt.Println("sub", sub)
}

func dualReturn(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
