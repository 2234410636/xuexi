package main

import "fmt"

func main11() {
	const tall = 175//自动推倒类型
	fmt.Println("身高",tall)
	//tall = 180  //不行 不允许修改
}

//func returnint() int{
//	return 10
//}

func main() {
	const tall int = 175
	//显式类型定义： const b string = "abc"
	//隐式类型定义： const b = "abc"

	const (//批量常量声名
		e  = 2.7182818
		pi = 3.1415926
	)
	fmt.Printf("%T",tall)

	//const talls = returnint() //这样不行  常量必须在编译之前确定 必须是一个可计算的值 不能是返回值
//变量定义不使用 出错 常量定义不使用 没问题
}


