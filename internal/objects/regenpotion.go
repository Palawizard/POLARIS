package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
	"time"
)

func effectRegenPotion(p *utils.Player) {
	const tick = 15.0
	const ticks = 4
	for i := 0; i < ticks; i++ {
		healed := utils.ApplyHeal(&p.Health, p.MaxHealth, tick)
		if healed <= 0 {
			break
		}
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
		fmt.Printf("Regeneration tick %d/%d: +%.0f HP %s\n", i+1, ticks, tick, utils.HPString(p.Health, p.MaxHealth))
		if p.Health >= p.MaxHealth {
			break
		}
		time.Sleep(800 * time.Millisecond)
	}
}
