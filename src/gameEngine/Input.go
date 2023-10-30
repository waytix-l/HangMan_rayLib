package gameEngine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) Input() {

	switch {

	case rl.IsKeyPressed(rl.KeyQ): //A
		g.guess = "A"
		

	case rl.IsKeyPressed(rl.KeyB): //B
		g.guess = "B"
		

	case rl.IsKeyPressed(rl.KeyC): //C
		g.guess = "C"
		

	case rl.IsKeyPressed(rl.KeyD): //D
		g.guess = "D"
		

	case rl.IsKeyPressed(rl.KeyE): //E
		g.guess = "E"
		

	case rl.IsKeyPressed(rl.KeyF): //F
		g.guess = "F"
		

	case rl.IsKeyPressed(rl.KeyG): //G
		g.guess = "G"
		

	case rl.IsKeyPressed(rl.KeyH): //H
		g.guess = "H"
		

	case rl.IsKeyPressed(rl.KeyI): //I
		g.guess = "I"
		

	case rl.IsKeyPressed(rl.KeyJ): //J
		g.guess = "J"
		

	case rl.IsKeyPressed(rl.KeyK): //K
		g.guess = "K"
		

	case rl.IsKeyPressed(rl.KeyL): //L
		g.guess = "L"
		

	case rl.IsKeyPressed(rl.KeySemicolon): //M
		g.guess = "M"
		

	case rl.IsKeyPressed(rl.KeyN): //N
		g.guess = "N"
		

	case rl.IsKeyPressed(rl.KeyO): //O
		g.guess = "O"
		

	case rl.IsKeyPressed(rl.KeyP): //P
		g.guess = "P"
		

	case rl.IsKeyPressed(rl.KeyA): //Q
		g.guess = "Q"
		

	case rl.IsKeyPressed(rl.KeyR): //R
		g.guess = "R"
		

	case rl.IsKeyPressed(rl.KeyS): //S
		g.guess = "S"
		

	case rl.IsKeyPressed(rl.KeyT): //T
		g.guess = "T"
		

	case rl.IsKeyPressed(rl.KeyU): //U
		g.guess = "U"
		

	case rl.IsKeyPressed(rl.KeyV): //V
		g.guess = "V"
		

	case rl.IsKeyPressed(rl.KeyZ): //W
		g.guess = "W"
		

	case rl.IsKeyPressed(rl.KeyX): //X
		g.guess = "X"
		

	case rl.IsKeyPressed(rl.KeyY): //Y
		g.guess = "Y"
		

	case rl.IsKeyPressed(rl.KeyW): //Z
		g.guess = "Z"
		

	default :
		g.guess = ""
	}


}
