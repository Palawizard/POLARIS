package objects

import "projet-red_POLARIS/utils"

func effectHealthPotion(p *utils.Player) {
	p.Health += 50
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
}

func TakePotion(p *utils.Player) {
	_ = ApplyItem("Potion", p)
}
