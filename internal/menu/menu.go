package menu

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/internal/shop"
	"projet-red_POLARIS/utils"
)

// ShowMenu is the main menu for the game. It displays a menu with three
// options: "Character Info", "Inventory", and "Quit". If the player chooses
// "Character Info", it displays the player's character info. If the player
// chooses "Inventory", it displays the player's inventory and allows them to
// shop. If the player chooses "Quit", the function will return.
func ShowMenu(player *utils.Player) {
	for {
		utils.Clearscreen()
		fmt.Println("Menu")
		fmt.Print("\n")
		fmt.Println("1. Character Info")
		fmt.Println("2. Inventory")
		fmt.Println("3. Black-Smith")
		fmt.Println("4. Quit")

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
			equipement.BlackSmith(player)
		case 4:
			return
		}
	}
}
