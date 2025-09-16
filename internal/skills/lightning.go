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
	m.Health -= dmg
	if m.Health < 0 {
		m.Health = 0
	}

	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "fire.mp3"))
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s casts Lightning Bolt\n", p.Name)
	fmt.Printf("%s takes %.0f damage\n", m.Name, dmg)
	fmt.Printf("%s HP: %.0f / %.0f\n", m.Name, m.Health, m.MaxHealth)

	p.Skills["Lightning Bolt"]--
}
