package utils

type Player struct {
	Name      string
	Class     string
	Level     int
	Money     int
	MaxHealth int
	Health    int
	Skills    map[string]int
	Inventory map[string]int
	Equipment map[string]int
}
