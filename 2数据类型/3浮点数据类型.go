package main

import "fmt"

func main13() {
	var tall float32 = 1.68
	fmt.Printf("身高是%f\n", tall)

	var money float32 = 4.2e10
	fmt.Printf("余额是%f", money)
}

func main() {

	var myfloat01 float32 = 1.11111111e10
	var myfloat02 float32 = 1.00000000187
	fmt.Printf("myfloat: %f\n", myfloat01)
	fmt.Printf("myfloat: %f\n", myfloat02)
	fmt.Println(myfloat02 == myfloat01) //精度有限 32 精度六位

	var myfloat03 float64 = 1.745645645645645645677e20
	var myfloat04 float64 = 1.00000000187
	fmt.Printf("myfloat: %f\n", myfloat03)
	fmt.Printf("myfloat: %f\n", myfloat04)
	fmt.Println(myfloat03 == myfloat04) //精度有限 64 精度14位

}
