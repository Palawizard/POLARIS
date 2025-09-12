package monsters

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

func InitGoblin() utils.Monster {
	name := "Goblin"
	atk := 5
	hp := 40
	maxhp := 40

	return utils.Monster{
		Name:      name,
		Health:    hp,
		MaxHealth: maxhp,
		ATK:       atk,
	}
}

func GoblinPattern(player *utils.Player, monster utils.Monster, turn int) {
	if turn%3 == 0 {
		fmt.Println("The goblin attacks you!")
		time.Sleep(2 * time.Second)
		player.Health -= monster.ATK * 2
		fmt.Println("Critial hit!")
		time.Sleep(1 * time.Second)
		fmt.Println("Ouch ! The goblin dealt", monster.ATK, "damage to", player.Name, "!")
		time.Sleep(1 * time.Second)
		if utils.IsDead(player) {
			fmt.Println("You have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			return
		}
		fmt.Println("Your HP is now at", player.Health, "/", player.MaxHealth, "hp.")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("The goblin attacks you!")
		time.Sleep(2 * time.Second)
		player.Health -= monster.ATK
		fmt.Println("Ouch ! The goblin dealt", monster.ATK, "damage to", player.Name, "!")
		time.Sleep(1 * time.Second)
		if utils.IsDead(player) {
			fmt.Println("You have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			return
		}
		fmt.Println("Your HP is now at", player.Health, "/", player.MaxHealth, "hp.")
		time.Sleep(2 * time.Second)
	}
}
