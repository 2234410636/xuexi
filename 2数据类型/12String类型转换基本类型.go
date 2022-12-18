package main

import (
	"fmt"
	"strconv"
)

func main789() { //Sscanf转换
	var string1, string2, string3, string4 = "12", "12.5", "true", "A"
	var a int
	var b float32
	var c bool
	var d byte

	fmt.Sscanf(string1, "%d", &a)
	fmt.Sscanf(string2, "%f", &b)
	fmt.Sscanf(string3, "%t", &c)
	fmt.Sscanf(string4, "%c", &d)
	fmt.Printf("%d\n%f\n%t\n%c\n", a, b, c, d)
	fmt.Printf("%T\n%T\n%T\n%T\n", a, b, c, d)
}

func main9879() {
	var string1, string2, string3, string4 = "12", "12.5", "true", "66"
	var a int
	var b float32
	var c bool
	var d byte = 'g'
	fmt.Println(d)
	a1, _ := strconv.ParseInt(string1, 10, 64)
	b1, _ := strconv.ParseFloat(string2, 64)
	c, _ = strconv.ParseBool(string3)
	d1, _ := strconv.ParseUint(string4, 10, 64)
	a = int(a1)
	b = float32(b1)
	d = byte(d1)

	fmt.Println(a, b, c, d)
	fmt.Printf("%T\n%T\n%T\n%T\n%c", a, b, c, d, d)
}

func main757() {
	num := 100
	str := strconv.Itoa(num)
	fmt.Printf("type:%T value:%#v\n", str, str)
}

func main() {
	str1 := "110"
	str2 := "s100"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败！", str1)
	} else {
		fmt.Printf("type:%T value:%#v\n", num1, num1)
	}
	num2, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Printf("%v 转换失败！", str2)
	} else {
		fmt.Printf("type:%T value:%#v\n", num2, num2)
	}
}
