package chapters

import (
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

// StartChapter1 runs the opening story beats and three sequential fights.
// Returns false if the player loses or exits during any encounter.
func StartChapter1(player *utils.Player) bool {
	utils.ClearScreen()

	// Opening cutscene: music + intro lines
	StartMusic()
	textToShow := "Goddess Polaris: No! Hey! Are you serious?! You’re peeing on my altar! >:|"
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "How dare you defy the goddess of life and death? Cursed be you! You gross creep! As punishment, you must rekindle the sacred flame in all my temples!"
	utils.ShowText(textToShow)

	// Nicolas enft
	utils.ClearScreen()
	utils.PrintASCII(`
                ▒▓▓▓▓▓▓█▓█░                
           ░▓▓▒▓██▓██████████▓▒█           
         ▓▓▓▓▓█████████████████░ ▓         
        ░▓▓▓██▓█████████████████░          
       ▓▒▓▓█▓█▓▓▓▓▓▓▓▓▓▓▓▓▓▓███████        
      ▓▓█▓█▓▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓█████       
     ░▓███▒░░░░▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓████▓      
     ▒▓██▒░░░░▒░▒▒░▒▒▒▒▒▒▓▓▓▓▓▓▓█████      
     ░███░░░░░▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓█████      
      ▓▓▓░░░▓▒▒█▓▒▓▒▓▓▓▓▓▓██████▓████      
      ▓▓▓▒▓██████▓▓▒▓▓████▓▓▓▓▓▓▓▓███      
     ░▓▒░▒▒▒▒▓▓█▓▓▒████▓▓████▓▓▓▓▓██▓      
      ░▓▒▒▓▒▒█▓██▓▒█▒▒▓████████▓▓▓▓█       
       █░░▒▒▒▒▓▓▒▒█░░▓▓█▓▓▓▓▓▓▓▓▓▓▓█       
       ▓█░░░▒▒▒▒▒█▒░▒▓▓▓▓▓██████▓▓▓█       
       ▒░░░▒▒▒▒▒▒▒░░▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓█       
        ░░▒░▒▒▒▓▓░░▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▒       
        ░░▒▒▒▒▒▓▒▒▓▓▒▓▓██▓▓▓▓▓▓▓▓▓▓        
         ▒▒▒▒▒▒▒▒▒▒░▓▓▒███▓▓▓▓▓▓▓██        
         ▒▒▒▒▒▒▒▒░▒▒▒▒▓▓▓▓▓▓█▓▓▓▓█▒        
         ▒▒▒▒▒▓█▒░▒▒▓▓▓▓█▓▓▓▓▓▓▓▓█▓▓▒░     
        ░▓▓▒▒▒▒░░░▒▒▒▓▓▓▓▓▓▓▓▓▓▓██▓██▓▓▓░  
       ░░▓▓▒▒▓▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓█▓████▓▓█▓▓▓▒ 
    ░░░░▒█▒░▓▓▒▒▒░░░▒▒▒▓▓▓▓█▓▓██▓██▓██▓▓▓▓▒
░░░░░░░░░▓▒░▒▓▓▒░▒▒▒▓▓▓▓▓▓█▓██▓▓▓██▓█▓▓▓▓▓▒
░▒▓▒░░░░░░█▒▒▒▓▒▓█▓▓▓▓▓▓▓███▓▓▓▓▓██▓▓▓█▓▓▓▒
▒▒▒░░▒▓░░░▒█▒▒█▒░░░▓▒████▓▓▓▓▓▓▓▓██▓▓▓▒▓▓▓▒
▒░░░▒▓▓▒░░░░▒▓▓█▓▒▒░░▒▒▒▒▒▒▒▓▓▓▓▓▓▒▒▓▒▓▒▓▒▓
`)

	textToShow = "You look at the goddess like this"
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "Goddess Polaris: Oh, I forgot to mention—you’ll have a few enemies to fight; otherwise it’d be too easy, owo."
	utils.ShowText(textToShow)

	// Fight 1 — Flower
	StopMusic()
	flower := monsters.New("Flower")
	if won, exit := fightsystem.RunFight(player, flower, false); exit || !won {
		return false
	}

	utils.ClearScreen()
	StartMusic()
	textToShow = "Goddess Polaris: Mmph! rolls eyes We’ll see if you handle what comes next just as well!"
	utils.ShowText(textToShow)

	// Fight 2 — Skeleton
	StopMusic()
	skeleton := monsters.New("Skeleton")
	if won, exit := fightsystem.RunFight(player, skeleton, false); exit || !won {
		return false
	}

	utils.ClearScreen()
	StartMusic()
	textToShow = "You near the temple… but a horrible monster, reeking of raclette, is guarding it…"
	utils.ShowText(textToShow)

	// Fight 3 — Boss Potato
	StopMusic()
	potato := monsters.New("Boss Potato")
	if won, exit := fightsystem.RunFight(player, potato, true); exit || !won {
		return false
	}

	// Chapter wrap-up and progression note
	StartWinningMusic()
	utils.ClearScreen()
	textToShow = "You reach the goddess’s altar and light the sacred flame."
	utils.ShowText(textToShow)

	utils.ClearScreen()
	textToShow = "You have 2 altars left to relight to regain the favor of Goddess Polaris."
	utils.ShowText(textToShow)
	StopMusic()

	return true
}
