package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/utils"
)

func Shop(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Shop")
	fmt.Print("\n\n")
	fmt.Println("=== Bienvenue chez le marchand ===")
	fmt.Println("1. Potion de vie (GRATUITE)")
	fmt.Println("2. Retour")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		character.AddInventory(player, "Potion")
		fmt.Println("Vous recevez 1 Potion. Total :", player.Inventory["Potion"])
	case 2:
		return
	}
}
