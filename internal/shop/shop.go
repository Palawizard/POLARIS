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
		fmt.Println("1. Potion Shop")
		fmt.Println("2. Spell Shop")
		fmt.Println("3. Equipment Shop")
		fmt.Println("4. Return")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Potionshop(player)
		case 2:
			Spellshop(player)
		case 3:
			EquipementShop(player)
		case 4:
			return
		}
	}
}
