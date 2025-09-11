package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

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
