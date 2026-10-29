//go:build !js

package utils

import (
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen clears the console on Windows, macOS, and Linux.
func ClearScreen() {
	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/c", "cls")
	} else {
		c = exec.Command("clear")
	}
	c.Stdout = os.Stdout
	_ = c.Run()
}
