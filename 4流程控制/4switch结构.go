package main

import (
	"fmt"
)

func main() {
	var x interface{} //可以接受任何类型
	var y byte = 'a'
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型：%T", i)
	case uint8:
		fmt.Printf("x 的类型uint8")
	case float64:
		fmt.Printf("x 的类型：%T", i)
	case func(int) float64:
		fmt.Printf("x 的类型：%T", i)
	case bool, string:
		fmt.Printf("x 的类型：%T", i)
	default:
		fmt.Printf("x 的类型：%T", i)
	}
}
func main3() {
	var char byte
	fmt.Scanf("%c", &char)
	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	default:
		fmt.Println("other")

	}
}
