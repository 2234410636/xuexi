package main

import "fmt"

//func main() {
//	var tall int
//	var heavy int
//	fmt.Scanf("%d",&tall)//&地址符号 取出变量地址进行赋值
//	fmt.Scanf("%d",&heavy)
//	fmt.Println(&tall)
//	fmt.Println(&heavy)
//	fmt.Printf("身高%d，体重%d",tall,heavy)
//}
func main() {
	var tall int
	var heavy int
	fmt.Scan(&tall)//&地址符号 取出变量地址进行赋值
	fmt.Scan(&heavy)
	fmt.Println(&tall)
	fmt.Println(&heavy)
	fmt.Printf("身高%d，体重%d %T.%T",tall,heavy,tall,heavy)
}
