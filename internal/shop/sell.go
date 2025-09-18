package shop

import (
	"fmt"
	"math"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/equipment"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

// SellShop lets the player sell owned items, equipment, or spellbooks.
// Entries are sorted alphabetically and sell for ~50% of buy price.
// Enter 0 to return.
func SellShop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.ClearScreen()
		fmt.Println("<=== Sell ===>")
		fmt.Printf("Coins: %.0f\n\n", player.Money)

		type entry struct {
			kind  string
			id    string
			label string
			qty   int
			price float64
		}
		var catalog []entry

		// Inventory items.
		for id, qty := range player.Inventory {
			if qty <= 0 {
				continue
			}
			if it, ok := objects.Items[id]; ok {
				catalog = append(catalog, entry{
					kind:  "item",
					id:    id,
					label: it.Label,
					qty:   qty,
					price: math.Round(it.Price * 0.5),
				})
			}
		}

		// Equipment.
		for id, qty := range player.Equipment {
			if qty <= 0 {
				continue
			}
			if e, ok := equipment.Equipments[id]; ok {
				catalog = append(catalog, entry{
					kind:  "equip",
					id:    id,
					label: e.Name,
					qty:   qty,
					price: math.Round(e.Price * 0.5),
				})
			}
		}

		// Spellbooks (only those that are purchasable).
		for id, qty := range player.Skills {
			if qty <= 0 {
				continue
			}
			if s, ok := skills.Skills[id]; ok && s.Price > 0 {
				catalog = append(catalog, entry{
					kind:  "spell",
					id:    id,
					label: s.Label,
					qty:   qty,
					price: math.Round(s.Price * 0.5),
				})
			}
		}

		if len(catalog) == 0 {
			fmt.Println("You have nothing to sell.")
			fmt.Println("\n0. Return")
			var _tmp int
			_, _ = fmt.Scanln(&_tmp)
			_ = audiosystem.PlaySFXCached("select")
			return
		}

		sort.Slice(catalog, func(i, j int) bool { return catalog[i].label < catalog[j].label })

		// List entries; tag equipped gear as "equiped".
		for i, e := range catalog {
			tag := ""
			if e.kind == "equip" && player.Equipped != nil {
				slot := equipment.SlotOf(e.id)
				if player.Equipped[slot] == e.id {
					tag = " [equiped]"
				}
			}
			fmt.Printf("%d. %s%s (x%d) — sell for %.0f coins each\n", i+1, e.label, tag, e.qty, e.price)
		}
		fmt.Println("0. Return")

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		if _, err := fmt.Scanln(&choice); err != nil {
			continue
		}
		_ = audiosystem.PlaySFXCached("select")

		if choice == 0 {
			return
		}
		if choice < 1 || choice > len(catalog) {
			lastMsg = "Invalid choice."
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		sel := catalog[choice-1]

		// Compute max sellable (can't sell an equipped copy).
		maxSell := sel.qty
		if sel.kind == "equip" && player.Equipped != nil {
			slot := equipment.SlotOf(sel.id)
			if player.Equipped[slot] == sel.id && maxSell > 0 {
				maxSell = sel.qty - 1
			}
		}
		if maxSell <= 0 {
			lastMsg = "Unequip it first before selling."
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		// Ask how many to sell.
		fmt.Printf("How many '%s' do you want to sell? (max %d, 0 to cancel): ", sel.label, maxSell)
		var amount int
		if _, err := fmt.Scanln(&amount); err != nil {
			continue
		}
		_ = audiosystem.PlaySFXCached("select")

		if amount == 0 {
			continue
		}
		if amount < 0 || amount > maxSell {
			lastMsg = "Invalid amount."
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		// Apply sale.
		switch sel.kind {
		case "item":
			if q, ok := player.Inventory[sel.id]; ok {
				if amount >= q {
					delete(player.Inventory, sel.id)
				} else {
					player.Inventory[sel.id] = q - amount
				}
			}

		case "equip":
			if q, ok := player.Equipment[sel.id]; ok {
				newQ := q - amount
				if newQ <= 0 {
					// newQ should never be 0 if one is equipped thanks to maxSell logic.
					delete(player.Equipment, sel.id)
				} else {
					player.Equipment[sel.id] = newQ
				}
			}

		case "spell":
			if q, ok := player.Skills[sel.id]; ok {
				if amount >= q {
					delete(player.Skills, sel.id)
				} else {
					player.Skills[sel.id] = q - amount
				}
			}
		}

		total := math.Round(sel.price * float64(amount))
		player.Money += total

		lastMsg = fmt.Sprintf("Sold %d x %s for %.0f coins.", amount, sel.label, total)
		_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "buy.wav"))
		time.Sleep(800 * time.Millisecond)
	}
}
