package objects

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func effectHealthPotion(p *utils.Player) {
	p.Health += 50
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	fmt.Println("You gained 50 hp ! You now have", p.Health, "hp.")
}
