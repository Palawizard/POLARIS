package skills

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

// rngLightning generates the variable portion of Lightning Bolt's damage.
var rngLightning = rand.New(rand.NewSource(time.Now().UnixNano()))

// effectLightningBolt deals lightning damage with a random range and a chance to crit.
func effectLightningBolt(p *utils.Player, m *monsters.Monster) {
	if m == nil {
		return
	}
	turn := utils.GetTurn()

	critical := rng.Intn(4) == 0

	dmg := 12.0 + rngLightning.Float64()*16.0
	if critical {
		dmg *= 1.75
	}
	dmg *= 1.0 + 0.2*float64(p.Level-1)

	applied := utils.ApplyDamage(&m.Health, dmg)

	utils.ClearScreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "fire.mp3"))
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s casts Lightning Bolt\n", p.Name)
	if critical {
		fmt.Println("Critical hit!")
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("%s takes %d damage\n", m.Name, applied)
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))

	p.Skills["Lightning Bolt"]--
}
