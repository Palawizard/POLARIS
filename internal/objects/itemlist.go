package objects

import "projet-red_POLARIS/utils"

type Item struct {
	ID    string
	Label string
	Price int
	Apply func(*utils.Player)
}

var Items = map[string]Item{
	"Potion": {
		ID:    "Potion",
		Label: "Potion de vie",
		Price: 0,
		Apply: effectHealthPotion,
	},
	"Poison": {
		ID:    "Poison",
		Label: "Potion de poison",
		Price: 10,
		Apply: effectPoisonPotion,
	},
}

func GetItem(id string) (Item, bool) {
	it, ok := Items[id]
	return it, ok
}

func ApplyItem(id string, p *utils.Player) bool {
	it, ok := Items[id]
	if !ok || it.Apply == nil {
		return false
	}
	it.Apply(p)
	return true
}
