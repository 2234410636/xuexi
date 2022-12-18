package main

import (
	"fmt"
	golang_tts "github.com/zhaopuyang/golang-tts"
)

func main1() {
	var age int
	fmt.Scanf("%d", &age)
	if age < 18 {
		fmt.Println("未成年禁止入内")
		golang_tts.SpeakText("未成年禁止入内")

	} else {
		fmt.Println("欢迎光临")
		golang_tts.SpeakText("欢迎光临")
	}
}

func main123() {
	var user, password string
	fmt.Scanf("%s%s", &user, &password)
	if user == "admin" || password == "mypass" {
		fmt.Println("取钱嘛")
		golang_tts.SpeakText("取钱嘛")

	} else {
		fmt.Println("密码错误")
		golang_tts.SpeakText("密码错误")
	}
}

func main() {
	var mo int
	fmt.Scanf("%d", &mo)
	if mo > 3000 {
		fmt.Println("贵宾欢迎你")
	} else if mo > 2000 {
		fmt.Println("一等座欢迎你")
	} else if mo > 1000 {
		fmt.Println("二等座欢迎你")
	} else {
		fmt.Println("你的钱不够")
	}
}
