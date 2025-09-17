package utils

import (
	"fmt"
	"math"
)

func ApplyDamage(hp *float64, dmg float64) int {
	if dmg <= 0 {
		return 0
	}
	applied := int(math.Round(dmg))
	if applied == 0 {
		applied = 1
	}
	*hp -= float64(applied)
	if *hp < 0 {
		*hp = 0
	}
	return applied
}

func ApplyHeal(hp *float64, max float64, heal float64) int {
	if heal <= 0 {
		return 0
	}
	applied := int(math.Round(heal))
	if applied == 0 {
		applied = 1
	}
	before := *hp
	*hp += float64(applied)
	if *hp > max {
		*hp = max
	}
	return int(math.Round(*hp - before))
}

func HPString(hp, max float64) string {
	return fmt.Sprintf("%.0f/%.0f", math.Round(hp), math.Round(max))
}
