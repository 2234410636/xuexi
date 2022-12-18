package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main1() {
	r := regexp.MustCompile("abc")
	fmt.Println(r.MatchString("abc"))
	fmt.Println(r.MatchString("abcd"))
	fmt.Println(r.MatchString("avcde"))

}
func main2() {
	r := regexp.MustCompile("^-?[0-9]+$")
	fmt.Println(r.MatchString("-0"))

}

func main3() {
	buf := "abc azc a7c aac 888 a9c tac"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile("a.c")
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func main4() {
	buf := "abc azc a7c aac 888 a9c tac"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile("(?:a|9|z)c")
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

func main5() {
	buf := "abc azc a7c aac 888 a9c tac Abc Azc A7c AAc 888 A9c tAc"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile("(?i:a)bc")

	fmt.Println(reg1.FindAllStringSubmatch(buf, -1))
}

func main6() {
	buf := "abc azc a7c aac 888 a9c tac Abc Azc A7c AAc 888 A9c tAc"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`\QA\E`)

	fmt.Println(reg1.FindAllStringSubmatch(buf, -1))
}

func main() {
	buf := "abc azc a7c aac 888 a9c tac Abc Azc A7c AAc 888 A9c tAc"
	reg1 := regexp.MustCompile(`\QA\E`)
	fmt.Println(reg1.Find([]byte(buf)))
	fmt.Println(reg1.FindString(buf))
	fmt.Println(reg1.FindAll([]byte(buf), -1))
	fmt.Println(reg1.FindAllString(buf, -1))
	fmt.Println(reg1.FindIndex([]byte(buf)))
	fmt.Println(reg1.FindSubmatchIndex([]byte(buf)))
	fmt.Println(reg1.FindStringIndex(buf))
	r := bytes.NewReader([]byte(buf))
	fmt.Println(reg1.FindReaderIndex(r))

}
