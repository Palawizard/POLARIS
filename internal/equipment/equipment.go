package equipment

import (
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

// ensureEquipped makes sure the Equipped map exists before we read/write it.
// Keeps callers simple and avoids nil-map panics.
func ensureEquipped(m *map[string]string) {
	if *m == nil {
		*m = make(map[string]string)
	}
}

// UnequipSlot removes the item currently worn in the given slot.
// It also rolls back the MaxHealth bonus from that piece and clamps current
// Health if it now sits above the reduced MaxHealth.
// Returns true if something was unequipped, false if the slot was empty.
func UnequipSlot(p *utils.Player, slot string) bool {
	ensureEquipped(&p.Equipped)
	prev, ok := p.Equipped[slot]
	if !ok || prev == "" {
		return false
	}
	p.MaxHealth -= BonusOf(prev)
	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}
	delete(p.Equipped, slot)
	return true
}

// Equip puts the specified equipment on the player, replacing any item already
// in that slot. It adjusts MaxHealth by removing the old piece's bonus (if any)
// then adding the new one, and clamps current Health if MaxHealth goes down.
// Returns false if the player doesn't own the item, the id is unknown, or the
// slot can't be determined.
func Equip(p *utils.Player, id string) bool {
	if p == nil {
		return false
	}
	if p.Equipment == nil || p.Equipment[id] <= 0 {
		return false
	}
	slot := SlotOf(id)
	if slot == "" {
		return false
	}

	ensureEquipped(&p.Equipped)

	// If something is already equipped in that slot, remove its bonus first.
	if cur, ok := p.Equipped[slot]; ok && cur != "" {
		p.MaxHealth -= BonusOf(cur)
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
	}

	// Apply new piece and its bonus.
	p.Equipped[slot] = id
	p.MaxHealth += BonusOf(id)

	_ = audiosystem.PlaySFX(filepath.Join("assets", "audio", "sfx", "equip.mp3"))
	return true
}

// EquippedBonus sums the MaxHealth bonus provided by all currently equipped items.
// Safe on nil players or empty equipment.
func EquippedBonus(p *utils.Player) float64 {
	if p == nil || p.Equipped == nil {
		return 0
	}
	b := 0.0
	for _, id := range p.Equipped {
		b += BonusOf(id)
	}
	return b
}
