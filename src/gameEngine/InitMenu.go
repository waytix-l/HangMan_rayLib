package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Menu) InitMenu() {

	m.menu = 0

	//----- Intro -----//

	m.timer_intro = 0

	//----- Menu -----//

	m.sound = rl.LoadMusicStream("assets/musique_h.mp3")

	m.Title = rl.LoadTexture("assets/title_hangman.png")

	m.background = rl.LoadTexture("assets/background_dark.jpg")
	m.Sr_background = rl.NewRectangle(0, 0, 564, 1222)
	m.Dr_background = rl.NewRectangle(555, 0, 800, 1222)
	m.Vector_background = rl.NewVector2(0, 0)
	m.color_background = rl.ColorAlpha(rl.White, 0)
	m.timer_background = 0

	m.couleur_gris = rl.ColorFromHSV(243, 0.23, 0.16)
	m.color_sep = rl.ColorAlpha(m.couleur_gris, 0)

	m.FontHorror = rl.LoadFont("assets/Another Danger - Demo.otf")
	m.color_title = rl.ColorAlpha(rl.White, 0)
	m.timer_title = 0

	m.Button = rl.LoadTexture("assets/bouton_test.png")
	m.ButtonHover = rl.LoadTexture("assets/buttonhover.png")
	m.color_button = rl.ColorAlpha(rl.White, 0)
	m.timer_button = 0

	m.color_text_button = rl.ColorAlpha(rl.White, 0)

	m.close_menu = false
	m.timer_close_menu = 0

	m.color_retry_text = rl.ColorAlpha(rl.White, 255)
	m.tiemr_retry_text = 0

	m.test_font = rl.LoadFont("assets/Rough Battle.ttf")
	m.rect_text = rl.NewRectangle(650, 800, 600, 170)
	m.color_rect = rl.ColorAlpha(rl.DarkGray, 0.7)

	m.tree1 = rl.LoadTexture("assets/Tree/arbre1.png")
	m.tree2 = rl.LoadTexture("assets/Tree/arbre2.png")
	m.tree3 = rl.LoadTexture("assets/Tree/arbre3.png")
	m.tree4 = rl.LoadTexture("assets/Tree/arbre4.png")
	m.tree5 = rl.LoadTexture("assets/Tree/arbre5.png")
	m.tree6 = rl.LoadTexture("assets/Tree/arbre6.png")
	m.tree7 = rl.LoadTexture("assets/Tree/arbre7.png")
	m.tree8 = rl.LoadTexture("assets/Tree/arbre8.png")
	m.tree9 = rl.LoadTexture("assets/Tree/arbre9.png")
	m.tree10 = rl.LoadTexture("assets/Tree/arbre10.png")

	m.timer_dead = 0
	m.background_color_dead = rl.ColorAlpha(rl.Black, 0)
	m.color_uaredead = rl.ColorAlpha(rl.Red, 0)
	m.color_quit = rl.ColorAlpha(rl.Red, 0)

	m.timer_win = 0
}
