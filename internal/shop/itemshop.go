package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/utils"
	"time"
)

// Potionshop displays a menu allowing the player to buy potions. It will
// display the player's current coins, and then display a list of potions
// available for purchase. The player is prompted to enter the number of the
// potion they wish to purchase. If the player enters a number that is not
// in the range of the options, or if they do not have enough coins, it will
// print out an error message and then loop back to the start of the menu.
// If the player chooses to purchase a potion, it will be added to their
// inventory and the cost will be deducted from their coins. After the
// potion is purchased, the player will be prompted to enter "1" to return
// to the previous menu.
func Potionshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Item Shop ===>")
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
