package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Menu struct {
	sound rl.Music

	Title     rl.Texture2D

	background        rl.Texture2D
	Sr_background     rl.Rectangle
	Dr_background     rl.Rectangle
	Vector_background rl.Vector2
	color_background rl.Color
	timer_background int

	FontHorror  rl.Font
	color_title rl.Color
	timer_title int

	Button rl.Texture2D
	color_button rl.Color
	timer_button int

	color_text_button rl.Color
}
