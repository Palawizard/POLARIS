package fightsystem

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/character"
	"projet-red_POLARIS/internal/equipment"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/internal/objects"
	"projet-red_POLARIS/internal/skills"
	"projet-red_POLARIS/utils"
	"time"
)

// RunFight drives a full encounter loop until someone dies or the player exits.
// It returns (won, exit): won=true if the enemy is defeated, exit=true if the
// player chose to leave the fight (training/escape).
func RunFight(player *utils.Player, enemy *monsters.Monster, boss bool) (won bool, exit bool) {
	turn := 1
	firstTurn := true

	// Choose intro SFX and background music (special-cases for some enemies).
	sfx := "fightstart.mp3"
	delay := 1368 * time.Millisecond
	music := "fight.mp3"
	if enemy.Name == "Annoying Dog" {
		music = "dog.mp3"
	}
	if boss {
		sfx = "bossstart.mp3"
		delay = 5840 * time.Millisecond
		music = "bossbattle.mp3"
	}
	if enemy.Name == "Polaris" {
		music = "finalboss.mp3"
	}

	_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", sfx))
	time.Sleep(delay)

	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("assets", "audio", "music", music)
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}

	// Core battle loop.
	for {
		// Early out if the player is already at 0 (e.g., scripted).
		if player.Health <= 0 {
			audiosystem.StopMusic()
			return false, false
		}

		// Enemy may act first on turn 1 depending on initiative.
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

		// Player turn (TurnMenu may return true when quitting/escaping).
		if exit := TurnMenu(player, enemy, turn); exit {
			audiosystem.StopMusic()
			return false, true
		}

		// Post-player checks.
		if enemy.Health <= 0 {
			grantVictoryRewards(player, enemy)
			return true, false
		}
		if player.Health <= 0 {
			audiosystem.StopMusic()
			return false, false
		}

		// Enemy turn.
		if ok := monsters.AttackPattern(player, enemy, turn); !ok {
			audiosystem.StopMusic()
			return false, false
		}
		turn++

		// Post-enemy checks.
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

// grantVictoryRewards handles coins, EXP, and any loot drop after a win.
func grantVictoryRewards(player *utils.Player, enemy *monsters.Monster) {
	fmt.Println("You won!")
	audiosystem.StopMusic()
	_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "music", "win.mp3"))

	player.Money += enemy.CoinsToGive
	fmt.Println("You received", enemy.CoinsToGive, "coins.")
	time.Sleep(1 * time.Second)
	character.AddEXP(player, enemy.EXPToGive)

	// Grant optional loot: items, equipment, or spellbooks.
	if enemy.Loot != "" {
		id := enemy.Loot
		if it, ok := objects.GetItem(id); ok {
			character.AddInventory(player, id)
			fmt.Println("You found", it.Label+".")
		} else if e, ok := equipment.Equipments[id]; ok {
			equipment.AddEquipment(id, player)
			fmt.Println("You obtained", e.Name+".")
		} else if s, ok := skills.Skills[id]; ok {
			skills.SpellBook(id, player)
			fmt.Println("You obtained", s.ID+".")
		}
		time.Sleep(4 * time.Second)
	}
}
