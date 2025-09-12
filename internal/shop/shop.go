package shop

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

// Shop is a menu allowing the player to access the different shops.
// It will display the player's current coins, and then display a list of shops
// available for purchase. The player is prompted to enter the number of the
// shop they wish to access. If the player enters a number that is not in the
// range of the options, it will simply loop back to the start of the menu.
// If the player chooses to access a shop, it will be run in a loop until the
// player chooses to return.
func Shop(player *utils.Player) {
	for {
		utils.Clearscreen()
		fmt.Println("Shop")
		fmt.Print("\n\n")
		fmt.Println("<=== Welcome to the Shop ===>")
		fmt.Println("1. Item Shop")
		fmt.Println("2. Spell Shop")
		fmt.Println("3. Inventory Shop")
		fmt.Println("4. Return")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Itemshop(player)
		case 2:
			Spellshop(player)
		case 3:
			Inventoryshop(player)
		case 4:
			return
		}
	}
}
