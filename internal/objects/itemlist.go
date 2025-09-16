package objects

import "projet-red_POLARIS/utils"

type Item struct {
	ID    string
	Label string
	Price float64
	Apply func(*utils.Player)
}

var Items = map[string]Item{
	"Potion":       {ID: "Potion", Label: "Healing Potion", Price: 3, Apply: effectHealthPotion},
	"Poison":       {ID: "Poison", Label: "Poisoning Potion", Price: 6, Apply: effectPoisonPotion},
	"Wolf Fur":     {ID: "Wolf Fur", Label: "Wolf Fur", Price: 4, Apply: nil},
	"Troll Skin":   {ID: "Troll Skin", Label: "Troll Skin", Price: 7, Apply: nil},
	"Boar Leather": {ID: "Boar Leather", Label: "Boar Leather", Price: 3, Apply: nil},
	"Crow Feather": {ID: "Crow Feather", Label: "Crow Feather", Price: 1, Apply: nil},
	"Chocolatine":  {ID: "Chocolatine", Label: "Chocolatine", Price: 40, Apply: effectChocolatine},
	"Bandage":      {ID: "Bandage", Label: "Bandage", Price: 2, Apply: effectBandage},
	"Hi-Potion":    {ID: "Hi-Potion", Label: "Hi-Potion", Price: 8, Apply: effectHiPotion},
	"Elixir":       {ID: "Elixir", Label: "Elixir", Price: 25, Apply: effectElixir},
	"Grandma Soup": {ID: "Grandma Soup", Label: "Grandma Soup", Price: 6, Apply: effectGrandmaSoup},
	"Regen Potion": {ID: "Regen Potion", Label: "Regeneration Potion", Price: 10, Apply: effectRegenPotion},
	"Baguette":     {ID: "Baguette", Label: "Curative Baguette", Price: 5, Apply: effectBaguette},
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
