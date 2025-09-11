package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/utils"
	"time"
)

func Potionshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Shop de potions ===>")
		fmt.Printf("Argent: %d$\n\n", player.Money)

		get := func(k string) int {
			if player.Inventory == nil {
				return 0
			}
			return player.Inventory[k]
		}

		fmt.Printf("1. Potion de heal (3$)     [x%d]\n", get("Potion"))
		fmt.Printf("2. Potion de poison (6$)   [x%d]\n", get("Poison"))
		fmt.Println("3. Retour")

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if !character.CheckInvSize(player) {
				lastMsg = "Votre inventaire est plein."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 3 {
				lastMsg = "Vous n'avez pas assez d'argent"
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 3
			character.AddInventory(player, "Potion")
			lastMsg = fmt.Sprintf("Vous recevez 1 potion de heal, total : %d", player.Inventory["Potion"])
			time.Sleep(1 * time.Second)

		case 2:
			if !character.CheckInvSize(player) {
				lastMsg = "Votre inventaire est plein."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 6 {
				lastMsg = "Vous n'avez pas assez d'argent"
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 6
			character.AddInventory(player, "Poison")
			lastMsg = fmt.Sprintf("Vous recevez 1 potion de poison, total : %d", player.Inventory["Poison"])
			time.Sleep(1 * time.Second)

		case 3:
			return
		}
	}
}
