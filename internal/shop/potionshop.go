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
		fmt.Println("<=== Potion Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)

		get := func(k string) int {
			if player.Inventory == nil {
				return 0
			}
			return player.Inventory[k]
		}

		fmt.Printf("1. Healing Potion (3 coins)     [x%d]\n", get("Potion"))
		fmt.Printf("2. Poisoning Potion (6 coins)   [x%d]\n", get("Poison"))
		fmt.Println("3. Return")

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if !character.CheckInvSize(player) {
				lastMsg = "Your inventory is full."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 3 {
				lastMsg = "You do no have enough money."
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 3
			character.AddInventory(player, "Potion")
			lastMsg = fmt.Sprintf("You received 1 Healing Potion, total : %d", player.Inventory["Potion"])
			time.Sleep(1 * time.Second)

		case 2:
			if !character.CheckInvSize(player) {
				lastMsg = "Your inventory is full."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 6 {
				lastMsg = "You do no have enough money."
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 6
			character.AddInventory(player, "Poison")
			lastMsg = fmt.Sprintf("You received 1 Poisoning Potion, total : %d", player.Inventory["Poison"])
			time.Sleep(1 * time.Second)

		case 3:
			return
		}
	}
}
