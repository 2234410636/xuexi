package main

import (
	"fmt"
)

func main() {
	var a interface{} = 123

	a = fmt.Sprintf("%v", a)
	switch a.(type) {
	case string:
		fmt.Println("str")
	}
	fmt.Println(a)
}
