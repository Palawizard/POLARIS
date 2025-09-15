package fightsystem

import (
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func TrainingFight(player *utils.Player) {
	goblin := monsters.New("Goblin")
	RunFight(player, goblin, false)
}
