package skills

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func effectIceShard(p *utils.Player, m *monsters.Monster) {
	if m == nil {
		return
	}
	turn := utils.GetTurn()
	dmg := 14.0
	applied := utils.ApplyDamage(&m.Health, dmg)

	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "sword.mp3"))
	utils.Shake(50, 1)
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s casts Ice Shard\n", p.Name)
	fmt.Printf("%s takes %d damage\n", m.Name, applied)
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))

	p.Skills["Ice Shard"]--
}
