package gameEngine

import rl "github.com/gen2brain/raylib-go/raylib"

func (m * Menu) InitMenu() {
	m.sound = rl.LoadMusicStream("assets/musique_h.mp3")

	m.Title = rl.LoadTexture("assets/title_hangman.png")

	m.background = rl.LoadTexture("assets/background_dark.jpg")
	m.Sr_background = rl.NewRectangle(0, 0, 564, 1222)
	m.Dr_background = rl.NewRectangle(555, 0, 800, 1222)
	m.Vector_background = rl.NewVector2(0, 0)
	m.color_background = rl.ColorAlpha(rl.White, 0)
	m.timer_background = 0

	m.FontHorror = rl.LoadFont("assets/Another Danger - Demo.otf")
	m.color_title = rl.ColorAlpha(rl.White, 0)
	m.timer_title = 0

	m.Button = rl.LoadTexture("assets/sprite_button.png")
	m.color_button = rl.ColorAlpha(rl.White, 0)
	m.timer_button = 0

	m.color_text_button = rl.ColorAlpha(rl.White, 0)
}