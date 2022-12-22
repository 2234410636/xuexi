package main

import "fmt"

type add struct {
	id   int
	name string
	age  int
}

func demo(a []add) []add {
	a[1].age = 999
	a = append(a, add{3, "xiaoping", 44})
	fmt.Printf("dome %p %p\n", &a, a)
	fmt.Println(a)
	return a
}

func main() {
	ce := []add{
		{1, "xiaoming", 22},
		{2, "xiaoling", 33},
	}

	fmt.Printf("main %p %p\n", &ce, ce)
	fmt.Println(ce)
	ce = demo(ce)
	fmt.Printf("main %p %p\n", &ce, ce)
	fmt.Println(ce)

}

//输出结果
//main 0xc000004078 0xc000062040
//[{1 xiaoming 22} {2 xiaoling 33}]
//dome 0xc0000040c0 0xc000020080
//[{1 xiaoming 22} {2 xiaoling 999} {3 xiaoping 44}]
//main 0xc000004078 0xc000020080
//[{1 xiaoming 22} {2 xiaoling 999} {3 xiaoping 44}]
