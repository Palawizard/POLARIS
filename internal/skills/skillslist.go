package skills

import (
	"math/rand"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type Skill struct {
	ID    string
	Label string
	Price float64
	Apply func(*utils.Player, *monsters.Monster)
}

var Skills = map[string]Skill{
	"Punch":          {ID: "Punch", Label: "Punch", Price: 0, Apply: effectPunch},
	"Fire Ball":      {ID: "Fire Ball", Label: "Fire Ball", Price: 20, Apply: effectFireball},
	"Heal":           {ID: "Heal", Label: "Heal", Price: 15, Apply: effectHeal},
	"Lightning Bolt": {ID: "Lightning Bolt", Label: "Lightning Bolt", Price: 18, Apply: effectLightningBolt},
	"Ice Shard":      {ID: "Ice Shard", Label: "Ice Shard", Price: 15, Apply: effectIceShard},
	"Vampiric Touch": {ID: "Vampiric Touch", Label: "Vampiric Touch", Price: 22, Apply: effectVampiricTouch},
	"Meteor":         {ID: "Meteor", Label: "Meteor", Price: 35, Apply: effectMeteor},
}

// SpellBook adds the given skill to the player's spellbook, incrementing its count by 1.
// If the player's spellbook is currently nil, it will be initialized.
// If the skill doesn't exist or its Apply func is nil, it returns false.
// Otherwise, it adds the skill to the player's spellbook and returns true.
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

// Cast casts the given skill
// If the skill doesn't exist or its Apply func is nil, it returns false.
// Otherwise, it casts the skill and returns true.
func Cast(id string, player *utils.Player, monster *monsters.Monster) bool {
	it, ok := Skills[id]
	if !ok || it.Apply == nil {
		return false
	}
	it.Apply(player, monster)
	return true
}
