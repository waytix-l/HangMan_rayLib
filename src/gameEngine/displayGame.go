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

	

	switch g.Stage {

	case 0:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree1, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 1:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree2, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 2:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree3, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 3:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree4, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 4:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree5, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 5:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree6, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 6:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree7, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 7:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree8, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 8:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree9, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 9:
		g.Input()
		g.verif(m)
		rl.DrawTextureEx(m.tree10, rl.NewVector2(500, -70), 0, 1.5, rl.White)

	case 10: // Dead Screen
		m.menu = 4
	}

	rl.EndDrawing()
}
