package main

import "fmt"

func g1(a ...int) int {

	return g2(a...)
}

func g2(a ...int) int {
	var c = 0
	for _, b := range a {
		c += b
	}
	return c
}

func main() {
	fmt.Println(g1(1, 2, 3, 4))
}
