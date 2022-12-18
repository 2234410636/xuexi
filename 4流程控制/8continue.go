package main

import (
	"fmt"
)

func main12() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func main() {
	//for i := 1; i <= 100; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	var i int
	var i1 int
	var i2 int
	for true {
		fmt.Println("请输入一个整数，0为退出")
		fmt.Scanf("%d\n", &i)
		fmt.Printf("%d\n", i)
		if i > 0 {
			i1++
			fmt.Println("正数有%d个", i1)
			fmt.Println("负数有%d个", i2)
			continue
		} else if i < 0 {
			i2++
			fmt.Println("负数有%d个", i2)
			fmt.Println("正数有%d个", i1)
			continue
		} else {
			break
		}

	}

}
