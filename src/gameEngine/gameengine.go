package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameEngine struct {
	ScreenWidth  int32
	ScreenHeight int32
	Title        string
}

func (g *GameEngine) InitGameEngine(x int32, y int32, title string) {
	g.ScreenWidth = x
	g.ScreenHeight = y
	g.Title = title
	rl.InitWindow(g.ScreenWidth, g.ScreenHeight, g.Title)
	rl.SetTargetFPS(60)
	rl.ToggleFullscreen()
}

func (g *GameEngine) RunningGameEngine(m *Menu) {
	m.InitMenu()
	rl.PlayMusicStream(m.sound)
	for !rl.WindowShouldClose() {
		switch m.menu {
		case 0:
			m.DisplayIntro()
		case 1:
			m.DisplayMenu()
		}
	}
}
