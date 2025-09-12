package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

// Itemshop displays the player's coins, and a list of items they can purchase
// with their coins. The player is prompted to enter the number of the item
// they wish to purchase. If the player enters a number that is not in the
// range of the options, or if they do not have enough money, it will print
// an error message and loop back to the start of the menu. If the player
// chooses to purchase an item, it will be added to their inventory, and the
// cost will be subtracted from their money. After the item is purchased, the
// player will be prompted to enter "1" to return to the previous menu.
func Itemshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Item Shop ===>")
		fmt.Printf("Coins: %.0f\n\n", player.Money)

		catalog := make([]string, 0, len(objects.Items))
		for id := range objects.Items {
			catalog = append(catalog, id)
		}
		sort.Slice(catalog, func(i, j int) bool {
			return objects.Items[catalog[i]].Label < objects.Items[catalog[j]].Label
		})

		getInv := func(k string) int {
			if player.Inventory == nil {
				return 0
			}
			return player.Inventory[k]
		}

		for i, id := range catalog {
			it := objects.Items[id]
			fmt.Printf("%d. %s (%.0f coins)     [x%d]\n", i+1, it.Label, it.Price, getInv(id))
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
		it := objects.Items[id]

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
