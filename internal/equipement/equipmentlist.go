package equipement

import "projet-red_POLARIS/utils"

// Equipment represents a craftable/equippable piece of gear.
type Equipment struct {
	ID      string  // unique id
	Name    string  // display name
	Type    string  // slot: Head, Body, Feet
	Price   float64 // coin cost to craft
	Defense float64 // max HP bonus
}

// Equipments lists all known equipment.
var Equipments = map[string]Equipment{
	"Adventurer's Hat":   {ID: "Adventurer's Hat", Name: "Adventurer's Hat", Type: "Head", Price: 5, Defense: 6},
	"Adventurer's Tunic": {ID: "Adventurer's Tunic", Name: "Adventurer's Tunic", Type: "Body", Price: 5, Defense: 18},
	"Adventurer's Boots": {ID: "Adventurer's Boots", Name: "Adventurer's Boots", Type: "Feet", Price: 5, Defense: 12},

	"Leather Cap":   {ID: "Leather Cap", Name: "Leather Cap", Type: "Head", Price: 6, Defense: 16},
	"Leather Armor": {ID: "Leather Armor", Name: "Leather Armor", Type: "Body", Price: 12, Defense: 40},
	"Leather Boots": {ID: "Leather Boots", Name: "Leather Boots", Type: "Feet", Price: 8, Defense: 22},

	"Iron Helm":    {ID: "Iron Helm", Name: "Iron Helm", Type: "Head", Price: 12, Defense: 24},
	"Iron Plate":   {ID: "Iron Plate", Name: "Iron Plate", Type: "Body", Price: 22, Defense: 64},
	"Iron Greaves": {ID: "Iron Greaves", Name: "Iron Greaves", Type: "Feet", Price: 14, Defense: 32},
}

// Recipes lists required materials per equipment.
var Recipes = map[string]map[string]int{
	"Adventurer's Hat":   {"Crow Feather": 1, "Boar Leather": 1},
	"Adventurer's Tunic": {"Wolf Fur": 2, "Troll Skin": 1},
	"Adventurer's Boots": {"Wolf Fur": 1, "Boar Leather": 1},

	"Leather Cap":   {"Wolf Fur": 1, "Boar Leather": 1, "Crow Feather": 2},
	"Leather Armor": {"Wolf Fur": 3, "Boar Leather": 3, "Troll Skin": 1, "Crow Feather": 2},
	"Leather Boots": {"Wolf Fur": 2, "Boar Leather": 2, "Crow Feather": 1},

	"Iron Helm":    {"Troll Skin": 2, "Wolf Fur": 2, "Crow Feather": 3},
	"Iron Plate":   {"Troll Skin": 5, "Wolf Fur": 4, "Boar Leather": 3},
	"Iron Greaves": {"Troll Skin": 3, "Wolf Fur": 3, "Crow Feather": 4},
}

// GetEquipment returns the Equipment by id (zero value if unknown).
func GetEquipment(id string) Equipment { return Equipments[id] }

// AddEquipment increments the owned count of an equipment for the player.
func AddEquipment(name string, p *utils.Player) {
	if p.Equipment == nil {
		p.Equipment = make(map[string]int)
	}
	p.Equipment[name]++
}

// BonusOf returns the max HP bonus for a given equipment id.
func BonusOf(id string) float64 {
	if e, ok := Equipments[id]; ok {
		return e.Defense
	}
	return 0
}

// SlotOf returns the slot (Head/Body/Feet) for a given equipment id.
func SlotOf(id string) string {
	if e, ok := Equipments[id]; ok {
		return e.Type
	}
	return ""
}
