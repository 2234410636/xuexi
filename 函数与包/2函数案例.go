package main

import "fmt"

const (
	SecondsPerMinute = 60
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerday    = 24 * SecondsPerHour
)

func main() {
	var asd int
	fmt.Scanf("%d", &asd)
	a, b, c := GetTime(asd)
	fmt.Println(a, b, c)
}

func GetTime(seconds int) (day, hour, minute int) {
	minute = seconds / SecondsPerMinute
	hour = seconds / SecondsPerHour
	day = seconds / SecondsPerday
	return
}
