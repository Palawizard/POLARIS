package equipment

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
	"sort"
	"strings"
	"time"
)

// BlackSmith shows all craftable gear, checks costs + mats, and crafts on confirm.
// Uses a small status line (lastMsg) to surface errors/success without leaving the menu.
func BlackSmith(player *utils.Player) {
	lastMsg := ""
	for {
		utils.ClearScreen()
		fmt.Println("<=== Black-Smith ===>")
		fmt.Printf("Coins: %.0f\n\n", player.Money)
		fmt.Println("Crafting requires resources:")

		// Build a stable, alpha-sorted catalog by display name.
		catalog := make([]string, 0, len(Equipments))
		for id := range Equipments {
			catalog = append(catalog, id)
		}
		sort.Slice(catalog, func(i, j int) bool {
			return Equipments[catalog[i]].Name < Equipments[catalog[j]].Name
		})

		// List entries with slot + Max HP bonus.
		for i, id := range catalog {
			eq := Equipments[id]
			fmt.Printf("%d. %s [%s] +%.0f MHP (%.0f coins)\n", i+1, eq.Name, eq.Type, eq.Defense, eq.Price)

			// Print recipe lines in a stable order.
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

		// Footer actions on a new line as requested.
		fmt.Printf("\n%d. Help\n", len(catalog)+1)
		fmt.Println("0. Return")

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scanln(&choice)
		_ = audiosystem.PlaySFXCached("select")

		if choice == 0 {
			return
		}
		if choice == len(catalog)+1 {
			blacksmithHelp()
			lastMsg = "" // clear any stale status after returning from Help
			continue
		}
		if choice < 1 || choice > len(catalog) {
			lastMsg = "Invalid choice."
			time.Sleep(1 * time.Second)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			continue
		}

		id := catalog[choice-1]
		price := Equipments[id].Price
		reqs := Recipes[id]

		// Coin check first.
		if player.Money < price {
			lastMsg = "You do not have enough coins."
			time.Sleep(1 * time.Second)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			continue
		}

		// Gather missing mats, if any.
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
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			continue
		}

		// Craft: pay coins, consume mats, add item.
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
		_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "buy.wav"))
	}
}

// blacksmithHelp renders costs, total MHP and total mats for each set, then waits for 0 to return.
func blacksmithHelp() {
	utils.ClearScreen()
	fmt.Println("<=== Help ===>\n")

	// Fixed order to avoid map iteration randomness.
	setOrder := []string{"Adventurer", "Leather", "Iron"}
	sets := map[string][]string{
		"Adventurer": {"Adventurer's Hat", "Adventurer's Tunic", "Adventurer's Boots"},
		"Leather":    {"Leather Cap", "Leather Armor", "Leather Boots"},
		"Iron":       {"Iron Helm", "Iron Plate", "Iron Greaves"},
	}

	for _, setName := range setOrder {
		parts := sets[setName]

		totalCoins := 0.0
		totalMHP := 0.0
		mat := map[string]int{}

		fmt.Printf("%s Set:\n", setName)
		for _, id := range parts {
			eq := Equipments[id]
			fmt.Printf(" - %s [%s] +%.0f MHP (%.0f coins)\n", eq.Name, eq.Type, eq.Defense, eq.Price)
			totalCoins += eq.Price
			totalMHP += eq.Defense
			for res, q := range Recipes[id] {
				mat[res] += q
			}
		}

		// Stable material listing.
		keys := make([]string, 0, len(mat))
		for k := range mat {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		fmt.Println("\n   Required materials (full set):")
		for _, k := range keys {
			fmt.Printf("   - %s x%d\n", k, mat[k])
		}
		fmt.Printf("   Total coins (full set): %.0f\n", totalCoins)
		fmt.Printf("   Total MHP bonus (full set): +%.0f\n\n", totalMHP)
	}

	fmt.Println("0. Return")
	var _tmp int
	fmt.Scanln(&_tmp)
	_ = audiosystem.PlaySFXCached("select")
}
