package main

import (
	"fmt"
	"strconv"
)

func main() {
	var users = []map[string]string{
		{
			"name": "张三",
			"age":  "18",
		},
		{
			"name": "李四",
			"age":  "22",
		},
		{
			"name": "王五",
			"age":  "20",
		},
		{
			"name": "赵六",
			"age":  "90",
		},
		{
			"name": "孙七",
			"age":  "60",
		},
		{
			"name": "周八",
			"age":  "10",
		},
	}
	c := sum(zhuanhuan(users, mapbool))
	fmt.Println(zhuanhuan(users, mapbool))
	fmt.Println(c)

}

func sum(str []string) int {
	sum1 := 0
	for _, v := range str {
		c, _ := strconv.Atoi(v)
		sum1 += c
	}
	return sum1
}

func mapbool(user map[string]string) bool {
	age, ok := user["age"]
	if !ok {
		return false
	}
	intAge, err := strconv.Atoi(age)
	if err != nil {
		return false
	}
	if intAge < 0 || intAge > 100 {
		return false
	}
	return true

}

func zhuanhuan(user []map[string]string, f func(map[string]string) bool) []string {
	userstr := make([]string, 2)
	for _, v := range user {
		if f(v) {
			userstr = append(userstr, v["age"])

		}
	}
	return userstr
}
