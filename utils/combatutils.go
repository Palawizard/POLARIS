package utils

import (
	"fmt"
	"math"
)

// ApplyDamage rounds incoming damage to the nearest int (min 1 if dmg > 0),
// subtracts it from hp, clamps hp at 0, and returns the applied amount.
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

// ApplyHeal rounds incoming heal to the nearest int (min 1 if heal > 0),
// adds it to hp, clamps at max, and returns the actual amount restored.
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

// HPString prints current/max HP with integer-friendly rounding so UI matches combat math.
func HPString(hp, max float64) string {
	return fmt.Sprintf("%.0f/%.0f", math.Round(hp), math.Round(max))
}
