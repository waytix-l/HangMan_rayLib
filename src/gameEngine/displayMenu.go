package gameEngine

import rl "github.com/gen2brain/raylib-go/raylib"



func (m *Menu) DisplayMenu() {

	rl.PlayMusicStream(m.sound)

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

	if m.timer_title > 200 && m.color_title.A < 255 {
		m.color_title.A += 5
	}

	rl.DrawTextEx(
		m.FontHorror, 
		"HANGMAN",
		rl.NewVector2(825, 100), 
		60, 
		20, 
		m.color_title,
	)

	if m.timer_button > 280 && m.color_button.A < 255 && m.color_text_button.A < 255 {
		m.color_button.A += 5
		m.color_text_button.A += 5
	}

	rl.DrawTextureEx(
		m.Button,
		rl.NewVector2(760, 300),
		0,
		1.5,
		m.color_button,
	)

	rl.DrawTextEx(
		m.FontHorror, 
		"PLAY",
		rl.NewVector2(880, 340), 
		70, 
		20, 
		m.color_text_button,
	)


	rl.DrawTextureEx(
		m.Button,
		rl.NewVector2(760, 500),
		0,
		1.5,
		m.color_button,
	)

	rl.DrawTextEx(
		m.FontHorror, 
		"QUIT",
		rl.NewVector2(880, 540), 
		70, 
		20, 
		m.color_text_button,
	)

	rl.EndDrawing()
}