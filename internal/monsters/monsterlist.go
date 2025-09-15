package monsters

import "fmt"

type Monster struct {
	Name           string
	Health         float64
	MaxHealth      float64
	ATK            float64
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
		Health:         40,
		MaxHealth:      40,
		ATK:            5,
		EXPtogive:      10,
		Coinstogive:    10,
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
		Health:         10,
		MaxHealth:      10,
		ATK:            1,
		EXPtogive:      5,
		Coinstogive:    5,
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
		Health:         20,
		MaxHealth:      20,
		ATK:            13,
		EXPtogive:      20,
		Coinstogive:    20,
		Initiative:     4,
		AttackMsg:      "The skeleton attacks you!",
		CritEvery:      3,
		CritMultiplier: 2,
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
		Health:         50,
		MaxHealth:      50,
		ATK:            15,
		EXPtogive:      50,
		Coinstogive:    50,
		Initiative:     10,
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
