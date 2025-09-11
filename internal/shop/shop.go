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
		fmt.Println("<=== Bienvenue chez le marchand ===>")
		fmt.Println("1. Shop de potions")
		fmt.Println("2. Shop de sort")
		fmt.Println("3. Shop d'équipement")
		fmt.Println("4. Retour")

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
