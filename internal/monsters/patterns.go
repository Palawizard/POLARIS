package monsters

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

func AttackPattern(player *utils.Player, monster *Monster, turn int) bool {
	utils.Clearscreen()
	fmt.Println("Turn", turn)
	utils.SendTurn(turn)
	PrintHeader(monster)
	fmt.Println("\n")

	msg := monster.AttackMsg
	if msg == "" {
		msg = fmt.Sprintf("The %s attacks you!", monster.Name)
	}
	fmt.Println(msg)
	time.Sleep(2 * time.Second)

	dmg := monster.ATK
	if monster.CritEvery > 0 && turn%monster.CritEvery == 0 {
		mult := monster.CritMultiplier
		if mult <= 0 {
			mult = 2
		}
		dmg *= mult
		fmt.Println("Critical hit!")
		time.Sleep(1 * time.Second)
	}

	player.Health -= dmg
	if player.Health < 0 {
		player.Health = 0
	}

	fmt.Printf("Ouch! %s deals %.0f damage to %s!\n", monster.Name, dmg, player.Name)
	time.Sleep(1 * time.Second)

	if utils.IsDead(player) {
		fmt.Printf("\nYou have been defeated by the %s.\n", monster.Name)
		time.Sleep(3 * time.Second)
		return true
	}

	fmt.Printf("\nYour HP is now %.0f / %.0f hp.\n", player.Health, player.MaxHealth)
	time.Sleep(2 * time.Second)
	return true
}
