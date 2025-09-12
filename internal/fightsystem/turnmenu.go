package fightsystem

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

func TurnMenu(player *utils.Player, monster *monsters.Monster, turn int) bool {
	for {
		utils.Clearscreen()
		fmt.Println(monster.Name, "HP:", monster.Health, "/", monster.MaxHealth)
		fmt.Println(player.Name, "HP:", player.Health, "/", player.MaxHealth)
		fmt.Println("Turn", turn)
		fmt.Println("It's your turn!\n")
		fmt.Println("1. Attack")
		fmt.Println("2. Inventory")
		fmt.Println("3. Return to menu")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			dmg := 5.0
			monster.Health -= dmg
			if monster.Health < 0 {
				monster.Health = 0
			}
			fmt.Printf("%s uses Basic Attack\n", player.Name)
			fmt.Printf("%s takes %.0f damage\n", monster.Name, dmg)
			fmt.Printf("%s HP: %.0f / %.0f\n", monster.Name, monster.Health, monster.MaxHealth)
			time.Sleep(2 * time.Second)
			return false

		case 2:
			type opt struct{ id string }
			var opts []opt
			for id, qty := range player.Inventory {
				if qty > 0 {
					if it, ok := objects.Items[id]; ok && it.Apply != nil {
						_ = it
						opts = append(opts, opt{id: id})
					}
				}
			}

			utils.Clearscreen()
			fmt.Println("Inventory (usable)\n")
			if len(opts) == 0 {
				fmt.Println("(none)")
				fmt.Println("\n1. Return")
				var _tmp int
				fmt.Scan(&_tmp)
				continue
			}

			sort.Slice(opts, func(i, j int) bool {
				return objects.Items[opts[i].id].Label < objects.Items[opts[j].id].Label
			})
			for i, o := range opts {
				it := objects.Items[o.id]
				fmt.Printf("%d. %s (x%d)\n", i+1, it.Label, player.Inventory[o.id])
			}
			fmt.Println("0. Cancel")

			var idx int
			fmt.Scan(&idx)
			if idx == 0 {
				continue
			}
			if idx < 0 || idx > len(opts) {
				continue
			}

			id := opts[idx-1].id
			it := objects.Items[id]
			fmt.Printf("You use %s\n", it.Label)
			if ok := objects.ApplyItem(id, player); ok {
				character.RemoveInventory(player, id)
			} else {
				fmt.Println("Nothing happens.")
			}
			time.Sleep(2 * time.Second)
			return false

		case 3:
			return true

		default:
			continue
		}
	}
}
