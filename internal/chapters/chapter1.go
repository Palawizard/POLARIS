package chapters

import (
	"fmt"
	"path/filepath"
	"projet-red_POLARIS/internal/audiosystem"
	"projet-red_POLARIS/internal/fightsystem"
	"projet-red_POLARIS/internal/monsters"
	"projet-red_POLARIS/utils"
)

func StartMusic() {
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "chapters.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}
}

func StartWinningMusic() {
	if err := audiosystem.Init(); err != nil {
		fmt.Println("audio init error:", err)
	}
	musicPath := filepath.Join("internal", "audiosystem", "music", "chapterwon.mp3")
	if err := audiosystem.PlayMusicLoop(musicPath); err != nil {
		fmt.Println("play loop error:", err)
	}
}

func StopMusic() {
	audiosystem.StopMusic()
}

func StartChapter1(player *utils.Player) bool {
	utils.Clearscreen()
	StartMusic()
	texttoshow := "Goddess Polaris: No! Hey! Are you serious?! You’re peeing on my altar! >:|"
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "How dare you defy the goddess of life and death? Cursed be you! You gross creep! As punishment, you must rekindle the sacred flame in all my temples!"
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	utils.PrintASCII(`

                                .=*#=#%-+%##                               
                            -+*%#%+%%@@@%@%@@@@%*:                         
                        .*=*@%%@%@@%@%@@@@@@@@@@@@#=                       
                       **+##@#@%@@@@@@@@@@@@@@@@##@  +                     
                      *%%@*%#%%@@@@%%%%%%%%%@@@@@@@%-=                     
                    =#=*#%*%#*+++=----=-==++**##%%@@@@@                    
                   -%*#%%%*-:....::::----===++**##%@@@@*                   
                   ++@%@*.   ..::::--:-----=+++***#@@@@@@                  
                   ##%@%.     . ..-:----::===+*****#@@@@@                  
                   *@@@:    . ..::--------==+**#***#@@@@@                  
                    #**.   ..:-:::--=====++##@%@%@#*@@@@@+                 
                    %#* ..+%*@@@*+++*+**####@#@@@@@@@@@@@=                 
                    ##%*:=-+++***+=@@@@@@####*+***++*%@@@                  
                    #* .:--.%@@@#*=:@-+@@%%%@@@@@@#++*#+                   
                     *..--*-=%%#*=:@ .-*@%%%%%#######**@                   
                     %    .:--==:. @ .=##@#%#####*****#@                   
                     =@    ....:.@=  :=*##%@@@@@@@@%***@                   
                     -:. . :-::--:. .---*##*#%####*****@                   
                      +  ...::=++= .::-=+*#####********%                   
                       . ..:--=+=-.-*-=+@++%##*********                    
                       ..::=-===:.. .+**#%%##**+++***#@                    
                       *..:::--:...=:+#++#%%##%*+++**#@                    
                        =..::=--. ::--++******##*++*#@*                    
                        *::..:*+:. .-=+***%%##**+++*%@*#*=                 
                       :*+-.... ...:-==+*++*****%**#@@%%%%*++.             
                       -%:*:+=:..:.:-==++*****#@%%@%#@%%@@@###+            
                   =   +@  ##.-.:.  ..::=+*####%%@@**@@%@##%**#=           
              .        -%. -+##--...=--==*#%%%%@@@#*#%@#%##%*#%*=          
                   =    =+ .=@ #.=*-*-+***#%@@@@##**%%@#%*#@**%+++         
           :  .:.       :+=.=+% .=@%#%%%@@@@@@**++*#%%%++*#+++#+++*=-      
          .::::   .=*    :=*-+@+    ..-+*#+====++*###%*=+*===%+=+*#*+=-    
         .--..   -=**=   . -+**@#:.....::---===+++*#*#-=*:==++==#***#-+    
     .  .-=:    -+**+:     .:--*@@=::.:::----=++==+*.:-+:==%+=%#*+=++=+    
    :   -+=.  .=+*#*-.    ..::+.:@@*=--::::::----=  .-.---#+=###=--===-    
       .+*:  -:-*##+:    # :- . ::-+=+**=-::::*@     :-=#=-#**++=:::---    
       :*+  +*::*##=.  . .%# .+==:#%*+=+@%@@@     . :-#=-#*+==+=-:.: -:    
       -*- -##:-*#+:.  .. %@@#. : *=*:--+=      .:=*+-@*+======-::..: -    
       -*- +##--##=::  ...@@ %+   +---+     :-:-**@@@@@----===--:.  ..=    
      :-+:.*#%=:**--..  ...#:.#:  :+:-=+++=+== --*@@@@=-:---=-:-:          
    :.=.-=+##%*.=++-:.  .:.@@%@-..==+--===-:::=-+ @@@:.:::-=-:..:     =
`)

	texttoshow = "You look at the goddess like this"
	utils.ShowText(texttoshow)

	utils.Clearscreen()
	texttoshow = "Goddess Polaris: Oh, I forgot to mention—you’ll have a few enemies to fight; otherwise it’d be too easy, owo."
	utils.ShowText(texttoshow)

	StopMusic()
	flower := monsters.New("Flower")
	if won, exit := fightsystem.RunFight(player, flower, false); exit || !won {
		return false
	}

	utils.Clearscreen()
	StartMusic()
	texttoshow = "Goddess Polaris: Mmph! rolls eyes We’ll see if you handle what comes next just as well!"
	utils.ShowText(texttoshow)

	StopMusic()
	skeleton := monsters.New("Skeleton")
	if won, exit := fightsystem.RunFight(player, skeleton, false); exit || !won {
		return false
	}

	utils.Clearscreen()
	StartMusic()
	texttoshow = "You near the temple… but a horrible monster, reeking of raclette, is guarding it…"
	utils.ShowText(texttoshow)

	StopMusic()
	potato := monsters.New("Boss Potato")
	if won, exit := fightsystem.RunFight(player, potato, true); exit || !won {
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
