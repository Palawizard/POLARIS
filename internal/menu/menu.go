package menu

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/chapters"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/equipment"
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/shop"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

// ShowMenu runs the main loop: prints options, reads choice, routes to screens.
// The select SFX is only played after a successful read to avoid a "ghost" sound.
func ShowMenu(player *utils.Player) {
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("assets", "audio", "music", "menu.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}

	for {
		utils.ClearScreen()

		fmt.Println("Main Menu")
		fmt.Print("\n")
		fmt.Println("1. Start next chapter")
		fmt.Println("2. Character Info")
		fmt.Println("3. Inventory")
		fmt.Println("4. Shop")
		fmt.Println("5. Black-Smith")
		fmt.Println("6. Training Fight")
		fmt.Println("7. Qui sont-ils ?")
		fmt.Println("9. Quit")

		var choice int
		// Only play the select SFX after a valid input was read.
		if _, err := fmt.Scanln(&choice); err != nil {
			continue
		}
		_ = audiosystem.PlaySFXCached("select")

		switch choice {
		case 1:
			// Advance story.
			audiosystem.StopMusic()
			chapters.StartNextChapter(player)
			if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
				fmt.Println("play loop error:", err)
			}
		case 2:
			// Character sheet.
			character.DisplayInfo(player)
		case 3:
			// Inventory browser (use/equip).
			character.AccessInventory(player)
		case 4:
			// Shops (items, spells, sell, inv upgrade).
			shop.Shop(player)
		case 5:
			// Craft equipment.
			equipment.BlackSmith(player)
		case 6:
			// Quick practice fight.
			audiosystem.StopMusic()
			fightsystem.TrainingFight(player)
			if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
				fmt.Println("play loop error:", err)
			}
		case 7:
			// Little easter-egg credits.
			showHiddenArtists()
		case 9:
			// Exit to desktop.
			return

		// Dev shortcuts for testing chapters (silent no-UI entries).
		case 991:
			chapters.ChangeChapter(1)
		case 992:
			chapters.ChangeChapter(2)
		case 993:
			chapters.ChangeChapter(3)
		case 994:
			chapters.ChangeChapter(4)

		// Simple admin cheat to speed up testing.
		case 25565:
			utils.ClearScreen()
			fmt.Println("Admin mode activated.")
			time.Sleep(2 * time.Second)
			player.MaxHealth = 999999
			player.Health = 999999
			player.MaxMana = 999999
			player.Mana = 999999
			player.Money = 999999
			for i := 0; i < 100; i++ {
				skills.SpellBook("Meteor", player)
			}
		}
	}
}

// showHiddenArtists is a small modal screen with hidden credits.
func showHiddenArtists() {
	utils.ClearScreen()
	fmt.Println("Artistes cachés :")
	fmt.Println(" - ABBA (partie 2)")
	fmt.Println(" - Steven Spielberg (partie 3)")
	fmt.Println("\n0. Retour")
	var _tmp int
	_, _ = fmt.Scanln(&_tmp)
	_ = audiosystem.PlaySFXCached("select")
}
