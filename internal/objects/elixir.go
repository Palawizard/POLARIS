package objects

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

func effectElixir(p *utils.Player) {
	before := p.Health
	p.Health = p.MaxHealth
	actual := p.Health - before
	if actual < 0 {
		actual = 0
	}
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "heal.mp3"))
	fmt.Printf("Elixir consumed! Fully restored (+%.0f HP). (%.0f/%.0f)\n", actual, p.Health, p.MaxHealth)
}
