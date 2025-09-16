package chapters

import (
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func StartFinalBoss(player *utils.Player) bool {
	utils.Clearscreen()
	texttoshow := "Silence falls. The air crackles with old magic."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: So you made it..."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: You rekindled my flames, felled my beasts… and defiled my altar (I haven’t forgotten)."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: One final trial: prove your will is greater than your luck."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "The ground hums. Constellations spin into a circle around you."
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: Kneel—or fight. Choose."
	utils.ShowText(texttoshow)

	StopMusic()
	polaris := monsters.New("Polaris")
	if won, exit := fightsystem.RunFight(player, polaris, true); exit || !won {
		return false
	}

	StartWinningMusic()
	utils.Clearscreen()
	texttoshow = "Well played ! You finished the game ! (For now)"
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Thank you for playing Polaris !"
	utils.ShowText(texttoshow)

	StopMusic()
	return true
}
