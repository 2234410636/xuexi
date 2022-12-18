package main

import (
	"fmt"
	"math/big"
	"time"
)

const LIM1 = 1000        //求第1000位的斐波那契数列
var fibs1 [LIM1]*big.Int //使用数组保存计算出来的数列的指针
func main() {
	result := big.NewInt(0)
	start := time.Now()
	for i := 0; i < LIM1; i++ {
		result = fibonacci1(i)
		fmt.Printf("数列第 %d 位: %d\n", i+1, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行完成，所耗时间为: %s\n", delta)
	a := int64(2100000000000000000)
	a = int64(9223372036854775807)
	fmt.Println(a)
}
func fibonacci1(n int) (res *big.Int) {
	if n <= 1 {
		res = big.NewInt(1)
	} else {
		temp := new(big.Int)
		res = temp.Add(fibs1[n-1], fibs1[n-2])
	}
	fibs1[n] = res
	return
}

//
//const LIM1 = 1000     //求第1000位的斐波那契数列
//var fibs1 [LIM1]int64 //使用数组保存计算出来的数列的指针
//func main() {
//	result := int64(0)
//	start := time.Now()
//	for i := 0; i < LIM1; i++ {
//		result = fibonacci1(i)
//		fmt.Printf("数列第 %d 位: %d\n", i+1, result)
//	}
//	end := time.Now()
//	delta := end.Sub(start)
//	fmt.Printf("执行完成，所耗时间为: %s\n", delta)
//}
//func fibonacci1(n int) (res int64) {
//	if n <= 1 {
//		res = 1
//	} else {
//
//		res = fibs1[n-1] + fibs1[n-2]
//	}
//	fibs1[n] = res
//	return
//}
