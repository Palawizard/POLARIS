package menu

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/shop"
	"projet-red_POLARIS/utils"
)

func ShowMenu(player *utils.Player) {
	for {
		utils.Clearscreen()
		fmt.Println("Menu")
		fmt.Print("\n")
		fmt.Println("1. Character Info")
		fmt.Println("2. Inventory")
		fmt.Println("3. Quit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			character.DisplayInfo(player)
		case 2:
			for {
				if character.AccessInventory(player) {
					shop.Shop(player)
					continue
				}
				break
			}
		case 3:
			return
		}
	}
}
