package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/utils"
	"time"
)

func EquipementShop(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("<=== Shop d'équipement ===>")
	fmt.Println("1. Fourrure de Loup (4$)")
	fmt.Println("2. Peau de Troll (7$)")
	fmt.Println("3. Cuir de Sanglier (3$)")
	fmt.Println("4. Plume de Corbeau (1$)")
	fmt.Println("5. Retour")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		if player.Money < 4 {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 4
		equipement.AddEquipment("Fourrure de loup", player)
		fmt.Println("Vous recevez 1 fourrure de loup, total :", player.Equipment["Fourrure de loup"])
		time.Sleep(2 * time.Second)
	case 2:
		if player.Money < 7 {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 7
		equipement.AddEquipment("Peau de Troll", player)
		fmt.Println("Vous recevez 1 Peau de Troll, total :", player.Equipment["Peau de Troll"])
		time.Sleep(2 * time.Second)
	case 3:
		if player.Money < 3 {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 3
		equipement.AddEquipment("Cuir de Sanglier", player)
		fmt.Println("Vous recevez 1 Cuir de Sanglier, total :", player.Equipment["Cuir de Sanglier"])
		time.Sleep(2 * time.Second)
	case 4:
		if player.Money < 1 {
			fmt.Println("Vous n'avez pas assez d'argent")
			time.Sleep(2 * time.Second)
			return
		}
		player.Money -= 1
		equipement.AddEquipment("Plume de Corbeau", player)
		fmt.Println("Vous recevez 1 Plume de Corbeau, total :", player.Equipment["Plume de Corbeau"])
		time.Sleep(2 * time.Second)
	case 5:
		return
	}
}
