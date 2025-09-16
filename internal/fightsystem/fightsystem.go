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

func RunFight(player *utils.Player, enemy *monsters.Monster, boss bool) (bool, bool) {
	turn := 1
	firstTurn := true

	sfx := "fightstart.mp3"
	delay := 1368 * time.Millisecond
	music := "fight.mp3"
	if boss {
		sfx = "bossstart.mp3"
		delay = 5840 * time.Millisecond
		music = "bossbattle.mp3"
	}

	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", sfx))
	time.Sleep(delay)

	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", music)
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}

	for {
		if player.Initiative < enemy.Initiative && firstTurn {
			if ok := monsters.AttackPattern(player, enemy, turn); !ok {
				audiosystem.StopMusic()
				return false, true
			}
			firstTurn = false

			if enemy.Health <= 0 {
				fmt.Println("You won!")
				audiosystem.StopMusic()
				_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "music", "win.mp3"))
				player.Money += enemy.Coinstogive
				fmt.Println("You received", enemy.Coinstogive, "coins.")
				time.Sleep(1 * time.Second)
				character.AddEXP(player, enemy.EXPtogive)
				time.Sleep(3 * time.Second)
				return true, false
			}
		}

		if exit := TurnMenu(player, enemy, turn); exit {
			audiosystem.StopMusic()
			return false, true
		}

		if enemy.Health <= 0 {
			fmt.Println("You won!")
			audiosystem.StopMusic()
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "music", "win.mp3"))
			player.Money += enemy.Coinstogive
			fmt.Println("You received", enemy.Coinstogive, "coins.")
			time.Sleep(1 * time.Second)
			character.AddEXP(player, enemy.EXPtogive)
			time.Sleep(3 * time.Second)
			return true, false
		}

		if ok := monsters.AttackPattern(player, enemy, turn); !ok {
			audiosystem.StopMusic()
			return false, true
		}
		turn++

		if enemy.Health <= 0 {
			fmt.Println("You won!")
			audiosystem.StopMusic()
			_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "music", "win.mp3"))
			player.Money += enemy.Coinstogive
			fmt.Println("You received", enemy.Coinstogive, "coins.")
			time.Sleep(1 * time.Second)
			character.AddEXP(player, enemy.EXPtogive)
			time.Sleep(3 * time.Second)
			return true, false
		}
	}
}
