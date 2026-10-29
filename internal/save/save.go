// Package save stores the current run so it can be resumed later.
// Only the browser build has a storage backend (localStorage); in the
// terminal build Load always reports "no save" and Store does nothing,
// so the original game flow is unchanged.
package save

import (
	"encoding/json"

	"projet-red_POLARIS/utils"
)

// version is bumped if the saved format ever changes incompatibly.
const version = 1

// Data is what gets persisted: the player and the story progress.
type Data struct {
	Version int          `json:"version"`
	Chapter int          `json:"chapter"`
	Player  utils.Player `json:"player"`
}

// Store persists the player and current chapter.
func Store(p *utils.Player, chapter int) {
	if p == nil {
		return
	}
	b, err := json.Marshal(Data{Version: version, Chapter: chapter, Player: *p})
	if err != nil {
		return
	}
	write(string(b))
}

// Load returns the saved run, if any.
func Load() (Data, bool) {
	raw := read()
	if raw == "" {
		return Data{}, false
	}
	var d Data
	if err := json.Unmarshal([]byte(raw), &d); err != nil || d.Version != version || d.Player.Name == "" {
		return Data{}, false
	}
	// Maps must never be nil: the game writes into them directly.
	if d.Player.Skills == nil {
		d.Player.Skills = map[string]int{}
	}
	if d.Player.Inventory == nil {
		d.Player.Inventory = map[string]int{}
	}
	if d.Player.Equipment == nil {
		d.Player.Equipment = map[string]int{}
	}
	if d.Player.Equipped == nil {
		d.Player.Equipped = map[string]string{}
	}
	if d.Chapter < 1 {
		d.Chapter = 1
	}
	return d, true
}

// Clear deletes the saved run.
func Clear() { remove() }
