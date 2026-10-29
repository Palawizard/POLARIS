package game

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/chapters"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/menu"
	"projet-red_POLARIS/internal/save"
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
	musicPath := filepath.Join("assets", "audio", "music", "titlescreen.mp3")
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

	// Browser build: offer to resume the saved run, if there is one.
	if d, ok := save.Load(); ok && askContinue(d) {
		chapters.ChangeChapter(d.Chapter)
		p := d.Player
		utils.ClearScreen()
		fmt.Println("Welcome back to Polaris,", p.Name+" !")
		time.Sleep(2 * time.Second)
		audiosystem.StopMusic()
		menu.ShowMenu(&p)
		return
	}

	c1 := character.InitCharacter()

	utils.ClearScreen()
	fmt.Println("Welcome to Polaris !")
	time.Sleep(2 * time.Second)
	audiosystem.StopMusic()
	menu.ShowMenu(&c1)
}

// askContinue shows the "Continue / New game" choice for a saved run.
// Returns true to resume it; a new game erases the save.
func askContinue(d save.Data) bool {
	chapter := fmt.Sprintf("Chapter %d", d.Chapter)
	if d.Chapter > 4 {
		chapter = "Game finished"
	}
	for {
		utils.ClearScreen()
		fmt.Println("A saved game was found:")
		fmt.Printf("  %s, %s, level %d, %s\n\n", d.Player.Name, character.ClassLabel(d.Player.Class), d.Player.Level, chapter)
		fmt.Println("1. Continue")
		fmt.Println("2. New game (erases this save)")

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			continue
		}
		_ = audiosystem.PlaySFXCached("select")
		switch choice {
		case 1:
			return true
		case 2:
			save.Clear()
			return false
		}
	}
}
