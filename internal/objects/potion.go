package objects

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

// effectHealthPotion increases the player's health by 50, capping it at
// MaxHealth if necessary. It then prints a message to the console
// indicating the new health value.
func effectHealthPotion(p *utils.Player) {
	p.Health += 50
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	fmt.Println("You gained 50 hp ! You now have", p.Health, "hp.")
}
