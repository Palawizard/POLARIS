package character

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func IsDead(player *utils.Player) bool {
	if player.Health <= 0 {
		player.Health = player.MaxHealth / 2
		fmt.Println("Oh no! You DIED! You respawned with", player.Health, "HP.")
		return true
	} else {
		return false
	}
}
