package fightsystem

import (
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

// TrainingFight runs a short practice encounter against a standard Goblin.
// Useful for quick balance checks without advancing chapters.
func TrainingFight(player *utils.Player) {
	goblin := monsters.New("Goblin")
	RunFight(player, goblin, false)
}
