package objects

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

// effectPoisonPotion applies the effect of the poison potion on the player.
// It will deal 30 damage to the player over 3 seconds.
func effectPoisonPotion(p *utils.Player) {
	p.Health -= 10
	fmt.Println("Ouch! You lost 10 hp ! You now have ", p.Health, " hp left.")
	time.Sleep(1 * time.Second)
	p.Health -= 10
	fmt.Println("Ouch! You lost 10 hp ! You now have ", p.Health, " hp left.")
	time.Sleep(1 * time.Second)
	p.Health -= 10
	fmt.Println("Ouch! You lost 10 hp ! You now have ", p.Health, " hp left.")
	time.Sleep(1 * time.Second)
	if utils.IsDead(p) {
	}
}
