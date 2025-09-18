package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

// effectBaguette heals the player for a flat amount and plays a heal SFX.
func effectBaguette(p *utils.Player) {
	const gain = 60.0 // raw heal before clamping to MaxHealth
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Printf("You nibble a curative baguette. Très bon! +%d HP: %s\n", healed, utils.HPString(p.Health, p.MaxHealth))
}
