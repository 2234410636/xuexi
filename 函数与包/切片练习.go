package main

import (
	"fmt"
)

func main111() {
	fb(1000)
}

func fb(n int) []uint64 {
	fb1 := make([]uint64, n)
	fb1[1] = 1
	fb1[0] = 1

	for i := 2; i < n; i++ {
		fb1[i] = fb1[i-1] + fb1[i-2]
	}
	fmt.Println(fb1)
	return fb1

}
