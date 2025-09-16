package monsters

import "fmt"

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
	Art            string
}

var Monsters = map[string]Monster{
	"Goblin": {
		Name:           "Goblin",
		Health:         45,
		MaxHealth:      45,
		MaxATK:         9,
		EXPtogive:      25,
		Coinstogive:    18,
		Initiative:     6,
		AttackMsg:      "The goblin attacks you!",
		CritEvery:      4,
		CritMultiplier: 1.8,
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
		AttackMsg:      "The flower attacks you!",
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
		Health:         30,
		MaxHealth:      30,
		MaxATK:         8,
		EXPtogive:      22,
		Coinstogive:    22,
		Initiative:     4,
		AttackMsg:      "The skeleton attacks you!",
		CritEvery:      3,
		CritMultiplier: 1.8,
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
		Health:         75,
		MaxHealth:      75,
		MaxATK:         12,
		EXPtogive:      60,
		Coinstogive:    60,
		Initiative:     9,
		AttackMsg:      "The Boss Potato attacks you!",
		CritEvery:      3,
		CritMultiplier: 2,
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
		Health:         140,
		MaxHealth:      140,
		MaxATK:         22,
		EXPtogive:      80,
		Coinstogive:    60,
		Initiative:     5,
		AttackMsg:      "The Chocoblast attacks you!",
		CritEvery:      3,
		CritMultiplier: 2.2,
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
		Health:         80,
		MaxHealth:      80,
		MaxATK:         14,
		EXPtogive:      35,
		Coinstogive:    28,
		Initiative:     3,
		AttackMsg:      "The Slime attacks you!",
		CritEvery:      3,
		CritMultiplier: 1.8,
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
		Health:         90,
		MaxHealth:      90,
		MaxATK:         24,
		EXPtogive:      55,
		Coinstogive:    55,
		Initiative:     8,
		AttackMsg:      "The flame burns you !",
		CritEvery:      5,
		CritMultiplier: 2.0,
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
		Health:         110,
		MaxHealth:      110,
		MaxATK:         28,
		EXPtogive:      70,
		Coinstogive:    70,
		Initiative:     10,
		AttackMsg:      "The annoying dog annoys you !",
		CritEvery:      4,
		CritMultiplier: 2.3,
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
		Health:         220,
		MaxHealth:      220,
		MaxATK:         45,
		EXPtogive:      180,
		Coinstogive:    180,
		Initiative:     12,
		AttackMsg:      "The flowers attack !",
		CritEvery:      3,
		CritMultiplier: 2.5,
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
		Health:         380,
		MaxHealth:      380,
		MaxATK:         58,
		EXPtogive:      360,
		Coinstogive:    300,
		Initiative:     13,
		AttackMsg:      "Polaris rends reality around you!",
		CritEvery:      3,
		CritMultiplier: 2.5,
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
	fmt.Printf("%s HP: %.0f / %.0f\n", m.Name, m.Health, m.MaxHealth)
}

func New(id string) *Monster {
	m := Monsters[id]
	return &m
}
