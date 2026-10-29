//go:build js

package utils

import "fmt"

// ClearScreen clears the browser terminal (xterm.js) and its scrollback
// with ANSI sequences, since there is no "clear" command to run.
func ClearScreen() {
	fmt.Print("\x1b[H\x1b[2J\x1b[3J")
}
