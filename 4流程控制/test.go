package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	for n := 10; n <= a; n++ {
		if sxhs(n, nlen(n)) != 0 {
			fmt.Println("水仙花数", sxhs(n, nlen(n)), nlen(n))
		}

	}

}

func nlen(x int) int {
	if x < 10 {
		return 1
	}
	n := 1
	for ; x >= 10; n++ {
		x = x / 10
	}
	return n
}

func sxhs(x int, y int) int {
	var xx int
	var xxx int

	for n := 1; n <= y; n++ { //几位数循环几次
		xx = x % 10
		var yy int = 1

		for n2 := n; n2 >= 1; n2-- { //位于第几位数循环几次
			yy = yy * xx
		}
		xxx += yy

		x = x / 10
	}

	return xxx

}
