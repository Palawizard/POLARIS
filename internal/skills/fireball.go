package skills

import (
	"fmt"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func effectFireball(p *utils.Player, m *monsters.Monster) {
	if m == nil {
		return
	}
	dmg := 18.0
	m.Health -= dmg
	if m.Health < 0 {
		m.Health = 0
	}
	fmt.Printf("%s casts Fire Ball\n", p.Name)
	fmt.Printf("%s takes %.0f damage\n", m.Name, dmg)
	fmt.Printf("%s HP: %.0f / %.0f\n", m.Name, m.Health, m.MaxHealth)
	p.Skills["Fire Ball"]--
}
