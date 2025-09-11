package equipement

import "projet-red_POLARIS/utils"

type Equipment struct {
	ID      string
	Name    string
	Type    string
	Price   int
	Defense int
}

var Equipments = map[string]Equipment{
	"Wolf Fur": {
		ID:      "Wolf Fur",
		Name:    "Wolf Fur",
		Type:    "Fur",
		Price:   4,
		Defense: 0,
	},
	"Troll Skin": {
		ID:      "Troll Skin",
		Name:    "Troll Skin",
		Type:    "Skin",
		Price:   7,
		Defense: 0,
	},
	"Boar Leather": {
		ID:      "Boar Leather",
		Name:    "Boar Leather",
		Type:    "Leather",
		Price:   3,
		Defense: 0,
	},
	"Crow Feather": {
		ID:      "Crow Feather",
		Name:    "Crow Feather",
		Type:    "Feather",
		Price:   1,
		Defense: 0,
	},
}

// GetEquipment returns an Equipment from the Equipments map.
// If the Equipment doesn't exist, it returns an empty Equipment.
func GetEquipment(id string) Equipment {
	return Equipments[id]
}

// AddEquipment adds the given equipment to the player's inventory, incrementing its
// count by 1. If the player's inventory is currently nil, it will be initialized.
func AddEquipment(name string, p *utils.Player) {
	if p.Equipment == nil {
		p.Equipment = make(map[string]int)
	}
	p.Equipment[name]++
}

// RemoveEquipment removes one instance of the given equipment from the player's
// inventory.
//
// If the equipment doesn't exist in the player's inventory, it does nothing.
// If the equipment is the last one in the player's inventory, it is removed
// from the inventory.
// Otherwise, it decrements the count of the equipment in the player's inventory
// by 1.
func RemoveEquipment(name string, p *utils.Player) {
	if p.Equipment == nil {
		return
	}
	if q, ok := p.Equipment[name]; ok {
		if q <= 1 {
			delete(p.Equipment, name)
		} else {
			p.Equipment[name] = q - 1
		}
	}
}
