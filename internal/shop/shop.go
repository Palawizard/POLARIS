package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/utils"
	"time"
)

func Shop(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Shop")
	fmt.Print("\n\n")
	fmt.Println("=== Bienvenue chez le marchand ===")
	fmt.Println("1. Potion de vie (GRATUITE)")
	fmt.Println("2. Potion de poison (10$)")
	fmt.Println("3. Retour")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		character.AddInventory(player, "Potion")
		fmt.Println("Vous recevez 1 Potion de heal. Total :", player.Inventory["Potion"])
		time.Sleep(2 * time.Second)
	case 2:
		character.AddInventory(player, "Poison")
		fmt.Println("Vous recevez 1 Potion de poison. Total :", player.Inventory["Poison"])
		time.Sleep(2 * time.Second)
	case 3:
		return
	}
}
