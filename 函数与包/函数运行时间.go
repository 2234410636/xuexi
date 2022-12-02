package main

import (
	"fmt"
	"time"
)

func test() {
	start := time.Now()
	var a int64 = 1
	for i := 0; i < 500000000; i++ {

		a = a + 2
	}
	fmt.Println(a)
	c := time.Now().Sub(start)
	fmt.Println(c)
}
func main() {
	test()
}
