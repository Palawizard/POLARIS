package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

// Spellshop is a menu allowing the player to buy a spellbook. It will
// display the player's current coins, and then display a list of spells
// available for purchase. The player is prompted to enter the number of the
// spell they wish to purchase. If the player enters a number that is not
// in the range of the options, or if they do not have enough coins, it will
// print out an error message and then loop back to the start of the menu.
// If the player chooses to purchase a spell, it will be added to their
// spellbook and the cost will be deducted from their coins. After the
// spell is purchased, the player will be prompted to enter "1" to return.
func Spellshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Spell Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)

		ownedFire := 0
		if player.Skills != nil {
			ownedFire = player.Skills["Fire Ball"]
		}

		fmt.Printf("1. Spellbook: Fire Ball (25$)  [owned: x%d]\n", ownedFire)
		fmt.Println("2. Return")

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if !character.CheckInvSize(player) {
				lastMsg = "Your inventory is full."
				time.Sleep(1 * time.Second)
				continue
			}
			if player.Money < 25 {
				lastMsg = "You do not have enough coins."
				time.Sleep(1 * time.Second)
				continue
			}
			player.Money -= 25
			if skills.SpellBook("Fire Ball", player) {
				lastMsg = fmt.Sprintf("You received 1 Spellbook: Fire Ball, total : %d", player.Skills["Fire Ball"])
			} else {
				lastMsg = "Can't buy that."
			}
			time.Sleep(1 * time.Second)
		case 2:
			return
		}
	}
}
