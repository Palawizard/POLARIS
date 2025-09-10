package objects

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

func effectPoisonPotion(p *utils.Player) {
	fmt.Println("Ouch! You lost 10 hp ! You now have ", p.Health, " hp left.")
	p.Health -= 10
	time.Sleep(1 * time.Second)
	fmt.Println("Ouch! You lost 10 hp ! You now have ", p.Health, " hp left.")
	p.Health -= 10
	time.Sleep(1 * time.Second)
	fmt.Println("Ouch! You lost 10 hp ! You now have ", p.Health, " hp left.")
	p.Health -= 10
	time.Sleep(1 * time.Second)
}
