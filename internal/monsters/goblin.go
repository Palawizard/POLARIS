package monsters

import "projet-red_POLARIS/utils"

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
