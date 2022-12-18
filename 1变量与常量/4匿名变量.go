package main

import "fmt"

func main1()  {
	a,b,c,_:=1,2,3,4//_匿名变量 不可以使用 抵消无用的东西
	fmt.Println(a,b,c)

}

func returngo() (int,int,int,int) {

	return 1,2,3,4

}

func main()  {
	a,b,c,_ := returngo()

	fmt.Println(a,b,c)

}