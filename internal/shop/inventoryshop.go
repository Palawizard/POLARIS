package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

func Inventoryshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Inventory Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)
		fmt.Println("1. Upgrade inventory slot")
		fmt.Println("2. Return")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if player.Money >= 30 {
				if character.UpgradeInventorySlot(player) {
					player.Money -= 30
					lastMsg = "Inventory slot upgraded."
				} else {
					lastMsg = "Upgrade limit reached!"
				}
				time.Sleep(1 * time.Second)
				continue
			}
			lastMsg = "You do not have enough coins."
			time.Sleep(1 * time.Second)
			continue
		case 2:
			return
		}
	}
}
