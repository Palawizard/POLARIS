package character

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func InitCharacter(name string, class string, level int, maxhealth int, health int, inventory map[string]int) utils.Player {
	return utils.Player{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHealth: maxhealth,
		Health:    health,
		Inventory: inventory,
	}
}

func DisplayInfo(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Character Info")
	fmt.Print("\n")
	fmt.Println("Name:  ", player.Name)
	fmt.Println("Class: ", player.Class)
	fmt.Println("Level: ", player.Level)
	fmt.Printf("Health: %d/%d\n\n", player.Health, player.MaxHealth)

	fmt.Println("1. Retour")
	var choice int
	fmt.Scan(&choice)
}
