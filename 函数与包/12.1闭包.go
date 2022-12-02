package main

import "fmt"

func add() func() int {
	x := 0
	//aa := func() int {
	//	x += 1
	//	return x
	//}

	return func() int {
		x += 1
		return x
	}
}

func main() {
	x := add()
	fmt.Println(x)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	y := add()
	fmt.Println(y)
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(x())
	fmt.Println(x()) //不管x在哪被调用都会继续叠加
}
