package main

import (
	"fmt"
)

func main3() {
	a := 10
	b := 2
	fmt.Println(a - b)
	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println("not" + "pad")
	a++
	b--
	//不能a=b++ 也没有--a
	fmt.Println(a, b) //不能直接打印a++ b
}
func main7() {
	var n = 88
	var n2 = 88
	var n3 = 87
	sum := n + n2 + n3
	avg := float64(sum) / 3
	fmt.Println(sum, avg)
	fmt.Println()
}
func main123() {
	n1, n2 := 10, 5
	fmt.Printf("%f", (float64(n1)+float64(n2))/4)

}
func main4564() {
	var r float64 = 7.2
	pi := 3.14
	fmt.Println(pi*r*r, 2*pi*r)
	T := 35.0
	Tro := 120.0
	fmt.Println(3*T + 2*Tro)
	fmt.Println((3*T + 2*Tro) * 0.88)
	day := 46
	fmt.Println("第", day%7, "天", "第", day/7, "周")
	m := 107653
	fmt.Println("第", m/86400, "天")
	fmt.Println("第", (m%86400)/3600, "小时")
	fmt.Println("第", ((m%86400)%3600)%60, "秒")
}

func main() {
	fmt.Println(10 / 4)
}
