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
	"Adventurer's Hat":   {ID: "Adventurer's Hat", Name: "Adventurer's Hat", Type: "Head", Price: 5, Defense: 1},
	"Adventurer's Tunic": {ID: "Adventurer's Tunic", Name: "Adventurer's Tunic", Type: "Body", Price: 5, Defense: 2},
	"Adventurer's Boots": {ID: "Adventurer's Boots", Name: "Adventurer's Boots", Type: "Feet", Price: 5, Defense: 1},
}

var Recipes = map[string]map[string]int{
	"Adventurer's Hat":   {"Crow Feather": 1, "Boar Leather": 1},
	"Adventurer's Tunic": {"Wolf Fur": 2, "Troll Skin": 1},
	"Adventurer's Boots": {"Wolf Fur": 1, "Boar Leather": 1},
}

func GetEquipment(id string) Equipment { return Equipments[id] }

// AddEquipment adds the given equipment to the player's equipment list, incrementing its count by 1.
// If the player's equipment list is currently nil, it will be initialized.
func AddEquipment(name string, p *utils.Player) {
	if p.Equipment == nil {
		p.Equipment = make(map[string]int)
	}
	p.Equipment[name]++
}

// RemoveEquipment removes one instance of the given equipment from the player's
// equipment list. If the player's equipment list is currently nil, it does
// nothing. If the equipment is not found in the player's equipment list, it does
// nothing. Otherwise, it decrements the count of the equipment in the player's
// equipment list by one. If the count is zero after decrementing, it removes the
// equipment from the player's equipment list.
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
