package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/utils"
	"time"
)

func EquipementShop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Equipment Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)

		get := func(k string) int {
			if player.Equipment == nil {
				return 0
			}
			return player.Equipment[k]
		}

		fmt.Printf("1. Wolf Fur (4 coins)     [x%d]\n", get("Wolf Fur"))
		fmt.Printf("2. Troll Skin (7 coins)        [x%d]\n", get("Troll Skin"))
		fmt.Printf("3. Boar Leather (3 coins)     [x%d]\n", get("Boar Leather"))
		fmt.Printf("4. Crow Feather (1 coin)     [x%d]\n", get("Crow Feather"))
		fmt.Println("5. Return")

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
			if player.Money < 4 {
				lastMsg = "You do no have enough money."
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 4
			equipement.AddEquipment("Wolf Fur", player)
			lastMsg = fmt.Sprintf("You received 1 Wolf Fur, total : %d", player.Equipment["Wolf Fur"])
			time.Sleep(1 * time.Second)

		case 2:
			if !character.CheckInvSize(player) {
				lastMsg = "Your inventory is full."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 7 {
				lastMsg = "You do no have enough money."
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 7
			equipement.AddEquipment("Troll Skin", player)
			lastMsg = fmt.Sprintf("You received 1 Troll Skin, total : %d", player.Equipment["Troll Skin"])
			time.Sleep(1 * time.Second)

		case 3:
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
			equipement.AddEquipment("Boar Leather", player)
			lastMsg = fmt.Sprintf("You received 1 Boar Leather, total : %d", player.Equipment["Boar Leather"])
			time.Sleep(1 * time.Second)

		case 4:
			if !character.CheckInvSize(player) {
				lastMsg = "Your inventory is full."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 1 {
				lastMsg = "You do no have enough money."
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 1
			equipement.AddEquipment("Crow Feather", player)
			lastMsg = fmt.Sprintf("You received 1 Crow Feather, total : %d", player.Equipment["Crow Feather"])
			time.Sleep(1 * time.Second)

		case 5:
			return
		}
	}
}
