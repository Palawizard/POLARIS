package chapters

import (
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

// StartChapter2 advances the story, presents two normal encounters (Goblin, Slime),
// then a boss (Chocoblast). Returns false if the player loses or exits at any point.
func StartChapter2(player *utils.Player) bool {
	StartMusic()
	utils.ClearScreen()

	textToShow := "A divine force floods through you."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "Goddess Polaris: You still have 2 temples left to clear… You filthy degenerate—this will be harder now!"
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "The goddess yeets a goblin at you. She wonders whose face is more tolerable between you or the goblin…"
	utils.ShowText(textToShow)

	// Fight #1 — Goblin
	StopMusic()
	goblin := monsters.New("Goblin")
	if won, exit := fightsystem.RunFight(player, goblin, false); exit || !won {
		return false
	}

	StartMusic()
	utils.ClearScreen()
	textToShow = "Goddess Polaris: Look, a monster just as smelly and slimy as you!"
	utils.ShowText(textToShow)

	// Fight #2 — Slime
	StopMusic()
	slime := monsters.New("Slime")
	if won, exit := fightsystem.RunFight(player, slime, false); exit || !won {
		return false
	}

	StartMusic()
	utils.ClearScreen()
	textToShow = "Goddess Polaris: Pfft! You’re boring me!"
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "The goddess leaves. You head toward the next temple."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "A delicious bakery smell tickles your nose."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "A giant CHOCOBLAST blocks your path! (Please protect your data…)"
	utils.ShowText(textToShow)

	// Boss — Chocoblast
	StopMusic()
	chocoblast := monsters.New("Chocoblast")
	if won, exit := fightsystem.RunFight(player, chocoblast, true); exit || !won {
		return false
	}

	StartWinningMusic()
	utils.ClearScreen()
	textToShow = "You reach the goddess’s altar and light the sacred flame."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "You have 1 altars left to relight to regain the favor of Goddess Polaris."
	utils.ShowText(textToShow)
	StopMusic()

	return true
}
