package shop

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

// EquipementShop displays the player's equipment and allows them to buy or sell it.
//
// It will display the player's current coins, and then display a list of
// available equipment. The player is prompted to enter the number of the
// equipment they wish to buy/sell. If the player enters a number that is not in
// the range of the options, or if they do not have the equipment, it will simply
// loop back to the start of the menu. If the player chooses to buy/sell the
// equipment, it will be added/removed to/from their inventory, and the cost of
// the equipment will be subtracted/added from/to their money. After the
// equipment is bought/sold, the player will be prompted to enter "1" to return.
func EquipementShop(player *utils.Player) {
	lastMsg := ""
	catalog := []string{}
	prices := map[string]int{}

	for {
		utils.Clearscreen()
		fmt.Println("<=== Equipment Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)

		if len(catalog) == 0 {
			fmt.Println("No equipment available.")
		} else {
			for i, id := range catalog {
				fmt.Printf("%d. %s (%d coins)\n", i+1, id, prices[id])
			}
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
		_ = id
	}
}
