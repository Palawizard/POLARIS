package menu

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/shop"
	"projet-red_POLARIS/utils"
)

// ShowMenu is the main menu for the game. It displays a menu with three
// options: "Character Info", "Inventory", and "Quit". If the player chooses
// "Character Info", it displays the player's character info. If the player
// chooses "Inventory", it displays the player's inventory and allows them to
// shop. If the player chooses "Quit", the function will return.
func ShowMenu(player *utils.Player) {

	//Plays the menu music
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "menu.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}

	for {
		utils.Clearscreen()

		fmt.Println("Main Menu")
		fmt.Print("\n")
		fmt.Println("1. Character Info")
		fmt.Println("2. Inventory")
		fmt.Println("3. Black-Smith")
		fmt.Println("4. Training Fight")
		fmt.Println("5. Quit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			character.DisplayInfo(player)
		case 2:
			for {
				if character.AccessInventory(player) {
					shop.Shop(player)
					continue
				}
				break
			}
		case 3:
			equipement.BlackSmith(player)
		case 4:
			audiosystem.StopMusic()
			fightsystem.TrainingFight(player)

			//relance la musique si sortie de combat
			if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
				fmt.Println("play loop error:", err)
			}

		case 5:
			return
		}
	}
}
