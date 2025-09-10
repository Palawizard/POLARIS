package objects

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

func poisonPotion(player utils.Player) {
	utils.Clearscreen()
	fmt.Println("You drink the poison.")
	time.Sleep(1 * time.Second)
	player.Health -= 10
	fmt.Println("Ouch! Your health is now", player.Health, "/", player.Maxhealh)
	time.Sleep(1 * time.Second)
	player.Health -= 10
	fmt.Println("Ouch! Your health is now", player.Health, "/", player.Maxhealh)
	time.Sleep(1 * time.Second)
	player.Health -= 10
	fmt.Println("Ouch! Your health is now", player.Health, "/", player.Maxhealh)
}
