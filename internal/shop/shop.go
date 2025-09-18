package shop

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/utils"
	"time"
)

// Shop shows the shop hub (items, spells, sell, inventory upgrade). 0 returns.
func Shop(player *utils.Player) {
	for {
		utils.ClearScreen()
		fmt.Println("Shop")
		fmt.Print("\n\n")
		fmt.Println("<=== Welcome to the Shop ===>")
		fmt.Println("1. Item Shop")
		fmt.Println("2. Spell Shop")
		fmt.Println("3. Sell")
		fmt.Println("4. Upgrade Inventory (7 coins)")
		fmt.Println("0. Return")

		var choice int
		fmt.Scanln(&choice)
		_ = audiosystem.PlaySFXCached("select")

		switch choice {
		case 1:
			ItemShop(player)
		case 2:
			SpellShop(player)
		case 3:
			SellShop(player)
		case 4:
			if player.Money >= 7 {
				if character.UpgradeInventorySlot(player) {
					player.Money -= 7
					fmt.Println("Inventory slot upgraded.")
					_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "buy.wav"))
					time.Sleep(2 * time.Second)
				} else {
					fmt.Println("Upgrade limit reached!")
					_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
					time.Sleep(2 * time.Second)
				}
				continue
			}
			fmt.Println("You do not have enough coins.")
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "miss.mp3"))
			time.Sleep(2 * time.Second)
			continue
		case 0:
			return
		}
	}
}
