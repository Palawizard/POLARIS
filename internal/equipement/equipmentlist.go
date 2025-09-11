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

func AddEquipment(id string, player *utils.Player) bool {
	_, ok := Equipments[id]
	if !ok {
		return false
	}
	return true
}
