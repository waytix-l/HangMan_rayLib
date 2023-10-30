package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

type GameEngine struct {
	ScreenWidth  int32
	ScreenHeight int32
	Title        string
}

func (ge *GameEngine) InitGameEngine(x int32, y int32, title string) {
	ge.ScreenWidth = x
	ge.ScreenHeight = y
	ge.Title = title
	rl.InitWindow(ge.ScreenWidth, ge.ScreenHeight, ge.Title)
	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()
}

func (ge *GameEngine) RunningGameEngine(m *Menu) {
	var g Game
	m.InitMenu()
	rl.PlayMusicStream(m.sound)
	g.RandomWord()
	g.attemptsLeft = 12
	fmt.Println(g.Mot_secret)
	fmt.Println(g.mot)
	fmt.Println(g.underscore_word)
	for !rl.WindowShouldClose() {
		switch m.menu {
		case 0:
			m.DisplayIntro()
		case 1:
			m.DisplayMenu()
		case 2:
			m.DisplayGame(&g)
		}
	}
}
