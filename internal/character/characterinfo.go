package character

import (
	"fmt"
	"math"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/utils"
	"strings"
	"time"
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
	for !namegood {
		fmt.Scan(&name)
		_ = audiosystem.PlaySFXCached("select")

		valid := true
		if len(name) == 0 {
			valid = false
		} else {
			for _, v := range name {
				if v < 'A' || (v > 'Z' && v < 'a') || v > 'z' {
					valid = false
					break
				}
			}
		}

		if !valid {
			fmt.Println("Please enter a valid name")
			continue
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
		_ = audiosystem.PlaySFXCached("select")
		if _, ok := Classes[classID]; ok {
			classgood = true
		} else {
			fmt.Println("Please enter a valid class")
		}
	}

	cls := GetClass(classID)
	maxhealth := cls.MAXHP
	health := cls.HP
	exp := 0.0
	expToNextLevel := 80.0
	level := 1
	money := 0.0
	initiative := 4.0

	maxMana := 30.0
	mana := maxMana
	manaRegen := 0.0
	switch classID {
	case "Dwarf":
		manaRegen = 4.0
	case "Human":
		manaRegen = 5.0
	case "Elf":
		manaRegen = 6.0
	}

	skills := map[string]int{"Punch": 1}
	equipment := map[string]int{}
	inventory := map[string]int{"Potion": 2}
	inventorymax := 10
	inventoryupgradesused := 0

	return utils.Player{
		Name:                  name,
		Class:                 classID,
		EXP:                   exp,
		EXPToNextLevel:        expToNextLevel,
		Level:                 level,
		Money:                 money,
		MaxHealth:             maxhealth,
		Health:                health,
		Skills:                skills,
		Equipment:             equipment,
		Inventory:             inventory,
		InventoryMax:          inventorymax,
		InventoryUpgradesUsed: inventoryupgradesused,
		Equipped:              map[string]string{},
		Initiative:            initiative,
		Mana:                  mana,
		MaxMana:               maxMana,
		ManaRegen:             manaRegen,
	}
}

func DisplayInfo(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Character Info\n")
	fmt.Println("Name:  ", player.Name)
	fmt.Println("Class: ", ClassLabel(player.Class))
	fmt.Println("Level: ", player.Level)
	fmt.Printf("EXP: %.0f / %.0f\n", player.EXP, player.EXPToNextLevel)
	fmt.Println("Money: ", player.Money)
	fmt.Println("Max Inventory Size: ", player.InventoryMax)

	armorBonus := equipement.EquippedBonus(player)
	fmt.Printf("Health: %.0f/%.0f (+%.0f max HP from armor)\n", player.Health, player.MaxHealth, armorBonus)
	fmt.Printf("Mana:   %.0f/%.0f (+%.0f/turn)\n\n", player.Mana, player.MaxMana, player.ManaRegen)

	fmt.Println("Equipped:")
	head := "(none)"
	body := "(none)"
	feet := "(none)"
	if id := player.Equipped["Head"]; id != "" {
		e := equipement.GetEquipment(id)
		head = fmt.Sprintf("%s (+%.0f)", e.Name, e.Defense)
	}
	if id := player.Equipped["Body"]; id != "" {
		e := equipement.GetEquipment(id)
		body = fmt.Sprintf("%s (+%.0f)", e.Name, e.Defense)
	}
	if id := player.Equipped["Feet"]; id != "" {
		e := equipement.GetEquipment(id)
		feet = fmt.Sprintf("%s (+%.0f)", e.Name, e.Defense)
	}
	fmt.Println(" Head:", head)
	fmt.Println(" Body:", body)
	fmt.Println(" Feet:", feet)

	fmt.Println("\n1. Retour")
	var choice int
	fmt.Scan(&choice)
	_ = audiosystem.PlaySFXCached("select")
}

func AddEXP(player *utils.Player, exp float64) {
	exp = math.Round(exp)
	player.EXP += exp
	fmt.Println("You have gained", exp, "EXP.")
	time.Sleep(1 * time.Second)
	for player.EXP >= player.EXPToNextLevel {
		player.EXP -= player.EXPToNextLevel
		player.Level++
		player.EXPToNextLevel = math.Round(player.EXPToNextLevel * 1.15)
		player.MaxHealth = math.Round(player.MaxHealth + 15)
		player.Health = player.MaxHealth
		player.Initiative += 2.0
		player.MaxMana += 5
		player.Mana = player.MaxMana
		utils.Clearscreen()
		fmt.Println("Level up!")
		time.Sleep(1 * time.Second)
		fmt.Println("Your level is now", player.Level)
		fmt.Println("Your max health is now", player.MaxHealth)
		fmt.Println("Your were fully healed.")
		fmt.Println("Your max mana is now", player.MaxMana)
		fmt.Println("Your EXP to next level is now", player.EXPToNextLevel)
		fmt.Println("You got stronger...")
		time.Sleep(4 * time.Second)
	}
}
