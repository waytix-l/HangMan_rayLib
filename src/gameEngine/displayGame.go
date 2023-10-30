package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Menu) DisplayGame(g *Game) {

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawTexturePro(
		m.background,
		m.Sr_background,
		m.Dr_background,
		m.Vector_background,
		0,
		m.color_background,
	)

	rl.DrawRectangle(540, 0, 15, 1080, m.color_sep)
	rl.DrawRectangle(1350, 0, 15, 1080, m.color_sep)
	rl.DrawRectangle(540, 0, 820, 15, m.color_sep)
	rl.DrawRectangle(540, 1065, 820, 15, m.color_sep)

	rl.DrawRectangleRec(m.rect_text, m.color_rect)
	rl.DrawTextEx(m.FontHorror, "Test a letter", rl.NewVector2(860, 820), 30, 10, rl.White)

	g.underscore_word = ""
	for i := 0; i < len(g.Mot_secret); i++ {
		g.underscore_word += g.Mot_secret[i]
	}

	rl.DrawTextEx(m.test_font, g.underscore_word, rl.NewVector2(float32(880 - (7*len(g.underscore_word))), 865), 50, 20, rl.White)

	rl.EndDrawing()

	switch g.Stage {

	case 0:
		g.Input()
		g.verif(m)

	case 1:
		g.Input()
		g.verif(m)

	case 2:
		g.Input()
		g.verif(m)

	case 3:
		g.Input()
		g.verif(m)
	case 4:
		g.Input()
		g.verif(m)
	case 5:
		g.Input()
		g.verif(m)
	case 6:
		g.Input()
		g.verif(m)
	case 7:
		g.Input()
		g.verif(m)
	case 8:
		g.Input()
		g.verif(m)
	case 9:
		g.Input()
		g.verif(m)
	case 10:
		g.Input()
		g.verif(m)
	case 11:
		g.Input()
		g.verif(m)
	case 12: // Dead Screen
		m.DisplayLose()
	}
}
