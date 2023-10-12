package gameEngine

import rl "github.com/gen2brain/raylib-go/raylib"



func (m *Menu) DisplayMenu() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.EndDrawing()
}