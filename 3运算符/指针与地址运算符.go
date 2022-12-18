package main

import "fmt"

func main() {
	n := 10
	fmt.Println(&n)
	var a *int = &n //a为内存地址 *a为地址上的数据  &n为n的内存地址
	fmt.Println(*a)
}
