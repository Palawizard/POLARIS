package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

func Shop(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Shop")
	fmt.Print("\n\n")
	fmt.Println("=== Bienvenue chez le marchand ===")
	fmt.Println("1. Potion de vie (3$)")
	fmt.Println("2. Potion de poison (6$)")
	fmt.Println("3. Livre de Sort: Boule de feu (25$)")
	fmt.Println("4. Retour")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		if player.Money < 3 {
			fmt.Println("Vous n'avez pas assez d'argent.")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 3
		character.AddInventory(player, "Potion")
		fmt.Println("Vous recevez 1 Potion de heal. Total :", player.Inventory["Potion"])
		time.Sleep(2 * time.Second)
	case 2:
		if player.Money < 6 {
			fmt.Println("Vous n'avez pas assez d'argent.")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 6
		character.AddInventory(player, "Poison")
		fmt.Println("Vous recevez 1 Potion de poison. Total :", player.Inventory["Poison"])
		time.Sleep(2 * time.Second)
	case 3:
		if player.Money < 25 {
			fmt.Println("Vous n'avez pas assez d'argent.")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 25
		skills.SpellBook("Boule de feu", player)
		fmt.Println("Vous recevez 1 Livre de Sort: Boule de feu. Total :", player.Inventory["Boule de feu"])
		time.Sleep(2 * time.Second)
	case 4:
		return
	}
}
