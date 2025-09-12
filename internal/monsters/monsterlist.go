package monsters

type Monster struct {
	Name       string
	Health     float64
	MaxHealth  float64
	ATK        float64
	EXPtogive  float64
	Initiative float64
}

var Monsters = map[string]Monster{
	"Goblin": {
		Name:       "Goblin",
		Health:     40,
		MaxHealth:  40,
		ATK:        5,
		EXPtogive:  10,
		Initiative: 5,
	},
}

func New(id string) *Monster {
	m := Monsters[id]
	return &m
}
