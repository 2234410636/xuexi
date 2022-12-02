package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf(strings.Replace("a aa a a a ", "a", "2", -1))
	fmt.Printf("%q", strings.Split("asd a fgh", " "))
}
