package chapters

import (
	"fmt"
	"projet-red_POLARIS/utils"
	"time"
)

var currentChapter = 1

type chapterFunc func(*utils.Player) bool

var registry = map[int]chapterFunc{
	1: StartChapter1,
}

func StartNextChapter(player *utils.Player) {
	switch currentChapter {
	case 1:
		if StartChapter1(player) {
			currentChapter++
		}
	default:
		fmt.Println("No more chapters yet.")
		time.Sleep(1 * time.Second)
	}
}
