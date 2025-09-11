package game

import (
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/menu"
)

func InitGame() {
	c1 := character.InitCharacter("GagaLionlion", "Template", 1, "Coup de point", map[string]int{"Potion": 3})
	menu.ShowMenu(&c1)
}
