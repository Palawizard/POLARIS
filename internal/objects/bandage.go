package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

// effectBandage restores a small amount of HP and plays the heal SFX.
func effectBandage(p *utils.Player) {
	const gain = 20.0 // flat heal before clamping to MaxHealth
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Printf("You patch yourself up. +%d HP: %s\n", healed, utils.HPString(p.Health, p.MaxHealth))
}
