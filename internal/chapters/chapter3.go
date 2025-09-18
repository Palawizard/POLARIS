package chapters

import (
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

// StartChapter3 runs three encounters (Flame, Annoying Dog, Flowers boss)
// with short narrative beats in between. Returns false if the player loses or exits.
func StartChapter3(player *utils.Player) bool {
	StartMusic()

	utils.ClearScreen()
	textToShow := "You feel a gentle magical warmth."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "Goddess Polaris: yawns You know what you have to do, right?"
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "The goddess leaves."
	utils.ShowText(textToShow)

	// Fight #1 — Flame
	StopMusic()
	flame := monsters.New("Flame")
	if won, exit := fightsystem.RunFight(player, flame, false); exit || !won {
		return false
	}

	StartMusic()
	utils.ClearScreen()
	textToShow = "You hear barking in the distance."
	utils.ShowText(textToShow)

	// Fight #2 — Annoying Dog
	StopMusic()
	dog := monsters.New("Annoying Dog")
	if won, exit := fightsystem.RunFight(player, dog, false); exit || !won {
		return false
	}

	StartMusic()
	utils.ClearScreen()
	textToShow = "You approach the temple, but the goddess steps in to block your way."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "Goddess Polaris: Wait! I think the flower is back for revenge… Good luck!"
	utils.ShowText(textToShow)

	// Boss — Flowers
	StopMusic()
	flowers := monsters.New("Flowers")
	if won, exit := fightsystem.RunFight(player, flowers, true); exit || !won {
		return false
	}

	StartWinningMusic()
	utils.ClearScreen()
	textToShow = "You reach the goddess’s altar and light the sacred flame."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "You have 0 altars left to relight to regain the favor of Goddess Polaris."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "She's waiting for you..."
	utils.ShowText(textToShow)
	StopMusic()

	return true
}
