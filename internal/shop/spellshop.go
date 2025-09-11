package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

func Spellshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Shop de sorts ===>")
		fmt.Printf("Argent: %d$\n\n", player.Money)

		ownedFire := 0
		if player.Skills != nil {
			ownedFire = player.Skills["Boule de feu"]
		}

		fmt.Printf("1. Livre de Sort: Boule de feu (25$)  [possédé: x%d]\n", ownedFire)
		fmt.Println("2. Retour")

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if !character.CheckInvSize(player) {
				lastMsg = "Votre inventaire est plein."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 25 {
				lastMsg = "Vous n'avez pas assez d'argent"
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 25
			if skills.SpellBook("Boule de feu", player) {
				lastMsg = fmt.Sprintf("Vous recevez 1 Livre de Sort: Boule de feu, total : %d", player.Skills["Boule de feu"])
			} else {
				lastMsg = "Achat impossible."
			}
			time.Sleep(1 * time.Second)
		case 2:
			return
		}
	}
}
