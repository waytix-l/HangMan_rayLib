package gameEngine

import rl "github.com/gen2brain/raylib-go/raylib"

func (m *Menu) DisplayIntro() {
	m.timer_intro++

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawTextEx(m.FontHorror, "HangMan", rl.NewVector2(840, 650), 60, 20, m.color_title)
	rl.DrawTextEx(m.FontHorror, "Made By Lucas and Sam", rl.NewVector2(605, 800), 60, 20, m.color_title)
	
	if m.timer_intro > 200 && m.color_title.A < 255 {
		m.color_title.A += 5
	}

	if m.timer_intro >= 400 && m.color_title.A > 0 {
		m.color_title.A -= 5
	}

	if m.timer_intro >= 600 {
		m.color_title.A = 0
		m.menu = 1
	}

	if rl.IsMouseButtonPressed(0) {
		m.color_title.A = 0
		m.menu = 1
	}


	rl.EndDrawing()


}