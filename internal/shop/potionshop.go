package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/utils"
	"time"
)

func Potionshop(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("<=== Shop de potions ===>")
	fmt.Printf("Argent: %d$\n\n", player.Money)

	// compteur sécurisé
	get := func(k string) int {
		if player.Inventory == nil {
			return 0
		}
		return player.Inventory[k]
	}

	fmt.Printf("1. Potion de heal (3$)     [x%d]\n", get("Potion"))
	fmt.Printf("2. Potion de poison (6$)   [x%d]\n", get("Poison"))
	fmt.Println("3. Retour")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if !character.CheckInvSize(player) {
			time.Sleep(2 * time.Second)
			return
		}
		if player.Money < 3 {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 3
		character.AddInventory(player, "Potion")
		fmt.Println("Vous recevez 1 potion de heal, total :", player.Inventory["Potion"])
		time.Sleep(2 * time.Second)

	case 2:
		if !character.CheckInvSize(player) {
			time.Sleep(2 * time.Second)
			return
		}
		if player.Money < 6 {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 6
		character.AddInventory(player, "Poison")
		fmt.Println("Vous recevez 1 potion de poison, total :", player.Inventory["Poison"])
		time.Sleep(2 * time.Second)

	case 3:
		return
	}
}
