package fightsystem

import (
	"fmt"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

func TrainingFight(player *utils.Player) {
	turn := 0
	for {
		goblin := monsters.InitGoblin()
		for TurnMenu(player, &goblin) {
			if utils.IsDead(player) {
				fmt.Println("You have been defeated by the goblin.")
				time.Sleep(3 * time.Second)
				return
			}
			if goblin.Health <= 0 {
				fmt.Println("You have defeated the goblin.")
				time.Sleep(3 * time.Second)
				return
			}
		}
		for monsters.GoblinPattern(player, goblin, turn) {
			if utils.IsDead(player) {
				fmt.Println("You have been defeated by the goblin.")
				time.Sleep(3 * time.Second)
				return
			}
			if goblin.Health <= 0 {
				fmt.Println("You have defeated the goblin.")
				time.Sleep(3 * time.Second)
				return
			}
		}
	}
}
