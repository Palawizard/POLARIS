package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

// effectHealthPotion heals the player for a fixed amount (up to max HP)
// and plays a heal sound, then prints the new HP state.
func effectHealthPotion(p *utils.Player) {
	const gain = 50.0
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Printf("You gained %d HP. HP: %s\n", healed, utils.HPString(p.Health, p.MaxHealth))
}
