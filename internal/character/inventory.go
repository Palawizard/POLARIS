package character

import (
	"fmt"
	"projet-red_POLARIS/internal/shop"
	"projet-red_POLARIS/utils"
)

func AccessInventory(player utils.Player) {
	utils.Clearscreen()
	fmt.Println("Inventory")
	fmt.Print("\n")
	for key, value := range player.Inventory {
		fmt.Println(key, ": ", value)
	}
	fmt.Print("\n")
	fmt.Println("1. Marchand")
	fmt.Println("2. Retour")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		shop.Shop(player)
	case 2:
		return
	}
}
