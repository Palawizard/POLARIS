package skills

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

// effectFireball applies Fire Ball: single-target damage with a flat crit chance.
func effectFireball(p *utils.Player, m *monsters.Monster) {
	turn := utils.GetTurn()
	if m == nil {
		return
	}

	// 25% crit: roll once per cast
	critical := rng.Intn(4) == 0

	// Base damage with crit and level scaling
	dmg := 18.0
	if critical {
		dmg *= 1.75
	}
	dmg *= 1.0 + 0.2*float64(p.Level-1)

	// Apply damage using the unified rounding/clamping helper
	applied := utils.ApplyDamage(&m.Health, dmg)

	// Presentation
	utils.ClearScreen()
	_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "fire.mp3"))
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s casts Fire Ball\n", p.Name)
	if critical {
		fmt.Println("Critical hit!")
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("%s takes %d damage\n", m.Name, applied)
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))

	// Consume one use of the spellbook
	p.Skills["Fire Ball"]--
}
