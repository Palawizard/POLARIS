package chapters

import (
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func StartChapter3(player *utils.Player) bool {
	StartMusic()

	utils.Clearscreen()
	texttoshow := "You feel a gentle magical warmth."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: yawns You know what you have to do, right?"
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "The goddess leaves."
	utils.ShowText(texttoshow)

	StopMusic()
	flame := monsters.New("Flame")
	if won, exit := fightsystem.RunFight(player, flame, false); exit || !won {
		return false
	}

	StartMusic()
	utils.Clearscreen()
	texttoshow = "You hear barking in the distance."
	utils.ShowText(texttoshow)

	StopMusic()
	dog := monsters.New("Annoying Dog")
	if won, exit := fightsystem.RunFight(player, dog, false); exit || !won {
		return false
	}

	StartMusic()
	utils.Clearscreen()
	texttoshow = "You approach the temple, but the goddess steps in to block your way."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: Wait! I think the flower is back for revenge… Good luck!"
	utils.ShowText(texttoshow)

	StopMusic()
	flowers := monsters.New("Flowers")
	if won, exit := fightsystem.RunFight(player, flowers, true); exit || !won {
		return false
	}

	StartWinningMusic()
	utils.Clearscreen()
	texttoshow = "You reach the goddess’s altar and light the sacred flame."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "You have 0 altars left to relight to regain the favor of Goddess Polaris."
	utils.ShowText(texttoshow)
	StopMusic()

	return true
}
