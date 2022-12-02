package main

import "fmt"

func fo1() int {
	fmt.Println("go1 up")

	fmt.Println("go1 down")
	return 1
}

type Funct func() int

func main() {
	var f Funct
	f = fo1
	s := f()
	fmt.Println(s)
}
