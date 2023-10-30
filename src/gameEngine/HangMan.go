package gameEngine

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func (g *Game) RandomWord() {
	g.Content, g.Err = os.ReadFile("gameEngine/mot.txt")
	if g.Err != nil {
		fmt.Print(g.Err)
	}
	lines := strings.Split(string(g.Content), "\n")

	rand.Seed(time.Now().UnixNano())
	g.Random_number = rand.Intn(len(lines))

	g.mot = lines[g.Random_number]
	g.Tableau_rune = []rune(g.mot)

	for i := 0; i < len(g.mot); i++ {
		g.Mot_secret = append(g.Mot_secret, "_")
	}
}

func (g *Game) verif(m *Menu) {

	if g.guess != "" {
		for i := 0; i < len(g.guesses); i++ {
			if g.guess == g.guesses[i] {
				fmt.Println("retry")
				g.guess = ""

				g.verif(m)
			}
		}

		g.guesses = append(g.guesses, g.guess)
		g.perdu = true
		for j := 0; j < len(g.Tableau_rune); j++ {
			fmt.Println(string(g.Tableau_rune[j]-32))
			if g.guess == string(g.Tableau_rune[j]-32) {
				fmt.Println("vrai")
				g.Mot_secret[j] = g.guess
				g.perdu = false
			}
		}
		if g.perdu {
			g.Stage++
			g.attemptsLeft--
			g.guess = ""
		}
	}


	g.win = true
	for _, letter := range g.Mot_secret {
		if letter == "_" {
			g.win = false
		}
	}
	if g.win {
		m.menu = 3
	}
}

