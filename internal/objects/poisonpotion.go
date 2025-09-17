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
		applied := utils.ApplyDamage(&p.Health, dmg)
		fmt.Printf("Ouch! You lost %d hp ! You now have %s hp.\n", applied, utils.HPString(p.Health, p.MaxHealth))
		_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "enemyatk.mp3"))
		time.Sleep(1 * time.Second)
		if utils.IsDead(p) {
			return
		}
	}
}
