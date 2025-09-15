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

func ShowMenu(player *utils.Player) {
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
		fmt.Println("3. Shop")
		fmt.Println("4. Black-Smith")
		fmt.Println("5. Training Fight")
		fmt.Println("6. Quit")

		var choice int
		fmt.Scan(&choice)
		_ = audiosystem.PlaySFXCached("select")

		switch choice {
		case 1:
			character.DisplayInfo(player)
		case 2:
			character.AccessInventory(player)
		case 3:
			shop.Shop(player)
		case 4:
			equipement.BlackSmith(player)
		case 5:
			audiosystem.StopMusic()
			fightsystem.TrainingFight(player)
			if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
				fmt.Println("play loop error:", err)
			}
		case 6:
			return
		}
	}
}
