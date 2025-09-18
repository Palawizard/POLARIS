package game

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/menu"
	"projet-red_POLARIS/utils"
	"time"
)

// InitGame boots the title screen, music, and then enters the main menu.
func InitGame() {
	utils.ClearScreen()

	// Title music (loops while the splash anim plays).
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "titlescreen.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}

	// Simple splash animation.
	fmt.Println("\n")
	fmt.Println("░▒▓███████▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓███████▓▒░")
	fmt.Println("░▒▓█▓▒░      ")
	fmt.Println("░▒▓█▓▒░      ")
	fmt.Println("░▒▓█▓▒░       ")
	fmt.Println("\n\n\n")
	time.Sleep(150 * time.Millisecond)
	utils.ClearScreen()
	fmt.Println("\n")
	fmt.Println("░▒▓███████▓▒░ ░▒▓██████▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░       ░▒▓██████▓▒░")
	fmt.Println("\n\n\n")
	time.Sleep(150 * time.Millisecond)
	utils.ClearScreen()
	fmt.Println("\n")
	fmt.Println("░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓████████▓▒░")
	fmt.Println("\n\n\n")
	time.Sleep(150 * time.Millisecond)
	utils.ClearScreen()
	fmt.Println("\n")
	fmt.Println("░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░       ░▒▓██████▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓████████▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("\n\n\n")
	time.Sleep(150 * time.Millisecond)
	utils.ClearScreen()
	fmt.Println("\n")
	fmt.Println("░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓███████▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓████████▓▒░▒▓███████▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░")
	fmt.Println("\n\n\n")
	time.Sleep(150 * time.Millisecond)
	utils.ClearScreen()
	fmt.Println("\n")
	fmt.Println("░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓███████▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓████████▓▒░▒▓███████▓▒░░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░")
	fmt.Println("\n\n\n")
	time.Sleep(150 * time.Millisecond)
	utils.ClearScreen()
	fmt.Println("\n")
	fmt.Println("░▒▓███████▓▒░ ░▒▓██████▓▒░░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓███████▓▒░░▒▓█▓▒░░▒▓███████▓▒░ ")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░        ")
	fmt.Println("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░        ")
	fmt.Println("░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓████████▓▒░▒▓███████▓▒░░▒▓█▓▒░░▒▓██████▓▒░  ")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░ ")
	fmt.Println("░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░ ")
	fmt.Println("░▒▓█▓▒░       ░▒▓██████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓███████▓▒░  ")
	fmt.Println("\n\n")
	time.Sleep(150 * time.Millisecond)
	fmt.Println("Press Enter to START !")

	// Wait for Enter, then transition to character creation.
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	_ = audiosystem.PlaySFXCached("select")

	c1 := character.InitCharacter()

	utils.ClearScreen()
	fmt.Println("Welcome to Polaris !")
	time.Sleep(2 * time.Second)
	audiosystem.StopMusic()
	menu.ShowMenu(&c1)
}
