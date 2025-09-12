package equipement

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"sort"
	"strings"
	"time"
)

// BlackSmith displays a menu with all the craftable equipment, and prompts the user
// to enter the number of the equipment they wish to craft. If the user does not have
// enough coins, or if they do not have all the required resources, it will display
// an error message. If the user chooses to craft the equipment, it will be added
// to their inventory, and the required resources and coins will be removed.
//
// The menu will loop until the user chooses to return.
func BlackSmith(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Black-Smith ===>")
		fmt.Printf("Coins: %.0f\n\n", player.Money)
		fmt.Println("Crafting requires resources:")

		catalog := make([]string, 0, len(Equipments))
		for id := range Equipments {
			catalog = append(catalog, id)
		}
		sort.Slice(catalog, func(i, j int) bool {
			return Equipments[catalog[i]].Name < Equipments[catalog[j]].Name
		})

		for i, id := range catalog {
			eq := Equipments[id]
			fmt.Printf("%d. %s (%.0f coins)\n", i+1, eq.Name, eq.Price)
			req := Recipes[id]
			keys := make([]string, 0, len(req))
			for k := range req {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, res := range keys {
				fmt.Printf("   %d %s\n", req[res], res)
			}
		}
		fmt.Printf("%d. Return\n", len(catalog)+1)

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)

		if choice == len(catalog)+1 {
			return
		}
		if choice < 1 || choice > len(catalog) {
			lastMsg = "Invalid choice."
			time.Sleep(1 * time.Second)
			continue
		}

		id := catalog[choice-1]
		price := Equipments[id].Price
		reqs := Recipes[id]

		if player.Money < price {
			lastMsg = "You do not have enough coins."
			time.Sleep(1 * time.Second)
			continue
		}

		missing := []string{}
		for res, q := range reqs {
			have := 0
			if player.Inventory != nil {
				have = player.Inventory[res]
			}
			if have < q {
				missing = append(missing, fmt.Sprintf("%s x%d", res, q-have))
			}
		}
		if len(missing) > 0 {
			sort.Strings(missing)
			lastMsg = "Missing: " + strings.Join(missing, ", ")
			time.Sleep(1 * time.Second)
			continue
		}

		player.Money -= price
		for res, q := range reqs {
			player.Inventory[res] -= q
			if player.Inventory[res] <= 0 {
				delete(player.Inventory, res)
			}
		}
		AddEquipment(id, player)
		lastMsg = fmt.Sprintf("Crafted %s. Owned: x%d", id, player.Equipment[id])
		time.Sleep(1 * time.Second)
	}
}
