package character

import (
	"strings"
)

type Class struct {
	ID    string
	Label string
	HP    int
	MAXHP int
}

var Classes = map[string]Class{
	"Template": {
		ID:    "Template",
		Label: "Template",
		HP:    10,
		MAXHP: 10,
	},
}

func GetClass(id string) Class {
	if c, ok := Classes[id]; ok {
		return c
	}
	return Classes["Template"]
}

func ClassLabel(id string) string { return GetClass(id).Label }

func Classlist() string {
	var list []string
	for _, c := range Classes {
		list = append(list, c.ID)
	}
	result := strings.Join(list, ", ")
	return result
}
