package shop

import (
	"fmt"
	"math"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
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

		// Collect inventory items.
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

		// Collect equipment.
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

		// Collect spellbooks (only those that are purchasable).
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

		for i, e := range catalog {
			fmt.Printf("%d. %s (x%d) — sell for %.0f coins\n", i+1, e.label, e.qty, e.price)
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
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		sel := catalog[choice-1]

		switch sel.kind {
		case "item":
			character.RemoveInventory(player, sel.id)
			player.Money += sel.price

		case "equip":
			slot := equipment.SlotOf(sel.id)
			if player.Equipped != nil && player.Equipped[slot] == sel.id {
				lastMsg = "Unequip it first before selling."
				_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
				time.Sleep(1 * time.Second)
				continue
			}
			if q, ok := player.Equipment[sel.id]; ok {
				if q <= 1 {
					delete(player.Equipment, sel.id)
				} else {
					player.Equipment[sel.id] = q - 1
				}
				player.Money += sel.price
			}

		case "spell":
			if q, ok := player.Skills[sel.id]; ok && q > 0 {
				if q == 1 {
					delete(player.Skills, sel.id)
				} else {
					player.Skills[sel.id] = q - 1
				}
				player.Money += sel.price
			}
		}

		lastMsg = fmt.Sprintf("Sold %s for %.0f coins.", sel.label, sel.price)
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "buy.wav"))
		time.Sleep(800 * time.Millisecond)
	}
}
