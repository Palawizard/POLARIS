package utils

type Player struct {
	Name      string
	Class     string
	Level     int
	Money     int
	MaxHealth int
	Health    int
	Skills    string
	Inventory map[string]int
}
