package skills

import "projet-red_POLARIS/utils"

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
}
