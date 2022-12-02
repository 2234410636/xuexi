package main

import "fmt"

/*
// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(f func(int) int, list int) {
	start := time.Now()
	f(list)
	fmt.Println(time.Since(start))
}
func main() {
	// 使用匿名函数打印切片内容
	visit(func(x int) (sum int) {
		for i := 0; i < x; i++ {
			sum += i
		}
		return sum
	}, 10000000000)
}
*/

//func test() func() int {
//	var x int
//	return func() int {
//		x++
//
//		return x
//	}
//}
//func main123() {
//
//	t := test()
//	fmt.Println(t())
//	fmt.Println(t())
//	fmt.Println(t())
//	fmt.Println(t())
//	fmt.Println(t())
//	y := test()
//	fmt.Println(t(), y())
//	fmt.Println(&t, &y, test())
//}

func wjm(hzm string) func(string) string {
	return func(str string) string {
		var s string
		if str == hzm {
			s = str
			asdd = "123"
			fmt.Println(asdd)
		} else {
			s = ".jpg"
			asdd = "123"
			fmt.Println(asdd)
		}

		return s

	}
}

var asdd string = "asda"

func main() {
	s := wjm("jpg")
	fmt.Printf(s("asdq"))
	fmt.Printf(s("asdqads"))
	fmt.Printf(s(".jpga"))
	fmt.Println(asdd)
}
