package main

import "fmt"

func main() {
	var n int8 = -10
	fmt.Printf("%08b\n%d\n", n, n)
	n = n >> 1

	fmt.Printf("%08b\n%d\n", n, n)
	n = n << 2
	fmt.Printf("%08b\n%d\n", n, n)

}
