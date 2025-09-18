package skills

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

// effectHeal restores a flat amount of HP.
func effectHeal(p *utils.Player, m *monsters.Monster) {
	turn := utils.GetTurn()
	const gain = 35.0

	// Apply healing with clamping/rounding handled by utils.
	healed := utils.ApplyHeal(&p.Health, p.MaxHealth, gain)

	// Presentation
	utils.ClearScreen()
	_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "heal.mp3"))
	fmt.Println("Turn", turn)
	if m != nil {
		monsters.PrintHeader(m)
	}
	fmt.Println("\n")
	fmt.Printf("%s casts Heal\n", p.Name)
	fmt.Printf("Recovered +%d HP. HP:%s\n", healed, utils.HPString(p.Health, p.MaxHealth))

	// Consume one use if tracked as an itemized spell
	p.Skills["Heal"]--
}
