package utils

import (
	"fmt"
)

// IsDead checks if the player is dead, and if yes, it "respawns" them with half
// of their max health. It returns true if the player is dead and false
// otherwise.
func IsDead(player *Player) bool {
	if player.Health <= 0 {
		player.Health = player.MaxHealth / 2
		fmt.Println("Oh no! You DIED! You respawned with", player.Health, "HP.")
		return true
	} else {
		return false
	}
}
