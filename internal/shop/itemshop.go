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
// then asks for a quantity with an inline hint about max affordable & capacity.
// Enter 0 to return.
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
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		id := catalog[choice-1]
		it := objects.Items[id]

		// Compute remaining capacity *before* asking quantity.
		totalHeld := 0
		for _, q := range player.Inventory {
			totalHeld += q
		}
		capLeft := player.InventoryMax - totalHeld
		if capLeft <= 0 {
			lastMsg = "Your inventory is full."
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		// Compute max affordable (price can theoretically be 0).
		maxAffordable := capLeft
		if it.Price > 0 {
			maxAffordable = int(math.Floor(player.Money / it.Price))
		}
		if maxAffordable <= 0 {
			lastMsg = "You do not have enough coins."
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		// Effective maximum is min(capacity, affordable).
		maxBuy := capLeft
		if maxAffordable < maxBuy {
			maxBuy = maxAffordable
		}

		// Ask for quantity with inline hint.
		fmt.Printf("How many %s do you want? (0 to cancel)\n", it.Label)
		fmt.Printf("You can afford up to %d. Inventory space: %d.\n> ",
			maxAffordable, capLeft)

		var qty int
		fmt.Scan(&qty)
		_ = audiosystem.PlaySFXCached("select")

		if qty == 0 {
			lastMsg = "Cancelled."
			continue
		}
		if qty < 0 || qty > maxBuy {
			lastMsg = fmt.Sprintf("Invalid amount. Enter a value between 0 and %d.", maxBuy)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		// Perform purchase.
		if player.Inventory == nil {
			player.Inventory = make(map[string]int)
		}
		totalCost := it.Price * float64(qty) // OK if price == 0
		player.Money -= totalCost
		player.Inventory[id] += qty

		lastMsg = fmt.Sprintf("You received %d %s, total: %d", qty, it.Label, player.Inventory[id])
		_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "buy.wav"))
		time.Sleep(1 * time.Second)
	}
}
