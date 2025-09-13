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

// InitGame initializes the game by creating a character and starting the main
// game loop, which allows the player to access the menu.
func InitGame() {
	utils.Clearscreen()

	//Plays the main menu music
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "titlescreen.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}

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
	utils.Clearscreen()
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
	utils.Clearscreen()
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
	utils.Clearscreen()
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
	utils.Clearscreen()
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
	utils.Clearscreen()
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
	utils.Clearscreen()
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

	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	c1 := character.InitCharacter()

	utils.Clearscreen()
	fmt.Println("Welcome to Polaris !")
	time.Sleep(2 * time.Second)
	audiosystem.StopMusic()
	menu.ShowMenu(&c1)
}
