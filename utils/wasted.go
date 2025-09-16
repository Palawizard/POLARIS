package utils

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"time"
)

// IsDead checks if the player is dead, and if yes, it "respawns" them with half
// of their max health. It returns true if the player is dead and false
// otherwise.
func IsDead(player *Player) bool {
	if player.Health <= 0 {
		player.Health = player.MaxHealth / 2
		fmt.Println("Oh no! You DIED! You respawned with", player.Health, "HP.")
		audiosystem.StopMusic()
		if err := audiosystem.Init(); err != nil {
			fmt.Println("audio init error:", err)
		}
		musicPath := filepath.Join("internal", "audiosystem", "music", "death.mp3")
		if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
			fmt.Println("play loop error:", err)
		}
		time.Sleep(3 * time.Second)
		return true
	} else {
		return false
	}
}
