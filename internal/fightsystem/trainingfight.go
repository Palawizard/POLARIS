package fightsystem

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

func TrainingFight(player *utils.Player) {
	turn := 1
	goblin := monsters.New("Goblin")
	firstturn := true
	for {
		if player.Initiative < goblin.Initiative && firstturn {
			monsters.AttackPattern(player, goblin, turn)
			firstturn = false

			if utils.IsDead(player) {
				fmt.Println("You have been defeated by the goblin.")
				time.Sleep(3 * time.Second)
				return
			}
			if goblin.Health <= 0 {
				fmt.Println("You have defeated the goblin.")
				player.Money += goblin.Coinstogive
				fmt.Println("You received", goblin.Coinstogive, "coins.")
				time.Sleep(1 * time.Second)
				character.AddEXP(player, goblin.EXPtogive)
				time.Sleep(3 * time.Second)
				return
			}
		}
		if exit := TurnMenu(player, goblin, turn); exit {
			return
		}
		if goblin.Health <= 0 {
			fmt.Println("You have defeated the goblin.")
			player.Money += goblin.Coinstogive
			fmt.Println("You received", goblin.Coinstogive, "coins.")
			time.Sleep(1 * time.Second)
			character.AddEXP(player, goblin.EXPtogive)
			time.Sleep(3 * time.Second)
			return
		}
		if utils.IsDead(player) {
			fmt.Println("You have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			return
		}

		monsters.AttackPattern(player, goblin, turn)
		turn++

		if utils.IsDead(player) {
			fmt.Println("You have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			return
		}
		if goblin.Health <= 0 {
			fmt.Println("You have defeated the goblin.")
			player.Money += goblin.Coinstogive
			fmt.Println("You received", goblin.Coinstogive, "coins.")
			time.Sleep(1 * time.Second)
			character.AddEXP(player, goblin.EXPtogive)
			time.Sleep(3 * time.Second)
			return
		}
	}
}
