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

	Title     rl.Texture2D

	background        rl.Texture2D
	Sr_background     rl.Rectangle
	Dr_background     rl.Rectangle
	Vector_background rl.Vector2
	color_background rl.Color
	timer_background int

	couleur_gris color.RGBA
	color_sep rl.Color

	FontHorror  rl.Font
	color_title rl.Color
	timer_title int

	Button rl.Texture2D
	color_button rl.Color
	timer_button int

	color_text_button rl.Color

	thunder rl.Texture2D
	Sr_thunder rl.Rectangle
	Dr_thunder rl.Rectangle
	Vector_thunder rl.Vector2
	timer_thunder int
}
