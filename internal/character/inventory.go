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

func AccessInventory(player *utils.Player) bool {
	for {
		utils.Clearscreen()
		fmt.Println("Inventory\n")

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
	for {
		utils.Clearscreen()
		fmt.Println("Use an object\n")

		type option struct {
			id    string
			kind  string
			label string
		}

		var itemOpts []option
		var equipOpts []option

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

		ch := opts[idx-1]
		if ch.kind == "item" {
			if ok := objects.ApplyItem(ch.id, p); !ok {
				fmt.Println("You can't use this object.")
			} else {
				RemoveInventory(p, ch.id)
			}
		} else {
			slot := equipement.SlotOf(ch.id)
			cur := ""
			if p.Equipped != nil {
				cur = p.Equipped[slot]
			}
			if cur == ch.id {
				if equipement.UnequipSlot(p, slot) {
					fmt.Printf("Unequipped %s\n", ch.id)
				} else {
					fmt.Println("You can't unequip this.")
				}
			} else {
				prev := cur
				if equipement.Equip(p, ch.id) {
					if prev != "" {
						fmt.Printf("Equipped %s, replaced %s\n", ch.id, prev)
					} else {
						fmt.Printf("Equipped %s\n", ch.id)
					}
				} else {
					fmt.Println("You can't equip this.")
				}
			}
		}

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

func CheckInvSize(player *utils.Player) bool {
	if invCount(player) >= player.InventoryMax {
		fmt.Println("Your inventory is full.")
		return false
	}
	return true
}

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
