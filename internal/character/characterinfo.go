package character

import (
	"fmt"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/utils"
	"strings"
)

// InitCharacter initializes the player by calling the CharacterCreation function.
// It returns the fully initialized utils.Player.
func InitCharacter() utils.Player {
	p1 := CharacterCreation()
	return p1
}

func CharacterCreation() utils.Player {
	utils.Clearscreen()
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
	money := 100
	skills := map[string]int{"Punch": 1}
	equipment := map[string]int{}
	inventory := map[string]int{"Potion": 3}

	return utils.Player{
		Name:      name,
		Class:     classID,
		Level:     level,
		Money:     money,
		MaxHealth: maxhealth,
		Health:    health,
		Skills:    skills,
		Equipment: equipment,
		Inventory: inventory,
		Equipped:  map[string]string{},
	}
}

func DisplayInfo(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Character Info\n")
	fmt.Println("Name:  ", player.Name)
	fmt.Println("Class: ", ClassLabel(player.Class))
	fmt.Println("Level: ", player.Level)
	fmt.Println("Money: ", player.Money)

	armorBonus := equipement.EquippedBonus(player)
	fmt.Printf("Health: %d/%d (+%d max HP from armor)\n\n", player.Health, player.MaxHealth, armorBonus)

	fmt.Println("Equipped:")
	head := "(none)"
	body := "(none)"
	feet := "(none)"
	if id := player.Equipped["Head"]; id != "" {
		e := equipement.GetEquipment(id)
		head = fmt.Sprintf("%s (+%d)", e.Name, e.Defense)
	}
	if id := player.Equipped["Body"]; id != "" {
		e := equipement.GetEquipment(id)
		body = fmt.Sprintf("%s (+%d)", e.Name, e.Defense)
	}
	if id := player.Equipped["Feet"]; id != "" {
		e := equipement.GetEquipment(id)
		feet = fmt.Sprintf("%s (+%d)", e.Name, e.Defense)
	}
	fmt.Println(" Head:", head)
	fmt.Println(" Body:", body)
	fmt.Println(" Feet:", feet)

	fmt.Println("\n1. Retour")
	var choice int
	fmt.Scan(&choice)
}
