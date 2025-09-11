package character

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func InitCharacter(name string, classID string, level int, skills string, inventory map[string]int) utils.Player {
	cls := GetClass(classID)
	maxhealth := cls.MAXHP
	health := cls.HP

	return utils.Player{
		Name:      name,
		Class:     classID,
		Level:     level,
		MaxHealth: maxhealth,
		Health:    health,
		Skills:    skills,
		Inventory: inventory,
	}
}

func DisplayInfo(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Character Info\n")
	fmt.Println("Name:  ", player.Name)
	fmt.Println("Class: ", ClassLabel(player.Class)) // <- label humain
	fmt.Println("Level: ", player.Level)
	fmt.Printf("Health: %d/%d\n\n", player.Health, player.MaxHealth)
	fmt.Println("Skills:", player.Skills)

	fmt.Println("\n1. Retour")
	var choice int
	fmt.Scan(&choice)
}
