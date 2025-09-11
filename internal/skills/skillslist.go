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
	"Coup de point": {
		ID:    "Coup de point",
		Label: "Coup de point",
		Apply: effectCoupDePoint,
	},
	"Boule de feu": {
		ID:    "Boule de feu",
		Label: "Boule de feu",
		Apply: effectBouleDeFeu,
	},
}

func SpellBook(id string, player *utils.Player) bool {
	it, ok := Skills[id]
	if !ok || it.Apply == nil {
		return false
	}
	it.Apply(player)
	return true
}
