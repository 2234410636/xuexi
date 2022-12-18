package main

import (
	"fmt"
	"strconv"
)

func main11156() {
	var n1 int = 99
	var n2 float64 = 23.456
	var b bool = true
	var char int8 = 'h'
	var str string

	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", char)
	fmt.Printf("str type %T str=%q\n", str, str)
}

func main() {
	var n1 int = 99
	var n2 float64 = 23.456
	var b bool = true
	var char int8 = 'h'

	fmt.Println(strconv.FormatInt(int64(n1), 10)) //有返回值
	fmt.Println(strconv.FormatBool(b))
	fmt.Println(strconv.FormatFloat(float64(n2), 'f', 10, 64))
	fmt.Println(strconv.FormatUint(uint64(char), 10))
	fmt.Printf("%T", strconv.FormatUint(uint64(char), 10))
	fmt.Printf("%T", strconv.FormatBool(b))
	fmt.Printf("%T", strconv.FormatInt(int64(n1), 10))
	fmt.Printf("%T", strconv.FormatFloat(float64(n2), 'f', 10, 64))

}
