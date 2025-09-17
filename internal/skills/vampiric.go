package skills

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func effectVampiricTouch(p *utils.Player, m *monsters.Monster) {
	if m == nil {
		return
	}
	turn := utils.GetTurn()
	dmg := 14.0
	dmg = dmg * (1.0 + 0.2*float64(p.Level-1))
	heal := dmg * 0.5

	applied := utils.ApplyDamage(&m.Health, dmg)
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, heal)

	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "punch1.mp3"))
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s uses Vampiric Touch\n", p.Name)
	fmt.Printf("%s takes %d damage\n", m.Name, applied)
	fmt.Printf("%s absorbs %d HP\n", p.Name, healed)
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))

	p.Skills["Vampiric Touch"]--
}
