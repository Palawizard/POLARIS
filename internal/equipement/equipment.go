package equipement

import (
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/utils"
)

func ensureEquipped(m *map[string]string) {
	if *m == nil {
		*m = make(map[string]string)
	}
}

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
	if cur, ok := p.Equipped[slot]; ok && cur != "" {
		p.MaxHealth -= BonusOf(cur)
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
	}
	p.Equipped[slot] = id
	p.MaxHealth += BonusOf(id)
	_ = audiosystem.PlaySFX(filepath.Join("internal", "audiosystem", "sfx", "equip.mp3"))
	return true
}

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
