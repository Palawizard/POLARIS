package utils

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"time"
)

// IsDead returns true if the player's HP is 0 or below.
func IsDead(player *Player) bool {
	if player == nil || player.Health > 0 {
		return false
	}

	// Respawn at half max HP and cue death sequence.
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

	time.Sleep(1 * time.Second)
	fmt.Println("You should try again after getting the right equipment!")
	time.Sleep(1 * time.Second)
	fmt.Println("You can also farm the training fight to get stronger.")
	time.Sleep(4 * time.Second)

	return true
}
