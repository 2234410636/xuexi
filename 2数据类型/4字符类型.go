package main

import "fmt"

func main5() {
	var ch byte
	ch = 'z'
	fmt.Printf("%c,%d", ch, ch)
}

func main123() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}

func main() {
	var a int = '我'
	var b byte = 'A'
	fmt.Printf("%T,%v", a, a)
	fmt.Printf("%T,%v"+"\n", b, b)
	fmt.Printf("%c", a)
	fmt.Printf("%c", b)
}
