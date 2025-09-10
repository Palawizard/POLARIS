package objects

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func TakePotion(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("You take a potion.")

	if n := player.Inventory["Potion"]; n > 0 {
		player.Inventory["Potion"] = n - 1
	}

	player.Health += 50
	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}

	fmt.Println("You now have ", player.Inventory["Potion"], " potions.")
	fmt.Println("You now have ", player.Health, "/", player.MaxHealth, " health.")
	fmt.Print("\n\n")
	fmt.Println("1. Retour")

	var choice int
	fmt.Scan(&choice)
}
