package main

import "fmt"

func asd(pn *int) {
	*pn = 1000
	fmt.Println(*pn, &(*pn))
}

func main() {
	ss := 100
	fmt.Println(ss, &ss)
	asd(&ss)
	fmt.Println(ss, &ss)
}
