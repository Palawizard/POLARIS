package character

import (
	"strings"
)

type Class struct {
	ID    string
	Label string
	HP    float64
	MAXHP float64
}

var Classes = map[string]Class{
	"Human": {
		ID:    "Human",
		Label: "Human",
		HP:    50,
		MAXHP: 100,
	},
	"Elf": {
		ID:    "Elf",
		Label: "Elf",
		HP:    40,
		MAXHP: 80,
	},
	"Dwarf": {
		ID:    "Dwarf",
		Label: "Dwarf",
		HP:    60,
		MAXHP: 120,
	},
}

// GetClass returns a Class from the given id string.
// If the Class doesn't exist, it returns the "Template" Class.
func GetClass(id string) Class {
	if c, ok := Classes[id]; ok {
		return c
	}
	return Classes["Template"]
}

// ClassLabel returns the label of the given class id.
// If the class doesn't exist, it returns an empty string.
func ClassLabel(id string) string { return GetClass(id).Label }

// Classlist returns a comma-separated string of all the Class IDs available in the Classes map.
func Classlist() string {
	var list []string
	for _, c := range Classes {
		list = append(list, c.ID)
	}
	result := strings.Join(list, ", ")
	return result
}
