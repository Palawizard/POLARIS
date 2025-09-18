package monsters

import (
	"fmt"
	"projet-red_POLARIS/utils"
)

// Monster is the runtime state for an enemy during a fight.
// Values like Health mutate; others (e.g., MaxHealth) are static per instance.
type Monster struct {
	Name           string  // Display name
	Health         float64 // Current HP
	MaxHealth      float64 // Max HP
	MaxATK         float64 // Upper bound for base attack damage roll
	EXPToGive      float64 // EXP awarded on defeat
	CoinsToGive    float64 // Coins awarded on defeat
	Initiative     float64 // Turn order priority
	AttackMsg      string  // Flavor text shown when attacking
	CritEvery      int     // Target crit cadence (used by crit logic)
	CritMultiplier float64 // Damage multiplier applied on crit
	Loot           string  // Item/skill/equipment ID dropped on defeat
	SinceLastCrit  int     // Internal counter updated by the attack pattern
	Art            string  // Optional ASCII art header
}

// Monsters holds the immutable templates used to spawn instances via New.
var Monsters = map[string]Monster{
	"Goblin": {
		Name:           "Goblin",
		Health:         42,
		MaxHealth:      42,
		MaxATK:         7,
		EXPToGive:      25,
		CoinsToGive:    18,
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
		EXPToGive:      6,
		CoinsToGive:    6,
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
		EXPToGive:      22,
		CoinsToGive:    22,
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
		EXPToGive:      60,
		CoinsToGive:    60,
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
		EXPToGive:      85,
		CoinsToGive:    65,
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
		EXPToGive:      35,
		CoinsToGive:    28,
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
		EXPToGive:      55,
		CoinsToGive:    55,
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
		EXPToGive:      70,
		CoinsToGive:    70,
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
		EXPToGive:      180,
		CoinsToGive:    180,
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
		EXPToGive:      360,
		CoinsToGive:    300,
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

// PrintHeader draws the monster header (ASCII art if any) and its HP line.
func PrintHeader(m *Monster) {
	if m == nil {
		return
	}
	if m.Art != "" {
		fmt.Println(m.Art)
	}
	fmt.Printf("%s HP: %s\n", m.Name, utils.HPString(m.Health, m.MaxHealth))
}

// New returns a fresh instance copied from the template map.
// Callers mutate the returned pointer without affecting the template.
func New(id string) *Monster {
	m := Monsters[id]
	return &m
}
