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
	Art            string
}

var Monsters = map[string]Monster{
	"Goblin": {
		Name:           "Goblin",
		Health:         35,
		MaxHealth:      35,
		MaxATK:         4,
		EXPtogive:      12,
		Coinstogive:    12,
		Initiative:     5,
		AttackMsg:      "The goblin attacks you!",
		CritEvery:      3,
		CritMultiplier: 2,
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
