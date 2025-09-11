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

// SpellBook adds the given skill to the player's spellbook, incrementing its count by 1.
// If the player's spellbook is currently nil, it will be initialized.
// If the skill doesn't exist or its Apply func is nil, it returns false.
// Otherwise, it adds the skill to the player's spellbook and returns true.
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

// Cast casts the given skill
// If the skill doesn't exist or its Apply func is nil, it returns false.
// Otherwise, it casts the skill and returns true.
func Cast(id string, player *utils.Player) bool {
	it, ok := Skills[id]
	if !ok || it.Apply == nil {
		return false
	}
	it.Apply(player)
	return true
}
