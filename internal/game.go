package game

import (
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/menu"
)

func InitGame() {
	c1 := character.InitCharacter("Palawi", "Elfe", 1, 100, 40, map[string]int{"Potion": 3})
	menu.ShowMenu(c1)
}
