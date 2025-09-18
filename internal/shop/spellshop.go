package shop

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"sort"
	"time"
)

// SpellShop shows purchasable spells; 0 returns.
func SpellShop(player *utils.Player) {
	lastMsg := ""
	for {
		utils.ClearScreen()
		fmt.Println("<=== Spell Shop ===>")
		fmt.Printf("Coins: %.0f\n\n", player.Money)

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
			fmt.Printf("%d. Spellbook: %s (%.0f coins, %.0f mana)  [owned: x%d]\n",
				i+1, s.Label, s.Price, s.ManaCost, owned)
		}
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
		if choice < 1 || choice > len(catalog) {
			lastMsg = "Invalid choice."
			time.Sleep(1 * time.Second)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			continue
		}

		id := catalog[choice-1]
		s := skills.Skills[id]

		if !character.CheckInvSize(player) {
			lastMsg = "Your inventory is full."
			time.Sleep(1 * time.Second)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			continue
		}
		if player.Money < s.Price {
			lastMsg = "You do not have enough coins."
			time.Sleep(1 * time.Second)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
			continue
		}

		player.Money -= s.Price
		if skills.SpellBook(id, player) {
			lastMsg = fmt.Sprintf("You received 1 Spellbook: %s, total : %d", s.Label, player.Skills[id])
			time.Sleep(1 * time.Second)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "buy.wav"))
		} else {
			lastMsg = "Can't buy that."
			time.Sleep(1 * time.Second)
			_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "miss.mp3"))
		}
	}
}
