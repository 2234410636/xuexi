package main

import "fmt"

type Student struct {
	id      int
	name    string
	chima   int
	age     int
	leixing string
}

//func printValue(stu Student) {
//	stu.id = 250
//	//printValue stu =  {250 Luffy 109 18 s EastSea}
//	fmt.Println("printValue stu = ", stu)
//}
//
//func main() {
//
//	var s Student = Student{1, "Luffy", 'm', 18, "EastSea"}
//
//
//	printValue(s)        //值传递，形参修改不会影响到实参值
//	fmt.Println("main s = ", s) //main s =  {1 Luffy 109 18 EastSea}
//}
func printPointer(p *Student) {
	p.id = 250
	//printPointer p =  &{250 Luffy 109 18 EastSea}
	fmt.Println("printPointer p = ", p.id)
}

func main() {
	var s Student = Student{1, "Luffy", 'm', 18, "EastSea"}
	fmt.Println("main s = ", s)
	printPointer(&s) //传引用(地址)，形参修改会影响到实参值
	//main s = {250 Luffy 109 18 EastSea}
}
