package main

import "os/exec"

func main() {

	exec.Command("notepad").Run()
	exec.Command("write").Run()
	exec.Command("calc").Run()

}
