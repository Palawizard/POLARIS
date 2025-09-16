package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
	"time"
)

// effectPoisonPotion applies the effect of the poison potion on the player.
// It will deal 30 damage to the player over 3 seconds.
func effectPoisonPotion(p *utils.Player) {
	for i := 0; i < 3; i++ {
		dmg := 10.0
		if p.Health < dmg {
			dmg = p.Health
		}
		p.Health -= dmg
		if p.Health < 1 {
			p.Health = 0
		}
		fmt.Printf("Ouch! You lost %.0f hp ! You now have %.0f hp.\n", dmg, p.Health)
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "enemyatk.mp3"))
		time.Sleep(1 * time.Second)
		if utils.IsDead(p) {
			return
		}
	}
}
