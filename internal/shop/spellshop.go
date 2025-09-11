package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

// Spellshop is a menu allowing the player to access the different spells
// available for purchase. It will display the player's current coins, and then
// display a list of spells available for purchase. The player is prompted to
// enter the number of the spell they wish to purchase. If the player enters a
// number that is not in the range of the options, it will simply loop back to the
// start of the menu. If the player chooses to purchase a spell, it will be added
// to their inventory, and the player's coins will be reduced by the price of
// the spell. After the player has purchased a spell, the player will be
// prompted to enter "1" to return.
func Spellshop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.Clearscreen()
		fmt.Println("<=== Spell Shop ===>")
		fmt.Printf("Coins: %d\n\n", player.Money)

		catalog := make([]string, 0, len(skills.Skills))
		for id, s := range skills.Skills {
			if s.Price > 0 {
				catalog = append(catalog, id)
			}
		}
		sort.Slice(catalog, func(i, j int) bool {
			return skills.Skills[catalog[i]].Label < skills.Skills[catalog[j]].Label
		})

		for i, id := range catalog {
			s := skills.Skills[id]
			owned := 0
			if player.Skills != nil {
				owned = player.Skills[id]
			}
			fmt.Printf("%d. Spellbook: %s (%d coins)  [owned: x%d]\n", i+1, s.Label, s.Price, owned)
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
		s := skills.Skills[id]

		if !character.CheckInvSize(player) {
			lastMsg = "Your inventory is full."
			time.Sleep(1 * time.Second)
			continue
		}
		if player.Money < s.Price {
			lastMsg = "You do not have enough coins."
			time.Sleep(1 * time.Second)
			continue
		}

		player.Money -= s.Price
		if skills.SpellBook(id, player) {
			lastMsg = fmt.Sprintf("You received 1 Spellbook: %s, total : %d", s.Label, player.Skills[id])
		} else {
			lastMsg = "Can't buy that."
		}
		time.Sleep(1 * time.Second)
	}
}
