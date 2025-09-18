package character

import (
	"fmt"
	"math"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/equipment"
	"projet-red_POLARIS/utils"
	"strings"
	"time"
)

// InitCharacter asks the player to create a character and returns the result.
func InitCharacter() utils.Player {
	p := CharacterCreation()
	return p
}

// CharacterCreation handles name entry, class selection, and initial stats.
// Light validation on the name keeps things tidy (letters only).
func CharacterCreation() utils.Player {
	utils.ClearScreen()

	nameOK := false
	classOK := false
	var name string

	fmt.Println("What is your name?")
	for !nameOK {
		fmt.Scan(&name)
		_ = audiosystem.PlaySFXCached("select")

		// Simple alphabetical check; avoids weird input and keeps display clean.
		valid := true
		if len(name) == 0 {
			valid = false
		} else {
			for _, r := range name {
				if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
					valid = false
					break
				}
			}
		}
		if !valid {
			fmt.Println("Please enter a valid name")
			continue
		}
		nameOK = true
	}

	// Normalize casing: capitalize first letter, keep the rest as typed.
	lowerName := strings.ToLower(name)
	runes := []rune(lowerName)
	newName := ""
	for i := 0; i < len(runes); i++ {
		if i == 0 {
			newName += strings.ToUpper(string(runes[i]))
			continue
		}
		newName += string(runes[i])
	}
	name = newName

	fmt.Println("What class do you want to choose ? You can choose between : ")
	fmt.Println(Classlist())

	var classID string
	for classOK == false {
		fmt.Scan(&classID)
		_ = audiosystem.PlaySFXCached("select")
		if _, ok := Classes[classID]; ok {
			classOK = true
		} else {
			fmt.Println("Please enter a valid class")
		}
	}

	cls := GetClass(classID)

	// Base progression & economy pacing.
	maxHealth := cls.MaxHP
	health := cls.HP
	exp := 0.0
	expToNextLevel := 80.0
	level := 1
	money := 0.0
	initiative := 4.0

	// Mana pool & class regen feel. Tweakable knobs for pacing.
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

	// Starter kit: one basic skill, empty gear, a couple of potions.
	skills := map[string]int{"Punch": 1}
	equipment := map[string]int{}
	inventory := map[string]int{"Potion": 2}
	inventoryMax := 10
	inventoryUpgradesUsed := 0

	return utils.Player{
		Name:                  name,
		Class:                 classID,
		EXP:                   exp,
		EXPToNextLevel:        expToNextLevel,
		Level:                 level,
		Money:                 money,
		MaxHealth:             maxHealth,
		Health:                health,
		Skills:                skills,
		Equipment:             equipment,
		Inventory:             inventory,
		InventoryMax:          inventoryMax,
		InventoryUpgradesUsed: inventoryUpgradesUsed,
		Equipped:              map[string]string{},
		Initiative:            initiative,
		Mana:                  mana,
		MaxMana:               maxMana,
		ManaRegen:             manaRegen,
	}
}

// DisplayInfo presents a compact snapshot of the character's state,
// including equipped pieces and their effective bonuses.
func DisplayInfo(player *utils.Player) {
	utils.ClearScreen()
	fmt.Println("Character Info\n")
	fmt.Println("Name:  ", player.Name)
	fmt.Println("Class: ", ClassLabel(player.Class))
	fmt.Println("Level: ", player.Level)
	fmt.Printf("EXP: %.0f / %.0f\n", player.EXP, player.EXPToNextLevel)
	fmt.Println("Money: ", player.Money)
	fmt.Println("Max Inventory Size: ", player.InventoryMax)

	armorBonus := equipment.EquippedBonus(player)
	fmt.Printf("Health: %.0f/%.0f (+%.0f max HP from armor)\n", player.Health, player.MaxHealth, armorBonus)
	fmt.Printf("Mana:   %.0f/%.0f (+%.0f/turn)\n\n", player.Mana, player.MaxMana, player.ManaRegen)

	// Show which items occupy each slot, if any.
	fmt.Println("Equipped:")
	head := "(none)"
	body := "(none)"
	feet := "(none)"
	if id := player.Equipped["Head"]; id != "" {
		e := equipment.GetEquipment(id)
		head = fmt.Sprintf("%s (+%.0f)", e.Name, e.Defense)
	}
	if id := player.Equipped["Body"]; id != "" {
		e := equipment.GetEquipment(id)
		body = fmt.Sprintf("%s (+%.0f)", e.Name, e.Defense)
	}
	if id := player.Equipped["Feet"]; id != "" {
		e := equipment.GetEquipment(id)
		feet = fmt.Sprintf("%s (+%.0f)", e.Name, e.Defense)
	}
	fmt.Println(" Head:", head)
	fmt.Println(" Body:", body)
	fmt.Println(" Feet:", feet)

	fmt.Println("\n0. Return")
	var choice int
	fmt.Scan(&choice)
	_ = audiosystem.PlaySFXCached("select")
}

// AddEXP grants experience and handles level-ups in-place.
// Rounds values to keep numbers neat and avoids fractional drift.
func AddEXP(player *utils.Player, exp float64) {
	exp = math.Round(exp)
	player.EXP += exp
	fmt.Println("You have gained", exp, "EXP.")
	time.Sleep(1 * time.Second)

	for player.EXP >= player.EXPToNextLevel {
		// Spend the bar, climb a level, and refresh core resources.
		player.EXP -= player.EXPToNextLevel
		player.Level++
		player.EXPToNextLevel = math.Round(player.EXPToNextLevel * 1.15)

		// Boosts on level: more health, refill, slight initiative bump, more mana.
		player.MaxHealth = math.Round(player.MaxHealth + 15)
		player.Health = player.MaxHealth
		player.Initiative += 2.0
		player.MaxMana += 5
		player.Mana = player.MaxMana

		// Small ceremony—clear the screen and show the upgrades cleanly.
		utils.ClearScreen()
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
