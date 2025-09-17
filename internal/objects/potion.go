package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

// effectHealthPotion increases the player's health by 50, capping it at
// MaxHealth if necessary. It then prints a message to the console
// indicating the new health value.
func effectHealthPotion(p *utils.Player) {
	const gain = 50.0
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Printf("You gained %d hp ! You now have %s hp.\n", healed, utils.HPString(p.Health, p.MaxHealth))
}
