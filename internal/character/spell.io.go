package character

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

func spellBook(player *utils.Player) {
	utils.Clearscreen()
	fmt.Println("Livre de Sort")
	fmt.Print("\n")
	fmt.Println("2. Livre de Sort : Boule de Feu")
	fmt.Println("1. Retour")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		return
	case 2:

	}
}
