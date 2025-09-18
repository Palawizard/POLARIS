package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
	"time"
)

// effectRegenPotion restores HP in several short pulses; stops early if already full.
func effectRegenPotion(p *utils.Player) {
	const perTick = 15.0
	const totalTicks = 4
	for i := 0; i < totalTicks; i++ {
		healed := utils.ApplyHeal(&p.Health, p.MaxHealth, perTick)
		if healed <= 0 {
			break
		}
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
		fmt.Printf("Regeneration tick %d/%d: +%d HP: %s\n", i+1, totalTicks, healed, utils.HPString(p.Health, p.MaxHealth))
		if p.Health >= p.MaxHealth {
			break
		}
		time.Sleep(800 * time.Millisecond)
	}
}
