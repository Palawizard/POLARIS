package main

import (
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/game"
)

func main() {

	_ = audiosystem.PreloadSFX("select", filepath.Join("internal", "audiosystem", "sfx", "select.wav"))

	game.InitGame()
}
