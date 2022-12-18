package main

import "fmt"

func main()  {
	a,b :=100,10
	fmt.Println(a,b)
	a=a+b  //  a=110
	b=a-b //a+b-b b=100
	a=a-b //a+b-a-b+b  a=10
	fmt.Println(a,b)
	a,b=b,a
	fmt.Println(a,b)

}
