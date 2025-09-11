package objects

import "projet-red_POLARIS/utils"

type Item struct {
	ID    string
	Label string
	Price int
	Apply func(*utils.Player)
}

var Items = map[string]Item{
	"Potion":       {ID: "Potion", Label: "Healing Potion", Price: 3, Apply: effectHealthPotion},
	"Poison":       {ID: "Poison", Label: "Poisoning Potion", Price: 6, Apply: effectPoisonPotion},
	"Wolf Fur":     {ID: "Wolf Fur", Label: "Wolf Fur", Price: 4, Apply: nil},
	"Troll Skin":   {ID: "Troll Skin", Label: "Troll Skin", Price: 7, Apply: nil},
	"Boar Leather": {ID: "Boar Leather", Label: "Boar Leather", Price: 3, Apply: nil},
	"Crow Feather": {ID: "Crow Feather", Label: "Crow Feather", Price: 1, Apply: nil},
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
