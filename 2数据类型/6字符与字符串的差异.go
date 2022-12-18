package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ch = 'a'
	var ch1 = "a"
	fmt.Printf("%c\n"+ch1+"\n", ch)
	fmt.Printf("%T\n%T\n", ch1, ch)
	fmt.Println(unsafe.Sizeof(ch1))
	fmt.Println(unsafe.Sizeof(ch)) //不论字符串的len有多大，sizeof始终返回16，这是为啥，字符串不是也是不可变的吗？
	// 实际上字符串类型对应一个 结构体，该结构体有两个域，第一个域是指向该字符串的指针，
	//第二个域是字符串的长度，每个域占8个字节，但是并不包含指针指向的字符串的内容，这也就是为什么sizeof始终返回的是16
	//————————————————
	//版权声明：本文为CSDN博主「小辣抓」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
	//原文链接：https://blog.csdn.net/haodawang/article/details/80005072
}
