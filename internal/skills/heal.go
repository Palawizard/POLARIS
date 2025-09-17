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
	const gain = 35.0
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)

	utils.Clearscreen()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Println("Turn", turn)
	if m != nil {
		monsters.PrintHeader(m)
	}
	fmt.Println("\n")
	fmt.Printf("%s casts Heal\n", p.Name)
	fmt.Printf("You gained %d hp ! You now have %s hp.\n", healed, utils.HPString(p.Health, p.MaxHealth))

	p.Skills["Heal"]--
}
