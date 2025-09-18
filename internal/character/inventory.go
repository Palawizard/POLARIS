package character

import (
	"fmt"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"sort"
	"strings"
)

// AccessInventory displays the player's inventory, skills, and equipment,
// then lets them use an item or equip/unequip gear. It loops until the
// player chooses to return.
func AccessInventory(player *utils.Player) bool {
	for {
		utils.ClearScreen()
		fmt.Println("Inventory\n")

		// --- Items (sorted by label) ---
		fmt.Println("Items:")
		if len(player.Inventory) == 0 {
			fmt.Println(" (none)")
		} else {
			type kv struct {
				label string
				qty   int
			}
			var list []kv
			for id, qty := range player.Inventory {
				lbl := id
				if it, ok := objects.Items[id]; ok && it.Label != "" {
					lbl = it.Label
				}
				list = append(list, kv{label: lbl, qty: qty})
			}
			sort.Slice(list, func(i, j int) bool { return list[i].label < list[j].label })
			for _, e := range list {
				fmt.Printf(" - %s x%d\n", e.label, e.qty)
			}
		}

		// --- Skills (sorted by label) ---
		fmt.Println("\nSkills:")
		if len(player.Skills) == 0 {
			fmt.Println(" (none)")
		} else {
			type kv struct {
				label string
				qty   int
			}
			var list []kv
			for id, qty := range player.Skills {
				lbl := id
				if s, ok := skills.Skills[id]; ok && s.Label != "" {
					lbl = s.Label
				}
				list = append(list, kv{label: lbl, qty: qty})
			}
			sort.Slice(list, func(i, j int) bool { return list[i].label < list[j].label })
			for _, e := range list {
				fmt.Printf(" - %s x%d\n", e.label, e.qty)
			}
		}

		// --- Equipment (sorted by slot then name) ---
		fmt.Println("\nEquipment:")
		if len(player.Equipment) == 0 {
			fmt.Println(" (none)")
		} else {
			type kv struct {
				name, slot string
				def        float64
				qty        int
				eq         bool
			}
			var list []kv
			for id, qty := range player.Equipment {
				e := equipement.GetEquipment(id)
				eq := player.Equipped != nil && player.Equipped[e.Type] == id
				list = append(list, kv{name: e.Name, slot: e.Type, def: e.Defense, qty: qty, eq: eq})
			}
			sort.Slice(list, func(i, j int) bool {
				if list[i].slot == list[j].slot {
					return list[i].name < list[j].name
				}
				return list[i].slot < list[j].slot
			})
			for _, e := range list {
				tag := ""
				if e.eq {
					tag = " [equipped]"
				}
				fmt.Printf(" - %s [%s] +%.0f (x%d)%s\n", e.name, e.slot, e.def, e.qty, tag)
			}
		}

		fmt.Print("\n")
		fmt.Println("1. Use an object")
		fmt.Println("0. Return")

		var choice int
		fmt.Scanln(&choice)
		_ = audiosystem.PlaySFXCached("select")

		switch choice {
		case 1:
			useItemMenu(player)
			continue
		case 0:
			return false
		default:
			continue
		}
	}
}

// invCount returns the total number of items in the player's inventory.
func invCount(p *utils.Player) int {
	if p == nil || p.Inventory == nil {
		return 0
	}
	total := 0
	for _, q := range p.Inventory {
		total += q
	}
	return total
}

// AddInventory adds one unit of item to the player's inventory,
// respecting the InventoryMax limit.
func AddInventory(player *utils.Player, item string) {
	if player == nil || item == "" {
		return
	}
	if player.Inventory == nil {
		player.Inventory = make(map[string]int)
	}
	if invCount(player) >= player.InventoryMax {
		fmt.Println("Your inventory is full.")
		return
	}
	player.Inventory[item]++
	fmt.Println(item, "added to inventory.")
}

// RemoveInventory removes one unit of item from the player's inventory.
// If it was the last one, the entry is deleted.
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

// useItemMenu lets the player use consumables or equip/unequip gear.
// Items and equipment are shown in separate alphabetical blocks, then merged.
func useItemMenu(p *utils.Player) {
	for {
		utils.ClearScreen()
		fmt.Println("Use an object\n")

		type option struct {
			id    string
			kind  string // "item" or "equip"
			label string
		}

		var itemOpts []option
		var equipOpts []option

		// Collect items (only those present in inventory) and sort by label.
		for id, qty := range p.Inventory {
			if qty > 0 {
				if it, ok := objects.GetItem(id); ok {
					lbl := it.Label
					if lbl == "" {
						lbl = id
					}
					itemOpts = append(itemOpts, option{id: id, kind: "item", label: lbl})
				}
			}
		}
		sort.Slice(itemOpts, func(i, j int) bool {
			return strings.ToLower(itemOpts[i].label) < strings.ToLower(itemOpts[j].label)
		})

		// Collect equipment (only those owned) and sort by name.
		for id, qty := range p.Equipment {
			if qty > 0 {
				if eq, ok := equipement.Equipments[id]; ok {
					equipOpts = append(equipOpts, option{id: id, kind: "equip", label: eq.Name})
				}
			}
		}
		sort.Slice(equipOpts, func(i, j int) bool {
			return strings.ToLower(equipOpts[i].label) < strings.ToLower(equipOpts[j].label)
		})

		// Merge the two lists for display/selection.
		opts := append([]option{}, itemOpts...)
		opts = append(opts, equipOpts...)

		if len(opts) == 0 {
			fmt.Println("No usable objects")
			fmt.Println("\n1. Return")
			var _tmp int
			fmt.Scanln(&_tmp)
			_ = audiosystem.PlaySFXCached("select")
			return
		}

		// Render options with counts; equipment shows slot and [equipped] tag if applicable.
		for i, o := range opts {
			if o.kind == "item" {
				fmt.Printf("%d. %s (x%d)\n", i+1, o.label, p.Inventory[o.id])
			} else {
				eq := equipement.Equipments[o.id]
				tag := ""
				if p.Equipped != nil && p.Equipped[eq.Type] == o.id {
					tag = " [equipped]"
				}
				fmt.Printf("%d. %s [%s] (x%d)%s\n", i+1, o.label, eq.Type, p.Equipment[o.id], tag)
			}
		}
		fmt.Println("0. Return")

		var idx int
		fmt.Scanln(&idx)
		_ = audiosystem.PlaySFXCached("select")
		if idx == 0 {
			return
		}
		if idx < 0 || idx > len(opts) {
			continue
		}

		choice := opts[idx-1]
		if choice.kind == "item" {
			// Try to apply the item; on success, consume one.
			if ok := objects.ApplyItem(choice.id, p); !ok {
				fmt.Println("You can't use this object.")
			} else {
				RemoveInventory(p, choice.id)
			}
		} else {
			// Toggle equipment in its slot.
			slot := equipement.SlotOf(choice.id)
			cur := ""
			if p.Equipped != nil {
				cur = p.Equipped[slot]
			}
			if cur == choice.id {
				if equipement.UnequipSlot(p, slot) {
					fmt.Printf("Unequipped %s\n", choice.id)
				} else {
					fmt.Println("You can't unequip this.")
				}
			} else {
				prev := cur
				if equipement.Equip(p, choice.id) {
					if prev != "" {
						fmt.Printf("Equipped %s, replaced %s\n", choice.id, prev)
					} else {
						fmt.Printf("Equipped %s\n", choice.id)
					}
				} else {
					fmt.Println("You can't equip this.")
				}
			}
		}

		// Stay in the use menu unless the player explicitly returns.
		fmt.Println("\n1. Continue")
		fmt.Println("0. Return")
		var cont int
		fmt.Scanln(&cont)
		_ = audiosystem.PlaySFXCached("select")
		if cont == 0 {
			return
		}
	}
}

// CheckInvSize reports whether the player can carry at least one more item.
func CheckInvSize(player *utils.Player) bool {
	if invCount(player) >= player.InventoryMax {
		fmt.Println("Your inventory is full.")
		return false
	}
	return true
}

// UpgradeInventorySlot increases InventoryMax by 10, up to a fixed number of upgrades.
func UpgradeInventorySlot(player *utils.Player) bool {
	if player.InventoryUpgradesUsed >= 10 {
		fmt.Println("You can't upgrade your inventory anymore!")
		return false
	}
	player.InventoryMax += 10
	player.InventoryUpgradesUsed++
	fmt.Println("Your inventory capacity has increased by 10!")
	return true
}
