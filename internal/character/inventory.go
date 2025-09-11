package character

import (
	"fmt"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
)

// AccessInventory displays the player's inventory and allows them to use an
// object or shop. It will return true if the player chooses to shop, and false
// otherwise.
func AccessInventory(player *utils.Player) bool {
	utils.Clearscreen()
	fmt.Println("Inventory")
	fmt.Print("\n")
	if len(player.Inventory) == 0 {
		fmt.Println("(empty)")
	} else {
		for key, value := range player.Inventory {
			fmt.Println(key, ": ", value)
		}
	}
	fmt.Print("\n")
	fmt.Println("1. Return")
	fmt.Println("2. Use an object")
	fmt.Println("3. Shop")

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

// AddInventory adds the given item to the player's inventory, incrementing its
// count by 1. If the player's inventory is currently nil, it will be initialized.
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

// RemoveInventory removes one instance of the given item from the player's
// inventory.
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

// useItemMenu displays the player's usable objects and allows them to use one.
// It will print out the player's usable objects, and then prompt them to enter
// the number of the object they wish to use. If the player enters a number that
// is not in the range of the options, or if they do not have the object, it will
// simply return. If the player chooses to use the object, it will be removed
// from their inventory, and the Apply function of the object will be called on
// the player. After the object is used, the player will be prompted to enter
// "1" to return.
func useItemMenu(p *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Use an object\n")

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
		fmt.Println("No usable objects")
		fmt.Println("\n1. Return")
		var _tmp int
		fmt.Scan(&_tmp)
		return
	}

	for i, o := range opts {
		it, _ := objects.GetItem(o.id)
		fmt.Printf("%d. %s (x%d)\n", i+1, it.Label, p.Inventory[o.id])
	}
	fmt.Println("0. Cancel")

	var idx int
	fmt.Scan(&idx)
	if idx <= 0 || idx > len(opts) {
		return
	}

	id := opts[idx-1].id

	if ok := objects.ApplyItem(id, p); !ok {
		fmt.Println("You can't use this object.")
	} else {
		RemoveInventory(p, id)
	}

	fmt.Println("\n1. Return")
	var _tmp int
	fmt.Scan(&_tmp)
}

// CheckInvSize returns true if the given player's inventory has less than 10
// items, and false otherwise. If the player's inventory is full, it will print
// a message to the console.
func CheckInvSize(player *utils.Player) bool {
	if len(player.Inventory) >= 10 {
		fmt.Println("Your inventory is full.")
		return false
	}
	return true
}
