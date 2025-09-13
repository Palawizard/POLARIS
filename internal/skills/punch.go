package skills

import (
	"fmt"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func effectPunch(p *utils.Player, m *monsters.Monster) {
	turn := utils.GetTurn()
	if m == nil {
		return
	}
	dmg := 8.0
	m.Health -= dmg
	if m.Health < 0 {
		m.Health = 0
	}
	utils.Clearscreen()
	fmt.Println("Turn", turn)
	monsters.PrintHeader(m)
	fmt.Println("\n")
	fmt.Printf("%s uses Punch\n", p.Name)
	fmt.Printf("%s takes %.0f damage\n", m.Name, dmg)
	fmt.Printf("%s HP: %.0f / %.0f\n", m.Name, m.Health, m.MaxHealth)
}
