package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"runtime"
	"strings"
	"time"
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

// ShowText renders text with a simple typewriter effect and short SFX.
// Blocks until the player presses Enter.
func ShowText(text string) {
	for _, r := range text {
		fmt.Print(string(r))
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "text.wav"))
		time.Sleep(20 * time.Millisecond)
	}
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "text_end.wav"))
	fmt.Print("\n\nPress Enter to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

// PrintASCII trims common leading indentation and prints multi-line ASCII art.
func PrintASCII(art string) {
	lines := strings.Split(art, "\n")
	if len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}
	if n := len(lines); n > 0 && strings.TrimSpace(lines[n-1]) == "" {
		lines = lines[:n-1]
	}
	indent := -1
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		n := 0
		for _, r := range l {
			if r == ' ' || r == '\t' {
				n++
			} else {
				break
			}
		}
		if indent == -1 || n < indent {
			indent = n
		}
	}
	if indent > 0 {
		for i, l := range lines {
			if len(l) >= indent {
				lines[i] = l[indent:]
			}
		}
	}
	fmt.Println(strings.Join(lines, "\n"))
}

// Flash performs a brief screen shake + flash effect.
// ms controls total duration; amp scales the shake intensity.
func Flash(ms int, amp int) {
	if ms <= 0 {
		ms = 120
	}
	if amp < 1 {
		amp = 1
	}
	frames := 3
	step := time.Duration(ms/(frames*4)) * time.Millisecond

	// Save cursor & hide for the effect; restore on exit.
	fmt.Print("\x1b7\x1b[?25l")
	defer fmt.Print("\x1b8\x1b[?25h")

	// Jitter cycle with a quick invert-flash (DEC private mode 5).
	for i := 0; i < frames; i++ {
		for a := 0; a < amp; a++ {
			fmt.Print("\x1b[1C") // right
			time.Sleep(step)
			fmt.Print("\x1b[1D") // left
			time.Sleep(step)
			fmt.Print("\x1b[1B") // down
			time.Sleep(step)
			fmt.Print("\x1b[1A") // up
			time.Sleep(step / 2)

			// brief flash
			fmt.Print("\x1b[?5h")
			time.Sleep(step / 2)
			fmt.Print("\x1b[?5l")
		}
	}
}
