package shop

import (
	"fmt"
	"math"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

// ItemShop shows all purchasable items, sorted by label, lets the player pick one,
// asks for a quantity, then buys that many if there is enough money and inventory space.
// Enter 0 to return. Shows the player's current coin total and owned quantity per item.
func ItemShop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.ClearScreen()
		fmt.Println("<=== Item Shop ===>")
		fmt.Printf("Coins: %.0f\n\n", player.Money)

		// Build a sorted catalog of item IDs by display label.
		catalog := make([]string, 0, len(objects.Items))
		for id := range objects.Items {
			catalog = append(catalog, id)
		}
		sort.Slice(catalog, func(i, j int) bool {
			return objects.Items[catalog[i]].Label < objects.Items[catalog[j]].Label
		})

		getInv := func(k string) int {
			if player.Inventory == nil {
				return 0
			}
			return player.Inventory[k]
		}

		for i, id := range catalog {
			it := objects.Items[id]
			fmt.Printf("%d. %s (%.0f coins)     [x%d]\n", i+1, it.Label, it.Price, getInv(id))
		}
		fmt.Println("0. Return")

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)
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

		id := catalog[choice-1]
		it := objects.Items[id]

		// Ask for quantity.
		fmt.Printf("How many %s do you want? (0 to cancel)\n> ", it.Label)
		var qty int
		fmt.Scan(&qty)
		_ = audiosystem.PlaySFXCached("select")

		if qty <= 0 {
			lastMsg = "Cancelled."
			continue
		}

		// Compute remaining capacity.
		total := 0
		for _, q := range player.Inventory {
			total += q
		}
		capLeft := player.InventoryMax - total
		if capLeft <= 0 {
			lastMsg = "Your inventory is full."
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}
		if qty > capLeft {
			lastMsg = fmt.Sprintf("Not enough space. You can carry at most %d more item(s).", capLeft)
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		// Check funds.
		totalCost := it.Price * float64(qty)
		if player.Money < totalCost {
			maxAffordable := int(math.Floor(player.Money / it.Price))
			if maxAffordable < 1 {
				lastMsg = "You do not have enough coins."
			} else {
				lastMsg = fmt.Sprintf("Not enough coins. You can afford up to %d.", maxAffordable)
			}
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		// Perform purchase.
		if player.Inventory == nil {
			player.Inventory = make(map[string]int)
		}
		player.Money -= totalCost
		player.Inventory[id] += qty

		lastMsg = fmt.Sprintf("You received %d %s, total: %d", qty, it.Label, player.Inventory[id])
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "buy.wav"))
		time.Sleep(1 * time.Second)
	}
}
