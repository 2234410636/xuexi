package main

import "fmt"

func main() {
	var n1 int8 = 10
	var n2 int32 = 265
	var n3 int8 = int8(n2)
	fmt.Print(n1, n2, n3)
}
