package main

import (
	"fmt"
	"unsafe"
)

func main1() {
	var data1 int = 10
	var data2 int = 0b10
	var data3 int = 0o10
	var data4 int = 0x10
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)
	fmt.Println(data4)
	fmt.Printf("二%b", data2)
	fmt.Printf("八%o", data3)
	fmt.Printf("十六%x,%d", data4, data4)
}

func main() {
	var data uint = 10 //无符号整数只能用正数赋值
	fmt.Println(data)
	fmt.Println(unsafe.Sizeof(data))
}
