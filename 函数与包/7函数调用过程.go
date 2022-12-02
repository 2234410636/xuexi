package main

import "fmt"

func fo2() int {
	fmt.Println("go2 up")
	fo1()
	fmt.Println("go2 down")
	return 2

}
func fo3() int {
	fmt.Println("go3 up")
	fo2()
	fmt.Println("go3 down")
	return 3
}

func main() {
	fo3()
}
