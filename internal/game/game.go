package game

import (
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/menu"
)

// InitGame initializes the game by creating a character and starting the main
// game loop, which allows the player to access the menu.
func InitGame() {
	c1 := character.InitCharacter()
	menu.ShowMenu(&c1)
}
