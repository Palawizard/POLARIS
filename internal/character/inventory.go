package character

import (
	"fmt"
	"projet-red_POLARIS/internal/shop"
	"projet-red_POLARIS/utils"
)

func AccessInventory(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Inventory")
	fmt.Print("\n")
	for key, value := range player.Inventory {
		fmt.Println(key, ": ", value)
	}
	fmt.Print("\n")
	fmt.Println("4. Marchand")
	fmt.Println("3. Ajouter un objet")
	fmt.Println("2. Retirer un objet")
	fmt.Println("1. Retour")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 4:
		shop.Shop(player)
	case 3:
		AddInventory(player)
	case 2:
		RemoveInventory(player)
	case 1:
		return
	}
}

func AddInventory(player *utils.Player) {
	fmt.Print("Nom de l'objet à ajouter : ")
	var item string
	fmt.Scan(&item)
	if player.Inventory == nil {
		player.Inventory = make(map[string]int)
	}
	player.Inventory[item]++
	fmt.Println("Objet ajouté !")
}

func RemoveInventory(player *utils.Player) {
	fmt.Print("Nom de l'objet à retirer : ")
	var item string
	fmt.Scan(&item)

	if count, ok := player.Inventory[item]; ok {
		if count <= 1 {
			delete(player.Inventory, item)
		} else {
			player.Inventory[item] = count - 1
		}
		fmt.Println("Objet retiré !")
	} else {
		fmt.Println("Objet introuvable.")
	}
}
