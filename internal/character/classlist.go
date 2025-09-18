package character

import (
	"strings"
)

// Class groups the static properties of a playable class.
// HP is the starting health; MaxHP is the cap before gear bonuses.
type Class struct {
	ID    string
	Label string
	HP    float64
	MaxHP float64
}

// Classes is the registry of playable classes and their base stats.
var Classes = map[string]Class{
	"Human": {
		ID:    "Human",
		Label: "Human (Normal)",
		HP:    50,
		MaxHP: 100,
	},
	"Elf": {
		ID:    "Elf",
		Label: "Elf (Hard)",
		HP:    40,
		MaxHP: 80,
	},
	"Dwarf": {
		ID:    "Dwarf",
		Label: "Dwarf (Easy)",
		HP:    60,
		MaxHP: 120,
	},
}

// GetClass returns the class matching id. If none is found,
// it returns the zero-value Class (all fields empty/zero).
func GetClass(id string) Class {
	if c, ok := Classes[id]; ok {
		return c
	}
	return Classes["Template"] // zero-value if "Template" is not set
}

// ClassLabel returns the user-facing label for a class id.
// Unknown ids return an empty string.
func ClassLabel(id string) string { return GetClass(id).Label }

// Classlist returns a comma-separated string of all class labels.
// This is meant for simple display in selection prompts.
func Classlist() string {
	var list []string
	for _, c := range Classes {
		list = append(list, c.Label)
	}
	return strings.Join(list, ", ")
}
