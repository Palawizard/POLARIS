package character

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"strings"
)

func InitCharacter() utils.Player {
	p1 := CharacterCreation()
	return p1
}

func CharacterCreation() utils.Player {
	namegood := false
	classgood := false
	var name string
	fmt.Println("What is your name?")
	for namegood == false {
		fmt.Scan(&name)
		rsn := []rune(name)
		for _, v := range rsn {
			if v < 'A' || (v > 'Z' && v < 'a') || v > 'z' {
				fmt.Println("Please enter a valid name")
				break
			}
		}
		namegood = true
	}
	lowername := strings.ToLower(name)
	rsln := []rune(lowername)
	newname := ""
	for i := 0; i < len(rsln); i++ {
		if i == 0 {
			newname += strings.ToUpper(string(rsln[i]))
			continue
		}
		newname += string(rsln[i])
	}
	name = newname

	fmt.Println("What class do you want to choose ? You can choose between : ")
	fmt.Println(Classlist())
	var classID string
	for classgood == false {
		fmt.Scan(&classID)
		if _, ok := Classes[classID]; ok {
			classgood = true
		} else {
			fmt.Println("Please enter a valid class")
		}
	}

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
