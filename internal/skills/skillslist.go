package skills

import (
	"projet-red_POLARIS/utils"
)

type Skill struct {
	ID    string
	Label string
	Apply func(*utils.Player)
}

var Skills = map[string]Skill{
	"Punch": {
		ID:    "Punch",
		Label: "Punch",
		Apply: effectPunch,
	},
	"Fire Ball": {
		ID:    "Fire Ball",
		Label: "Fire Ball",
		Apply: effectFireball,
	},
}

func SpellBook(id string, player *utils.Player) bool {
	it, ok := Skills[id]
	if !ok || it.Apply == nil {
		return false
	}
	if player.Skills == nil {
		player.Skills = make(map[string]int)
	}
	player.Skills[id]++
	return true
}

func Cast(id string, player *utils.Player) bool {
	it, ok := Skills[id]
	if !ok || it.Apply == nil {
		return false
	}
	it.Apply(player)
	return true
}
