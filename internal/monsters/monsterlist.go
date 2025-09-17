package monsters

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

type Monster struct {
	Name           string
	Health         float64
	MaxHealth      float64
	MaxATK         float64
	EXPtogive      float64
	Coinstogive    float64
	Initiative     float64
	AttackMsg      string
	CritEvery      int
	CritMultiplier float64
	Loot           string
	SinceLastCrit  int
	Art            string
}

var Monsters = map[string]Monster{
	"Goblin": {
		Name:           "Goblin",
		Health:         42,
		MaxHealth:      42,
		MaxATK:         7,
		EXPtogive:      25,
		Coinstogive:    18,
		Initiative:     6,
		AttackMsg:      "The goblin darts in with a rusty shiv!",
		CritEvery:      4,
		CritMultiplier: 1.7,
		Loot:           "Boar Leather",
		Art: `
             ,      ,
            /(.-""-.)\
        |\  \/      \/  /|
        | \ / =.  .= \ / |
        \( \   o\/o   / )/
         \_, '-/  \-' ,_/
           /   \__/   \
           \ \__/\__/ /
         ___\ \|--|/ /___
       /     \      /     \
      /       '----'       \
              `,
	},
	"Flower": {
		Name:           "Flower",
		Health:         12,
		MaxHealth:      12,
		MaxATK:         1.2,
		EXPtogive:      6,
		Coinstogive:    6,
		Initiative:     1,
		AttackMsg:      "Razor petals whirl toward you!",
		CritEvery:      3,
		CritMultiplier: 1.5,
		Loot:           "Crow Feather",
		Art: `
              .--.
            .'_\/_'.
            '. /\ .'
              "||"
               || /\
            /\ ||//\)
           (/\\||/
        ______\||/_______
`,
	},
	"Skeleton": {
		Name:           "Skeleton",
		Health:         26,
		MaxHealth:      26,
		MaxATK:         6,
		EXPtogive:      22,
		Coinstogive:    22,
		Initiative:     4,
		AttackMsg:      "Bones clack—its chipped blade carves the air!",
		CritEvery:      4,
		CritMultiplier: 1.6,
		Loot:           "Potion",
		Art: `
              .-.
             (o.o)
              |=|
             __|__
           //.=|=.\\
          // .=|=. \\
          \\ .=|=. //
           \\(_=_)//
            (:| |:)
             || ||
             () ()
             || ||
             || ||
            ==' '==
`,
	},
	"Boss Potato": {
		Name:           "Boss Potato",
		Health:         70,
		MaxHealth:      70,
		MaxATK:         10,
		EXPtogive:      60,
		Coinstogive:    60,
		Initiative:     8,
		AttackMsg:      "The Boss Potato belly-flops with starchy fury!",
		CritEvery:      4,
		CritMultiplier: 1.8,
		Loot:           "Poison",
		Art: `

⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠚⠛⠉⠉⠉⠳⢦⡀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⣿⠁⠀⠀⠀⠀⠀⠀⠀⢹⡄⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣟⣿⠀⠀⠀⠀⠀⠀⠀⠀⠈⣷⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡟⠘⠉⠱⣿⠀⠀⠀⠀⢠⣶⡄⣿⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠃⠀⠀⠀⠀⠀⠉⠙⠃⠀⠀⠀⢸⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢸⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⡇⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣾⣤⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣷⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⣤⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⠀⠶⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⠇⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⣶⣇⣰⡄⠀⠀⠀⠀⢀⣠⣾⡿⠏⠀⢤⡀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⠛⠿⠿⠿⠿⡿⢿⡛⡩⠤⠴⠒⠛⠁`,
	},
	"Chocoblast": {
		Name:           "Chocoblast",
		Health:         120,
		MaxHealth:      120,
		MaxATK:         16,
		EXPtogive:      85,
		Coinstogive:    65,
		Initiative:     5,
		AttackMsg:      "A sugary shockwave detonates—cocoa shrapnel flies!",
		CritEvery:      4,
		CritMultiplier: 2.0,
		Loot:           "Chocolatine",
		Art: `
                       ░▒▒▒▒▒░░░                  
                   ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░              
               ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▒▒▒▒▒▒           
            ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▒▒▒▒         
          ░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░       
         ▒░▓▓▒░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▒▒▒▒▒▒▒▒▒▒░       
        ░▒▒▒░░░░▓▒░▒▒▒▒▓▓▒▒▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒░       
        ▒░░░░▒▒▒░░░▓░░▒▒▒▒▓▒▒▒▒▒▒▒▒▒▒▒▒▒▒         
        ░▓▓░░▒░░▒▒▒░░▒░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒░          
         ░▒▓▓░░░▒░░▒▒░░░░▒▒▒▒▒▒▒▒▒▒▒▒             
           ░░░▓▓▓░▒░▒░░▒░▒▒▒▒▒▒▒▒▒░               
             ░▓▓▓▓▒▒▓▒░▓░░▒▒▒▒▒                   
                  ░░▒░░░░░░`,
	},
	"Slime": {
		Name:           "Slime",
		Health:         70,
		MaxHealth:      70,
		MaxATK:         12,
		EXPtogive:      35,
		Coinstogive:    28,
		Initiative:     3,
		AttackMsg:      "A glob of acidic ooze splashes toward you!",
		CritEvery:      4,
		CritMultiplier: 1.6,
		Loot:           "Potion",
		Art: `
		⠀⠀⠀⠀⠀⠀⠀⠀⣤⣤⣤⣤⣤⣤⣤⣤⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀
		⠀⠀⠀⢀⣠⣤⣤⣾⠟⠛⠛⠛⠛⠛⠛⠛⠻⣦⣤⣤⣄⠀⠀⠀⠀⠀
		⠀⠀⣠⣾⡿⠟⠛⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⠟⠿⣿⣦⣄⠀⠀⠀
		⣠⣼⡿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⠀⠀⠀
		⣿⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢿⣿⣶⡄⠀
		⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⡀⠀⠀⠀⣀⣀⠀⠀⠀⣿⣿⡇⠀
		⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀⠀⣿⣿⠀⠀⠀⣯⣿⡇⠀
		⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⠿⠇⠀⠀⠀⠿⣿⠤⣤⣧⣧⣿⡇⡀
		⠿⢿⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠂⠀⡀⠀⠀⠀⠁⣠⣼⣿⠟⠃
		⠀⠀⠿⢷⣄⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣀⣁⣀⣀⣀⣤⣿⠿⠁⠀⠀
		⠀⠀⠀⠘⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠁⠁`,
	},
	"Flame": {
		Name:           "Flame",
		Health:         85,
		MaxHealth:      85,
		MaxATK:         18,
		EXPtogive:      55,
		Coinstogive:    55,
		Initiative:     7,
		AttackMsg:      "The flame burns you !",
		CritEvery:      5,
		CritMultiplier: 1.8,
		Loot:           "Fire Ball",
		Art: `
		   )
		  ) \
		 / ) (
		 \(_)/  
`,
	},
	"Annoying Dog": {
		Name:           "Annoying Dog",
		Health:         95,
		MaxHealth:      95,
		MaxATK:         20,
		EXPtogive:      70,
		Coinstogive:    70,
		Initiative:     9,
		AttackMsg:      "The annoying dog annoys you !",
		CritEvery:      5,
		CritMultiplier: 2.0,
		Loot:           "Potion",
		Art: `
		░░░░░░░░░░░░░░░░░░░░
		░▄▀▄▀▀▀▀▄▀▄░░░░░░░░░
		░█░░░░░░░░▀▄░░░░░░▄░
		█░░▀░░▀░░░░░▀▄▄░░█░█
		█░▄░█▀░▄░░░░░░░▀▀░░█
		█░░▀▀▀▀░░░░░░░░░░░░█
		█░░░░░░░░░░░░░░░░░░█
		█░░░░░░░░░░░░░░░░░░█
		░█░░▄▄░░▄▄▄▄░░▄▄░░█░
		░█░▄▀█░▄▀░░█░▄▀█░▄▀░
		░░▀░░░▀░░░░░▀░░░▀░░░
`,
	},
	"Flowers": {
		Name:           "Flowers",
		Health:         180,
		MaxHealth:      180,
		MaxATK:         36,
		EXPtogive:      180,
		Coinstogive:    180,
		Initiative:     11,
		AttackMsg:      "A tide of thorns lashes out in unison!",
		CritEvery:      4,
		CritMultiplier: 2.2,
		Loot:           "Poison",
		Art: `
                    _
                  _(_)_                          wWWWw   _
      @@@@       (_)@(_)   vVVVv     _     @@@@  (___) _(_)_
     @@()@@ wWWWw  (_)\    (___)   _(_)_  @@()@@   Y  (_)@(_)
      @@@@  (___)     |/    Y    (_)@(_)  @@@@   \|/   (_)\
       /      Y       \|    \|/    /(_)    \|      |/      |
    \ |     \ |/       | / \ | /  \|/       |/    \|      \|/
    \\|//   \\|///  \\\|//\\\|/// \|///  \\\|//  \\|//  \\\|// 
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
`,
	},
	"Polaris": {
		Name:           "Polaris",
		Health:         320,
		MaxHealth:      320,
		MaxATK:         40,
		EXPtogive:      360,
		Coinstogive:    300,
		Initiative:     13,
		AttackMsg:      "Polaris rends reality around you!",
		CritEvery:      4,
		CritMultiplier: 2.3,
		Loot:           "Chocolatine",
		Art: `
                 )
               /  )
              /  / )
         -   /  / /
            '  / / -
           / _/ / /
     _    / _/_, /          ,
   + $$$ / _/_/_/          \       |
 /- + $$/ _/_/_/      /
 \ _ $$/'_/_/    .    ______   _
		\ (  / ___,_____ _ _____,
		|   (|/_,_,__ ________/
    |.   |''_,_______)
     \   (_
      \  / |-._
       \.' /|/ \_._
       /_/   _/    /-'__
         \     \'       \.___
          '.   /,     |_/_   |._
            \ / )   '.     '_/, )
             (_(     -\_   /  \ \
                \__      |-'   |/
                  \._  /_/_
                     \_/\' )
                         \ |
                         |/`,
	},
}

func PrintHeader(m *Monster) {
	if m == nil {
		return
	}
	if m.Art != "" {
		fmt.Println(m.Art)
	}
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))
}

func New(id string) *Monster {
	m := Monsters[id]
	return &m
}
