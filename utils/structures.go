package utils

type Player struct {
	Name      string
	Class     string
	Level     int
	MaxHealth int
	Health    int
	Inventory map[string]int
}
