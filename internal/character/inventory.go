package character

import (
	"fmt"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
)

func AccessInventory(player *utils.Player) bool {
	utils.Clearscreen()
	fmt.Println("Inventory")
	fmt.Print("\n")
	if len(player.Inventory) == 0 {
		fmt.Println("(vide)")
	} else {
		for key, value := range player.Inventory {
			fmt.Println(key, ": ", value)
		}
	}
	fmt.Print("\n")
	fmt.Println("1. Retour")
	fmt.Println("2. Utiliser un objet")
	fmt.Println("3. Marchand")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 2:
		useItemMenu(player)
		return false
	case 3:
		return true
	default:
		return false
	}
}

func AddInventory(player *utils.Player, item string) {
	if player == nil || item == "" {
		return
	}
	if player.Inventory == nil {
		player.Inventory = make(map[string]int)
	}
	player.Inventory[item]++
	fmt.Println(item, "added to inventory.")
}

func RemoveInventory(player *utils.Player, item string) {
	if player == nil || item == "" || player.Inventory == nil {
		return
	}
	if qty, ok := player.Inventory[item]; ok {
		if qty <= 1 {
			delete(player.Inventory, item)
		} else {
			player.Inventory[item] = qty - 1
		}
		fmt.Println(item, "removed from inventory.")
	} else {
		fmt.Println(item, "not found in inventory.")
	}
}

func useItemMenu(p *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Utiliser un objet\n")

	type option struct{ id string }
	var opts []option
	for id, qty := range p.Inventory {
		if qty > 0 {
			if _, ok := objects.GetItem(id); ok {
				opts = append(opts, option{id: id})
			}
		}
	}

	if len(opts) == 0 {
		fmt.Println("Aucun objet utilisable.")
		fmt.Println("\n1. Retour")
		var _tmp int
		fmt.Scan(&_tmp)
		return
	}

	for i, o := range opts {
		it, _ := objects.GetItem(o.id)
		fmt.Printf("%d. %s (x%d)\n", i+1, it.Label, p.Inventory[o.id])
	}
	fmt.Println("0. Annuler")

	var idx int
	fmt.Scan(&idx)
	if idx <= 0 || idx > len(opts) {
		return
	}

	id := opts[idx-1].id

	if ok := objects.ApplyItem(id, p); !ok {
		fmt.Println("Impossible d'utiliser cet objet.")
	} else {
		RemoveInventory(p, id)
	}

	fmt.Println("\n1. Retour")
	var _tmp int
	fmt.Scan(&_tmp)
}

func CheckInvSize(player *utils.Player) bool {
	if len(player.Inventory) >= 10 {
		fmt.Println("Votre inventaire est plein.")
		return false
	}
	return true
}
