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

func GetEquipment(id string) Equipment {
	return Equipments[id]
}

func AddEquipment(name string, p *utils.Player) {
	if p.Equipment == nil {
		p.Equipment = make(map[string]int)
	}
	p.Equipment[name]++
}

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
