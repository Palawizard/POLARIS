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

var _rngLightning = rand.New(rand.NewSource(time.Now().UnixNano()))

func effectLightningBolt(p *utils.Player, m *monsters.Monster) {
	if m == nil {
		return
	}
	turn := utils.GetTurn()
	dmg := 12.0 + _rngLightning.Float64()*16.0
	applied := utils.ApplyDamage(&m.Health, dmg)

	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "fire.mp3"))
	utils.Shake(50, 1)
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s casts Lightning Bolt\n", p.Name)
	fmt.Printf("%s takes %d damage\n", m.Name, applied)
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))

	p.Skills["Lightning Bolt"]--
}
