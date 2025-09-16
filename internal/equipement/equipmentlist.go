package equipement

import "projet-red_POLARIS/utils"

type Equipment struct {
	ID      string
	Name    string
	Type    string
	Price   float64
	Defense float64
}

var Equipments = map[string]Equipment{
	"Adventurer's Hat":   {ID: "Adventurer's Hat", Name: "Adventurer's Hat", Type: "Head", Price: 5, Defense: 6},
	"Adventurer's Tunic": {ID: "Adventurer's Tunic", Name: "Adventurer's Tunic", Type: "Body", Price: 5, Defense: 18},
	"Adventurer's Boots": {ID: "Adventurer's Boots", Name: "Adventurer's Boots", Type: "Feet", Price: 5, Defense: 12},
}

var Recipes = map[string]map[string]int{
	"Adventurer's Hat":   {"Crow Feather": 1, "Boar Leather": 1},
	"Adventurer's Tunic": {"Wolf Fur": 2, "Troll Skin": 1},
	"Adventurer's Boots": {"Wolf Fur": 1, "Boar Leather": 1},
}

func GetEquipment(id string) Equipment { return Equipments[id] }

func AddEquipment(name string, p *utils.Player) {
	if p.Equipment == nil {
		p.Equipment = make(map[string]int)
	}
	p.Equipment[name]++
}

func BonusOf(id string) float64 {
	if e, ok := Equipments[id]; ok {
		return e.Defense
	}
	return 0
}

func SlotOf(id string) string {
	if e, ok := Equipments[id]; ok {
		return e.Type
	}
	return ""
}
