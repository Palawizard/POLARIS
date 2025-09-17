package skills

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func effectFireball(p *utils.Player, m *monsters.Monster) {
	turn := utils.GetTurn()
	if m == nil {
		return
	}
	dmg := 18.0
	dmg = dmg * (1.0 + 0.2*float64(p.Level-1))
	applied := utils.ApplyDamage(&m.Health, dmg)
	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "fire.mp3"))
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s casts Fire Ball\n", p.Name)
	fmt.Printf("%s takes %d damage\n", m.Name, applied)
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))
	p.Skills["Fire Ball"]--
}
