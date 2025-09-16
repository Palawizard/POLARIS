package fightsystem

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/equipement"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

func RunFight(player *utils.Player, enemy *monsters.Monster, boss bool) (won bool, exit bool) {
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
		if player.Health <= 0 {
			audiosystem.StopMusic()
			return false, false
		}

		if player.Initiative < enemy.Initiative && firstTurn {
			if ok := monsters.AttackPattern(player, enemy, turn); !ok {
				audiosystem.StopMusic()
				return false, false
			}
			firstTurn = false
			if enemy.Health <= 0 {
				grantVictoryRewards(player, enemy)
				return true, false
			}
		}

		if exit := TurnMenu(player, enemy, turn); exit {
			audiosystem.StopMusic()
			return false, true
		}

		if enemy.Health <= 0 {
			grantVictoryRewards(player, enemy)
			return true, false
		}
		if player.Health <= 0 {
			audiosystem.StopMusic()
			return false, false
		}

		if ok := monsters.AttackPattern(player, enemy, turn); !ok {
			audiosystem.StopMusic()
			return false, false
		}
		turn++

		if enemy.Health <= 0 {
			grantVictoryRewards(player, enemy)
			return true, false
		}
		if player.Health <= 0 {
			audiosystem.StopMusic()
			return false, false
		}
	}
}

func grantVictoryRewards(player *utils.Player, enemy *monsters.Monster) {
	fmt.Println("You won!")
	audiosystem.StopMusic()
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "music", "win.mp3"))

	player.Money += enemy.Coinstogive
	fmt.Println("You received", enemy.Coinstogive, "coins.")
	time.Sleep(1 * time.Second)
	character.AddEXP(player, enemy.EXPtogive)

	if enemy.Loot != "" {
		id := enemy.Loot
		if it, ok := objects.GetItem(id); ok {
			character.AddInventory(player, id)
			fmt.Println("You found", it.Label+".")
		} else if e, ok := equipement.Equipments[id]; ok {
			equipement.AddEquipment(id, player)
			fmt.Println("You obtained", e.Name+".")
		} else if s, ok := skills.Skills[id]; ok {
			skills.SpellBook(id, player)
			fmt.Println("You obtained", s.ID+".")
		}
		time.Sleep(4 * time.Second)
	}
}
