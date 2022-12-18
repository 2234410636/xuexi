package main

import "fmt"

//func main1(){
//	var N int
//	N = 10
//	fmt.Printf("我有%d个女朋友",N)//%d 整数
//}
//
//
//func main2(){
//	var 王大锤女朋友数量 int = 100//默认值0
//	//王大锤女朋友数量 = 10 //变量赋值 见名知意
//	var 王大锤女朋友数量 = 10 // 自动推断 根据赋值的默认值推断类型
//	fmt.Printf("王大锤有%d个女朋友",王大锤女朋友数量)
//}

//func main(){
//	 王大锤女朋友数量  := 100//修改默认值位100 :=初始化
//	王大锤女朋友数量 = 1000
//	fmt.Printf("王大锤有%d个女朋友",王大锤女朋友数量)
//
//}

func main(){
	/*
		var girlFriend1  string= "红红"
		var girlFriend2  string= "蓝蓝"
		var girlFriend3  string= "绿绿"

	*/

	/*
	var (
		girlFriend1  string= "红红"
		girlFriend2  string= "蓝蓝"
		girlFriend3  string= "绿绿"
			)
	*/


		var(
			girlFriend1  = "红红"
			girlFriend2  = "蓝蓝"
			girlFriend3  = "绿绿"
		)



	fmt.Printf(girlFriend1+girlFriend2+girlFriend3)
}