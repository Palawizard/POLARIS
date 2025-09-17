package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

func effectHiPotion(p *utils.Player) {
	const gain = 120.0
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Printf("Glug glug! +%d HP %s\n", healed, utils.HPString(p.Health, p.MaxHealth))
}
