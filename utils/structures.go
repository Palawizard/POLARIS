package utils

type Player struct {
	Name      string
	Class     string
	Level     int
	Maxhealh  int
	Health    int
	Inventory map[string]int
}
