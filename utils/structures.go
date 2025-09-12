package utils

type Player struct {
	Name                  string
	Class                 string
	EXP                   float64
	EXPToNextLevel        float64
	Level                 int
	Money                 float64
	MaxHealth             float64
	Health                float64
	Skills                map[string]int
	Inventory             map[string]int
	InventoryMax          int
	InventoryUpgradesUsed int
	Equipment             map[string]int
	Equipped              map[string]string
}
