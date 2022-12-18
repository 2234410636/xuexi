package main

import "fmt"

func main4561() {
	str := "steven"
	//%T类型 %v标记数值 %#V展示数据格式，语法展示
	fmt.Printf("%T, %#v \n", str, str)
	var a rune = '一'
	fmt.Printf("%T, %#v \n", a, a)
	var b byte = 'b'
	fmt.Printf("%T, %#v \n", b, b)
	var c int32 = 98
	fmt.Printf("%T, %#v \n", c, c)

}

func main222() {
	n := 123
	fmt.Printf("%8d\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%q\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%U\n", n)

}

func main789789() {
	pi := 3.1415926
	fmt.Printf("%b\n", pi) //无小数部分、二进制指数的科学计数法
	fmt.Printf("%e\n", pi) //（=%.6e）有 6 位小数部分的科学计数法
	fmt.Printf("%.10e\n", pi)
	fmt.Printf("%E\n", pi) //E为大写
	fmt.Printf("%.10E\n", pi)
	fmt.Printf("%f\n", pi) //小数点展示默认6位
	fmt.Printf("%.10f\n", pi)
	fmt.Printf("%g\n", pi)
	fmt.Printf("%g\n", 100000000000000000000.123)
	fmt.Printf("%G\n", pi)
	fmt.Printf("%G\n", 100000000000000000000.123)
}

func main() {

	arr := []byte{'x', 'y', 'z', 'z'}
	fmt.Printf("%s \n", "欢迎访问清华尹成大神")
	fmt.Printf("%q \n", "欢迎访问清华尹成大神")
	fmt.Printf("%x \n", "欢")
	fmt.Printf("%X \n", "欢迎访问清华尹成大神")
	fmt.Printf("%T, %s \n", arr, arr)
	fmt.Printf("%T, %q \n", arr, arr)
	fmt.Printf("%T, %x \n", arr, arr)
	fmt.Printf("%T, %X \n", arr, arr)

}
