package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Menu) DisplayMenu() {

	rl.UpdateMusicStream(m.sound)

	m.timer_title++
	m.timer_button++
	m.timer_background++

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	if m.timer_background > 100 && m.color_background.A < 255 {
		m.color_background.A += 5
	}

	rl.DrawTexturePro(
		m.background,
		m.Sr_background,
		m.Dr_background,
		m.Vector_background,
		0,
		m.color_background,
	)


	if m.timer_background > 100 && m.color_sep.A < 255 {
		m.color_sep.A += 5
	}

	rl.DrawRectangle(540, 0, 15, 1080, m.color_sep)
	rl.DrawRectangle(1350, 0, 15, 1080, m.color_sep)
	rl.DrawRectangle(540, 0, 820, 15, m.color_sep)
	rl.DrawRectangle(540, 1065, 820, 15, m.color_sep)


	if m.timer_title > 200 && m.color_title.A < 255 && m.timer_title < 280 {
		m.color_title.A += 5
	}

	
	rl.DrawTextEx(
		m.FontHorror,
		"HANGMAN",
		rl.NewVector2(830, 140),
		60,
		20,
		m.color_title,
	)

	if m.timer_button > 280 && m.color_button.A < 255 && m.color_text_button.A < 255 && m.timer_button < 380{
		m.color_button.A += 5
		m.color_text_button.A += 5
	}

	rl.DrawTextureEx(
		m.Button,
		rl.NewVector2(680, 875),
		0,
		2,
		m.color_button,
	)


	rl.DrawTextureEx(
		m.Button,
		rl.NewVector2(980, 875),
		0,
		2,
		m.color_button,
	)


	x_mouse := rl.GetMouseX()
	y_mouse := rl.GetMouseY()

	//----- Quit Button -----//

	if x_mouse > 980 && x_mouse < 1230 && y_mouse > 875 && y_mouse < 975 {
		rl.DrawTextureEx(
			m.ButtonHover,
			rl.NewVector2(980, 875),
			0,
			2,
			m.color_button,
		)
		if rl.IsMouseButtonPressed(0) {
			rl.CloseWindow()
		}
	}

	//----- Play Button -----//

	if x_mouse > 680 && x_mouse < 930 && y_mouse > 875 && y_mouse < 975 {
		rl.DrawTextureEx(
			m.ButtonHover,
			rl.NewVector2(680, 875),
			0,
			2,
			m.color_button,
		)
		if rl.IsMouseButtonPressed(0) {
			m.close_menu = true
		}
	}

	if m.close_menu {
		m.timer_close_menu++
		if m.color_button.A > 0 && m.color_button.A <= 255 {
			m.color_button.A -= 5
		}
		
		if m.color_text_button.A > 0 && m.color_text_button.A <= 255{
			m.color_text_button.A -= 5
		}
		if m.color_title.A > 0 && m.color_title.A <= 255{
			m.color_title.A -= 5
		}
		if m.timer_close_menu > 130 {
			m.close_menu = false
			m.timer_close_menu = 0
			m.menu = 2
		}
	}

	
	rl.DrawTextEx(
		m.FontHorror,
		"PLAY",
		rl.NewVector2(735, 890),
		70,
		20,
		m.color_text_button,
	)

	rl.DrawTextEx(
		m.FontHorror,
		"QUIT",
		rl.NewVector2(1035, 890),
		70,
		20,
		m.color_text_button,
	)

	rl.EndDrawing()
}
