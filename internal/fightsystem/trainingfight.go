package fightsystem

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

func TrainingFight(player *utils.Player) {
	turn := 1
	goblin := monsters.New("Goblin")
	firstturn := true

	//Plays the fight music
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "fight.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}

	for {
		if player.Initiative < goblin.Initiative && firstturn {
			monsters.AttackPattern(player, goblin, turn)
			firstturn = false

			if utils.IsDead(player) {
				fmt.Println("You have been defeated by the goblin.")
				time.Sleep(3 * time.Second)
				audiosystem.StopMusic()
				return
			}
			if goblin.Health <= 0 {
				fmt.Println("You have defeated the goblin.")
				player.Money += goblin.Coinstogive
				fmt.Println("You received", goblin.Coinstogive, "coins.")
				time.Sleep(1 * time.Second)
				character.AddEXP(player, goblin.EXPtogive)
				time.Sleep(3 * time.Second)
				audiosystem.StopMusic()
				return
			}
		}
		if exit := TurnMenu(player, goblin, turn); exit {
			audiosystem.StopMusic()
			return
		}
		if goblin.Health <= 0 {
			fmt.Println("You have defeated the goblin.")
			player.Money += goblin.Coinstogive
			fmt.Println("You received", goblin.Coinstogive, "coins.")
			time.Sleep(1 * time.Second)
			character.AddEXP(player, goblin.EXPtogive)
			time.Sleep(3 * time.Second)
			audiosystem.StopMusic()
			return
		}
		if utils.IsDead(player) {
			fmt.Println("You have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			audiosystem.StopMusic()
			return
		}

		monsters.AttackPattern(player, goblin, turn)
		turn++

		if utils.IsDead(player) {
			fmt.Println("You have been defeated by the goblin.")
			time.Sleep(3 * time.Second)
			audiosystem.StopMusic()
			return
		}
		if goblin.Health <= 0 {
			fmt.Println("You have defeated the goblin.")
			player.Money += goblin.Coinstogive
			fmt.Println("You received", goblin.Coinstogive, "coins.")
			time.Sleep(1 * time.Second)
			character.AddEXP(player, goblin.EXPtogive)
			time.Sleep(3 * time.Second)
			audiosystem.StopMusic()
			return
		}
	}
}
