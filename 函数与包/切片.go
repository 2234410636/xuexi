package main

import "fmt"

func main8() {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)

}

func main7() {
	var a = []int{1, 3, 4, 5, 6, 5, 8, 9, 5, 4, 1, 2, 1, 5, 8, 9, 4}
	fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
	b := a[0:1]
	fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
	c := b[0:4:10]
	fmt.Printf("slice c : %v , len(c) : %d\n %d\n", c, len(c), cap(c))
}

func main9() {
	d := [5]struct {
		x int
	}{{60}, {0}, {0}, {15}, {10}}
	s := d[:]
	d[1].x = 10
	s[2].x = 20
	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
}

func main10() {
	s1 := []int{1, 2, 3, 4, 5}
	s := []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	copy(s2, s)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6, 1, 5, 6, 2)
	fmt.Printf("last slice s3 : %v\n", s3)
}

func main11() {
	slice := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[5:8]
	d1 = append(d1, 1, 3)
	fmt.Println(d1, len(d1), cap(d1))
	fmt.Println(slice)
	d2 := slice[:10:10]
	fmt.Println(d2, len(d2), cap(d2))
}

func main12() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:8]
	s2 = append(s2, 100)
	fmt.Println(slice)
	s2 = append(s2, 200)
	fmt.Println(slice)
	s1[2] = 20
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

func main13() {
	var str []string = []string{"tom", "jack", "mary"}
	fmt.Println(cap(str))
	fmt.Println(len(str))
	fmt.Println(str)

}

func main() {
	var array = [...]int{1, 2, 3, 4, 6, 4, 7, 8, 9, 7, 8, 8}
	s1 := array[2:12]
	s2 := array[2:10]
	s3 := array[2:9]
	s4 := array[:]
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(len(s3), cap(s3))
	fmt.Println(len(s4), cap(s4))

}
