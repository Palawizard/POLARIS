package shop

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

func Itemshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Item Shop ===>")
		fmt.Printf("Coins: %.0f\n\n", player.Money)

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
		fmt.Printf("%d. Return\n", len(catalog)+1)

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)
		_ = audiosystem.PlaySFXCached("select")

		if choice == len(catalog)+1 {
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

		if !character.CheckInvSize(player) {
			lastMsg = "Your inventory is full."
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}
		if player.Money < it.Price {
			lastMsg = "You do not have enough money."
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
			time.Sleep(1 * time.Second)
			continue
		}

		player.Money -= it.Price
		character.AddInventory(player, id)
		lastMsg = fmt.Sprintf("You received 1 %s, total : %d", it.Label, player.Inventory[id])
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "buy.wav"))
		time.Sleep(1 * time.Second)
	}
}
