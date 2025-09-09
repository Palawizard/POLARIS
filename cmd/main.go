package main

import (
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/menu"
)

func main() {
	c1 := character.InitCharacter("Palawi", "Elfe", 1, 100, 40, map[string]int{"Potion": 3})
	menu.ShowMenu(c1)
}
