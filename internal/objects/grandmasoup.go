package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

// effectGrandmaSoup heals a large chunk of HP and plays the heal SFX.
func effectGrandmaSoup(p *utils.Player) {
	const gain = 80.0
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)
	_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "heal.mp3"))
	fmt.Printf("You sip Grandma’s soup. Cozy! +%d HP: %s\n", healed, utils.HPString(p.Health, p.MaxHealth))
}
