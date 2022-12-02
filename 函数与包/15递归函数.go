package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]int

func fibonacci(n int) (res int) {
	// 记忆化：检查数组中是否已知斐波那契（n）
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
func main2() {
	//fmt.Println(fbnq(60)) //很慢
	a, b, c := 1, 1, 0
	for i := 3; i <= 60; i++ {
		c = a + b
		a = b
		b = c
	}
	fmt.Println(c)
}
func main22() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
}
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}
func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}
func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
func fn(x int) int {
	if x == 1 {
		return 3
	} else {
		return 2*fn(x-1) + 1
	}
}

func main() {
	//fmt.Println(fibonacci(45))
	//start := time.Now()
	//for i := 1; i < 30000; i++ {
	//	hz(60)
	//}
	//end := time.Now()
	//cc := end.Sub(start)
	//fmt.Printf("%s\n", cc)
	fmt.Println(fibonacciTail(11, 0, 1))
	fmt.Println(tao1(11, 1))

}

func fibonacciTail(n, first, second int) int {
	if n < 2 {
		return first
	}
	return fibonacciTail(n-1, second, first+second)
}

func tao1(n, first int) int {
	if n < 2 {
		return first
	}
	return tao1(n-1, 2*(first+1))
}

func js(f func(x int) int) func(x int) time.Duration {

	return func(x int) time.Duration {
		start := time.Now()
		for i := 0; i < 10000000; i++ {
			f(x)
		}
		fmt.Println(f(x))
		stop := time.Now().Sub(start)
		return stop

	}
}

func hz1(y, x int) int {
	if y == 1 {
		return x
	} else {
		return hz1(y-1, 2*(x+1))
	}
}

func tao(a int) int {
	if a == 1 {
		return 1
	} else {
		return (tao(a-1) + 1) * 2
	}
}
