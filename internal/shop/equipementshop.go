package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/utils"
	"time"
)

func EquipementShop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Shop d'équipement ===>")
		fmt.Printf("Argent: %d$\n\n", player.Money)

		get := func(k string) int {
			if player.Equipment == nil {
				return 0
			}
			return player.Equipment[k]
		}

		fmt.Printf("1. Fourrure de Loup (4$)     [x%d]\n", get("Fourrure de loup"))
		fmt.Printf("2. Peau de Troll (7$)        [x%d]\n", get("Peau de Troll"))
		fmt.Printf("3. Cuir de Sanglier (3$)     [x%d]\n", get("Cuir de Sanglier"))
		fmt.Printf("4. Plume de Corbeau (1$)     [x%d]\n", get("Plume de Corbeau"))
		fmt.Println("5. Retour")

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
			if player.Money < 4 {
				lastMsg = "Vous n'avez pas assez d'argent"
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 4
			equipement.AddEquipment("Fourrure de loup", player)
			lastMsg = fmt.Sprintf("Vous recevez 1 Fourrure de loup, total : %d", player.Equipment["Fourrure de loup"])
			time.Sleep(1 * time.Second)

		case 2:
			if !character.CheckInvSize(player) {
				lastMsg = "Votre inventaire est plein."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 7 {
				lastMsg = "Vous n'avez pas assez d'argent"
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 7
			equipement.AddEquipment("Peau de Troll", player)
			lastMsg = fmt.Sprintf("Vous recevez 1 Peau de Troll, total : %d", player.Equipment["Peau de Troll"])
			time.Sleep(1 * time.Second)

		case 3:
			if !character.CheckInvSize(player) {
				lastMsg = "Votre inventaire est plein."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 3 {
				lastMsg = "Vous n'avez pas assez d'argent"
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 3
			equipement.AddEquipment("Cuir de Sanglier", player)
			lastMsg = fmt.Sprintf("Vous recevez 1 Cuir de Sanglier, total : %d", player.Equipment["Cuir de Sanglier"])
			time.Sleep(1 * time.Second)

		case 4:
			if !character.CheckInvSize(player) {
				lastMsg = "Votre inventaire est plein."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 1 {
				lastMsg = "Vous n'avez pas assez d'argent"
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 1
			equipement.AddEquipment("Plume de Corbeau", player)
			lastMsg = fmt.Sprintf("Vous recevez 1 Plume de Corbeau, total : %d", player.Equipment["Plume de Corbeau"])
			time.Sleep(1 * time.Second)

		case 5:
			return
		}
	}
}
