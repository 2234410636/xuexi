package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {
	rand.Seed(time.Now().Unix())
	for {
		n := rand.Int() % 100
		if n != 99 {

			fmt.Println(n)
		} else {
			fmt.Println(n)
			break
		}
	}

}
func main() {
	jihui := 3
asd:
	for {
		fmt.Println("请输入用户名和密码")
		for ; jihui > 0; jihui-- {
			var user string
			var password string
			fmt.Scanf("%s%s\n", &user, &password)
			if user == "admin" || password == "123456" {
				fmt.Println("欢迎光临")
				break asd
			} else {
				fmt.Printf("密码错误请重输，您还有%d次机会\n", jihui-1)
			}
		}
		fmt.Println("机会用完请30秒后再试")
		time.Sleep(time.Second * 30)
		jihui = 3

	}
}
