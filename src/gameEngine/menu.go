package gameEngine

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Menu struct {
	menu int

	//----- Intro -----//

	timer_intro int

	//----- Menu -----//

	sound rl.Music

	Title rl.Texture2D

	background        rl.Texture2D
	Sr_background     rl.Rectangle
	Dr_background     rl.Rectangle
	Vector_background rl.Vector2
	color_background  rl.Color
	timer_background  int

	couleur_gris color.RGBA
	color_sep    rl.Color

	FontHorror  rl.Font
	color_title rl.Color
	timer_title int

	Button      rl.Texture2D
	ButtonHover rl.Texture2D

	color_button rl.Color
	timer_button int

	color_text_button rl.Color

	close_menu       bool
	timer_close_menu int

	test_font        rl.Font
	rect_text        rl.Rectangle
	color_rect       rl.Color
	tiemr_retry_text int
	color_retry_text rl.Color

	tree1  rl.Texture2D
	tree2  rl.Texture2D
	tree3  rl.Texture2D
	tree4  rl.Texture2D
	tree5  rl.Texture2D
	tree6  rl.Texture2D
	tree7  rl.Texture2D
	tree8  rl.Texture2D
	tree9  rl.Texture2D
	tree10 rl.Texture2D

	timer_dead            int
	background_color_dead rl.Color
	color_uaredead        rl.Color
	color_quit            rl.Color

	timer_win int
}
