package gameEngine

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (m *Menu) DisplayWin() {

	m.timer_win++

	rl.BeginDrawing()

	if m.timer_win > 100 && m.timer_win < 300 && m.background_color_dead.A < 150 {
		m.background_color_dead.A += 2
	}

	rl.DrawRectangle(530, 0, 850, 1222, m.background_color_dead)

	if m.timer_win > 200 && m.timer_win < 350 && m.color_uaredead.A <= 250 {
		m.color_uaredead.A += 2
	}

	if m.timer_win > 350 && m.timer_win < 500 && m.color_quit.A <= 250 {
		m.color_quit.A += 2
	}

	rl.DrawTextEx(m.FontHorror, "THOMAS IS ALIVE", rl.NewVector2(690, 500), 70, 20, m.color_uaredead)
	rl.DrawTextEx(m.FontHorror, "Press enter to quit", rl.NewVector2(650, 600), 50, 20, m.color_quit)

	rl.EndDrawing()

	if rl.IsKeyPressed(rl.KeyEnter) {
		os.Exit(0)
	}

}
