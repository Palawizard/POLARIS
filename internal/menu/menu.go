package menu

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/chapters"
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
		fmt.Println("1. Start next chapter")
		fmt.Println("2. Character Info")
		fmt.Println("3. Inventory")
		fmt.Println("4. Shop")
		fmt.Println("5. Black-Smith")
		fmt.Println("6. Training Fight")
		fmt.Println("7. Qui sont-ils ?")
		fmt.Println("8. Quit")

		var choice int
		fmt.Scanln(&choice)
		_ = audiosystem.PlaySFXCached("select")

		switch choice {
		case 1:
			audiosystem.StopMusic()
			chapters.StartNextChapter(player)
			if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
				fmt.Println("play loop error:", err)
			}
		case 2:
			character.DisplayInfo(player)
		case 3:
			character.AccessInventory(player)
		case 4:
			shop.Shop(player)
		case 5:
			equipement.BlackSmith(player)
		case 6:
			audiosystem.StopMusic()
			fightsystem.TrainingFight(player)
			if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
				fmt.Println("play loop error:", err)
			}
		case 7:
			showHiddenArtists()
		case 8:
			return
		}
	}
}

func showHiddenArtists() {
	utils.Clearscreen()
	fmt.Println("Artistes cachés :")
	fmt.Println(" - ABBA (partie 2)")
	fmt.Println(" - Steven Spielberg (partie 3)")
	fmt.Println("\n1. Retour")
	var _tmp int
	fmt.Scanln(&_tmp)
	_ = audiosystem.PlaySFXCached("select")
}
