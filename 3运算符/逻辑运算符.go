package main

import "fmt"

func main113() {
	var n bool
	fmt.Println(!n)
	fmt.Println(n)
	fmt.Println(!n && n)
	fmt.Println(!n && !n)
	fmt.Println(n && !n)
	fmt.Println(n && n)
	fmt.Println(!n || !n)
	fmt.Println(n || !n)
	fmt.Println(!n || n)
	fmt.Println(n || n)
}
func test() bool {

	fmt.Println("tesr...")
	return true
}
func main() {
	var i int = 10
	if i < 9 && test() {
		fmt.Println("ok...")

	}
	if test() && i < 9 {
		fmt.Println("ok...")

	}
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true \n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true \n")
	}
	// 修改 a 和 b 的值
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true \n")
	} else {
		fmt.Printf("第三行 - 条件为 false \n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true \n")
	}
}
