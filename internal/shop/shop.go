package shop

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/utils"
	"time"
)

// Shop is a menu allowing the player to access the different shops.
// It will display the player's current coins, and then display a list of shops
// available for purchase. The player is prompted to enter the number of the
// shop they wish to access. If the player enters a number that is not in the
// range of the options, it will simply loop back to the start of the menu.
// If the player chooses to access a shop, it will be run in a loop until the
// player chooses to return.
func Shop(player *utils.Player) {
	for {
		utils.Clearscreen()
		fmt.Println("Shop")
		fmt.Print("\n\n")
		fmt.Println("<=== Welcome to the Shop ===>")
		fmt.Println("1. Item Shop")
		fmt.Println("2. Spell Shop")
		fmt.Println("3. Upgrade Inventory (7 coins)")
		fmt.Println("0. Return")

		var choice int
		fmt.Scan(&choice)
		_ = audiosystem.PlaySFXCached("select")
		switch choice {
		case 1:
			Itemshop(player)
		case 2:
			Spellshop(player)
		case 3:
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
