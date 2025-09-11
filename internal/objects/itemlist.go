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
		Label: "Healing Potion",
		Price: 0,
		Apply: effectHealthPotion,
	},
	"Poison": {
		ID:    "Poison",
		Label: "Potion of poison",
		Price: 10,
		Apply: effectPoisonPotion,
	},
}

// GetItem returns an Item from the given id string.
// If the Item doesn't exist, it returns an empty Item and false.
func GetItem(id string) (Item, bool) {
	it, ok := Items[id]
	return it, ok
}

// ApplyItem applies the given item to the player.
// If the item doesn't exist or its Apply func is nil, it returns false.
// Otherwise, it applies the item and returns true.
func ApplyItem(id string, p *utils.Player) bool {
	it, ok := Items[id]
	if !ok || it.Apply == nil {
		return false
	}
	it.Apply(p)
	return true
}
