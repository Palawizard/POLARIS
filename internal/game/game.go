package game

import (
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/menu"
)

func InitGame() {
	c1 := character.InitCharacter()
	menu.ShowMenu(&c1)
}
