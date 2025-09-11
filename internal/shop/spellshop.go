package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

// Spellshop displays the player's spellbook and allows them to buy spells.
// It will display the player's current coins, and then display a list of
// available spells. The player is prompted to enter the number of the
// spell they wish to buy. If the player enters a number that is not in the
// range of the options, or if they do not have enough coins, or if their
// inventory is full, it will simply loop back to the start of the menu. If
// the player chooses to buy a spell, it will be added to their spellbook,
// the cost of the spell will be subtracted from their money, and the player
// will be prompted to enter "1" to return.
func Spellshop(player *utils.Player) {
	lastMsg := ""
	catalog := []string{"Fire Ball"}
	prices := map[string]int{"Fire Ball": 25}

	for {
		utils.Clearscreen()
		fmt.Println("<=== Spell Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)

		for i, id := range catalog {
			label := id
			if s, ok := skills.Skills[id]; ok && s.Label != "" {
				label = s.Label
			}
			owned := 0
			if player.Skills != nil {
				owned = player.Skills[id]
			}
			fmt.Printf("%d. Spellbook: %s (%d coins)  [owned: x%d]\n", i+1, label, prices[id], owned)
		}
		fmt.Printf("%d. Return\n", len(catalog)+1)

		if lastMsg != "" {
			fmt.Println()
			fmt.Println(lastMsg)
		}

		var choice int
		fmt.Scan(&choice)

		if choice == len(catalog)+1 {
			return
		}
		if choice < 1 || choice > len(catalog) {
			lastMsg = "Invalid choice."
			time.Sleep(1 * time.Second)
			continue
		}

		id := catalog[choice-1]
		price := prices[id]
		label := id
		if s, ok := skills.Skills[id]; ok && s.Label != "" {
			label = s.Label
		}

		if !character.CheckInvSize(player) {
			lastMsg = "Your inventory is full."
			time.Sleep(1 * time.Second)
			continue
		}
		if player.Money < price {
			lastMsg = "You do not have enough coins."
			time.Sleep(1 * time.Second)
			continue
		}

		player.Money -= price
		if skills.SpellBook(id, player) {
			lastMsg = fmt.Sprintf("You received 1 Spellbook: %s, total : %d", label, player.Skills[id])
		} else {
			lastMsg = "Can't buy that."
		}
		time.Sleep(1 * time.Second)
	}
}
