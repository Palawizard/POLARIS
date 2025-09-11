package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func Clearscreen() { // Pris sur internet
	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/c", "cls")
	} else {
		c = exec.Command("clear")
	}
	c.Stdout = os.Stdout
	_ = c.Run()
}

var Money = 0
