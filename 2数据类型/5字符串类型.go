package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var cmd string
	fmt.Scanf("%s", &cmd)
	exec.Command(cmd).Run()
}
