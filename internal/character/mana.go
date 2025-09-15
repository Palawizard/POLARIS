package character

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func UseSkill(player *utils.Player, skillName string) {
	switch skillName {
	case "Fireball":
		manaCost := 20
		if player.Mana >= float64(manaCost) {
			player.Mana -= float64(manaCost)
			fmt.Println(player.Name, "casts Fireball!")
			// appliquer les dégâts...
		} else {
			fmt.Println("Not enough mana!")
		}
	}
}

func ManaRegen(player *utils.Player, turn int) {
	if player.Mana < player.MaxMana && turn%2 == 0 {
		regen := 10.0
		player.Mana += regen
		if player.Mana > player.MaxMana {
			player.Mana = player.MaxMana
		}
		fmt.Printf("%s récupère %.0f mana (%.0f/%.0f)\n", player.Name, regen, player.Mana, player.MaxMana)
	}
}

var SkillManaCost = map[string]float64{
	"Fireball": 20,
	"Punch":    0,
}
