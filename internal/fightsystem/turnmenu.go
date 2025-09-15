package fightsystem

import (
	"fmt"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

func TurnMenu(player *utils.Player, monster *monsters.Monster, turn int) bool {
	for {
		utils.Clearscreen()
		fmt.Println("Turn", turn)
		utils.SendTurn(turn)
		monsters.PrintHeader(monster)
		fmt.Println(player.Name, "HP:", player.Health, "/", player.MaxHealth)
		fmt.Println("It's your turn!\n")
		fmt.Println("1. Skills")
		fmt.Println("2. Inventory")
		fmt.Println("3. Return to menu")

		var choice int
		fmt.Scan(&choice)
		_ = audiosystem.PlaySFXCached("select")

		switch choice {
		case 1:
			type sopt struct{ id string }
			var sopts []sopt
			for id, qty := range player.Skills {
				if qty > 0 {
					if sk, ok := skills.Skills[id]; ok && sk.Apply != nil {
						sopts = append(sopts, sopt{id: id})
					}
				}
			}

			utils.Clearscreen()
			fmt.Println("Turn", turn)
			utils.SendTurn(turn)
			monsters.PrintHeader(monster)
			fmt.Println(player.Name, "HP:", player.Health, "/", player.MaxHealth)
			fmt.Println("It's your turn!\n")
			fmt.Println("Skills\n")
			fmt.Println("Mana:", player.Mana, "/", player.MaxMana)
			if len(sopts) == 0 {
				fmt.Println("(none)")
				fmt.Println("\n1. Return")
				var _tmp int
				fmt.Scan(&_tmp)
				_ = audiosystem.PlaySFXCached("select")
				continue
			}

			sort.Slice(sopts, func(i, j int) bool {
				return skills.Skills[sopts[i].id].Label < skills.Skills[sopts[j].id].Label
			})
			for i, o := range sopts {
				sk := skills.Skills[o.id]
				cost := character.SkillManaCost[o.id]
				fmt.Printf("%d. %s (x%d, Mana: %.0f)\n", i+1, sk.Label, player.Skills[o.id], cost)
			}
			fmt.Println("0. Cancel")

			var sidx int
			fmt.Scan(&sidx)
			_ = audiosystem.PlaySFXCached("select")
			if sidx == 0 || sidx < 0 || sidx > len(sopts) {
				continue
			}

			sel := sopts[sidx-1].id
			cost := character.SkillManaCost[sel]
			if player.Mana >= cost {
				player.Mana -= cost
				skills.Cast(sel, player, monster)
			} else {
				fmt.Println("Pas assez de mana !")
			}
			time.Sleep(2 * time.Second)
			character.ManaRegen(player, turn)
			return false

		case 2:
			type opt struct{ id string }
			var opts []opt
			for id, qty := range player.Inventory {
				if qty > 0 {
					if it, ok := objects.Items[id]; ok && it.Apply != nil {
						opts = append(opts, opt{id: id})
					}
				}
			}

			utils.Clearscreen()
			fmt.Println("Turn", turn)
			utils.SendTurn(turn)
			monsters.PrintHeader(monster)
			fmt.Println(player.Name, "HP:", player.Health, "/", player.MaxHealth)
			fmt.Println("It's your turn!\n")
			fmt.Println("Inventory (usable)\n")
			if len(opts) == 0 {
				fmt.Println("(none)")
				fmt.Println("\n1. Return")
				var _tmp int
				fmt.Scan(&_tmp)
				_ = audiosystem.PlaySFXCached("select")
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
			_ = audiosystem.PlaySFXCached("select")
			if idx == 0 || idx < 0 || idx > len(opts) {
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
			character.ManaRegen(player, turn)
			return false

		case 3:
			audiosystem.StopMusic()
			return true

		default:
			continue
		}
	}
}
