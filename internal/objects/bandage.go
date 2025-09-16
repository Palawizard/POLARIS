package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

func effectBandage(p *utils.Player) {
	const gain = 20.0
	before := p.Health
	p.Health += gain
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	actual := p.Health - before
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Printf("You patch yourself up. +%.0f HP (%.0f/%.0f)\n", actual, p.Health, p.MaxHealth)
}
