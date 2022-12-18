package main

import "fmt"

func main() {

	//str := "hello, world"
	//str1 := str[:5] // 获取索引5（不含）之前的子串
	//str2 := str[7:] // 获取索引7（含）之后的子串
	//str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	//fmt.Println("str1:", str1)
	//fmt.Println("str2:", str2)
	//fmt.Println("str3:", str3)

	str := "Hello, 世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，ch 类型为 byte，汉字三个字节
		fmt.Println(i, ch)
	}
	for i, ch := range str {
		fmt.Println(i, ch) // ch 的类型为 rune
	}

}
