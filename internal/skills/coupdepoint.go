package skills

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func effectCoupDePoint(p *utils.Player) {
	p.Health -= 20
	fmt.Println("You have been hit by a point! You lost 20 hp!")
}
