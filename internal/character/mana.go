package character

import "projet-red_POLARIS/utils"

func ManaRegen(player *utils.Player, turn int) {
	if player.Mana < player.MaxMana {
		if turn%2 == 0 {
			player.Mana += 10
			if player.Mana > player.MaxMana {
				player.Mana = player.MaxMana
			}
		}
	}
}
