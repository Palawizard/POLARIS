package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

// effectElixir fully restores the player's HP and plays the heal SFX.
func effectElixir(p *utils.Player) {
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, p.MaxHealth)
	_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "heal.mp3"))
	fmt.Printf("Elixir consumed! Fully restored (+%d HP). HP: %s\n", healed, utils.HPString(p.Health, p.MaxHealth))
}
