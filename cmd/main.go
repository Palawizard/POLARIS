package main

import (
	"path/filepath"

	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/game"
)

// main is the entry point of the application.
func main() {
	_ = audiosystem.PreloadSFX("select",
		filepath.Join("assets", "audio", "sfx", "select.wav"))

	game.InitGame()
}
