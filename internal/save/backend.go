//go:build !js

package save

// The terminal build has no save system (unchanged from the original game).

func read() string { return "" }
func write(string) {}
func remove()      {}
