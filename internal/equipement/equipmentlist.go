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
	"Fourrure de loup": {
		ID:      "Fourrure de loup",
		Name:    "Fourrure de loup",
		Type:    "Fourrure",
		Price:   4,
		Defense: 0,
	},
	"Peau de Troll": {
		ID:      "Peau de Troll",
		Name:    "Peau de Troll",
		Type:    "Peau",
		Price:   7,
		Defense: 0,
	},
	"Cuir de Sanglier": {
		ID:      "Cuir de Sanglier",
		Name:    "Cuir de Sanglier",
		Type:    "Cuir",
		Price:   3,
		Defense: 0,
	},
	"Plume de Corbeau": {
		ID:      "Plume de Corbeau",
		Name:    "Plume de Corbeau",
		Type:    "Plume",
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
