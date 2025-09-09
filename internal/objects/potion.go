package objects

import (
	"fmt"
	"projet-red_POLARIS/internal/menu"
	"projet-red_POLARIS/utils"
)

func TakePotion(player utils.Player) {
	utils.Clearscreen()
	fmt.Println("You take a potion.")
	player.Inventory["Potion"]--
	player.Health += 50
	if player.Health > player.Maxhealh {
		player.Health = player.Maxhealh
	}
	fmt.Println("You now have ", player.Inventory["Potion"], " potions.")
	fmt.Println("You now have ", player.Health, "/", player.Maxhealh, " health.")
	fmt.Print("\n\n")
	fmt.Println("1. Retour")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		menu.ShowMenu(player)
	}
}
