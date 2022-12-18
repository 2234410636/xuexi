package main

import (
	"fmt"
)

type 学历 int
//设计择偶工具
const (
	幼儿园 学历=iota
	小学
	初中
	高中
	大学
	硕士
	博士

)

func main21() {

	fmt.Println(幼儿园)
	fmt.Println(博士)
	fmt.Println(大学)
	fmt.Println(初中)
	var 我的学历 学历=大学
	fmt.Println("我的学历",我的学历)
}

func main31() {
	const (
		flognone = 1 << iota
		flogred
		floggreen
		flogblue
	)
	fmt.Printf("%d %d %d", flogred, flogblue, floggreen)
}

type ChipType int
const (
	None ChipType = iota
	CPU    // 中央处理器
	GPU    // 图形处理器
)
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}
func main() {
	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU)
	fmt.Printf("%s %d", GPU, GPU)
}








