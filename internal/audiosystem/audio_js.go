//go:build js

// Browser build of the audio system: the game keeps calling the same
// functions, but playback is delegated to the page (Web Audio for SFX,
// <audio> for looping music) through the global `polarisAudio` object.
package audiosystem

import (
	"path/filepath"
	"syscall/js"
)

// call invokes polarisAudio[method](args...) if the page provides it.
// A missing bridge or a JS exception must never crash the game.
func call(method string, args ...any) {
	defer func() { _ = recover() }()
	a := js.Global().Get("polarisAudio")
	if a.IsUndefined() || a.IsNull() {
		return
	}
	a.Call(method, args...)
}

// url turns a game asset path ("assets/audio/...") into a URL path.
func url(path string) string { return filepath.ToSlash(path) }

// Init is a no-op in the browser: the page unlocks audio on the first click.
func Init() error { return nil }

// PreloadSFX asks the page to fetch and decode an effect under the given id.
func PreloadSFX(id, path string) error {
	call("preload", id, url(path))
	return nil
}

// PlaySFXCached plays a preloaded SFX by id.
func PlaySFXCached(id string) error {
	call("play", id)
	return nil
}

// PlaySFX plays an SFX by path, preloading it on first use.
func PlaySFX(path string) error {
	call("playPath", url(path))
	return nil
}

// PlayMusicLoop starts background music and loops it, replacing any current track.
func PlayMusicLoop(path string) error {
	call("music", url(path))
	return nil
}

// StopMusic halts the current background track.
func StopMusic() {
	call("stopMusic")
}
