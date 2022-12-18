package main

import "fmt"
import _ "./mypackage"//只调用init函数 其他不调用

func main() {
	fmt.Println("gogogo")

}
