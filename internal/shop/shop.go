package shop

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

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
