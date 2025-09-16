package chapters

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
	"time"
)

var currentChapter = 1

func ChangeChapter(i int) {
	currentChapter = i
}

func StartMusic() {
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "chapters.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}
}

func StartWinningMusic() {
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "chapterwon.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}
}

func StopMusic() {
	audiosystem.StopMusic()
}

func StartNextChapter(player *utils.Player) {
	switch currentChapter {
	case 1:
		if StartChapter1(player) {
			currentChapter++
		}
	case 2:
		if StartChapter2(player) {
			currentChapter++
		}
	case 3:
		if StartChapter3(player) {
			currentChapter++
		}
	case 4:
		if StartFinalBoss(player) {
			currentChapter++
		}
	default:
		fmt.Println("No more chapters yet.")
		time.Sleep(1 * time.Second)
	}
}
