package main

import "fmt"

func main() {
	var mystr01 string = "\\r\\n" //需要转义字符\

	var mystr02 string = `\r\n` //不需要转义字符
	fmt.Println(mystr01)
	fmt.Println(mystr02)
}
