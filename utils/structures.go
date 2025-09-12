package utils

type Player struct {
	Name                  string
	Class                 string
	Level                 int
	Money                 int
	MaxHealth             int
	Health                int
	Skills                map[string]int
	Inventory             map[string]int
	InventoryMax          int
	InventoryUpgradesUsed int
	Equipment             map[string]int
	Equipped              map[string]string
}

type Monster struct {
	Name      string
	Health    int
	MaxHealth int
	ATK       int
}
