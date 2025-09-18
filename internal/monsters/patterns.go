package monsters

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
	"time"
)

// rng is a package-level PRNG seeded once.
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// AttackPattern resolves one monster action:
// - shows the header/turn,
// - plays attack SFX + screen flash,
// - rolls damage with a crit chance that trends toward CritEvery,
// - applies damage and reports whether the player survived.
func AttackPattern(player *utils.Player, monster *Monster, turn int) bool {
	utils.ClearScreen()
	fmt.Println("Turn", turn)
	utils.SendTurn(turn)
	PrintHeader(monster)
	fmt.Println("\n")

	msg := monster.AttackMsg
	if msg == "" {
		msg = fmt.Sprintf("The %s attacks you!", monster.Name)
	}
	fmt.Println(msg)
	time.Sleep(2 * time.Second)
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "enemyatk.mp3"))
	utils.Flash(50, 1)

	var dmg float64
	crit := false

	// Crit model: target average cadence is CritEvery.
	// Base chance is 1/CritEvery, with a small “pity” boost growing with SinceLastCrit.
	// Capped, and forced after ~2× cadence without a crit.
	if monster.CritEvery > 0 {
		base := 1.0 / float64(monster.CritEvery)
		boost := float64(monster.SinceLastCrit) / float64(monster.CritEvery)
		p := base * (1.0 + boost)
		if p > 0.95 {
			p = 0.95
		}
		if rng.Float64() < p || monster.SinceLastCrit >= monster.CritEvery*2 {
			crit = true
		}
	}

	if crit {
		mult := monster.CritMultiplier
		if mult <= 0 {
			mult = 2
		}
		dmg = monster.MaxATK * mult
		fmt.Println("Critical hit!")
		time.Sleep(1 * time.Second)
		monster.SinceLastCrit = 0
	} else {
		// Non-crit damage rolls in [0.5, 1.0] × MaxATK.
		dmg = monster.MaxATK * (0.5 + rng.Float64()*0.5)
		if monster.CritEvery > 0 {
			monster.SinceLastCrit++
		}
	}

	applied := utils.ApplyDamage(&player.Health, dmg)
	fmt.Printf("Ouch! %s deals %d damage to %s!\n", monster.Name, applied, player.Name)

	if utils.IsDead(player) {
		fmt.Printf("\nYou have been defeated by the %s.\n", monster.Name)
		time.Sleep(3 * time.Second)
		return false
	}

	fmt.Printf("\nYour HP is now %s hp.\n", utils.HPString(player.Health, player.MaxHealth))
	time.Sleep(2 * time.Second)
	return true
}
