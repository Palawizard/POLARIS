package skills

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func effectHeal(p *utils.Player, m *monsters.Monster) {
	turn := utils.GetTurn()
	heal := 35.0
	before := p.Health
	p.Health += heal
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	gained := p.Health - before

	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Println("Turn", turn)
	if m != nil {
		monsters.PrintHeader(m)
	}
	fmt.Println("\n")
	fmt.Printf("%s casts Heal\n", p.Name)
	fmt.Printf("%s recovers %.0f HP\n", p.Name, gained)
	fmt.Printf("%s HP: %.0f / %.0f\n", p.Name, p.Health, p.MaxHealth)

	p.Skills["Heal"]--
}
