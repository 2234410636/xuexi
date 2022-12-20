package main

import (
	"fmt"
)

func main5() {
	var mystr [26]byte
	for i := 0; i < 26; i++ {
		mystr[i] = 'A' + byte(i)
		fmt.Printf("%c\n", mystr[i])
	}
}

func main6() {
	var erfen [1000]int
	for i := 0; i < 1000; i++ {
		erfen[i] = 2 * i

	}
	var aaa int
	fmt.Println("请输入要找的数")
	fmt.Scanln(&aaa)

}
func BinaryFind(arr *[1000]int, leftIndex int, rightIndex int, findVal int) {

	//判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {

		fmt.Println("找不到")
		return

	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2

	if (*arr)[middle] > findVal {
		//说明我们要查找的数，应该在        leftIndex --- middel-1
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//说明我们要查找的数，应该在        middel+1 --- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了，下标为%v \n", middle)

	}
}

func main() {

	var erfen [1000]int
	for i := 0; i < 1000; i++ {
		erfen[i] = 2 * i

	}
	var aaa int
	fmt.Println("请输入要找的数")
	fmt.Scanln(&aaa)

	//测试一把
	BinaryFind(&erfen, 0, len(erfen)-1, aaa)

}
