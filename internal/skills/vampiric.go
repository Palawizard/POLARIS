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
	heal := dmg * 0.5

	m.Health -= dmg
	if m.Health < 0 {
		m.Health = 0
	}

	before := p.Health
	p.Health += heal
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	gained := p.Health - before

	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "punch1.mp3"))
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s uses Vampiric Touch\n", p.Name)
	fmt.Printf("%s takes %.0f damage\n", m.Name, dmg)
	fmt.Printf("%s absorbs %.0f HP\n", p.Name, gained)
	fmt.Printf("%s HP: %.0f / %.0f\n", m.Name, m.Health, m.MaxHealth)

	p.Skills["Vampiric Touch"]--
}
