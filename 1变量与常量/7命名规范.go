package main

import (
	"./mypackage"
	"fmt"
)
//func main() {
//	break := 100//标识符不能作为变量
//	fmt.Println(break)
//}

func main() {
	fmt.Println(mypackage.North)//大写外部可以用 小写不可以
}