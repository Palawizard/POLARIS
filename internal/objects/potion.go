package objects

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

// effectHealthPotion increases the player's health by 50, capping it at
// MaxHealth if necessary. It then prints a message to the console
// indicating the new health value.
func effectHealthPotion(p *utils.Player) {
	const gain = 50.0
	before := p.Health
	p.Health += gain
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	actual := p.Health - before
	fmt.Printf("You gained %.0f hp ! You now have %.0f hp.\n", actual, p.Health)
}
