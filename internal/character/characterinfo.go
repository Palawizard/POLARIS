package character

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func InitCharacter() utils.Player {
	p1 := CharacterCreation()
	return p1
}

func CharacterCreation() utils.Player {
	fmt.Println("What is your name?")
	var name string
	fmt.Scan(&name)
	fmt.Println("What class do you want to choose ?")
	var classID string
	fmt.Scan(&classID)

	cls := GetClass(classID)
	maxhealth := cls.MAXHP
	health := cls.HP
	level := 1
	skills := "Coup de point"
	inventory := map[string]int{"Potion": 3}

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
	fmt.Println("Class: ", ClassLabel(player.Class))
	fmt.Println("Level: ", player.Level)
	fmt.Printf("Health: %d/%d\n\n", player.Health, player.MaxHealth)
	fmt.Println("Skills:", player.Skills)

	fmt.Println("\n1. Retour")
	var choice int
	fmt.Scan(&choice)
}
