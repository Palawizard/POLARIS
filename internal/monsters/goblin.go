package monsters

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

func InitGoblin() *Monster {
	return New("Goblin")
}

func GoblinPattern(player *utils.Player, monster *Monster, turn int) bool {
	if turn%3 == 0 {
		utils.Clearscreen()
		fmt.Println("The goblin attacks you!")
		time.Sleep(2 * time.Second)
		player.Health -= monster.ATK * 2
		fmt.Println("Critial hit!")
		time.Sleep(1 * time.Second)
		fmt.Println("Ouch ! The goblin dealt", monster.ATK, "damage to", player.Name, "!")
		time.Sleep(1 * time.Second)
		if utils.IsDead(player) {
			fmt.Println("\nYou have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			return true
		}
		fmt.Println("\nYour HP is now at", player.Health, "/", player.MaxHealth, "hp.")
		time.Sleep(2 * time.Second)
	} else {
		utils.Clearscreen()
		fmt.Println("The goblin attacks you!")
		time.Sleep(2 * time.Second)
		player.Health -= monster.ATK
		fmt.Println("Ouch ! The goblin dealt", monster.ATK, "damage to", player.Name, "!")
		time.Sleep(1 * time.Second)
		if utils.IsDead(player) {
			fmt.Println("\nYou have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			return true
		}
		fmt.Println("\nYour HP is now at", player.Health, "/", player.MaxHealth, "hp.")
		time.Sleep(2 * time.Second)
	}
	return true
}
