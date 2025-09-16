package chapters

import (
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func StartChapter2(player *utils.Player) bool {
	StartMusic()
	utils.Clearscreen()
	texttoshow := "A divine force floods through you."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: You still have 2 temple left to clear… You filthy degenerate—this will be harder now!"
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "The goddess yeets a goblin at you. She wonders whose face is more tolerable between you or the goblin…"
	utils.ShowText(texttoshow)

	StopMusic()
	goblin := monsters.New("Goblin")
	if won, exit := fightsystem.RunFight(player, goblin, false); exit || !won {
		return false
	}

	StartMusic()
	utils.Clearscreen()
	texttoshow = "Goddess Polaris: Look, a monster just as smelly and slimy as you!"
	utils.ShowText(texttoshow)

	StopMusic()
	slime := monsters.New("Slime")
	if won, exit := fightsystem.RunFight(player, slime, false); exit || !won {
		return false
	}

	StartMusic()
	utils.Clearscreen()
	texttoshow = "Goddess Polaris: Pfft! You’re boring me!"
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "The goddess leaves. You head toward the next temple."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "A delicious bakery smell tickles your nose."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "A giant CHOCOBLAST blocks your path! (Please protect your data…)"
	utils.ShowText(texttoshow)

	StopMusic()
	chocoblast := monsters.New("Chocoblast")
	if won, exit := fightsystem.RunFight(player, chocoblast, true); exit || !won {
		return false
	}

	StartWinningMusic()
	utils.Clearscreen()
	texttoshow = "You reach the goddess’s altar and light the sacred flame."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "You have 1 altars left to relight to regain the favor of Goddess Polaris."
	utils.ShowText(texttoshow)
	StopMusic()

	return true
}
