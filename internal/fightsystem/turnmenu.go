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

// TurnMenu runs the player's turn UI.
// Returns true if the player chose to exit the fight, false otherwise.
func TurnMenu(player *utils.Player, monster *monsters.Monster, turn int) bool {
	for {
		// Passive mana regeneration at the start of each menu loop.
		if player.Mana < player.MaxMana {
			player.Mana += player.ManaRegen
			if player.Mana > player.MaxMana {
				player.Mana = player.MaxMana
			}
		}

		// Header + basic HUD
		utils.ClearScreen()
		fmt.Println("Turn", turn)
		utils.SendTurn(turn)
		monsters.PrintHeader(monster)
		fmt.Printf("%s HP: %s\n", player.Name, utils.HPString(player.Health, player.MaxHealth))
		fmt.Printf("%s MP: %.0f/%.0f\n", player.Name, player.Mana, player.MaxMana)
		fmt.Println("It's your turn!\n")
		fmt.Println("1. Skills")
		fmt.Println("2. Inventory")
		fmt.Println("0. Return to menu")

		var choice int
		fmt.Scanln(&choice)
		_ = audiosystem.PlaySFXCached("select")

		switch choice {
		case 1:
			// Build the list of castable skills the player owns (x>0).
			type skillOpt struct{ id string }
			var skillOpts []skillOpt
			for id, qty := range player.Skills {
				if qty > 0 {
					if sk, ok := skills.Skills[id]; ok && sk.Apply != nil {
						_ = sk
						skillOpts = append(skillOpts, skillOpt{id: id})
					}
				}
			}

			// Skills submenu
			utils.ClearScreen()
			fmt.Println("Turn", turn)
			utils.SendTurn(turn)
			monsters.PrintHeader(monster)
			fmt.Printf("%s HP: %s\n", player.Name, utils.HPString(player.Health, player.MaxHealth))
			fmt.Printf("%s MP: %.0f/%.0f\n", player.Name, player.Mana, player.MaxMana)
			fmt.Println("It's your turn!\n")
			fmt.Println("Skills\n")
			if len(skillOpts) == 0 {
				fmt.Println("(none)")
				fmt.Println("\n1. Return")
				var _tmp int
				fmt.Scanln(&_tmp)
				_ = audiosystem.PlaySFXCached("select")
				continue
			}

			// Sort by label for a stable, readable list.
			sort.Slice(skillOpts, func(i, j int) bool {
				return skills.Skills[skillOpts[i].id].Label < skills.Skills[skillOpts[j].id].Label
			})

			for i, o := range skillOpts {
				sk := skills.Skills[o.id]
				cost := sk.ManaCost
				tag := ""
				if player.Mana < cost {
					tag = " [not enough mana]"
				}
				fmt.Printf("%d. %s (x%d) - %.0f MP%s\n", i+1, sk.Label, player.Skills[o.id], cost, tag)
			}
			fmt.Println("0. Cancel")

			// Selection + cast
			var skillIdx int
			fmt.Scanln(&skillIdx)
			_ = audiosystem.PlaySFXCached("select")
			if skillIdx == 0 {
				continue
			}
			if skillIdx < 0 || skillIdx > len(skillOpts) {
				continue
			}
			sel := skillOpts[skillIdx-1].id
			cost := skills.Skills[sel].ManaCost
			if player.Mana < cost {
				fmt.Println("Not enough mana.")
				time.Sleep(1 * time.Second)
				continue
			}
			if ok := skills.Cast(sel, player, monster); !ok {
				// Defensive: in case a spell fails its own checks.
				fmt.Println("The spell fizzles.")
				time.Sleep(1 * time.Second)
				continue
			}
			time.Sleep(2 * time.Second)
			return false

		case 2:
			// Build list of usable inventory items (those with Apply != nil and qty > 0).
			type itemOpt struct{ id string }
			var itemOpts []itemOpt
			for id, qty := range player.Inventory {
				if qty > 0 {
					if it, ok := objects.Items[id]; ok && it.Apply != nil {
						_ = it
						itemOpts = append(itemOpts, itemOpt{id: id})
					}
				}
			}

			// Items submenu
			utils.ClearScreen()
			fmt.Println("Turn", turn)
			utils.SendTurn(turn)
			monsters.PrintHeader(monster)
			fmt.Printf("%s HP: %s\n", player.Name, utils.HPString(player.Health, player.MaxHealth))
			fmt.Printf("%s MP: %.0f/%.0f\n", player.Name, player.Mana, player.MaxMana)
			fmt.Println("It's your turn!\n")
			fmt.Println("Inventory (usable)\n")
			if len(itemOpts) == 0 {
				fmt.Println("(none)")
				fmt.Println("\n1. Return")
				var _tmp int
				fmt.Scanln(&_tmp)
				_ = audiosystem.PlaySFXCached("select")
				continue
			}

			// Stable alphabetical listing by item label.
			sort.Slice(itemOpts, func(i, j int) bool {
				return objects.Items[itemOpts[i].id].Label < objects.Items[itemOpts[j].id].Label
			})
			for i, o := range itemOpts {
				it := objects.Items[o.id]
				fmt.Printf("%d. %s (x%d)\n", i+1, it.Label, player.Inventory[o.id])
			}
			fmt.Println("0. Cancel")

			// Use item
			var itemIdx int
			fmt.Scanln(&itemIdx)
			_ = audiosystem.PlaySFXCached("select")
			if itemIdx == 0 {
				continue
			}
			if itemIdx < 0 || itemIdx > len(itemOpts) {
				continue
			}

			id := itemOpts[itemIdx-1].id
			it := objects.Items[id]
			fmt.Printf("You use %s\n", it.Label)
			if ok := objects.ApplyItem(id, player); ok {
				character.RemoveInventory(player, id)
			} else {
				fmt.Println("Nothing happens.")
			}
			time.Sleep(2 * time.Second)
			return false

		case 0:
			// Exit to main menu (ends the fight loop).
			audiosystem.StopMusic()
			return true

		default:
			// Any other input re-displays the menu.
			continue
		}
	}
}
