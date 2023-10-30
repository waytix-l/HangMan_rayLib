package gameEngine

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)



func (m *Menu) DisplayLose() {

	m.timer_dead++
	
	rl.BeginDrawing()

	if m.timer_dead > 100 && m.timer_dead < 300 && m.background_color_dead.A < 150 {
		m.background_color_dead.A += 2
	}
	
	rl.DrawRectangle(530, 0, 850, 1222, m.background_color_dead)

	if m.timer_dead > 200 && m.timer_dead < 350 && m.color_uaredead.A <= 250 {
		m.color_uaredead.A += 2
	}

	if m.timer_dead > 350 && m.timer_dead < 500 && m.color_quit.A <= 250 {
		m.color_quit.A += 2
	}


	rl.DrawTextEx(m.FontHorror, "THOMAS IS DEAD", rl.NewVector2(710, 500), 70, 20, m.color_uaredead)
	rl.DrawTextEx(m.FontHorror, "Press enter to quit", rl.NewVector2(650, 600), 50, 20, m.color_quit)

	rl.EndDrawing()


	if rl.IsKeyPressed(rl.KeyEnter) {
		os.Exit(0)
	}
}