package shop

import (
	"fmt"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/utils"
)

/*
/
Shop feature allowing you to access the shop and choose an item, either paid or free
*/

func Shop(player utils.Player) { // initialisation of shop function
	utils.Clearscreen()                               // clear the screen (cmd)
	fmt.Println("Shop")                               //
	fmt.Print("\n\n")                                 //
	fmt.Println("=== Bienvenue chez le marchand ===") //
	fmt.Println("11 Potion de vie (GRATUITE)")        //
	fmt.Println("2. Retour")                          //
	var choice int                                    //
	fmt.Scan(&choice)                                 //
	switch choice {                                   //
	case 1: //
		objects.TakePotion(player) //
	case 2:
		return //
	} //
} //
