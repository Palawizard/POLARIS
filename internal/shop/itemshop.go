package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
	"time"
)

// Itemshop displays the player's current money and the items in the catalog,
// along with their price and the number of items the player has in their
// inventory. The player is prompted to enter the number of the item they wish
// to buy. If the player enters a number that is not in the range of the options,
// or if they do not have enough money, or if their inventory is full, it will
// simply loop back to the start of the menu. If the player chooses to buy an
// item, it will be added to their inventory, the cost of the item will be
// subtracted from their money, and the player will be prompted to enter "1" to
// return.
func Itemshop(player *utils.Player) {
	lastMsg := ""
	catalog := []string{"Potion", "Poison", "Wolf Fur", "Troll Skin", "Boar Leather", "Crow Feather"}

	for {
		utils.Clearscreen()
		fmt.Println("<=== Item Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)

		getInv := func(k string) int {
			if player.Inventory == nil {
				return 0
			}
			return player.Inventory[k]
		}

		for i, id := range catalog {
			it, _ := objects.GetItem(id)
			fmt.Printf("%d. %s (%d coins)     [x%d]\n", i+1, it.Label, it.Price, getInv(id))
		}
		fmt.Printf("%d. Return\n", len(catalog)+1)

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)

		if choice == len(catalog)+1 {
			return
		}
		if choice < 1 || choice > len(catalog) {
			lastMsg = "Invalid choice."
			time.Sleep(1 * time.Second)
			continue
		}

		id := catalog[choice-1]
		it, _ := objects.GetItem(id)

		if !character.CheckInvSize(player) {
			lastMsg = "Your inventory is full."
			time.Sleep(1 * time.Second)
			continue
		}
		if player.Money < it.Price {
			lastMsg = "You do not have enough money."
			time.Sleep(1 * time.Second)
			continue
		}

		player.Money -= it.Price
		character.AddInventory(player, id)
		lastMsg = fmt.Sprintf("You received 1 %s, total : %d", it.Label, player.Inventory[id])
		time.Sleep(1 * time.Second)
	}
}
