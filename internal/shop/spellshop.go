package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

func Spellshop(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("<=== Shop de sorts ===>")
	fmt.Println("1. Livre de Sort: Boule de feu (25$)")
	fmt.Println("2. Retour")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		if player.Money < 25 {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 25
		skills.SpellBook("Boule de feu", player)
		fmt.Println("Vous recevez 1 Livre de Sort: Boule de feu, total :", player.Skills["Boule de feu"])
		time.Sleep(2 * time.Second)
	case 2:
		return
	}
}
