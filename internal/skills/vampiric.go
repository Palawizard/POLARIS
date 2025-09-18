package skills

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

// effectVampiricTouch drains life: deals damage (level-scaled, 25% crit at 1.75x)
// and heals the caster for 50% of the damage actually dealt.
func effectVampiricTouch(p *utils.Player, m *monsters.Monster) {
	if m == nil {
		return
	}
	turn := utils.GetTurn()

	critical := rng.Intn(4) == 0

	// Base damage with level scaling; crit boosts final damage.
	dmg := 14.0
	if critical {
		dmg *= 1.75
	}
	dmg = dmg * (1.0 + 0.2*float64(p.Level-1))

	// Apply damage first, then heal for 50% of the attempted damage.
	applied := utils.ApplyDamage(&m.Health, dmg)
	heal := dmg * 0.5
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, heal)

	// Feedback & SFX.
	utils.ClearScreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "punch1.mp3"))
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s uses Vampiric Touch\n", p.Name)
	if critical {
		fmt.Println("Critical hit!")
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("%s takes %d damage\n", m.Name, applied)
	fmt.Printf("%s absorbs %d HP\n", p.Name, healed)
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))

	// Consume one use of the spellbook.
	p.Skills["Vampiric Touch"]--
}
