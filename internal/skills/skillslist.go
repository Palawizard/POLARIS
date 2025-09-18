package skills

import (
	"math/rand"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

// rng is a package-local PRNG used for crits and random effects.
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Skill represents a learnable action: a name, price, mana cost, and effect.
type Skill struct {
	ID       string                                 // internal ID
	Label    string                                 // name shown to players
	Price    float64                                // shop price (0 = not sold)
	ManaCost float64                                // MP required to cast
	Apply    func(*utils.Player, *monsters.Monster) // effect on use
}

// Skills lists all skills available in the game.
var Skills = map[string]Skill{
	"Punch":          {ID: "Punch", Label: "Punch", Price: 0, ManaCost: 0, Apply: effectPunch},
	"Fire Ball":      {ID: "Fire Ball", Label: "Fire Ball", Price: 20, ManaCost: 8, Apply: effectFireball},
	"Heal":           {ID: "Heal", Label: "Heal", Price: 15, ManaCost: 10, Apply: effectHeal},
	"Lightning Bolt": {ID: "Lightning Bolt", Label: "Lightning Bolt", Price: 18, ManaCost: 12, Apply: effectLightningBolt},
	"Ice Shard":      {ID: "Ice Shard", Label: "Ice Shard", Price: 15, ManaCost: 10, Apply: effectIceShard},
	"Vampiric Touch": {ID: "Vampiric Touch", Label: "Vampiric Touch", Price: 22, ManaCost: 14, Apply: effectVampiricTouch},
	"Meteor":         {ID: "Meteor", Label: "Meteor", Price: 35, ManaCost: 25, Apply: effectMeteor},
}

// SpellBook grants one copy of the given skill to the player.
func SpellBook(id string, player *utils.Player) bool {
	it, ok := Skills[id]
	if !ok || it.Apply == nil {
		return false
	}
	if player.Skills == nil {
		player.Skills = make(map[string]int)
	}
	player.Skills[id]++
	return true
}

// Cast spends mana and executes the selected skill’s effect.
func Cast(id string, player *utils.Player, monster *monsters.Monster) bool {
	it, ok := Skills[id]
	if !ok || it.Apply == nil {
		return false
	}
	if player.Mana < it.ManaCost {
		return false
	}
	player.Mana -= it.ManaCost
	it.Apply(player, monster)
	return true
}
