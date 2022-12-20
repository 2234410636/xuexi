package main

import (
	"fmt"
)

func main1() {
	str := [...]string{"asdsaf", "adasdasdas", "sda"}
	for i, v := range str {
		fmt.Println(i, v)
	}

	fmt.Println(str)
}

func main2() {
	a := [...]int{1, 2, 3, 4, 8, 6, 9, 7, 6, 51, 66, 78}
	max := a[0]
	min := a[0]
	sun := 0
	for _, i := range a {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
		sun += i
	}
	fmt.Println(max, min, sun, sun/len(a))
}

func main3() { //冒泡排序
	a := [...]int{1, 2, 3, 4, 8, 6, 9, 7, 6, 51, 666, 78}
	for f := 0; f < len(a)-2; f++ {
		zj := 0
		for i := f + 1; i < len(a)-1; i++ {
			if a[f] > a[i] {
				zj = a[i]
				a[i] = a[f]
				a[f] = zj
			}
		}
	}

	fmt.Println(a)

}

func main() {

	a := [...]int{1, 2, 3, 4, 8, 6, 9, 7, 6, 51, 666, 78}
	mx(a)
	fmt.Println(a)
	fmt.Println(&a[0])

}

func mx(x [12]int) {
	x[0] = 555555
	fmt.Println(x)
	fmt.Println(&x[0])

}
